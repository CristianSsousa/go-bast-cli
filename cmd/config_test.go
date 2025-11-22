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

	var buf bytes.Buffer
	configListCmd.SetOut(&buf)
	configListCmd.SetArgs([]string{})

	err := configListCmd.Execute()
	require.NoError(t, err)

	output := buf.String()
	assert.Contains(t, output, "Configurações do bast CLI")
	assert.Contains(t, output, "App:")
	assert.Contains(t, output, "Logging:")
	assert.Contains(t, output, "Server:")
}

func TestConfigGet(t *testing.T) {
	config.Init("")

	var buf bytes.Buffer
	configGetCmd.SetOut(&buf)
	configGetCmd.SetArgs([]string{"app.name"})

	err := configGetCmd.Execute()
	require.NoError(t, err)

	output := buf.String()
	assert.Contains(t, output, "bast")
}

func TestConfigSet(t *testing.T) {
	// Usar diretório temporário para testes
	tmpDir := t.TempDir()
	originalHome := os.Getenv("HOME")
	os.Setenv("HOME", tmpDir)
	defer os.Setenv("HOME", originalHome)

	config.Init("")

	var buf bytes.Buffer
	configSetCmd.SetOut(&buf)
	configSetCmd.SetArgs([]string{"app.name", "test-app"})

	err := configSetCmd.Execute()
	require.NoError(t, err)

	output := buf.String()
	assert.Contains(t, output, "Configuração 'app.name' definida")
}

func TestConfigInit(t *testing.T) {
	// Usar diretório temporário para testes
	tmpDir := t.TempDir()
	originalHome := os.Getenv("HOME")
	os.Setenv("HOME", tmpDir)
	defer os.Setenv("HOME", originalHome)

	config.Init("")

	// Remover arquivo se existir
	configPath, _ := utils.GetConfigPath()
	os.Remove(configPath)

	var buf bytes.Buffer
	configInitCmd.SetOut(&buf)
	configInitCmd.SetArgs([]string{})

	err := configInitCmd.Execute()
	require.NoError(t, err)

	output := buf.String()
	assert.Contains(t, output, "Arquivo de configuração criado")
}

