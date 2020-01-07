package main

import (
	"errors"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
	"context"
	"strings"
)


func verifyToken(authorization []string) bool {
	if len(authorization) < 1 {
		return false
	}
	token := strings.TrimPrefix(authorization[0], "Bearer ")
	// Perform the token validation here. For the sake of this example, the code
	// here forgoes any of the usual OAuth2 token validation and instead checks
	// for a token matching an arbitrary string.
	return token == "some-secret-token"
}

func unaryInterceptor(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return nil,  errors.New("Missing metadata from incomming request")
	}
	if !verifyToken(md["authorization"]) {
		return nil, errors.New("Invalid token")
	}

	m, err := handler(ctx, req)
	if err != nil {
		Info("RPC failed with error %v", err)
	}
	return m, err
}


/*
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