package config

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/CristianSsousa/go-bast-cli/internal/constants"
	"github.com/spf13/viper"
)

// Config representa a estrutura de configuração da aplicação
type Config struct {
	App      AppConfig      `mapstructure:"app"`
	Logging  LoggingConfig  `mapstructure:"logging"`
	Server   ServerConfig   `mapstructure:"server"`
	Features FeaturesConfig `mapstructure:"features"`
}

// AppConfig configurações gerais da aplicação
type AppConfig struct {
	Name        string `mapstructure:"name"`
	Version     string `mapstructure:"version"`
	Description string `mapstructure:"description"`
	Author      string `mapstructure:"author"`
}

// LoggingConfig configurações de logging
type LoggingConfig struct {
	Level  string `mapstructure:"level"`  // debug, info, warn, error
	Format string `mapstructure:"format"` // text, json
}

// ServerConfig configurações do servidor
type ServerConfig struct {
	DefaultPort int    `mapstructure:"default_port"`
	DefaultHost string `mapstructure:"default_host"`
	Timeout     int    `mapstructure:"timeout"`
}

// FeaturesConfig configurações de features
type FeaturesConfig struct {
	AutoUpdate bool `mapstructure:"auto_update"`
	Verbose    bool `mapstructure:"verbose"`
}

var (
	// Cfg é a instância global de configuração
	Cfg *Config
)

// Init inicializa a configuração usando Viper
func Init(configPath string) error {
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")

	// Adicionar caminhos de busca
	if configPath != "" {
		viper.SetConfigFile(configPath)
	} else {
		// Caminhos padrão
		home, err := os.UserHomeDir()
		if err == nil {
			viper.AddConfigPath(filepath.Join(home, constants.ConfigDirName))
		}
		viper.AddConfigPath(".")
		viper.AddConfigPath("./config")
	}

	// Variáveis de ambiente
	viper.SetEnvPrefix(constants.EnvPrefix)
	viper.AutomaticEnv()

	// Valores padrão
	setDefaults()

	// Ler arquivo de configuração
	if err := viper.ReadInConfig(); err != nil {
		// Se o arquivo não existe, usar defaults (não é erro)
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			// Arquivo não encontrado é ok, usar defaults
		} else if os.IsNotExist(err) {
			// Arquivo não existe, usar defaults
		} else {
			// Outro tipo de erro ao ler arquivo
			return fmt.Errorf("erro ao ler arquivo de configuração: %w", err)
		}
	}

	// Carregar configuração na estrutura
	Cfg = &Config{}
	if err := viper.Unmarshal(Cfg); err != nil {
		return fmt.Errorf("erro ao fazer unmarshal da configuração: %w", err)
	}

	return nil
}

// setDefaults define valores padrão
func setDefaults() {
	viper.SetDefault("app.name", constants.AppName)
	viper.SetDefault("app.version", constants.AppVersion)
	viper.SetDefault("app.description", constants.AppDescription)
	viper.SetDefault("app.author", constants.AppAuthor)

	viper.SetDefault("logging.level", constants.LogLevelInfo)
	viper.SetDefault("logging.format", constants.LogFormatText)

	viper.SetDefault("server.default_port", constants.DefaultPort)
	viper.SetDefault("server.default_host", constants.DefaultHost)
	viper.SetDefault("server.timeout", constants.DefaultTimeout)

	viper.SetDefault("features.auto_update", false)
	viper.SetDefault("features.verbose", false)
}

// Get retorna a configuração atual
func Get() *Config {
	if Cfg == nil {
		Init("")
	}
	return Cfg
}

// Save salva a configuração atual em arquivo
func Save() error {
	home, err := os.UserHomeDir()
	if err != nil {
		return fmt.Errorf("erro ao obter diretório home: %w", err)
	}

	configDir := filepath.Join(home, constants.ConfigDirName)
	if err := os.MkdirAll(configDir, constants.ConfigDirPerm); err != nil {
		return fmt.Errorf("erro ao criar diretório de configuração: %w", err)
	}

	configFile := filepath.Join(configDir, constants.ConfigFileName)
	return viper.WriteConfigAs(configFile)
}

// Set define um valor de configuração
func Set(key string, value interface{}) {
	viper.Set(key, value)
	if Cfg != nil {
		viper.Unmarshal(Cfg)
	}
}

// GetString retorna um valor string da configuração
func GetString(key string) string {
	return viper.GetString(key)
}

// GetInt retorna um valor int da configuração
func GetInt(key string) int {
	return viper.GetInt(key)
}

// GetBool retorna um valor bool da configuração
func GetBool(key string) bool {
	return viper.GetBool(key)
}
