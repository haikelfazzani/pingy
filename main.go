package main

import (
	"flag"
	"fmt"
	"time"

	"github.com/haikelfazzani/ping/pingy"
)

var (
	url     = flag.String("u", "", "URL to ping (exp: google.com)")
	addr    = flag.String("a", "", "IP address to ping (exp: 142.250.184.78)")
	version = flag.Bool("v", false, "App version")
)

func main() {
	flag.Parse()

	if *version {
		fmt.Println("Pingy v1.0.0")
		return
	}

	responseMessage, rtt, err := pingy.Ping(*url, time.Second, *addr)

	if err != nil {
		fmt.Println(err)
	} else if rtt > 0 {
		protocol := responseMessage.Type.Protocol()
		fmt.Printf("Ping succeeded, round-trip time: %s \nProtocol: %d\n", rtt, protocol)
	} else {
		fmt.Println("Ping failed")
	}
}
