package main

import (
	. "maestro/api"
)

type topicEnvelope struct {
	req chan *TopicReq
	resp chan *TopicResp
}


