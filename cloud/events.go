package cloud

import "time"

// EventType representa o tipo de evento serverless
type EventType string

const (
	HTTPEvent    EventType = "http"
	MessageEvent EventType = "message"
)

// HTTPRequest representa uma requisição HTTP normalizada
type HTTPRequest struct {
	Method      string
	Path        string
	Headers     map[string]string
	QueryParams map[string]string
	Body        string
	IsBase64    bool
}

// CloudMessage representa uma mensagem genérica que funciona para ambos os provedores
type CloudMessage struct {
	ID          string            // ID da mensagem
	Body        string            // Corpo da mensagem
	Attributes  map[string]string // Atributos da mensagem
	PublishTime time.Time         // Hora de publicação
	Source      string            // Origem da mensagem (tópico/fila)
}

// CloudRequest representa uma requisição normalizada independente do provedor
type CloudRequest struct {
	Provider    CloudProvider
	EventType   EventType
	HTTPRequest *HTTPRequest  // Preenchido se for um evento HTTP
	Message     *CloudMessage // Preenchido se for um evento de mensagem
	RawEvent    interface{}   // Evento original não processado
}

// CloudResponse representa uma resposta normalizada
type CloudResponse struct {
	StatusCode int
	Headers    map[string]string
	Body       string
	IsBase64   bool
}
