package main

import (
	"fmt"

	"github.com/ericfialkowski/exitOn"
)

func main() {
	fmt.Println("Hit any key to exit")
	//time.AfterFunc(time.Second*5, func() {
	//	fmt.Println("Canceling")
	//	exitOn.Cancel()
	//})
	_ = exitOn.AnyKeyWait()
}
