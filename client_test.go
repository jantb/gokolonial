package gokolonial

import (
	"testing"
)

func TestNewCLient(t *testing.T) {
	err, _ := NewClient()
	if err != nil {
		t.Fail()
	}
}
