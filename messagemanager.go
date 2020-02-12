package main

import (
	"fmt"
	"github.com/google/uuid"
	. "maestro/api"
	"time"
)

func (a *App) messageManager() {
	signalmsgReQ := false
	signalTopicSub := false
	signalTopicUnSub := false
	a.log("messageManager, Entering processing loop")
loop:
	for {
		select {
		case env, ok := <-a.msgQ:
			if ok {
				for _, m := range env.messages {
					_, ok := a.msg[m.Topic]
					if !ok {
						if len(a.msg) < a.cfg.MAX_NUMBER_OF_TOPICS {
							topic  := &Topic{
								Id:                   uuid.New().String(),
								Tag:                  m.Topic,
							}
							a.topics[topic] = Status_NEW
							a.msg[m.Topic] = make(map[string]*Message, 0)
						} else {
							env.Status = Status_MAXIMUN_NUMBER_OF_TOPICS_REACHED
							env.resp <- notify{}
							continue
						}
					}

					if len(a.msg[m.Topic]) < a.cfg.MAX_NUMBER_OF_MESSAGES_PER_TOPIC {
						m.Set(DIRTY)
						a.msg[m.Topic][m.Uuid] = m
						subscriptions, ok := a.subscriptions[m.Topic]
						if ok {
							for _, user := range subscriptions {
								user.Lock()
								user.timeLine = append(user.timeLine, m)
								user.status.Set(DIRTY)
								user.Unlock()
							}
						}
						env.Status = Status_SUCCESS
					} else {
						env.Status = Status_MAXIMUN_NUMBER_OF_MESSAGES_PEER_TOPIC_REACHED
					}
				}
				env.resp <- notify{}
			} else {
				signalmsgReQ = true
				if signalmsgReQ && signalTopicSub && signalTopicUnSub {
					break loop
				}
			}

		case <-time.After(a.cfg.WRITE_LATENCY * time.Millisecond):
			//fmt.Printf("messageManager: Looking for changes in message database...\n")
			for _, messages := range a.msg {
				for _, aMessage := range messages {
					aMessage.Lock()
					if aMessage.Is(DIRTY) == true {
						a.mdirty = append(a.mdirty, aMessage)
						a.mdirtyCounter += 1
						aMessage.Clear(DIRTY)
					}
					aMessage.Unlock()
				}
			}
			if a.mdirtyCounter > 0 {
				a.log(fmt.Sprintf("messageManager, dirty messages found [%d]", len(a.mdirty)))
				select {
				case a.MsgDBQ <- a.mdirty:
					a.mdirty = make([]*Message, 0)
					a.mdirtyCounter = 0
				case <-time.After(2 * time.Second):
					a.log("messageManager: database server blocked")
				}
			}

		case env, ok := <-a.topicSubQ:
			if ok {
				newEnv := newUserEnvelope()
				newEnv.Username = env.Username
				a.userQ <- newEnv
				<-newEnv.resp
				if newEnv.Status == Status_SUCCESS {
					for _, t := range env.List {
						_, ok := a.subscriptions[t.Tag]
						if ok {
							a.subscriptions[t.Tag] = append(a.subscriptions[t.Tag], newEnv.User)
						} else {
							a.subscriptions[t.Tag] = append(make([]*User, 0), newEnv.User)
							a.topics[t] = Status_NEW
						}
					}
					env.Status = Status_SUCCESS
				} else {
					env.Status = newEnv.Status
				}
				env.resp <- notify{}
			} else {
				signalTopicSub = true
			}
		case env, ok := <-a.topicUnSubQ:
			if ok {
				newEnv := newUserEnvelope()
				newEnv.Username = env.Username
				a.userQ <- newEnv
				<-newEnv.resp
				if newEnv.Status == Status_SUCCESS {
					for _, t := range env.List {
						subscribers, ok := a.subscriptions[t.Tag]
						if ok {
							for i,user := range subscribers {
								if user == newEnv.User {
									subscribers[i] = subscribers[len(subscribers)-1] // Copy last element to index i
									subscribers[len(subscribers)-1] = nil   // Erase last element (write zero value)
									subscribers = subscribers[:len(subscribers)-1]   // Truncate slice
								}
							}
						}
					}
					env.Status = Status_SUCCESS
				} else {
					env.Status = newEnv.Status
				}
				env.resp <- notify{}
			} else {
				signalTopicUnSub = true
			}

		case env, ok := <-a.topicListQ:
			if ok {
				topics := make([]*Topic,0)
				for k,_ := range a.topics {
					topics = append(topics,k)
				}
				env.Status = Status_SUCCESS
				env.List = topics
				env.resp <- notify{}
			} else {
				signalTopicUnSub = true
			}




		case <-a.quit:
			break loop
		}
	}
	a.log("Exiting messageManger")
}

func (a *App) presistMessage(messages []*Message) {
	tx, err := a.Begin()
	handleError(err)
	a.log(fmt.Sprintf("presistMessage: Presisting %d messages ", len(messages)))
	for i := 0; i < len(messages); i++ {
		m := messages[i]
		m.RLock()

		_, err := tx.Exec("INSERT OR REPLACE INTO messages (mid, topic ,Pic,parentid,status,stamp) VALUES (?,?,?,?,?,?)",
			m.GetUuid(), m.GetTopic(), m.GetPic(), m.GetParentId(), m.Get(), 0)
		if err != nil {
			tx.Rollback()
		}

		m.RUnlock()
	}
	handleError(tx.Commit())
	/*
		for _, m := range messages {
			fmt.Printf("Presisted message[%s]\n", m.Id)
		}

	*/
	a.log(fmt.Sprintf("Pressised %d messages in batch", len(messages)))
}
