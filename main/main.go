package main

import (
	"fmt"
	"github.com/elsejj/dzhyunsdk"
)

func main() {
	server := "ws://gw.yundzh.com/ws?token=00000003:1481769809:088c4d1cc313524828215d27b0dc6f642376d76e"
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
