# Etapa 1: Build
FROM golang:1.23-alpine AS builder

# Defina o diretório de trabalho
WORKDIR /app

# Copie o arquivo go.mod e go.sum para a construção
COPY go.mod go.sum ./

# Baixe as dependências
RUN go mod tidy

# Copie o restante do código para o diretório de trabalho
COPY . .

# Compile o binário Go, especificando o arquivo index.go
RUN go build -o app ./index.go

# Etapa 2: Produção
FROM alpine:latest

# Instale dependências mínimas (se necessário)
RUN apk --no-cache add ca-certificates

# Defina o diretório de trabalho
WORKDIR /root/

# Copie o binário compilado da etapa de build
COPY --from=builder /app/app .

# Exponha a porta em que o app irá rodar (por exemplo, 8080)
EXPOSE 8080

# Comando para rodar o aplicativo Go
CMD ["./app"]
