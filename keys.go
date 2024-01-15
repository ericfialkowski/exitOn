package exitOn

import (
	"errors"
	"log"
	"os"
	"sync/atomic"
	"time"

	"github.com/eiannone/keyboard"
)

var (
	handlerSet         atomic.Bool
	MultipleHandlerErr = errors.New("only one key exit handler can be set at a time")
	AttachErr          = errors.New("error attaching to keyboard for input")
	running            = true
)

func Cancel() {
	running = false
}

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

	keysEvents, err := keyboard.GetKeys(10)
	if err != nil {
		return errors.Join(AttachErr, err)
	}

	go func() {
		defer func() {
			handlerSet.Store(false)
			_ = keyboard.Close()
		}()
		for running {
			select {
			case event := <-keysEvents:
				if event.Err != nil {
					log.Printf("Error checking for keys %v", err) //TODO something better
				}
				if anyKey || event.Key == key {
					os.Exit(0)
				}
			case <-time.After(time.Second):
				// noop/break to check the running flag
			}

		}
	}()

	if wait {
		for running {
			<-time.After(time.Second) // break to check the running flag
		}
	}
	return nil
}
