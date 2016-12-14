package dzhyunsdk

import (
	"C"

	"github.com/elsejj/dzhyunsdk/local"
)

// StartSDK connect to server and make a local listen
// export StartSDK
func StartSDK(server, localAddr string) {
	s := local.NewServer(server, localAddr)
	go s.Run()
}
