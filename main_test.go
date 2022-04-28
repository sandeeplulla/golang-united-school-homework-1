package main

import (
	"testing"
)

func TestGetMessage(t *testing.T) {
	message := GetMessage()
	if message != "Hello 🗺️ !" {
		t.Errorf("Incorrect String returned")
	}
}
