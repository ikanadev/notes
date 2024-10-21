package middleware

import (
	"books/internal/pkg/logger"
	"context"
	"runtime/debug"

	grpc_recovery "github.com/grpc-ecosystem/go-grpc-middleware/recovery"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// to handle panic and log stack trace
// also returns error from panic to send error-messag to gateway
func PanicTracer(ctx context.Context, p interface{}) error {
	logger.WithContext(ctx).Errorf("StackTrace: %s", string(debug.Stack()))
	return status.Errorf(codes.Unknown, "Panic triggered: %v", p)
}

func AddRecovery(uInterceptors *[]grpc.UnaryServerInterceptor, sInterceptors *[]grpc.StreamServerInterceptor) {
	opts := []grpc_recovery.Option{
		grpc_recovery.WithRecoveryHandlerContext(PanicTracer),
	}

	*uInterceptors = append(*uInterceptors, grpc_recovery.UnaryServerInterceptor(opts...))
	*sInterceptors = append(*sInterceptors, grpc_recovery.StreamServerInterceptor(opts...))
}
