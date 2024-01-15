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
	return key(keyboard.KeyEsc, true, false)
}

func AnyKeyWait() error {
	return key(keyboard.KeyEsc, true, true)
}

func EscKey() error {
	return key(keyboard.KeyEsc, false, false)
}

func EscKeyWait() error {
	return key(keyboard.KeyEsc, false, true)
}

func SpaceKey() error {
	return key(keyboard.KeySpace, false, false)
}

func SpaceKeyWait() error {
	return key(keyboard.KeySpace, false, true)
}

func EnterKey() error {
	return key(keyboard.KeyEnter, false, false)
}

func EnterKeyWait() error {
	return key(keyboard.KeyEnter, false, true)
}

func key(key keyboard.Key, anyKey, wait bool) error {
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
			if anyKey || k == key {
				os.Exit(0)
			}
		}
	}()

	if wait {
		select {}
	}
	return nil
}
