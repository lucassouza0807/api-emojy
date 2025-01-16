# Etapa 1: Build
FROM golang:1.23-alpine AS builder

# Defina o diretório de trabalho
WORKDIR /app

# Copiar arquivos de dependências
COPY go.mod go.sum ./

# Baixar as dependências
RUN go mod tidy

# Copiar o restante do código
COPY . .

# Compilar o binário Go
RUN go build -o app ./index.go

# Etapa 2: Produção
FROM alpine:latest

# Instale dependências mínimas
RUN apk --no-cache add ca-certificates

# Defina o diretório de trabalho
WORKDIR /root/

# Copiar o binário compilado
COPY --from=builder /app/app .

# Copiar o arquivo .env para o diretório de trabalho
COPY .env /root/.env

# Expor a porta
EXPOSE 8080

# Comando para rodar o aplicativo
CMD ["./app"]
