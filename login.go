package main

import (
	"context"
	"errors"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
	. "maestro/api"
)

type loginEnvelope struct {
	resp     chan *LoginResp
	token    chan *string
	username *string
	password *string
	device   *string
}

func newLoginEnvelope() *loginEnvelope {
	return &loginEnvelope{make(chan *LoginResp, 1), make(chan *string, 1), nil, nil, nil}
}

type loginService struct {
	name  string
	Q     chan *loginEnvelope
	cfg   *ServerConfig
	stats *metrics
}

func newLoginService(Q chan *loginEnvelope, cfg *ServerConfig) *loginService {
	return &loginService{"loginService", Q, cfg, newMetrics()}
}

func (l *loginService) Name() string {
	return l.name
}

func (l *loginService) Authenticate(ctx context.Context, req *LoginReq) (*LoginResp, error) {

	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return nil, errors.New("Missing metadata from incomming request")
	}

	/*
	if !verifyToken(md["authorization"]) {
		return nil, errors.New("Invalid token")
	}
	 */
	username := md.Get("username")
	password := md.Get("password")
	device := req.GetDevice()

	fmt.Printf("Got Auth request for %s,%s\n", username, password)

	if len(username) == 0 || len(password) == 0 {
		l.stats.invalidCalls += 1
		return &LoginResp{Status: Status_FAIL}, nil
	} else if len(password[0]) < l.cfg.MINIMUM_PASSWORD_LENGTH || len(password) > l.cfg.NAME_LENGTH_LIMIT {
		l.stats.invalidCalls += 1
		fmt.Printf("Inavild Password\n")
		return &LoginResp{Status: Status_FAIL}, nil
	}

	env := newLoginEnvelope()
	env.username = &username[0]
	env.password = &password[0]
	env.device = &device

	select {
	case l.Q <- env:
	case <-ctx.Done():
		err := ctx.Err()
		fmt.Printf("Authenticate, in error: %+v", err)
		l.stats.timeouts += 1
		return &LoginResp{Status: Status_TIMEOUT}, nil

	}

	select {
	case <-ctx.Done():
		err := ctx.Err()
		fmt.Printf("Authenticate, in error: %+v", err)
		l.stats.timeouts += 1
		return &LoginResp{Status: Status_TIMEOUT}, nil
	case res := <-env.resp:
		switch res.Status {
		case Status_ERROR:
			fmt.Printf("Password Error\n")
			l.stats.errors += 1
			return &LoginResp{Status: Status_ERROR}, errors.New(Status_name[int32(Status_ERROR)])
		case Status_NOTFOUND:
			fmt.Printf("Password Not found \n")
			l.stats.success += 1
			return res, nil
		case Status_SUCCESS:
			l.stats.success += 1
			token := *<-env.token
			fmt.Printf("token is set to: %s\n", token)
			grpc.SendHeader(ctx, metadata.New(map[string]string{"bearer-bin": token, "app": l.cfg.APP_NAME}))
			//ctx = metadata.(ctx, "app", l.cfg.APP_NAME, "bearer",env.token)
			return res, nil
		default:
			return res, nil
		}
	}
}
