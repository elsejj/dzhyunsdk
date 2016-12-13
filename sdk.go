package dzhyunsdk

import (
	"github.com/elsejj/dzhyunsdk/local"
)

// StartSDK connect to server and make a local listen
func StartSDK(server, localAddr string) {
	s := local.NewServer(server, localAddr)
	go s.Run()
}
