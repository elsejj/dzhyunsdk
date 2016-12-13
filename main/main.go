package main

import (
	"fmt"
	"github.com/elsejj/dzhyunsdk"
)

func main() {
	server := "ws://gw.yundzh.com/ws?token=00000003:1481705008:48ad77788e2e588c9eab9c52618173d51a600abe"
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
