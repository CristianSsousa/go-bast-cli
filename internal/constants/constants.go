package constants

// App constants
const (
	AppName        = "bast"
	AppVersion     = "1.0.0"
	AppDescription = "Uma CLI moderna construída com Go e Cobra"
	AppAuthor      = "CristianSsousa"
)

// Config constants
const (
	// ConfigDirName nome do diretório de configuração
	ConfigDirName = ".bast"

	// ConfigFileName nome do arquivo de configuração
	ConfigFileName = "config.yaml"

	// ConfigFileExample nome do arquivo de exemplo
	ConfigFileExample = "config.yaml.example"
)

// Logging constants
const (
	// LogLevelDebug nível de log debug
	LogLevelDebug = "debug"

	// LogLevelInfo nível de log info
	LogLevelInfo = "info"

	// LogLevelWarn nível de log warn
	LogLevelWarn = "warn"

	// LogLevelError nível de log error
	LogLevelError = "error"

	// LogFormatText formato de log texto
	LogFormatText = "text"

	// LogFormatJSON formato de log JSON
	LogFormatJSON = "json"
)

// Server constants
const (
	// DefaultPort porta padrão do servidor
	DefaultPort = 8080

	// DefaultHost host padrão do servidor
	DefaultHost = "0.0.0.0"

	// DefaultTimeout timeout padrão em segundos
	DefaultTimeout = 30

	// MinPort porta mínima válida
	MinPort = 1

	// MaxPort porta máxima válida
	MaxPort = 65535
)

// Network constants
const (
	// DefaultNetworkTimeout timeout padrão de rede em segundos
	DefaultNetworkTimeout = 3

	// TCPProtocol protocolo TCP
	TCPProtocol = "tcp"
)

// Environment variable prefixes
const (
	// EnvPrefix prefixo para variáveis de ambiente
	EnvPrefix = "BAST"
)

// File permissions
const (
	// ConfigDirPerm permissões do diretório de configuração
	ConfigDirPerm = 0755

	// ConfigFilePerm permissões do arquivo de configuração
	ConfigFilePerm = 0644
)

// Messages
const (
	// WelcomeMessage mensagem de boas-vindas
	WelcomeMessage = "Bem-vindo ao %s!"

	// HelpMessage mensagem de ajuda
	HelpMessage = "Use '%s --help' para ver os comandos disponíveis."

	// ConfigSavedMessage mensagem de configuração salva
	ConfigSavedMessage = "Configuração salva com sucesso!"

	// ConfigResetMessage mensagem de configuração resetada
	ConfigResetMessage = "Configurações resetadas para valores padrão"
)

// Error messages
const (
	// ErrConfigNotFound erro quando configuração não encontrada
	ErrConfigNotFound = "arquivo de configuração não encontrado"

	// ErrInvalidPort erro quando porta é inválida
	ErrInvalidPort = "porta deve estar entre %d e %d"

	// ErrConfigSave erro ao salvar configuração
	ErrConfigSave = "erro ao salvar configuração"

	// ErrConfigLoad erro ao carregar configuração
	ErrConfigLoad = "erro ao carregar configuração"
)

// Success messages
const (
	// SuccessPortAvailable porta disponível
	SuccessPortAvailable = "Porta %d em %s está DISPONÍVEL"

	// SuccessPortInUse porta em uso
	SuccessPortInUse = "Porta %d em %s está EM USO"

	// SuccessConfigCreated configuração criada
	SuccessConfigCreated = "Arquivo de configuração criado: %s"

	// SuccessConfigSet configuração definida
	SuccessConfigSet = "Configuração '%s' definida como '%v'"
)

// Info messages
const (
	// InfoConfigExists arquivo de configuração já existe
	InfoConfigExists = "Arquivo de configuração já existe: %s"

	// InfoConfigResetHint dica para resetar configuração
	InfoConfigResetHint = "Use '%s config reset' para resetar para valores padrão"

	// InfoConfigEditHint dica para editar configuração
	InfoConfigEditHint = "Você pode editar este arquivo ou usar '%s config set' para modificar valores"
)
