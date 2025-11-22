# bast

Uma CLI moderna construída com Go e Cobra. O **bast** fornece uma interface de linha de comando poderosa e extensível para diversas tarefas.

## 🚀 Características

- ✅ **CLI moderna** construída com [Cobra](https://github.com/spf13/cobra)
- ✅ **Comandos extensíveis** e fáceis de adicionar
- ✅ **Autocompletar** para bash, zsh, fish e PowerShell
- ✅ **CI/CD** com GitHub Actions
- ✅ **CodeQL** para análise de segurança
- ✅ **Dependabot** para atualizações automáticas de dependências

## 📋 Pré-requisitos

- Go 1.23 ou superior
- Git instalado e configurado

## 🔧 Instalação

### Instalação Local

```bash
# Clone o repositório
git clone https://github.com/CristianSsousa/go-bast-cli.git
cd go-bast-cli

# Instale as dependências
go mod download

# Compile o projeto
go build -o bast .

# No Windows
go build -o bast.exe .

# Adicione ao PATH (opcional)
# Linux/macOS
sudo mv bast /usr/local/bin/

# Windows: Adicione o diretório ao PATH do sistema
```

### Instalação via Go Install

```bash
go install github.com/CristianSsousa/go-bast-cli@latest
```

## 📖 Uso

### Comandos Disponíveis

```bash
# Ver ajuda geral
bast --help

# Ver versão
bast version

# Cumprimentar alguém
bast greet --name "João" --greeting "Olá"

# Iniciar servidor HTTP
bast serve --port 8080 --host 0.0.0.0

# Informações do sistema
bast info

# Verificar se porta está em uso
bast port 8080

# Gerenciar configurações
bast config list
bast config set default_port 3000
```

### Comandos Detalhados

#### `bast version`

Mostra a versão atual da aplicação.

```bash
bast version
```

#### `bast greet`

Cumprimenta uma pessoa pelo nome.

**Flags:**

- `--name, -n`: Nome da pessoa a ser cumprimentada
- `--greeting, -g`: Saudação personalizada

**Exemplos:**

```bash
bast greet --name "Maria"
bast greet -n "Pedro" -g "Bem-vindo"
```

#### `bast serve`

Inicia um servidor HTTP.

**Flags:**

- `--port, -p`: Porta do servidor (padrão: 8080)
- `--host, -H`: Host do servidor (padrão: 0.0.0.0)
- `--endpoint, -e`: Endpoint principal (padrão: /)

**Exemplos:**

```bash
bast serve
bast serve --port 3000
bast serve -p 3000 -H localhost
```

**Endpoints disponíveis:**

- `GET /`: Página principal
- `GET /health`: Health check

#### `bast info`

Mostra informações detalhadas do sistema operacional, Go e variáveis de ambiente.

**Flags:**

- `--os`: Mostra apenas informações do sistema operacional
- `--go`: Mostra apenas informações do Go
- `--env`: Mostra apenas variáveis de ambiente importantes

**Exemplos:**

```bash
bast info
bast info --os
bast info --go
bast info --env
```

#### `bast port`

Verifica se uma porta está em uso ou disponível.

**Flags:**

- `--host, -H`: Host para verificar a porta (padrão: localhost)
- `--timeout, -t`: Timeout em segundos (padrão: 3)

**Exemplos:**

```bash
bast port 8080
bast port 3000 --host google.com
bast port 22 --timeout 5
```

#### `bast config`

Gerencia configurações persistentes do bast CLI.

**Subcomandos:**

- `list`: Lista todas as configurações
- `get <chave>`: Obtém valor de uma configuração específica
- `set <chave> <valor>`: Define uma configuração
- `reset`: Reseta todas as configurações para valores padrão

**Chaves disponíveis:**

- `default_port`: Porta padrão para o servidor
- `default_host`: Host padrão para o servidor
- `editor`: Editor de texto preferido
- `theme`: Tema de interface
- `auto_update`: Atualização automática (true/false)

**Exemplos:**

```bash
bast config list
bast config get default_port
bast config set default_port 3000
bast config set auto_update true
bast config reset
```

**Localização do arquivo de configuração:**

- Linux/macOS: `~/.bast/config.json`
- Windows: `%USERPROFILE%\.bast\config.json`

### Autocompletar

O Cobra gera automaticamente scripts de autocompletar para vários shells:

```bash
# Bash
bast completion bash > /etc/bash_completion.d/bast

# Zsh
bast completion zsh > "${fpath[1]}/_bast"

# Fish
bast completion fish > ~/.config/fish/completions/bast.fish

# PowerShell
bast completion powershell | Out-String | Invoke-Expression
```

## 🏗️ Estrutura do Projeto

```
.
├── cmd/                  # Comandos CLI
│   ├── root.go          # Comando raiz
│   ├── version.go       # Comando version
│   ├── greet.go         # Comando greet
│   └── serve.go         # Comando serve
├── .github/             # GitHub Actions e templates
│   └── workflows/       # Workflows de CI/CD
├── go.mod
├── go.sum
├── main.go              # Ponto de entrada
└── README.md
```

## 🧪 Desenvolvimento

### Executar testes

```bash
# Executar todos os testes
go test ./...

# Executar com coverage
go test -v -race -coverprofile=coverage.out ./...
go tool cover -html=coverage.out

# Executar linter
golangci-lint run
```

### Adicionar um novo comando

1. Crie um novo arquivo em `cmd/` (ex: `cmd/novo-comando.go`)
2. Defina o comando usando Cobra:

```go
package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
)

var novoComandoCmd = &cobra.Command{
	Use:   "novo-comando",
	Short: "Descrição curta do comando",
	Long:  `Descrição longa do comando`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Executando novo comando!")
	},
}

func init() {
	rootCmd.AddCommand(novoComandoCmd)
	// Adicione flags aqui se necessário
}
```

3. O comando será automaticamente adicionado ao CLI!

### Compilar

```bash
# Compilar para o sistema atual
go build -o bast .

# Compilar para diferentes plataformas
GOOS=linux GOARCH=amd64 go build -o bast-linux-amd64 .
GOOS=windows GOARCH=amd64 go build -o bast-windows-amd64.exe .
GOOS=darwin GOARCH=amd64 go build -o bast-darwin-amd64 .
GOOS=darwin GOARCH=arm64 go build -o bast-darwin-arm64 .
```

## 📝 Conventional Commits

Este projeto usa [Conventional Commits](https://www.conventionalcommits.org/) para padronizar mensagens de commit:

```
<type>(<scope>): <subject>

<body>

<footer>
```

### Tipos de commit

- `feat`: Nova feature
- `fix`: Correção de bug
- `docs`: Mudanças na documentação
- `style`: Formatação, ponto e vírgula faltando, etc (não afeta código)
- `refactor`: Refatoração de código
- `perf`: Melhoria de performance
- `test`: Adiciona ou corrige testes
- `chore`: Mudanças no build, dependências, etc
- `ci`: Mudanças em CI/CD

### Exemplos

```bash
feat(cli): adiciona comando de configuração
fix(serve): corrige timeout do servidor
docs(readme): atualiza instruções de instalação
chore(deps): atualiza dependências
```

## 🔄 Branch Strategy

Este projeto usa a estratégia **GitHub Flow**:

- **`main`**: Branch principal, sempre estável e deployável
- **Feature branches**: Criadas a partir de `main` para novas features/fixes
- **Pull Requests**: Todas as mudanças passam por PR com revisão obrigatória

### Workflow recomendado

```bash
# Criar branch para feature
git checkout -b feat/nova-feature

# Fazer mudanças e commits
git add .
git commit -m "feat: adiciona nova feature"

# Push e criar PR
git push -u origin feat/nova-feature
```

## 🔧 Configurações

### Flags Globais

- `--verbose, -v`: Modo verboso (disponível em todos os comandos)

### Variáveis de Ambiente

O comando `serve` também pode usar variáveis de ambiente:

```bash
export PORT=3000
bast serve
```

## 📚 Recursos Úteis

- [Go Documentation](https://go.dev/doc/)
- [Effective Go](https://go.dev/doc/effective_go)
- [Cobra Documentation](https://github.com/spf13/cobra)
- [Conventional Commits](https://www.conventionalcommits.org/)
- [GitHub Actions Documentation](https://docs.github.com/en/actions)

## 🤝 Contribuindo

Contribuições são bem-vindas! Por favor:

1. Faça fork do projeto
2. Crie uma branch para sua feature (`git checkout -b feat/AmazingFeature`)
3. Commit suas mudanças (`git commit -m 'feat: Add some AmazingFeature'`)
4. Push para a branch (`git push origin feat/AmazingFeature`)
5. Abra um Pull Request

## 📄 Licença

Este projeto está sob a licença CC0 1.0 Universal. Veja o arquivo [LICENSE](LICENSE) para mais detalhes.

## 🆘 Suporte

Se encontrar problemas ou tiver dúvidas:

1. Verifique a documentação acima
2. Abra uma [Issue](https://github.com/CristianSsousa/go-bast-cli/issues)
3. Consulte a documentação do [Cobra](https://github.com/spf13/cobra)

---

**Feito com ❤️ usando Go e Cobra**
