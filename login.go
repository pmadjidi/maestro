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
	stats *metrics
	system *Server
}

func newLoginService(s *Server) *loginService {
	return &loginService{"loginService", newMetrics(),s}
}

func (l *loginService) Name() string {
	return l.name
}

func (l *loginService) Authenticate(ctx context.Context, req *LoginReq) (*LoginResp, error) {

	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return nil, fmt.Errorf(Status_NOAUTH.String())
	}

	/*
	if !verifyToken(md["authorization"]) {
		return nil, errors.New("Invalid token")
	}
	 */



	username := md.Get("username")
	password := md.Get("password")
	appName := md.Get("app")
	device := req.GetDevice()

	if len(appName) == 0 || appName[0] == "" {
		return nil, fmt.Errorf(Status_INVALID_APPNAME.String())
	}

	app,err  := l.system.GetApp(appName[0])
	if err != nil {
		return nil,err
	}



	fmt.Printf("Got Auth request for %s,%s\n", username, password)

	if len(username) == 0 || len(password) == 0 {
		l.stats.invalidCalls += 1
		return &LoginResp{Status: Status_FAIL}, nil
	} else if len(password[0]) < l.system.cfg.MINIMUM_PASSWORD_LENGTH || len(password) > l.system.cfg.NAME_LENGTH_LIMIT {
		l.stats.invalidCalls += 1
		fmt.Printf("Inavild Password\n")
		return &LoginResp{Status: Status_FAIL}, nil
	}

	env := newLoginEnvelope()
	env.username = &username[0]
	env.password = &password[0]
	env.device = &device

	select {
	case app.loginQ <- env:
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
			grpc.SendHeader(ctx, metadata.New(map[string]string{"bearer-bin": token, "app": appName[0]}))
			//ctx = metadata.(ctx, "app", l.cfg.APP_NAME, "bearer",env.token)
			return res, nil
		default:
			return res, nil
		}
	}
}
