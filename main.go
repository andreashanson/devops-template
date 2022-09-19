package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("HELLLLLLO DEVOPS")
	for i := 0; i < 10000; i++ {
		fmt.Println(fmt.Sprintf("HELLO %d", i))
		time.Sleep(time.Second * 10)
	}
	fmt.Println("GOOD BYE")
}
