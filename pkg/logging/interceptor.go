package logging

import (
	"context"
	"log/slog"

	"github.com/grpc-ecosystem/go-grpc-middleware/v2/interceptors/logging"
	"google.golang.org/grpc/metadata"
)

var DefaultOptions = []logging.Option{
	logging.WithLogOnEvents(logging.StartCall, logging.FinishCall),
}

func Interceptor(l *slog.Logger) logging.Logger {
	return logging.LoggerFunc(func(ctx context.Context, lvl logging.Level, msg string, fields ...any) {
		md, ok := metadata.FromIncomingContext(ctx)
		if ok && md["user_id"] != nil {
			fields = append(fields, "user_id", md["user_id"][0])
		}

		if ok && md["request_uuid"] != nil {
			fields = append(fields, "request_uuid", md["request_uuid"][0])
		}

		l.Log(ctx, slog.Level(lvl), msg, fields...)
	})
}
