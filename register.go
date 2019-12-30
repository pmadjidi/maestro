package main

import (
	"context"
	"errors"
	"fmt"
	. "maestro/api"
)

type registerEnvelope struct {
	req  chan *RegisterReq
	resp chan *RegisterResp
}


func newRegisterEnvelope() *registerEnvelope {
	return &registerEnvelope{make(chan *RegisterReq, 1), make(chan *RegisterResp, 1)}
}



type registerService struct {
	name string
	Q chan *registerEnvelope
	cfg *ServerConfig
	timeout int64
	error int64
	collision int64
	success int64
}

func (l *registerService) Name() string {
	return l.name
}

func newRegisterService(Q chan *registerEnvelope,cfg *ServerConfig) *registerService {
	return &registerService{"registerService",Q,cfg,0,0,0,0}
}


func (r *registerService) validateReq(req *RegisterReq) Status {
	Username := req.GetUserName()
	passWord := req.GetPassWord()
	FirstName := req.GetFirstName()
	LastName := req.GetLastName()
	Email	:= req.GetEmail()
	Phone := req.GetPhone()
	Adress := req.GetAddress()
	Device := req.GetDevice()

	if len(Username) == 0 || len(Username) > r.cfg.NAME_LENGTH_LIMIT {
		return Status_INVALID_USERNAME
	} else if len(passWord) == 0 || len(passWord) > r.cfg.NAME_LENGTH_LIMIT {
		return Status_INVALID_PASSWORD
	} else if len(FirstName) == 0 || len(FirstName) > r.cfg.NAME_LENGTH_LIMIT {
		return Status_INVALID_FIRSTNAME
	}else if len(LastName) == 0 || len(LastName) > r.cfg.NAME_LENGTH_LIMIT {
		return Status_INVALID_LASTNAME
	}else if len(Email) == 0 || len(Email) > r.cfg.NAME_LENGTH_LIMIT {
		return Status_INVALID_EMAIL
	}else if len(Phone) == 0 || len(Phone) > r.cfg.NAME_LENGTH_LIMIT {
		return Status_INVALID_PHONE
	}else if len(Device) == 0 || len(Device) > r.cfg.NAME_LENGTH_LIMIT {
		return Status_INVALID_DEVICE
	}else if len(Adress.State) == 0 || len(Adress.State)  > r.cfg.NAME_LENGTH_LIMIT {
		return Status_INVALID_ADRESS
	}else if len(Adress.City) == 0 || len(Adress.City)  > r.cfg.NAME_LENGTH_LIMIT {
		return Status_INVALID_ADRESS
	}else if len(Adress.Street) == 0 || len(Adress.Street)  > r.cfg.NAME_LENGTH_LIMIT {
		return Status_INVALID_ADRESS
	}else if len(Adress.Zip) == 0 || len(Adress.Zip)  > r.cfg.NAME_LENGTH_LIMIT {
		return Status_INVALID_ADRESS
	} else {
		return Status_VALIDATED
	}
}



func (r *registerService) Register(ctx context.Context, req *RegisterReq) (*RegisterResp, error) {


	validate := r.validateReq(req)

	if validate  != Status_VALIDATED {
		return &RegisterResp{Status: validate}, nil
	}


	env := newRegisterEnvelope()
	env.req <- req

	select {
	case r.Q <- env:
	case <-ctx.Done():
		err := ctx.Err()
		fmt.Printf("Register, in error to kernel: %+v\n", err)
		r.timeout += 1
		return &RegisterResp{Status: Status_TIMEOUT}, nil

	}

	select {
	case <-ctx.Done():
		err := ctx.Err()
		fmt.Printf("Register, in error from kernel: %+v\n", err)
		r.timeout += 1
		return &RegisterResp{Status: Status_TIMEOUT}, nil
	case res := <-env.resp:
		switch res.Status {
		case   Status_ERROR:
			r.error += 1
			return  nil, errors.New(Status_name[int32(Status_ERROR)])
		case    Status_EXITSTS:
			r.collision += 1
			return res,nil
		default:
			r.success += 1
			return res,nil
		}
	}
}



