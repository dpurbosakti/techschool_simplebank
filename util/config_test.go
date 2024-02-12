package util

import (
	"testing"

	"github.com/stretchr/testify/require"
)

const (
	dbDriver          = "postgres"
	dbSource          = "postgresql://root:mokopass@localhost:5432/simple_bank?sslmode=disable"
	httpServerAddress = "0.0.0.0:8080"
)

func TestLoadConfigFail(t *testing.T) {
	config, err := LoadConfig("/error")
	require.Error(t, err)
	require.Empty(t, config)
}

func TestLoadConfigSuccess(t *testing.T) {
	config, err := LoadConfig("..")
	require.NoError(t, err)
	require.NotEmpty(t, config)
	require.Equal(t, config.DBSource, dbSource)
	require.Equal(t, config.HTTPServerAddress, httpServerAddress)
}
