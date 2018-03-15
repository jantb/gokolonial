package gokolonial

import (
	"log"
	"testing"
	"os"
)

func TestLogin(t *testing.T) {
	err, client := NewClient()
	user, err := client.Login(os.Getenv("kolonialUsername"), os.Getenv("kolonialPassword"))
	if err != nil {
		log.Fatal(err)
	}
	if !user.IsAuthenticated {
		t.Fail()
	}
}
