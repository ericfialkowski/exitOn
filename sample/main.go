package main

import (
	"fmt"

	"github.com/ericfialkowski/exitOn"
)

func main() {
	fmt.Println("Hit any key to exit")
	_ = exitOn.AnyKey()
	select {}
}
