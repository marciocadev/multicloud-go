# Exemplo de Uso do Pacote de Filas

Este é um exemplo simples que demonstra como usar o pacote de filas multicloud.

## Pré-requisitos

- Go 1.21 ou superior
- Credenciais AWS configuradas
- URL de uma fila SQS AWS

## Configuração

1. Configure suas credenciais AWS:
   ```bash
   export AWS_ACCESS_KEY_ID=sua_access_key
   export AWS_SECRET_ACCESS_KEY=sua_secret_key
   ```

2. Configure a URL da fila:
   ```bash
   export QUEUE_URL=https://sqs.region.amazonaws.com/account-id/queue-name
   ```

## Executando o exemplo

```bash
go run main.go
```

## O que o exemplo faz

1. Cria um cliente de fila para AWS SQS
2. Envia uma mensagem de teste para a fila
3. Exibe o resultado da operação

## Estrutura do Projeto

- `main.go`: Código principal com o exemplo de uso
- `go.mod`: Arquivo de dependências do Go
- `README.md`: Este arquivo com instruções 