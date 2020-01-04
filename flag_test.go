package main

import (
	"fmt"
	"testing"
	. "maestro/api"
)

func TestFlagInit(t *testing.T) {
	u := newUser(&RegisterReq{})
	if !u.status.Is(DIRTY) {
		t.Errorf("Flag should be dirty at user creation %d",u.status.Get())
	}
}

func TestFlagClearDirty(t *testing.T) {
	u := newUser(&RegisterReq{})
	u.status.Clear(DIRTY)
	if u.status.Get() != 0 {
		t.Errorf("Flag should be zero %d",u.status.Get())
	}
}


func TestFlagDirtyAndBlocked(t *testing.T) {
	u := newUser(&RegisterReq{})
	u.status.Set(BLOCKED)
	fmt.Printf("flag set to blocked  %t \n",u.status.Is(BLOCKED))

	if !u.status.Is(BLOCKED) {
		t.Errorf("Flag should be blocked and dirty %d",u.status.Get())
	}

	if !u.status.Is(DIRTY) {
		t.Errorf("Flag should be blocked and dirty %d",u.status.Get())
	}
}


func TestFlagGeneral(t *testing.T) {
	u := newUser(&RegisterReq{})
	u.status.Set(BLOCKED)
	u.status.Set(DELETED)
	//u.status.Clear(DIRTY)

	if !u.status.Is(BLOCKED) || !u.status.Is(DELETED) {
		t.Errorf("Flag should be blocked and deleted %d",u.status.Get())
	}

	u.status.Clear(BLOCKED)
	u.status.Clear(DELETED)

	if u.status.Is(BLOCKED) || u.status.Is(DELETED) {
		t.Errorf("Flag should not be blocked and deleted %d",u.status.Get())
	}


}


