package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("start timer")
	timer := time.After(10 * time.Second)
	<-timer
	fmt.Println("timer is finished")
}
