package middleware

import (
	"books/internal/pkg/request"
	"context"
	"time"

	grpc_zap "github.com/grpc-ecosystem/go-grpc-middleware/logging/zap"
	"github.com/grpc-ecosystem/go-grpc-middleware/logging/zap/ctxzap"
	grpc_ctxtags "github.com/grpc-ecosystem/go-grpc-middleware/tags"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
)

func codeToLevel(code codes.Code) zapcore.Level {
	if code == codes.OK {
		// its a debug
		return zap.DebugLevel
	}
	return grpc_zap.DefaultCodeToLevel(code)
}

func extractFields(fullMethod string, req interface{}) map[string]interface{} {
	return make(map[string]interface{})
}

func messageProducer(ctx context.Context, msg string, level zapcore.Level, code codes.Code, err error, duration zapcore.Field) {
	ctxzap.AddFields(ctx, zap.String(request.RequestIDKey, request.GetContextRequestID(ctx)))
	ctxzap.Extract(ctx).Check(level, msg).Write(
		zap.Error(err),
		zap.String("grpc.code", code.String()),
		duration,
	)
}

func AddLogging(logger *zap.Logger, uInterceptors *[]grpc.UnaryServerInterceptor, sInterceptors *[]grpc.StreamServerInterceptor) {
	// shared options for the logger, with a custom gRPC code to log level function
	o := []grpc_zap.Option{
		grpc_zap.WithLevels(codeToLevel),
		grpc_zap.WithMessageProducer(messageProducer),
		grpc_zap.WithDurationField(func(duration time.Duration) zapcore.Field {
			return zap.Int64("grpc.time_ns", duration.Nanoseconds())
		}),
	}

	// make sure that log statements internal to gRPC library are logge dusing the zaplogger as well
	grpc_zap.ReplaceGrpcLogger(logger)

	*uInterceptors = append(*uInterceptors, grpc_ctxtags.UnaryServerInterceptor(grpc_ctxtags.WithFieldExtractor(extractFields)))
	*uInterceptors = append(*uInterceptors, grpc_zap.UnaryServerInterceptor(logger, o...))

	*sInterceptors = append(*sInterceptors, grpc_ctxtags.StreamServerInterceptor(grpc_ctxtags.WithFieldExtractor(grpc_ctxtags.CodeGenRequestFieldExtractor)))
	*sInterceptors = append(*sInterceptors, grpc_zap.StreamServerInterceptor(logger, o...))
}
