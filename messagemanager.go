package main

import (
	"fmt"
	. "maestro/api"
	"time"
)

func (a *App) messageManager() {
	signalmsgReQ := false
	a.log("messageManager, Entering processing loop")
loop:
	for {
		select {
		case env, ok := <-a.msgRecQ:
			if ok {
				for _, m := range env.messages {
					_, ok := a.msg[m.Topic]
					if !ok {
						if len(a.msg) < a.cfg.MAX_NUMBER_OF_TOPICS {
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
				if signalmsgReQ {
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
