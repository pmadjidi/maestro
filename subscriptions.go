package main

import (
	"context"
	"fmt"
	. "maestro/api"
	"sync"
)



/*
type SubscriptionsServer interface {
	Sub(context.Context, *TopicReq) (*TopicResp, error)
	Unsub(context.Context, *TopicReq) (*TopicResp, error)
	List(context.Context, *Empty) (*TopicResp, error)
}
 */



type topicService struct {
	*sync.RWMutex
	name string
	stats *metrics
	system *Server
}

func (t *topicService) getname() string {
	return t.name
}

func newTopicService(s *Server) *topicService {
	return &topicService{&sync.RWMutex{},"subscriptionService",newMetrics(),s}
}


func (t *topicService) Sub(ctx context.Context, req *TopicReq) (*TopicResp, error) {

	appName := ctx.Value("appName").(string)
	userName := ctx.Value("userName").(string)

	app, err := t.system.GetOrCreateApp(appName, false)

	if err != nil {
		return nil, err
	}

	env := newTopicEnvelope()
	env.TopicReq = req
	env.Username = userName

	select {
	case app.topicSubQ <- env:
	case <-ctx.Done():
		err := ctx.Err()
		fmt.Printf(": %+v\n", err)
		return nil, fmt.Errorf(Status_TIMEOUT.String())

	}

	select {
	case <-env.resp:
	case <-ctx.Done():
		err := ctx.Err()
		fmt.Printf("Register, in error to kernel: %+v\n", err)
		t.stats.timeouts += 1
		return nil, fmt.Errorf(Status_TIMEOUT.String())
	}

	if env.Status != Status_SUCCESS {
		return nil,fmt.Errorf(env.Status.String())
	}

	return &TopicResp{List: env.List},nil
}

func (t *topicService) Unsub(ctx context.Context, req *TopicReq) (*TopicResp, error) {

	appName := ctx.Value("appName").(string)
	userName := ctx.Value("userName").(string)
	app, err := t.system.GetOrCreateApp(appName, false)

	if err != nil {
		return nil, err
	}

	env := newTopicEnvelope()
	env.TopicReq = req
	env.Username = userName


	select {
	case app.topicUnSubQ <- env:
	case <-ctx.Done():
		err := ctx.Err()
		fmt.Printf(": %+v\n", err)
		return nil, fmt.Errorf(Status_TIMEOUT.String())

	}

	select {
	case <-env.resp:
	case <-ctx.Done():
		err := ctx.Err()
		fmt.Printf("Register, in error to kernel: %+v\n", err)
		t.stats.timeouts += 1
		return nil, fmt.Errorf(Status_TIMEOUT.String())
	}

	if env.Status != Status_SUCCESS {
		return nil,fmt.Errorf(env.Status.String())
	}

	return &TopicResp{List: env.List},nil
}


func (t *topicService) List(ctx context.Context, none *Empty) (*TopicResp, error) {

	appName := ctx.Value("appName").(string)
	userName := ctx.Value("userName").(string)
	app, err := t.system.GetOrCreateApp(appName, false)

	if err != nil {
		return nil, err
	}

	env := newTopicEnvelope()
	env.Username = userName


	select {
	case app.topicListQ <- env:
	case <-ctx.Done():
		err := ctx.Err()
		fmt.Printf(": %+v\n", err)
		return nil, fmt.Errorf(Status_TIMEOUT.String())

	}

	select {
	case <-env.resp:
	case <-ctx.Done():
		err := ctx.Err()
		fmt.Printf("Register, in error to kernel: %+v\n", err)
		t.stats.timeouts += 1
		return nil, fmt.Errorf(Status_TIMEOUT.String())
	}

	if env.Status != Status_SUCCESS {
		return nil,fmt.Errorf(env.Status.String())
	}

	return &TopicResp{List: env.List},nil
}


