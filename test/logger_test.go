package test

import (
	"log/slog"
	"testing"

	"github.com/domsnail/core-api-lib/cfg"
	"github.com/domsnail/core-api-lib/pkg/logging"
	"github.com/stretchr/testify/require"
)

func TestLogging(t *testing.T) {
	t.Run("test source logging", func(t *testing.T) {
		err := logging.SetDefaultLogger(cfg.LogsConfig{
			AddSource: true,
			Level:     0,
			Format:    "json",
		})

		require.NoError(t, err)

		slog.Info("test source")
	})
}
