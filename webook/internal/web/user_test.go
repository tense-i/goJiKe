package web

import (
	"golang.org/x/crypto/bcrypt"
	"testing"
)

func TestEncrpt(t *testing.T) {
	passwd := "#12345wzc"
	encrypt, err := bcrypt.GenerateFromPassword([]byte(passwd), bcrypt.DefaultCost)
	if err != nil {
		t.Fatal(err)
	}
	err = bcrypt.CompareHashAndPassword(encrypt, []byte(passwd))

}
