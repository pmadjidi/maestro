package main

import (
	. "maestro/api"
)



type msgEnvelope struct {
	req  chan *MsgReq
	resp chan *MsgResp
}


func newMsgEnvelope() *msgEnvelope {
	return &msgEnvelope{make(chan *MsgReq, 1), make(chan *MsgResp, 1)}
}


type msgService struct {
	name string
	Q chan *msgEnvelope
	cfg *ServerConfig
	stats *metrics
}

func newMsgService (Q chan *msgEnvelope,cfg *ServerConfig) *msgService {
	return &msgService{"msgService",Q,cfg,newMetrics()}
}



