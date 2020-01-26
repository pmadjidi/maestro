package main

import (
	"context"
	"errors"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
	. "maestro/api"
	"sync"
)

type registerEnvelope struct {
	req  chan *RegisterReq
	resp chan struct{}
	token  *string
	status Status
}


func newRegisterEnvelope() *registerEnvelope {
	return &registerEnvelope{make(chan *RegisterReq, 1), make(chan struct{}, 1),nil,Status_NEW}
}


type registerService struct {
	*sync.RWMutex
	name string
	stats *metrics
	system *Server
}

func (l *registerService) Name() string {
	return l.name
}

func newRegisterService(s *Server) *registerService {
	return &registerService{&sync.RWMutex{},"registerService",newMetrics(),s}
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
	appName := req.GetAppName()

	if len(Username) == 0 || len(Username) > r.system.cfg.NAME_LENGTH_LIMIT {
		return Status_INVALID_USERNAME
	} else if len(passWord) == 0 || len(passWord) > r.system.cfg.NAME_LENGTH_LIMIT {
		return Status_INVALID_PASSWORD
	} else if len(FirstName) == 0 || len(FirstName) > r.system.cfg.NAME_LENGTH_LIMIT {
		return Status_INVALID_FIRSTNAME
	}else if len(LastName) == 0 || len(LastName) > r.system.cfg.NAME_LENGTH_LIMIT {
		return Status_INVALID_LASTNAME
	}else if len(Email) == 0 || len(Email) > r.system.cfg.NAME_LENGTH_LIMIT {
		return Status_INVALID_EMAIL
	}else if len(Phone) == 0 || len(Phone) > r.system.cfg.NAME_LENGTH_LIMIT {
		return Status_INVALID_PHONE
	}else if len(Device) == 0 || len(Device) > r.system.cfg.NAME_LENGTH_LIMIT {
		return Status_INVALID_DEVICE
	}else if len(Adress.State) == 0 || len(Adress.State)  > r.system.cfg.NAME_LENGTH_LIMIT {
		return Status_INVALID_ADRESS
	}else if len(Adress.City) == 0 || len(Adress.City)  > r.system.cfg.NAME_LENGTH_LIMIT {
		return Status_INVALID_ADRESS
	}else if len(Adress.Street) == 0 || len(Adress.Street)  > r.system.cfg.NAME_LENGTH_LIMIT {
		return Status_INVALID_ADRESS
	}else if len(Adress.Zip) == 0 || len(Adress.Zip)  > r.system.cfg.NAME_LENGTH_LIMIT {
		return Status_INVALID_ADRESS

	} else if len(appName) == 0 || len(appName)  > r.system.cfg.NAME_LENGTH_LIMIT {
		return Status_INVALID_APPNAME
	} else {
		return Status_VALIDATED
	}
}



func (r *registerService) Register(ctx context.Context, req *RegisterReq) (*Empty,error) {



	validate := r.validateReq(req)

	if validate  != Status_VALIDATED {
		return &Empty{},fmt.Errorf(validate.String())
	}


	app,err :=  r.system.GetApp(req.AppName)

	if err != nil {
		app, err = r.system.NewApp(req.GetAppName())
		app.start()
		if err != nil {
			fmt.Printf("Error here...%s",err)
			return &Empty{}, err
		}
	}


	env := newRegisterEnvelope()
	env.req <- req

	select {
	case app.registerQ <- env:
	case <-ctx.Done():
		err := ctx.Err()
		fmt.Printf("Register, in error to kernel: %+v\n", err)
		r.stats.timeouts += 1
		return &Empty{}, fmt.Errorf(Status_TIMEOUT.String())

	}

	select {
	case  <-env.resp:
		switch env.status {
		case   Status_ERROR:
			r.Lock()
			r.stats.errors += 1
			r.Unlock()
			return  &Empty{}, errors.New(Status_ERROR.String())
		case  Status_EXITSTS:
			r.Lock()
			r.stats.success +=1
			r.Unlock()
			return &Empty{},errors.New(Status_EXITSTS.String())
		case Status_SUCCESS:
			r.Lock()
			r.stats.success += 1
			r.Unlock()
			header := metadata.Pairs("bearer-bin", *env.token,"app",app.cfg.APP_NAME,)
			grpc.SendHeader(ctx, header)
			return &Empty{},nil
		default:
			return &Empty{},nil
		}
	}
}



