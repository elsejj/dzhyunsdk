package main

import (
	"flag"
	"fmt"
	"github.com/elsejj/dzhyunsdk"
)

var (
	server string
	token  string
	local  string
)

func main() {

	flag.StringVar(&server, "server", "ws://gw.yundzh.com/ws", "the server url")
	flag.StringVar(&token, "token", "", "the server access token")
	flag.StringVar(&local, "local", "127.0.0.1:9999", "the local address")

	flag.Parse()

	dzhyunsdk.StartSDK(fmt.Sprintf("%s?token=%s", server, token), local)

	var cmd string
	for {
		fmt.Scanln(&cmd)
		if cmd == "quit" {
			break
		}
	}
}
