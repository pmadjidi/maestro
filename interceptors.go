package main

/*
import (
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
	"context"
	. "maestro/api"
)

func (a *App) authInterceptorUnary(ctx context.Context, req interface{}, _ *grpc.UnaryServerInfo, handler grpc.UnaryServerInterceptor) (interface{}, error) {
	meta, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return nil, status.Error(codes.Unauthenticated, Status_NOAUTH.String())
	}
	if len(meta["jwt"]) != 1 {
		return nil, status.Error(codes.Unauthenticated, Status_ERROR.String())
	}

	// if code here to verify jwt is correct. if not return nil and error by accessing meta["jwt"][0]

	return handler(ctx, req,nil),nil // go to function.
}


func (a *App) authInterceptorStream(ctx context.Context, req interface{}, _ *grpc.UnaryServerInfo, handler grpc.StreamServerInterceptor) (interface{}, error) {
	meta, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return nil, status.Error(codes.Unauthenticated, Status_NOAUTH.String())
	}
	if len(meta["jwt"]) != 1 {
		return nil, status.Error(codes.Unauthenticated, Status_ERROR.String())
	}

	// if code here to verify jwt is correct. if not return nil and error by accessing meta["jwt"][0]

	return handler(ctx, req),nil // go to function.
}


 */