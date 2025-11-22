# Build stage
FROM golang:1.23-alpine AS builder

WORKDIR /build

# Copiar arquivos de dependências
COPY go.mod go.sum ./
RUN go mod download

# Copiar código fonte
COPY . .

# Build
RUN CGO_ENABLED=0 GOOS=linux go build -ldflags="-w -s" -o bast .

# Runtime stage
FROM alpine:latest

RUN apk --no-cache add ca-certificates tzdata

WORKDIR /app

# Copiar binário do builder
COPY --from=builder /build/bast /app/bast

# Copiar arquivo de configuração exemplo
COPY --from=builder /build/config.yaml.example /app/config.yaml.example

# Criar diretório para configuração
RUN mkdir -p /app/config

# Expor porta padrão
EXPOSE 8080

# Comando padrão
ENTRYPOINT ["/app/bast"]

