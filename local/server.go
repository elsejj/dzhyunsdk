package local

import (
	"fmt"
	"github.com/elsejj/dzhyunsdk/dzhyun"
	"github.com/gogo/protobuf/proto"
	"github.com/gorilla/websocket"
	"log"
	"net/http"
)

type Server struct {
	remote          string
	local           string
	upgrader        websocket.Upgrader
	router          *Router
	remoteClient    *websocket.Conn
	remoteSendQueue chan []byte
}

func api1(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "It's works")
}

func NewServer(remoteAddr, localAddr string) *Server {
	s := &Server{
		local:           localAddr,
		remote:          remoteAddr,
		upgrader:        websocket.Upgrader{},
		router:          NewRouter(),
		remoteClient:    nil,
		remoteSendQueue: make(chan []byte),
	}
	return s
}

func (s *Server) Run() {

	go s.ConnectRemote()

	mux := http.NewServeMux()

	mux.HandleFunc("/", api1)
	mux.Handle("/ws", s)

	http.ListenAndServe(s.local, mux)
}

func (s *Server) ConnectRemote() {
	if s.remoteClient != nil {
		return
	}
	conn, _, err := websocket.DefaultDialer.Dial(s.remote, make(http.Header))
	if err != nil {
		log.Println("connect to remote", s.remote, "failed:", err)
		return
	}
	s.remoteClient = conn
	defer s.remoteClient.Close()

	// every data in remoteSendQueue will send to remote
	go func() {
		for data := range s.remoteSendQueue {
			s.remoteClient.WriteMessage(websocket.TextMessage, data)
		}
	}()

	log.Println("connet to", s.remote, "success")

	for {
		_, buf, err := conn.ReadMessage()
		if err != nil {
			log.Println("recv from remote failed:", err)
			break
		}
		var ua dzhyun.UAResponse
		err = proto.Unmarshal(buf, &ua)
		if err != nil {
			log.Println("parse remote incoming failed", err)
			continue
		}
		client := s.router.GetByQid(ua.Qid)
		if client != nil {
			client.Send(&ua)
		}
	}

}

func (s *Server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	c, err := s.upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println("client connect ws faield:", err)
		return
	}
	defer c.Close()

	peer := NewClient(s.router.NextCid(), c)
	s.router.Add(peer)

	for {
		_, msg, err := c.ReadMessage()
		if err != nil {
			log.Println("read ws client failed:", err)
			break
		}
		//TODO: process req here
		path := peer.ReBuildQS(msg)
		s.remoteSendQueue <- []byte(path)
	}
}
