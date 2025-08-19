package logging

import (
	"fmt"
	"io"
	"log/slog"
	"os"
	"time"

	"github.com/domsnail/core-api-lib/cfg"
)

// SetDefaultLogger uses logging config to create default slog handler
// https://betterstack.com/community/guides/logging/logging-in-go/
func SetDefaultLogger(config cfg.LogsConfig) (err error) {
	var logger *slog.Logger
	var output io.Writer

	opts := &slog.HandlerOptions{
		AddSource: config.AddSource,
		Level:     slog.Level(config.Level),
		ReplaceAttr: func(groups []string, a slog.Attr) slog.Attr {
			if a.Key == slog.TimeKey {
				return slog.Attr{
					Key:   "timestamp",
					Value: slog.StringValue(a.Value.Time().Format(time.RFC3339)),
				}
			}

			if a.Key == slog.SourceKey {
				return slog.Attr{
					Key:   "source",
					Value: slog.StringValue(a.Value.Any().(*slog.Source).Function),
				}
			}

			return a
		},
	}

	if config.File == "" {
		output = os.Stdout
	} else {
		output, err = os.OpenFile(config.File, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
		if err != nil {
			return fmt.Errorf("failed to open log file: %w", err)
		}
	}

	switch config.Format {
	case "txt", "text":
		logger = slog.New(slog.NewTextHandler(output, opts))
	case "json":
		logger = slog.New(slog.NewJSONHandler(output, opts))
	default:
		return fmt.Errorf("unknown log format: " + config.Format)
	}

	slog.SetDefault(logger)
	return
}
