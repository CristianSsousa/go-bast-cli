package cmd

import (
	"bytes"
	"os"
	"testing"

	"github.com/CristianSsousa/go-bast-cli/internal/config"
	"github.com/CristianSsousa/go-bast-cli/pkg/utils"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestConfigList(t *testing.T) {
	config.Init("")

	// Capturar stdout e stderr
	oldStdout := os.Stdout
	oldStderr := os.Stderr
	r, w, _ := os.Pipe()
	os.Stdout = w
	os.Stderr = w

	// Executar através do rootCmd com argumentos completos
	rootCmd.SetArgs([]string{"config", "list"})
	err := rootCmd.Execute()

	w.Close()
	os.Stdout = oldStdout
	os.Stderr = oldStderr

	var buf bytes.Buffer
	buf.ReadFrom(r)
	output := buf.String()

	require.NoError(t, err)
	assert.Contains(t, output, "Configurações do bast CLI")
	assert.Contains(t, output, "App:")
	assert.Contains(t, output, "Logging:")
	assert.Contains(t, output, "Server:")
}

func TestConfigGet(t *testing.T) {
	config.Init("")

	// Capturar stdout e stderr
	oldStdout := os.Stdout
	oldStderr := os.Stderr
	r, w, _ := os.Pipe()
	os.Stdout = w
	os.Stderr = w

	// Executar através do rootCmd
	rootCmd.SetArgs([]string{"config", "get", "app.name"})
	err := rootCmd.Execute()

	w.Close()
	os.Stdout = oldStdout
	os.Stderr = oldStderr

	var buf bytes.Buffer
	buf.ReadFrom(r)
	output := buf.String()

	require.NoError(t, err)
	assert.Contains(t, output, "bast")
}

func TestConfigSet(t *testing.T) {
	// Usar diretório temporário para testes
	tmpDir := t.TempDir()
	originalHome := os.Getenv("HOME")
	os.Setenv("HOME", tmpDir)
	defer func() {
		if originalHome != "" {
			os.Setenv("HOME", originalHome)
		} else {
			os.Unsetenv("HOME")
		}
	}()

	config.Init("")

	// Capturar stdout e stderr
	oldStdout := os.Stdout
	oldStderr := os.Stderr
	r, w, _ := os.Pipe()
	os.Stdout = w
	os.Stderr = w

	// Executar através do rootCmd
	rootCmd.SetArgs([]string{"config", "set", "app.name", "test-app"})
	err := rootCmd.Execute()

	w.Close()
	os.Stdout = oldStdout
	os.Stderr = oldStderr

	var buf bytes.Buffer
	buf.ReadFrom(r)
	output := buf.String()

	require.NoError(t, err)
	assert.Contains(t, output, "Configuração")
	assert.Contains(t, output, "app.name")
}

func TestConfigInit(t *testing.T) {
	// Usar diretório temporário para testes
	tmpDir := t.TempDir()
	originalHome := os.Getenv("HOME")
	os.Setenv("HOME", tmpDir)
	defer func() {
		if originalHome != "" {
			os.Setenv("HOME", originalHome)
		} else {
			os.Unsetenv("HOME")
		}
	}()

	config.Init("")

	// Remover arquivo se existir
	configPath, _ := utils.GetConfigPath()
	os.Remove(configPath)

	// Capturar stdout e stderr
	oldStdout := os.Stdout
	oldStderr := os.Stderr
	r, w, _ := os.Pipe()
	os.Stdout = w
	os.Stderr = w

	// Executar através do rootCmd
	rootCmd.SetArgs([]string{"config", "init"})
	err := rootCmd.Execute()

	w.Close()
	os.Stdout = oldStdout
	os.Stderr = oldStderr

	var buf bytes.Buffer
	buf.ReadFrom(r)
	output := buf.String()

	require.NoError(t, err)
	assert.Contains(t, output, "configuração criado")
}
