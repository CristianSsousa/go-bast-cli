package logger

import (
	"os"

	"github.com/sirupsen/logrus"
)

var (
	// Log é a instância global do logger estruturado
	Log *logrus.Logger
)

// Init inicializa o logger com configurações padrão
func Init(level string, format string) {
	Log = logrus.New()

	// Configurar nível de log
	logLevel, err := logrus.ParseLevel(level)
	if err != nil {
		logLevel = logrus.InfoLevel
	}
	Log.SetLevel(logLevel)

	// Configurar formato
	if format == "json" {
		Log.SetFormatter(&logrus.JSONFormatter{
			TimestampFormat: "2006-01-02T15:04:05Z07:00",
		})
	} else {
		Log.SetFormatter(&logrus.TextFormatter{
			FullTimestamp:   true,
			TimestampFormat: "2006-01-02 15:04:05",
			ForceColors:     true,
		})
	}

	// Output padrão
	Log.SetOutput(os.Stdout)
}

// GetLogger retorna a instância do logger
func GetLogger() *logrus.Logger {
	if Log == nil {
		Init("info", "text")
	}
	return Log
}

