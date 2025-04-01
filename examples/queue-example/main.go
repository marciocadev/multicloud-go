package main

import (
	"context"
	"log"
	"os"

	"github.com/marciocadev/multicloud-go/queue"
)

func main() {
	// Configurações
	queueURL := os.Getenv("QUEUE_URL")
	if queueURL == "" {
		log.Fatal("QUEUE_URL environment variable is required")
	}

	// Criar cliente da fila
	client, err := queue.GetQueueClient(queue.AWS, "us-east-1", queueURL)
	if err != nil {
		log.Fatalf("Erro ao criar cliente da fila: %v", err)
	}

	// Contexto para a operação
	ctx := context.Background()

	// Enviar mensagem
	message := "Olá, esta é uma mensagem de teste!"
	err = client.SendMessage(ctx, message)
	if err != nil {
		log.Fatalf("Erro ao enviar mensagem: %v", err)
	}

	log.Println("Mensagem enviada com sucesso!")
} 