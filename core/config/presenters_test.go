package config

import (
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/GoPlugin/pluginV2/core/logger"
)

func TestNewConfigPrinter(t *testing.T) {
	cfg := NewGeneralConfig(logger.TestLogger(t))
	printer := NewConfigPrinter(cfg)
	require.NotNil(t, printer)
}
