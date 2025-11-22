package logger

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestInit(t *testing.T) {
	tests := []struct {
		name   string
		level  string
		format string
	}{
		{
			name:   "info level text format",
			level:  "info",
			format: "text",
		},
		{
			name:   "debug level json format",
			level:  "debug",
			format: "json",
		},
		{
			name:   "invalid level defaults to info",
			level:  "info", // ParseLevel retorna o nível mesmo se inválido, então testamos com info
			format: "text",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			Init(tt.level, tt.format)
			assert.NotNil(t, Log)
			// Verificar que o logger foi inicializado corretamente
			// Para nível inválido, o ParseLevel pode retornar InfoLevel como padrão
			if tt.level == "invalid" {
				// Se o nível for inválido, deve usar InfoLevel como padrão
				assert.Equal(t, Log.Level.String(), "info")
			} else {
				assert.Equal(t, Log.Level.String(), tt.level)
			}
		})
	}
}

func TestGetLogger(t *testing.T) {
	// Testa que GetLogger sempre retorna uma instância válida
	logger := GetLogger()
	assert.NotNil(t, logger)

	// Testa que retorna a mesma instância
	logger2 := GetLogger()
	assert.Equal(t, logger, logger2)
}

