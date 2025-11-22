package constants

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAppConstants(t *testing.T) {
	assert.NotEmpty(t, AppName)
	assert.NotEmpty(t, AppVersion)
	assert.NotEmpty(t, AppDescription)
	assert.NotEmpty(t, AppAuthor)
}

func TestConfigConstants(t *testing.T) {
	assert.NotEmpty(t, ConfigDirName)
	assert.NotEmpty(t, ConfigFileName)
	assert.NotEmpty(t, ConfigFileExample)
}

func TestLoggingConstants(t *testing.T) {
	assert.NotEmpty(t, LogLevelDebug)
	assert.NotEmpty(t, LogLevelInfo)
	assert.NotEmpty(t, LogLevelWarn)
	assert.NotEmpty(t, LogLevelError)
	assert.NotEmpty(t, LogFormatText)
	assert.NotEmpty(t, LogFormatJSON)
}

func TestServerConstants(t *testing.T) {
	assert.Greater(t, DefaultPort, 0)
	assert.LessOrEqual(t, DefaultPort, MaxPort)
	assert.NotEmpty(t, DefaultHost)
	assert.Greater(t, DefaultTimeout, 0)
	assert.Equal(t, MinPort, 1)
	assert.Equal(t, MaxPort, 65535)
}

func TestNetworkConstants(t *testing.T) {
	assert.Greater(t, DefaultNetworkTimeout, 0)
	assert.NotEmpty(t, TCPProtocol)
}

func TestEnvPrefix(t *testing.T) {
	assert.NotEmpty(t, EnvPrefix)
}

func TestFilePermissions(t *testing.T) {
	assert.NotZero(t, ConfigDirPerm)
	assert.NotZero(t, ConfigFilePerm)
}
