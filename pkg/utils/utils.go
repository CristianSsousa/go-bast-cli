package utils

import (
	"fmt"
	"os"
	"path/filepath"
	"runtime"

	"github.com/CristianSsousa/go-bast-cli/internal/constants"
)

// GetConfigDir retorna o diretório de configuração baseado no OS
func GetConfigDir() (string, error) {
	home, err := os.UserHomeDir()
	if err != nil {
		return "", fmt.Errorf("erro ao obter diretório home: %w", err)
	}

	var configDir string
	switch runtime.GOOS {
	case "windows":
		configDir = filepath.Join(home, constants.ConfigDirName)
	default:
		configDir = filepath.Join(home, constants.ConfigDirName)
	}

	return configDir, nil
}

// EnsureConfigDir cria o diretório de configuração se não existir
func EnsureConfigDir() error {
	configDir, err := GetConfigDir()
	if err != nil {
		return err
	}

	return os.MkdirAll(configDir, constants.ConfigDirPerm)
}

// GetConfigPath retorna o caminho completo do arquivo de configuração
func GetConfigPath() (string, error) {
	configDir, err := GetConfigDir()
	if err != nil {
		return "", err
	}

	return filepath.Join(configDir, constants.ConfigFileName), nil
}

// FileExists verifica se um arquivo existe
func FileExists(path string) bool {
	_, err := os.Stat(path)
	return !os.IsNotExist(err)
}

// IsDir verifica se o caminho é um diretório
func IsDir(path string) bool {
	info, err := os.Stat(path)
	if err != nil {
		return false
	}
	return info.IsDir()
}
