package main

import (
	"fmt"
	"github.com/elsejj/dzhyunsdk"
)

func main() {
	server := "ws://gw.yundzh.com/ws?token=00000003:1481529152:b1ef39173bc5149c724186de9fdada26c40b9fd5"
	local := "127.0.0.1:9999"
	dzhyunsdk.StartSDK(server, local)

	var cmd string
	for {
		fmt.Scanln(&cmd)
		if cmd == "quit" {
			break
		}
	}
}
