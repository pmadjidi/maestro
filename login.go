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
	req  chan *LoginReq
	resp chan *LoginResp
	token chan *string
}


func newLoginEnvelope() *loginEnvelope {
	return &loginEnvelope{make(chan *LoginReq, 1), make(chan *LoginResp, 1),make(chan *string,1)}
}


type loginService struct {
	name string
	error int64
	success int64
	notfound int64
	timeout int64
	invalidPassword int64
	invalidUsername int64
	Q chan *loginEnvelope
	cfg *ServerConfig
}

func newLoginService(Q chan *loginEnvelope,cfg *ServerConfig) *loginService {
	 return &loginService{"loginService",0,0,0,0,0,
	 	0,Q,cfg,
	 }
}


func (l *loginService) Name() string {
	return l.name
}

func (l *loginService) Authenticate(ctx context.Context, req *LoginReq) (*LoginResp, error) {

	username := req.GetUserName()
	passWord := req.GetPassWord()

	fmt.Printf("Got Auth request for %s,%s\n",username,passWord)

	if username == "" {
		l.invalidUsername += 1
		return &LoginResp{Status: Status_FAIL}, nil
	}    else if len(passWord) < l.cfg.MINIMUM_PASSWORD_LENGTH || len(passWord) > l.cfg.NAME_LENGTH_LIMIT {
		   l.invalidPassword += 1
		   fmt.Printf("Inavild Password\n")
		return &LoginResp{Status: Status_FAIL}, nil
	}

	env := newLoginEnvelope()
	env.req <- req

	select {
	case l.Q <- env:
	case <-ctx.Done():
		err := ctx.Err()
		fmt.Printf("Authenticate, in error: %+v", err)
		l.timeout += 1
		return &LoginResp{Status: Status_TIMEOUT}, nil

	}

	select {
	case <-ctx.Done():
		err := ctx.Err()
		fmt.Printf("Authenticate, in error: %+v", err)
		l.timeout += 1
		return &LoginResp{Status: Status_TIMEOUT}, nil
	case res := <-env.resp:
		switch res.Status {
		case   Status_ERROR:
			fmt.Printf("Password Error\n")
			l.error += 1
			return &LoginResp{Status: Status_ERROR}, errors.New(Status_name[int32(Status_ERROR)])
		case    Status_NOTFOUND:
			fmt.Printf("Password Not found \n")
			l.notfound += 1
			return res,nil
		case Status_SUCCESS:
			l.success += 1
			token := *<- env.token
			fmt.Printf("token is set to: %s\n",token)
			grpc.SendHeader(ctx, metadata.New(map[string]string{"bearer-bin": token ,"app": l.cfg.APP_NAME}))
			//ctx = metadata.(ctx, "app", l.cfg.APP_NAME, "bearer",env.token)
			return res,nil
		default:
			return res,nil
		}


	}
}

