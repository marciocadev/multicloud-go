package gcp

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	"github.com/marciocadev/multicloud-go/cloud"
	"github.com/marciocadev/multicloud-go/function/event"
	"github.com/marciocadev/multicloud-go/function/handler"
)

// GCPWrapper implementa o wrapper para Google Cloud Functions
type GCPWrapper struct {
	handler handler.HandlerFunc
}

// Handle processa eventos GCP Cloud Functions
func (w *GCPWrapper) Handle(ctx context.Context, event interface{}) (interface{}, error) {
	req, err := w.parseGCPEvent(event)
	if err != nil {
		return nil, fmt.Errorf("erro ao fazer parse do evento GCP: %w", err)
	}

	resp, err := w.handler(ctx, req)
	if err != nil {
		return nil, err
	}

	return w.convertToGCPResponse(resp)
}

func (w *GCPWrapper) parseGCPEvent(rawEvent interface{}) (*event.CloudRequest, error) {
	// Estrutura para eventos HTTP do GCP
	type GCPHTTPRequest struct {
		Method      string            `json:"method"`
		URL         string            `json:"url"`
		Headers     map[string]string `json:"headers"`
		Body        string            `json:"body"`
		QueryParams map[string]string `json:"query"`
	}

	// Estrutura para eventos PubSub do GCP
	type PubSubMessage struct {
		Data        string            `json:"data"`
		Attributes  map[string]string `json:"attributes"`
		MessageID   string            `json:"messageId"`
		PublishTime string            `json:"publishTime"`
	}

	// Tentar converter para JSON se necessário
	var eventJSON []byte
	switch e := rawEvent.(type) {
	case []byte:
		eventJSON = e
	case string:
		eventJSON = []byte(e)
	default:
		var err error
		eventJSON, err = json.Marshal(e)
		if err != nil {
			return nil, err
		}
	}

	// Tentar parse como HTTP request
	var httpReq GCPHTTPRequest
	if err := json.Unmarshal(eventJSON, &httpReq); err == nil && httpReq.Method != "" {
		return &event.CloudRequest{
			Provider:  cloud.GCP,
			EventType: event.HTTPEvent,
			HTTPRequest: &event.HTTPRequest{
				Method:      httpReq.Method,
				Path:        httpReq.URL,
				Headers:     httpReq.Headers,
				QueryParams: httpReq.QueryParams,
				Body:        httpReq.Body,
				IsBase64:    false,
			},
			RawEvent: rawEvent,
		}, nil
	}

	// Tentar parse como PubSub message
	var pubsubMsg PubSubMessage
	if err := json.Unmarshal(eventJSON, &pubsubMsg); err == nil && pubsubMsg.MessageID != "" {
		publishTime, _ := time.Parse(time.RFC3339, pubsubMsg.PublishTime)

		return &event.CloudRequest{
			Provider:  cloud.GCP,
			EventType: event.MessageEvent,
			Message: &event.CloudMessage{
				ID:          pubsubMsg.MessageID,
				Body:        pubsubMsg.Data,
				Attributes:  pubsubMsg.Attributes,
				PublishTime: publishTime,
			},
			RawEvent: rawEvent,
		}, nil
	}

	return nil, fmt.Errorf("tipo de evento GCP não suportado")
}

func (w *GCPWrapper) convertToGCPResponse(resp *event.CloudResponse) (interface{}, error) {
	if resp == nil {
		return nil, nil
	}

	// GCP espera uma estrutura específica para respostas HTTP
	return map[string]interface{}{
		"statusCode":      resp.StatusCode,
		"headers":         resp.Headers,
		"body":            resp.Body,
		"isBase64Encoded": resp.IsBase64,
	}, nil
}
