package main

import (
	"fmt"
	"log"
	. "maestro/api"
	"time"
)

func (a *App) messageManager() {
	fmt.Println("messageManager, Entering processing loop...")
	for {
		select {
		case env := <-a.msgQ:
			m := <-env.req
			_,ok := a.messages.msg[m.Topic]
			if !ok {
				a.messages.msg[m.Topic] = make(map[string]*Message)
			}
			msg := newMessage(m)
			a.messages.msg[m.Topic][m.Id] = msg
			env.resp <- &MsgResp{Id:msg.GetId(),Status: Status_SUCCESS}

		case <-time.After(20 * time.Second):
			fmt.Printf("messageManager: Looking for changes in message database...\n")
			for _, messages := range a.messages.msg {
				for _, aMessage := range messages {
					aMessage.Lock()
					if aMessage.Is(DIRTY) == true {
						a.messages.dirty = append(a.messages.dirty,aMessage)
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
			} else {
				fmt.Printf("messageManager: No change to message database...\n")
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
	fmt.Printf("presistMessage: Presisting %d messages ",len(messages))
	for i := 0 ; i <  len(messages) ; i++ {
		m := messages[i]
		m.RLock()

		_, err := tx.Exec("INSERT OR REPLACE INTO messages (mid, topic ,Pic,parentid,status,stamp) VALUES (?,?,?,?,?,?)",
			m.GetId(), m.GetTopic(),m.GetPic(),m.GetParentId(),m.Get(),m.GetTimeName())
		if err != nil {
			tx.Rollback()
			log.Fatal(err)
		}

		m.RUnlock()
	}
	handleError(tx.Commit())
	for _, m := range messages {
		fmt.Printf("Presisted message[%s]\n", m.Id)
	}
}


