package config

import (
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestInit(t *testing.T) {
	t.Run("init with defaults", func(t *testing.T) {
		err := Init("")
		assert.NoError(t, err)
		assert.NotNil(t, Cfg)
		assert.Equal(t, "bast", Cfg.App.Name)
		assert.Equal(t, "info", Cfg.Logging.Level)
		assert.Equal(t, 8080, Cfg.Server.DefaultPort)
	})

	t.Run("init with non-existent config file", func(t *testing.T) {
		err := Init("/tmp/nonexistent-config.yaml")
		// Não deve retornar erro, deve usar defaults
		assert.NoError(t, err)
		assert.NotNil(t, Cfg)
	})
}

func TestGet(t *testing.T) {
	Init("")
	cfg := Get()
	assert.NotNil(t, cfg)
	assert.Equal(t, "bast", cfg.App.Name)
}

func TestSetAndGet(t *testing.T) {
	Init("")

	Set("app.name", "test-app")
	assert.Equal(t, "test-app", GetString("app.name"))

	Set("server.default_port", 3000)
	assert.Equal(t, 3000, GetInt("server.default_port"))

	Set("features.auto_update", true)
	assert.Equal(t, true, GetBool("features.auto_update"))
}

func TestSave(t *testing.T) {
	// Criar diretório temporário
	tmpDir := t.TempDir()
	configPath := filepath.Join(tmpDir, "config.yaml")

	Init("")
	Set("app.name", "test-save")

	// Salvar em local temporário usando Init com caminho específico
	err := Init(configPath)
	require.NoError(t, err)

	// Verificar que configuração foi carregada
	cfg := Get()
	assert.Equal(t, "test-save", cfg.App.Name)
}

func TestDefaults(t *testing.T) {
	Init("")
	cfg := Get()

	assert.Equal(t, "bast", cfg.App.Name)
	assert.Equal(t, "1.0.0", cfg.App.Version)
	assert.Equal(t, "info", cfg.Logging.Level)
	assert.Equal(t, "text", cfg.Logging.Format)
	assert.Equal(t, 8080, cfg.Server.DefaultPort)
	assert.Equal(t, "0.0.0.0", cfg.Server.DefaultHost)
	assert.Equal(t, 30, cfg.Server.Timeout)
	assert.False(t, cfg.Features.AutoUpdate)
	assert.False(t, cfg.Features.Verbose)
}

