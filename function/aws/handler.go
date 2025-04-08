package aws

import (
	"context"
	"encoding/json"
	"fmt"
	"strconv"
	"time"

	"github.com/aws/aws-lambda-go/events"

	"github.com/marciocadev/multicloud-go/cloud"
	"github.com/marciocadev/multicloud-go/function/event"
	"github.com/marciocadev/multicloud-go/function/handler"
)

// AWSWrapper implementa o wrapper para AWS Lambda
type AWSWrapper struct {
	Handler handler.HandlerFunc
}

// Handle processa eventos AWS Lambda
func (w *AWSWrapper) Handle(ctx context.Context, event interface{}) (interface{}, error) {
	// Converter o evento para um CloudRequest
	req, err := w.parseAWSEvent(event)
	if err != nil {
		return nil, fmt.Errorf("erro ao fazer parse do evento AWS: %w", err)
	}

	// Executar o handler
	resp, err := w.Handler(ctx, req)
	if err != nil {
		return nil, err
	}

	// Converter a resposta de volta para o formato AWS
	return w.convertToAWSResponse(resp)
}

func (w *AWSWrapper) parseAWSEvent(rawEvent interface{}) (*event.CloudRequest, error) {
	// Tentar converter para JSON primeiro se for um map
	var eventJSON []byte
	switch e := rawEvent.(type) {
	case map[string]interface{}:
		var err error
		eventJSON, err = json.Marshal(e)
		if err != nil {
			return nil, err
		}
	}

	// Tentar parse como API Gateway event
	var apiGatewayEvent events.APIGatewayProxyRequest
	if err := json.Unmarshal(eventJSON, &apiGatewayEvent); err == nil && apiGatewayEvent.HTTPMethod != "" {
		return &event.CloudRequest{
			Provider:  cloud.AWS,
			EventType: event.HTTPEvent,
			HTTPRequest: &event.HTTPRequest{
				Method:      apiGatewayEvent.HTTPMethod,
				Path:        apiGatewayEvent.Path,
				Headers:     apiGatewayEvent.Headers,
				QueryParams: apiGatewayEvent.QueryStringParameters,
				Body:        apiGatewayEvent.Body,
				IsBase64:    apiGatewayEvent.IsBase64Encoded,
			},
			RawEvent: rawEvent,
		}, nil
	}

	// Tentar parse como SQS event
	var sqsEvent events.SQSEvent
	if err := json.Unmarshal(eventJSON, &sqsEvent); err == nil && len(sqsEvent.Records) > 0 {
		messages := make([]event.CloudMessage, len(sqsEvent.Records))
		for i, record := range sqsEvent.Records {
			messages[i] = event.CloudMessage{
				ID:          record.MessageId,
				Body:        record.Body,
				Attributes:  record.Attributes,
				PublishTime: getAWSTimestamp(record.Attributes["SentTimestamp"]),
				Source:      record.EventSourceARN,
			}
		}

		return &event.CloudRequest{
			Provider:  cloud.AWS,
			EventType: event.MessageEvent,
			Message:   &messages[0], // Retorna primeira mensagem
			RawEvent:  rawEvent,
		}, nil
	}

	return nil, fmt.Errorf("tipo de evento AWS não suportado")
}

func (w *AWSWrapper) convertToAWSResponse(resp *event.CloudResponse) (interface{}, error) {
	if resp == nil {
		return nil, nil
	}

	return &events.APIGatewayProxyResponse{
		StatusCode:      resp.StatusCode,
		Headers:         resp.Headers,
		Body:            resp.Body,
		IsBase64Encoded: resp.IsBase64,
	}, nil
}

// Função auxiliar para converter timestamp AWS para time.Time
func getAWSTimestamp(timestamp string) time.Time {
	if ts, err := strconv.ParseInt(timestamp, 10, 64); err == nil {
		return time.Unix(ts/1000, 0)
	}
	return time.Time{}
}
