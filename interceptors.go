package main

import (
	"context"
	"errors"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
//	"github.com/grpc-ecosystem/go-grpc-middleware"
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
		Info("RPC failed with error %v\n", err)
	}
	return m, err
}


/*

func StreamServerInterceptor(authFunc AuthFunc) grpc.StreamServerInterceptor {
	return func(srv interface{}, stream grpc.ServerStream, info *grpc.StreamServerInfo, handler grpc.StreamHandler) error {
		var newCtx context.Context
		var err error
		if overrideSrv, ok := srv.(ServiceAuthFuncOverride); ok {
			newCtx, err = overrideSrv.AuthFuncOverride(stream.Context(), info.FullMethod)
		} else {
			newCtx, err = authFunc(stream.Context())
		}
		if err != nil {
			return err
		}
		wrapped := grpc_middleware.WrapServerStream(stream)
		wrapped.WrappedContext = newCtx
		return handler(srv, wrapped)
	}
}

 */