package restful

import (
	"fmt"
	"net/http"
	"time"

	"github.com/go-chi/chi/v5/middleware"
	"go.uber.org/zap"
)

func ZapToRequestLoggerAdaptor(logger *zap.Logger) middleware.LogFormatter {
	return &zapLogFormatter{logger}
}

type zapLogFormatter struct {
	logger *zap.Logger
}

func (z *zapLogFormatter) NewLogEntry(r *http.Request) middleware.LogEntry {
	logger := z.logger.With(zap.String("req_id", middleware.GetReqID(r.Context())))

	return &zapEntry{logger: logger, request: r}
}

type zapEntry struct {
	logger  *zap.Logger
	request *http.Request
}

func (e *zapEntry) Write(status, bytes int, _ http.Header, elapsed time.Duration, _ interface{}) {
	e.logger.With(
		zap.String("http_proto", e.request.Proto),
		zap.String("http_method", e.request.Method),
		zap.String("request_uri", e.request.RequestURI),
		zap.String("remote_addr", e.request.RemoteAddr),
		zap.String("user_agent", e.request.UserAgent()),

		zap.Int("status", status),
		zap.Int("bytes", bytes),
		zap.Duration("elapsed", elapsed.Round(time.Microsecond)),
	).Info("requested")
}

func (e *zapEntry) Panic(v interface{}, stack []byte) {
	e.logger = e.logger.With(
		zap.String("stack", string(stack)),
		zap.String("panic", fmt.Sprintf("%+v", v)),
	)
}

func getLogger(r *http.Request) *zap.Logger {
	return middleware.GetLogEntry(r).(*zapEntry).logger
}
