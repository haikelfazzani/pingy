package main

import (
	"fmt"
	"time"
)

func main() {
	responseMessage, rtt, err := Ping("google.com", time.Second)
	if err != nil {
		fmt.Println(err)
	} else if rtt > 0 {
		buff, err := responseMessage.Body.Marshal(responseMessage.Type.Protocol())
		if err != nil {
			println(err)
		}
		fmt.Println("Ping succeeded, round-trip time:", rtt, string(buff))
	} else {
		fmt.Println("Ping failed")
	}
}
