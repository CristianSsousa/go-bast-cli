# Go Default Template

Este é um template completo para projetos Go no GitHub, incluindo CI/CD, CodeQL, Dependabot e todas as configurações necessárias para começar rapidamente.

## 🚀 O que está incluído

- ✅ **CI/CD** com GitHub Actions
  - Build e testes automatizados
  - Linting com golangci-lint
  - Code coverage
- ✅ **CodeQL** para análise de segurança
- ✅ **Dependabot** para atualizações automáticas de dependências
- ✅ **CODEOWNERS** para revisão de código
- ✅ **Templates** de Pull Request e Issues
- ✅ **Branch Protection** (configurar via UI)
- ✅ Configurações de editor (`.editorconfig`)
- ✅ `.gitignore` otimizado para Go

## 📋 Pré-requisitos

- Git instalado e configurado
- Conta GitHub com permissões para criar repositórios
- Go 1.23 ou superior (recomendado)

## 🎯 Como usar este template

### 1. Criar um novo repositório a partir do template

1. No GitHub, abra este repositório
2. Clique no botão **"Use this template"** (ou acesse: `https://github.com/SEU_USUARIO/go-default-template/generate`)
3. Preencha o nome do novo repositório
4. Escolha se será público ou privado
5. Clique em **"Create repository from template"**

### 2. Clonar o novo repositório

```bash
git clone git@github.com:SEU_USUARIO/novo-repo.git
cd novo-repo
```

### 3. Configurar o projeto

#### Atualizar informações do projeto

1. **Atualizar `dependabot.yml`**:

   - Substitua `YOUR_GITHUB_USERNAME` pelo seu usuário ou time

2. **Atualizar `CODEOWNERS`**:

   - Substitua `YOUR_GITHUB_USERNAME` pelo seu usuário ou time

3. **Criar `go.mod`** (se ainda não existir):

   ```bash
   go mod init github.com/SEU_USUARIO/novo-repo
   ```

4. **Criar estrutura básica do projeto**:
   ```bash
   mkdir -p cmd/api internal/handlers internal/models
   ```

### 4. Configurar Branch Protection (via UI do GitHub)

1. Vá em **Settings** → **Branches**
2. Clique em **Add rule** ou **Add branch protection rule**
3. Em **Branch name pattern**, digite: `main`
4. Marque as opções:
   - ✅ **Require a pull request before merging**
     - Require approvals: `1`
   - ✅ **Require status checks to pass before merging**
     - Selecione: `CI / build-test` e `CodeQL / Analyze`
   - ✅ **Require linear history** (opcional)
   - ✅ **Include administrators** (opcional)
5. Clique em **Create**

### 5. Configurar Secrets (se necessário)

Se você planeja usar:

- **Semantic Release**: Adicione `GH_TOKEN` (Personal Access Token com permissões de escrita)
- **Deploy automático**: Adicione tokens conforme necessário (AWS, GCP, Docker Hub, etc.)

**Como adicionar secrets:**

1. Vá em **Settings** → **Secrets and variables** → **Actions**
2. Clique em **New repository secret**
3. Adicione o nome e valor do secret

### 6. Fazer o primeiro commit

```bash
git add .
git commit -m "chore: initial commit from template"
git push -u origin main
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
feat(auth): adiciona autenticação JWT
fix(api): corrige validação de email
docs(readme): atualiza instruções de instalação
chore(deps): atualiza dependências
```

## 🔄 Branch Strategy

Este template usa a estratégia **GitHub Flow**:

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

## 🧪 Executar testes localmente

```bash
# Executar todos os testes
go test ./...

# Executar com coverage
go test -v -race -coverprofile=coverage.out ./...
go tool cover -html=coverage.out

# Executar linter
golangci-lint run
```

## 📦 Estrutura de diretórios recomendada

```
.
├── cmd/              # Aplicações principais
│   └── api/         # API server
├── internal/         # Código privado da aplicação
│   ├── handlers/    # HTTP handlers
│   ├── models/      # Modelos de dados
│   └── services/    # Lógica de negócio
├── pkg/             # Código público (bibliotecas)
├── test/            # Testes de integração
├── .github/         # GitHub Actions e templates
├── go.mod
├── go.sum
└── README.md
```

## 🔧 Configurações opcionais

### Semantic Release (Releases automáticas)

Para configurar releases automáticas baseadas em Conventional Commits:

1. Instale as dependências:

   ```bash
   go install github.com/golangci/golangci-lint/cmd/golangci-lint@latest
   ```

2. Crie `.github/workflows/release.yml`:

   ```yaml
   name: Release
   on:
     push:
       branches: [main]
   jobs:
     release:
       runs-on: ubuntu-latest
       steps:
         - uses: actions/checkout@v4
         - name: Create Release
           # Adicione sua lógica de release aqui
   ```

3. Adicione `GH_TOKEN` como secret (Personal Access Token)

### Code Coverage

O workflow de CI já está configurado para enviar coverage para Codecov. Para habilitar:

1. Conecte seu repositório ao [Codecov](https://codecov.io)
2. O workflow já está configurado para enviar automaticamente

## 📚 Recursos úteis

- [Go Documentation](https://go.dev/doc/)
- [Effective Go](https://go.dev/doc/effective_go)
- [Go Code Review Comments](https://github.com/golang/go/wiki/CodeReviewComments)
- [Conventional Commits](https://www.conventionalcommits.org/)
- [GitHub Actions Documentation](https://docs.github.com/en/actions)

## 🤝 Contribuindo

1. Faça fork do projeto
2. Crie uma branch para sua feature (`git checkout -b feat/AmazingFeature`)
3. Commit suas mudanças (`git commit -m 'feat: Add some AmazingFeature'`)
4. Push para a branch (`git push origin feat/AmazingFeature`)
5. Abra um Pull Request

## 📄 Licença

Este projeto está sob a licença CC0 1.0 Universal. Veja o arquivo [LICENSE](LICENSE) para mais detalhes.

## ⚠️ Notas importantes

- **Substitua `YOUR_GITHUB_USERNAME`** nos arquivos `dependabot.yml` e `CODEOWNERS`
- **Configure Branch Protection** após criar o repositório
- **Adicione secrets** se for usar features que requerem autenticação
- **Ajuste os workflows** conforme suas necessidades específicas

## 🆘 Suporte

Se encontrar problemas ou tiver dúvidas:

1. Verifique se seguiu todos os passos de configuração
2. Abra uma [Issue](https://github.com/SEU_USUARIO/go-default-template/issues)
3. Consulte a documentação do GitHub Actions

---

**Feito com ❤️ para a comunidade Go**
