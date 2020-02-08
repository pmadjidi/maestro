package main

import (
	"context"
	"errors"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
	"github.com/dgrijalva/jwt-go"
	//	"github.com/grpc-ecosystem/go-grpc-middleware"
	. "maestro/api"
	"strings"
)

func getJwtToken(authorization []string,secret string) (jwt.MapClaims,error) {
	if len(authorization) < 1 {
		return nil,fmt.Errorf(Status_NOAUTH.String())
	}
	token := strings.TrimPrefix(authorization[0], "Bearer ")
	// Perform the token validation here. For the sake of this example, the code
	// here forgoes any of the usual OAuth2 token validation and instead checks
	// for a token matching an arbitrary string.
	tokenMap,err := decodeToken(token,secret)
	if err != nil {
		return nil,err
	}


	return tokenMap,nil
}


//type UnaryClientInterceptor func(ctx context.Context, method string, req, reply interface{}, cc *ClientConn, invoker UnaryInvoker, opts ...CallOption) error


func createUniaryInterCeptor(cfg *ServerConfig) grpc.UnaryServerInterceptor {
	return func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
		method := info.FullMethod
		if method != "/api.Register/Register" && method != "/api.Login/Authenticate" && method != "/api.Message/Msg" {

			md, ok := metadata.FromIncomingContext(ctx)
			if !ok {
				return nil, errors.New(Status_INVALID_TOKEN.String())
			}
			token, err := getJwtToken(md["bearer-bin"],cfg.SYSTEM_SECRET)
			if err != nil {
				return nil,err
			}
			ctx = context.WithValue(ctx, "appName", token["appname"])
			ctx = context.WithValue(ctx, "userName", token["username"])
			}

		m, err := handler(ctx, req)
		if err != nil {
			Info("RPC failed with error %v\n", err)
		}
		return m, err
	}
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