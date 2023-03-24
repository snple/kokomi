package metadata

import (
	"context"
	"net"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/peer"
	"google.golang.org/grpc/status"
)

func SetMetadata(ctx context.Context, kv ...string) context.Context {
	return metadata.AppendToOutgoingContext(ctx, kv...)
}

func GetMetadata(ctx context.Context, key string) (string, error) {
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return "", status.Errorf(codes.InvalidArgument, "metadata is not provided")
	}

	value := md[key]
	if len(value) == 0 {
		return "", status.Errorf(codes.InvalidArgument, "metadata key is not provided")
	}

	return value[0], nil
}

func HasMetadata(ctx context.Context, key string) bool {
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return false
	}

	value := md[key]
	if len(value) == 0 {
		return false
	}

	return true
}

func SetToken(ctx context.Context, token string) context.Context {
	return metadata.AppendToOutgoingContext(ctx, "token", token)
}

func GetToken(ctx context.Context) (string, error) {
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return "", status.Errorf(codes.Unauthenticated, "metadata is not provided")
	}

	token := md["token"]
	if len(token) == 0 {
		return "", status.Errorf(codes.Unauthenticated, "authorization token is not provided")
	}

	return token[0], nil
}

func SetSync(ctx context.Context) context.Context {
	return metadata.AppendToOutgoingContext(ctx, "sync", "")
}

func IsSync(ctx context.Context) bool {
	if md, ok := metadata.FromIncomingContext(ctx); ok {
		value := md["sync"]
		if len(value) > 0 {
			return true
		}
	}

	if md, ok := metadata.FromOutgoingContext(ctx); ok {
		value := md["sync"]
		if len(value) > 0 {
			return true
		}
	}

	return false
}

// GetPeerAddr get peer addr
func GetPeerAddr(ctx context.Context) string {
	var addr string
	if pr, ok := peer.FromContext(ctx); ok {
		if tcpAddr, ok := pr.Addr.(*net.TCPAddr); ok {
			addr = tcpAddr.IP.String()
		} else {
			addr = pr.Addr.String()
		}
	}
	return addr
}

func ContextError(ctx context.Context) error {
	switch ctx.Err() {
	case context.Canceled:
		return status.Error(codes.Canceled, "request is canceled")
	case context.DeadlineExceeded:
		return status.Error(codes.DeadlineExceeded, "deadline is exceeded")
	default:
		return nil
	}
}
