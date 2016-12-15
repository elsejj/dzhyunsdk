package local

import (
	"log"
	"net/http"

	"github.com/elsejj/dzhyunsdk/dzhyun"
	"github.com/golang/protobuf/proto"
	"github.com/golang/snappy"
	"github.com/gorilla/websocket"
)

type Remote struct {
	Online    bool
	SendBytes int
	RecvBytes int
	server    string
	conn      *websocket.Conn
	sendQueue chan []byte
	events    chan int
}

func NewRemote(url string) *Remote {
	r := &Remote{
		Online:    false,
		SendBytes: 0,
		RecvBytes: 0,
		server:    url,
		conn:      nil,
		sendQueue: make(chan []byte),
		events:    make(chan int),
	}
	return r
}

func (r *Remote) Start(router *Router) {
	if r.conn != nil {
		return
	}
	go r.startSend()
	conn, _, err := websocket.DefaultDialer.Dial(r.server, make(http.Header))
	if err != nil {
		log.Println("connect to remote", r.server, "failed:", err)
		return
	}
	r.conn = conn
	defer r.Shutdown()

	log.Println("connet to", r.server, "success")

	r.Online = true
	r.events <- 1

	for {
		_, buf, err := conn.ReadMessage()
		if err != nil {
			log.Println("recv from remote failed:", err)
			break
		}
		r.RecvBytes += len(buf)

		buf, err = snappy.Decode(nil, buf)
		if err != nil {
			log.Println("decompress remote data failed:", err)
			continue
		}

		var ua dzhyun.UAResponse
		err = proto.Unmarshal(buf, &ua)
		if err != nil {
			log.Println("parse remote incoming failed", err)
			continue
		}
		client := router.GetByQid(ua.Qid)
		if client != nil {
			client.Send(&ua)
		}
	}
}

func (r *Remote) startSend() {
	beforeConnectedCache := make([][]byte, 1024)
	cacheUsed := 0
	online := false
	for {
		select {
		case event := <-r.events:
			if event == -1 {
				online = false
				break
			} else if event == 1 {
				online = true
				for i, data := range beforeConnectedCache[0:cacheUsed] {
					r.SendBytes += len(data)
					r.conn.WriteMessage(websocket.TextMessage, data)
					beforeConnectedCache[i] = nil
				}
			}
		case data := <-r.sendQueue:
			if online {
				r.SendBytes += len(data)
				r.conn.WriteMessage(websocket.TextMessage, data)
			} else {
				beforeConnectedCache[cacheUsed] = data
				cacheUsed++
			}
		}
	}
}

func (r *Remote) Shutdown() {
	if r.Online {
		r.events <- -1
		r.conn.Close()
		r.Online = false
	}
}
