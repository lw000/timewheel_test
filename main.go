// timewheel_test project main.go
package main

import (
	"fmt"
	"time"
)

func main() {
	// WheelTest_0()
	time.AfterFunc(time.Second*time.Duration(1), func() {
		fmt.Println("time.AfterFunc()")
	})

	WheelTest_1()
}
