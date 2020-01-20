package main

import (
	"testing"
)

func TestEncrypt(t *testing.T) {

	secret := RandomString(3000)
	cfg := createServerConfig("TEST")
	e := encrypt([]byte(secret),cfg.SYSTEM_SECRET)
	d := decrypt(e,cfg.SYSTEM_SECRET)
	t.Logf("secret is %s",string(d))
	if string(d ) != secret {
		t.Error("Encrypt and decrypt do not match")
		t.Fail()
	} else {
		t.Logf("secret is %s",string(d))
	}
}
