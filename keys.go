package exitOn

import (
	"errors"
	"os"
	"sync/atomic"

	"github.com/eiannone/keyboard"
)

var (
	handlerSet         atomic.Bool
	MultipleHandlerErr = errors.New("only one key exit handler can be set at a time")
	AttachErr          = errors.New("error attaching to keyboard for input")
)

func AnyKey() error {
	if !handlerSet.CompareAndSwap(false, true) {
		return MultipleHandlerErr
	}
	if err := keyboard.Open(); err != nil {
		return errors.Join(AttachErr, err)
	}

	go func() {
		defer func() {
			handlerSet.Store(false)
			_ = keyboard.Close()
		}()
		_, _, _ = keyboard.GetSingleKey()
		os.Exit(0)
	}()
	return nil
}

func EscKey() error {
	return singleKey(keyboard.KeyEsc)
}
func SpaceKey() error {
	return singleKey(keyboard.KeySpace)
}

func EnterKey() error {
	return singleKey(keyboard.KeyEnter)
}

func singleKey(key keyboard.Key) error {
	if !handlerSet.CompareAndSwap(false, true) {
		return MultipleHandlerErr
	}

	if err := keyboard.Open(); err != nil {
		return errors.Join(AttachErr, err)
	}

	go func() {
		defer func() {
			handlerSet.Store(false)
			_ = keyboard.Close()
		}()
		for {
			_, k, _ := keyboard.GetSingleKey()
			if k == key {
				os.Exit(0)
			}
		}
	}()
	return nil
}
