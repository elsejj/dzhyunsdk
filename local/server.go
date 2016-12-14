package local

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/websocket"
	"log"
	"net"
	"net/http"
)

type Server struct {
	local    string
	upgrader websocket.Upgrader
	router   *Router
	remote   *Remote
}

func (s *Server) index(w http.ResponseWriter, r *http.Request) {
	peer := NewClient(s.router.NextCid(), nil)
	s.router.Add(peer)

	fmt.Println("http request", r.URL.String())
	path := peer.ReBuildQS([]byte(r.URL.String()))
	s.remote.sendQueue <- []byte(path)
	data := <-peer.SendQueue

	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.Write(data)
	s.router.Rm(peer.ID())
}

func NewServer(remoteAddr, localAddr string) *Server {
	s := &Server{
		local:    localAddr,
		upgrader: websocket.Upgrader{},
		router:   NewRouter(),
		remote:   NewRemote(remoteAddr),
	}
	return s
}

func (s *Server) Run() {

	ln, err := net.Listen("tcp", s.local)
	if err != nil {
		log.Println("start local server faield:", err)
		return
	}

	go s.remote.Start(s.router)

	mux := http.NewServeMux()
	mux.HandleFunc("/", s.index)
	mux.HandleFunc("/ws", s.ws)
	mux.HandleFunc("/rs", s.remoteStatus)
	err = http.Serve(ln, mux)
	if err != nil {
		log.Println("local http server faield:", err)
		s.remote.Shutdown()
	}
}

func (s *Server) ws(w http.ResponseWriter, r *http.Request) {
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
		path := peer.ReBuildQS(msg)
		s.remote.sendQueue <- []byte(path)
	}
	s.router.Rm(peer.ID())
}

func (s *Server) remoteStatus(w http.ResponseWriter, r *http.Request) {
	data, err := json.MarshalIndent(s.remote, "", "  ")
	if err != nil {
		fmt.Fprintln(w, err)
	}
	w.Write(data)
}
