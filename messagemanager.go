package main

import (
	"fmt"
	"log"
	. "maestro/api"
	"sync"
	"time"
)

func (a *App) messageManager(wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Println("messageManager, Entering processing loop...")
	for {
		select {
		case env := <-a.msgQ:
			m := <-env.req
			_, ok := a.messages.msg[m.Topic]
			if !ok {
				if len(a.messages.msg) < a.cfg.MAX_NUMBER_OF_TOPICS {
					a.messages.msg[m.Topic] = make([]*Message,10)
				} else {
					env.Status = Status_MAXIMUN_NUMBER_OF_TOPICS_REACHED
				}
			}

			 if env.Status != Status_MAXIMUN_NUMBER_OF_TOPICS_REACHED {
				if len(a.messages.msg[m.Topic]) < a.cfg.MAX_NUMBER_OF_MESSAGES_PER_TOPIC {
					msg := newMessage(m)
					msg.Set(DIRTY)
					a.messages.msg[m.Topic] = append(a.messages.msg[m.Topic], msg)
					subscriptions,ok :=  a.messages.subscriptions[msg.Topic]
					if ok {
						for _, user := range subscriptions {
							user.Lock()
							user.timeLine = append(user.timeLine,msg)
							user.status.Set(DIRTY)
							user.Unlock()
						}
					}
					env.Status = Status_SUCCESS
				} else {
					env.Status = Status_MAXIMUN_NUMBER_OF_MESSAGES_PEER_TOPIC_REACHED
				}
			}
			//fmt.Printf("messageManager status is %s\n",res.Status.String())
			env.resp <- struct{}{}


		case <-time.After(a.cfg.WRITE_LATENCY * time.Millisecond):
			//fmt.Printf("messageManager: Looking for changes in message database...\n")
			for _, messages := range a.messages.msg {
				for _, aMessage := range messages {
					aMessage.Lock()
					if aMessage.Is(DIRTY) == true {
						a.messages.dirty = append(a.messages.dirty, aMessage)
						a.messages.dirtyCounter += 1
						aMessage.Clear(DIRTY)
					}
					aMessage.Unlock()
				}
			}
			if a.messages.dirtyCounter > 0 {
				fmt.Printf("messageManager, dirty messages found [%d]...\n", len(a.messages.dirty))
				select {
				case a.MsgDBQ <- a.messages.dirty:
					a.messages.dirty = make([]*Message, 0)
					a.messages.dirtyCounter = 0
				case <-time.After(2 * time.Second):
					fmt.Printf("messageManager: database server blocked ...\n")
				}
			}

		case <-a.quit:
			break
		}

	}

	fmt.Println("LoginServer, Exit processing loop...")
}

func (a *App) presistMessage(messages []*Message) {
	tx, err := a.DATABASE.Begin()
	handleError(err)
	fmt.Printf("presistMessage: Presisting %d messages ", len(messages))
	for i := 0; i < len(messages); i++ {
		m := messages[i]
		m.RLock()

		_, err := tx.Exec("INSERT OR REPLACE INTO messages (mid, topic ,Pic,parentid,status,stamp) VALUES (?,?,?,?,?,?)",
			m.GetId(), m.GetTopic(), m.GetPic(), m.GetParentId(), m.Get(), 0)
		if err != nil {
			tx.Rollback()
			log.Fatal(err)
		}

		m.RUnlock()
	}
	handleError(tx.Commit())
	/*
		for _, m := range messages {
			fmt.Printf("Presisted message[%s]\n", m.Id)
		}

	*/
	fmt.Printf("Pressised %d messages in batch....\n", len(messages))
}
