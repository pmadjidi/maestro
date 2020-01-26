package main

import (
	"context"
	"errors"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
 	. "maestro/api"
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

func AuthInterceptor(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {

	method := info.FullMethod
	if method != "/api.Register/Register" && method != "/api.Login/Authenticate" && method != "/api.Message/Msg" {

		md, ok := metadata.FromIncomingContext(ctx)
		if !ok {
			return nil, errors.New(Status_INVALID_TOKEN.String())
		}
		if !verifyToken(md["bearer-bin"]) {
			return nil, fmt.Errorf(Status_NOAUTH.String())
		}
	}

	m, err := handler(ctx, req)
	if err != nil {
		Info("RPC failed with error %v", err)
	}
	return m, err
}


//func(srv interface{}, ss ServerStream, info *StreamServerInfo, handler StreamHandler) error

/*
func AuthInterceptorStream(ctx context.Context, req interface{}, _ *grpc.UnaryServerInfo, handler grpc.StreamServerInterceptor) (interface{}, error) {

	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return nil, errors.New(api.Status_INVALID_TOKEN.String())
	}
	if !verifyToken(md["bearer-bin"]) {
		return nil, fmt.Errorf(api.Status_NOAUTH.String())
	}

	err := handler(ctx, req)
	if err != nil {
		Info("RPC failed with error %v", err)
	}
	return m, err
}
*/

/*

func (a *App) AutnInterceptorStream(srv interface{}, stream grpc.ServerStream, info *grpc.StreamServerInfo, handler grpc.StreamHandler) error {
	newCtx := newTagsForCtx(stream.Context()
	if o.requestFieldsFunc == nil {
		// Short-circuit, don't do the expensive bit of allocating a wrappedStream.
		wrappedStream := grpc_middleware.WrapServerStream(stream)
		wrappedStream.WrappedContext = newCtx
		return handler(srv, wrappedStream)
	}
	wrapped := &wrappedStream{stream, info, o, newCtx, true}
	err := handler(srv, wrapped)
	return err
}


 */