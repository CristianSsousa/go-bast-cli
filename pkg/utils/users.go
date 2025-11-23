package utils

import (
	"os"
	"os/user"
	"runtime"
)

// GetCurrentUser retorna o nome do usuário atual
func GetCurrentUser() (string, error) {
	currentUser, err := user.Current()
	if err != nil {
		// Fallback para variáveis de ambiente
		if runtime.GOOS == "windows" {
			if username := os.Getenv("USERNAME"); username != "" {
				return username, nil
			}
		} else {
			if username := os.Getenv("USER"); username != "" {
				return username, nil
			}
		}
		return "", err
	}
	return currentUser.Username, nil
}

// GetCurrentUserInfo retorna todas as informações do usuário
func GetCurrentUserInfo() (*user.User, error) {
	return user.Current()
}

// GetCurrentUserHome retorna o diretório home do usuário atual
func GetCurrentUserHome() (string, error) {
	currentUser, err := user.Current()
	if err != nil {
		// Fallback para UserHomeDir
		return os.UserHomeDir()
	}
	return currentUser.HomeDir, nil
}
