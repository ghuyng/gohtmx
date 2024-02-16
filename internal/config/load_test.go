package config

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestLoad(t *testing.T) {
	t.Setenv("SERVER_ADDR", ":3000")
	fileName := "../../env/test.env"
	expectedCfg := &Config{
		Server: ServerConfig{
			Addr:                       ":3000",
			ReadTimeoutInSeconds:       20,
			ReadHeaderTimeoutInSeconds: 10,
			WriteTimeoutInSeconds:      30,
			IdleTimeoutInSeconds:       0,
			AllowOrigins:               []string{"*"},
			AllowMethods:               []string{"GET", "POST", "PUT", "PATCH", "DELETE", "OPTIONS", "HEAD"},
		},
	}

	// Call the Load function
	cfg, err := Load(fileName)
	require.NoError(t, err)

	// Compare the loaded config with the expected config
	require.Equal(t, expectedCfg, cfg)
}
