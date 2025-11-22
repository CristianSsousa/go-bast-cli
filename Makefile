.PHONY: help build test clean install run lint fmt vet coverage release docker

# Variáveis
BINARY_NAME=bast
VERSION?=1.0.0
BUILD_DIR=bin
MAIN_PATH=.
GO_FILES=$(shell find . -name '*.go' -not -path './vendor/*')

# Cores para output
GREEN=\033[0;32m
YELLOW=\033[1;33m
NC=\033[0m # No Color

help: ## Mostra esta mensagem de ajuda
	@echo "$(GREEN)Comandos disponíveis:$(NC)"
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | awk 'BEGIN {FS = ":.*?## "}; {printf "  $(YELLOW)%-15s$(NC) %s\n", $$1, $$2}'

build: ## Compila o projeto
	@echo "$(GREEN)Compilando...$(NC)"
	@mkdir -p $(BUILD_DIR)
	@go build -ldflags "-X main.version=$(VERSION)" -o $(BUILD_DIR)/$(BINARY_NAME) $(MAIN_PATH)
	@echo "$(GREEN)Build concluído: $(BUILD_DIR)/$(BINARY_NAME)$(NC)"

build-all: ## Compila para múltiplas plataformas
	@echo "$(GREEN)Compilando para múltiplas plataformas...$(NC)"
	@mkdir -p $(BUILD_DIR)
	@GOOS=linux GOARCH=amd64 go build -ldflags "-X main.version=$(VERSION)" -o $(BUILD_DIR)/$(BINARY_NAME)-linux-amd64 $(MAIN_PATH)
	@GOOS=linux GOARCH=arm64 go build -ldflags "-X main.version=$(VERSION)" -o $(BUILD_DIR)/$(BINARY_NAME)-linux-arm64 $(MAIN_PATH)
	@GOOS=windows GOARCH=amd64 go build -ldflags "-X main.version=$(VERSION)" -o $(BUILD_DIR)/$(BINARY_NAME)-windows-amd64.exe $(MAIN_PATH)
	@GOOS=darwin GOARCH=amd64 go build -ldflags "-X main.version=$(VERSION)" -o $(BUILD_DIR)/$(BINARY_NAME)-darwin-amd64 $(MAIN_PATH)
	@GOOS=darwin GOARCH=arm64 go build -ldflags "-X main.version=$(VERSION)" -o $(BUILD_DIR)/$(BINARY_NAME)-darwin-arm64 $(MAIN_PATH)
	@echo "$(GREEN)Builds concluídos em $(BUILD_DIR)/$(NC)"

test: ## Executa os testes
	@echo "$(GREEN)Executando testes...$(NC)"
	@go test -v -race -coverprofile=coverage.out ./...

test-coverage: test ## Executa testes com coverage
	@echo "$(GREEN)Gerando relatório de coverage...$(NC)"
	@go tool cover -html=coverage.out -o coverage.html
	@echo "$(GREEN)Relatório gerado: coverage.html$(NC)"

test-short: ## Executa testes rápidos (sem race)
	@go test -v -short ./...

clean: ## Remove arquivos gerados
	@echo "$(GREEN)Limpando...$(NC)"
	@rm -rf $(BUILD_DIR)
	@rm -f coverage.out coverage.html
	@go clean
	@echo "$(GREEN)Limpeza concluída$(NC)"

install: build ## Instala o binário no sistema
	@echo "$(GREEN)Instalando...$(NC)"
	@cp $(BUILD_DIR)/$(BINARY_NAME) /usr/local/bin/$(BINARY_NAME)
	@echo "$(GREEN)Instalado em /usr/local/bin/$(BINARY_NAME)$(NC)"

run: build ## Compila e executa o projeto
	@echo "$(GREEN)Executando...$(NC)"
	@./$(BUILD_DIR)/$(BINARY_NAME)

lint: ## Executa o linter
	@echo "$(GREEN)Executando linter...$(NC)"
	@if command -v golangci-lint > /dev/null; then \
		golangci-lint run; \
	else \
		echo "$(YELLOW)golangci-lint não encontrado. Instale com: go install github.com/golangci/golangci-lint/cmd/golangci-lint@latest$(NC)"; \
	fi

fmt: ## Formata o código
	@echo "$(GREEN)Formatando código...$(NC)"
	@go fmt ./...

vet: ## Executa go vet
	@echo "$(GREEN)Executando go vet...$(NC)"
	@go vet ./...

coverage: test-coverage ## Alias para test-coverage

release: ## Cria release usando goreleaser (dry-run)
	@echo "$(GREEN)Criando release...$(NC)"
	@if command -v goreleaser > /dev/null; then \
		goreleaser release --snapshot --clean; \
	else \
		echo "$(YELLOW)goreleaser não encontrado. Instale com: go install github.com/goreleaser/goreleaser@latest$(NC)"; \
	fi

release-prod: ## Cria release de produção usando goreleaser
	@echo "$(GREEN)Criando release de produção...$(NC)"
	@if command -v goreleaser > /dev/null; then \
		goreleaser release --clean; \
	else \
		echo "$(YELLOW)goreleaser não encontrado. Instale com: go install github.com/goreleaser/goreleaser@latest$(NC)"; \
	fi

docker: ## Cria imagem Docker
	@echo "$(GREEN)Construindo imagem Docker...$(NC)"
	@if [ -f Dockerfile ]; then \
		docker build -t $(BINARY_NAME):$(VERSION) .; \
	else \
		echo "$(YELLOW)Dockerfile não encontrado$(NC)"; \
	fi

deps: ## Baixa dependências
	@echo "$(GREEN)Baixando dependências...$(NC)"
	@go mod download
	@go mod tidy

update-deps: ## Atualiza dependências
	@echo "$(GREEN)Atualizando dependências...$(NC)"
	@go get -u ./...
	@go mod tidy

check: fmt vet lint test ## Executa todas as verificações (fmt, vet, lint, test)

ci: check build ## Executa pipeline CI completo

