package utils

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestGetConfigDir(t *testing.T) {
	configDir, err := GetConfigDir()
	assert.NoError(t, err)
	assert.NotEmpty(t, configDir)
	assert.Contains(t, configDir, ".bast")
}

func TestEnsureConfigDir(t *testing.T) {
	tmpDir := t.TempDir()
	originalHome := os.Getenv("HOME")

	// Mock HOME para teste
	os.Setenv("HOME", tmpDir)
	defer os.Setenv("HOME", originalHome)

	err := EnsureConfigDir()
	assert.NoError(t, err)

	// Verificar que diretório foi criado
	configDir, _ := GetConfigDir()
	_, err = os.Stat(configDir)
	assert.NoError(t, err)
}

func TestGetConfigPath(t *testing.T) {
	configPath, err := GetConfigPath()
	assert.NoError(t, err)
	assert.NotEmpty(t, configPath)
	assert.Contains(t, configPath, ".bast")
	assert.Contains(t, configPath, "config.yaml")
}

func TestFileExists(t *testing.T) {
	t.Run("file exists", func(t *testing.T) {
		tmpFile := filepath.Join(t.TempDir(), "test.txt")
		err := os.WriteFile(tmpFile, []byte("test"), 0644)
		require.NoError(t, err)

		assert.True(t, FileExists(tmpFile))
	})

	t.Run("file does not exist", func(t *testing.T) {
		assert.False(t, FileExists("/tmp/nonexistent-file-12345.txt"))
	})
}

func TestIsDir(t *testing.T) {
	t.Run("is directory", func(t *testing.T) {
		tmpDir := t.TempDir()
		assert.True(t, IsDir(tmpDir))
	})

	t.Run("is not directory", func(t *testing.T) {
		tmpFile := filepath.Join(t.TempDir(), "test.txt")
		err := os.WriteFile(tmpFile, []byte("test"), 0644)
		require.NoError(t, err)

		assert.False(t, IsDir(tmpFile))
	})

	t.Run("path does not exist", func(t *testing.T) {
		assert.False(t, IsDir("/tmp/nonexistent-dir-12345"))
	})
}
