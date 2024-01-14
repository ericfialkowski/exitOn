package exitOn

import (
	"errors"
	"testing"
)

func TestMultipleHandlers(t *testing.T) {
	err := AnyKey()
	if err != nil {
		t.Fatalf("Unexpected error :%v", err)
	}

	err2 := EscKey()
	if err2 == nil {
		t.Fatalf("Expected an error")
	}
	if !errors.Is(err2, MultipleHandlerErr) {
		t.Fatalf("Received wrong error :%v", err)
	}
}

func TestMultipleSingleKeyHandlers(t *testing.T) {
	err := EnterKey()
	if err != nil {
		t.Fatalf("Unexpected error :%v", err)
	}

	err2 := SpaceKey()
	if err2 == nil {
		t.Fatalf("Expected an error")
	}
	if !errors.Is(err2, MultipleHandlerErr) {
		t.Fatalf("Received wrong error :%v", err)
	}

}
