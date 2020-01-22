package main

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
	. "maestro/api"
	"sync"
)

type loginEnvelope struct {
	resp     chan notify
	token    *string
	username *string
	password *string
	device   *string
	Status
}

func newLoginEnvelope() *loginEnvelope {
	return &loginEnvelope{make(chan notify, 1), nil, nil, nil, nil, Status_NEW}
}

type loginService struct {
	*sync.RWMutex
	name   string
	stats  *metrics
	system *Server
}

func newLoginService(s *Server) *loginService {
	return &loginService{&sync.RWMutex{}, "loginService", newMetrics(), s}
}

func (l *loginService) Name() string {
	return l.name
}

func (l *loginService) Authenticate(ctx context.Context, req *LoginReq) (*Empty, error) {

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

	app, err := l.system.GetApp(appName[0])
	if err != nil {
		return nil, err
	}

	fmt.Printf("Got Auth request for %s\n", username)

	if len(username) == 0 || len(password) == 0 {
		l.stats.invalidCalls += 1
		return &Empty{}, fmt.Errorf(Status_FAIL.String())
	} else if len(password[0]) < l.system.cfg.MINIMUM_PASSWORD_LENGTH || len(password) > l.system.cfg.NAME_LENGTH_LIMIT {
		l.stats.invalidCalls += 1
		fmt.Printf("Inavild Password\n")
		return &Empty{}, fmt.Errorf(Status_FAIL.String())
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
		l.Lock()
		l.stats.timeouts += 1
		l.Unlock()
		return &Empty{}, fmt.Errorf(Status_TIMEOUT.String())
	}

	select {
	case <-ctx.Done():
		err := ctx.Err()
		fmt.Printf("Authenticate, in error: %+v", err)
		l.Lock()
		l.stats.timeouts += 1
		l.Unlock()
	case <-env.resp:
		switch env.Status {
		case Status_ERROR:
			fmt.Printf("Password Error\n")
			l.Lock()
			l.stats.errors += 1
			l.Unlock()
		case Status_NOTFOUND:
			fmt.Printf("Password Not found \n")
			l.Lock()
			l.stats.success += 1
			l.Unlock()
		case Status_SUCCESS:
			l.Lock()
			l.stats.success += 1
			l.Unlock()
			//token := env.token
			//fmt.Printf("\n\n## token is set to: %s\n", *token)
			header := metadata.Pairs("bearer-bin", *env.token)
			grpc.SendHeader(ctx, header)
			//ctx = metadata.AppendToOutgoingContext(ctx, "bearer-bin", *env.token)
			//err = grpc.SetHeader(ctx, metadata.New(map[string]string{"bearer-bin": *env.token, "app": appName[0]}))
			//grpc.S
			//ctx = metadata.NewIncomingContext(ctx,"bearer-bin": token, "app": appName[0])
			//ctx = metadata.(ctx, "app", l.cfg.APP_NAME, "bearer",env.token)
			//ctx = metadata.AppendToOutgoingContext(ctx, "bearer-bin", *token)
			return &Empty{},nil
		}
	}
		return &Empty{}, fmt.Errorf(env.Status.String())

}
