package local

import (
	"fmt"
	"sync"
)

type Router struct {
	table  map[string]*Client
	locker *sync.Mutex
}

func NewRouter() *Router {
	r := &Router{
		make(map[string]*Client),
		&sync.Mutex{},
	}
	return r
}

func (r *Router) NextCid() string {
	r.locker.Lock()
	defer r.locker.Unlock()
	return fmt.Sprintf("%04d", len(r.table))
}

func (r *Router) Add(client *Client) {
	r.locker.Lock()
	defer r.locker.Unlock()
	r.table[client.ID()] = client
}

func (r *Router) Rm(cid string) {
	r.locker.Lock()
	defer r.locker.Unlock()
	delete(r.table, cid)
}

func (r *Router) Get(cid string) *Client {
	r.locker.Lock()
	defer r.locker.Unlock()
	client, _ := r.table[cid]
	return client
}

func (r *Router) RmByQid(qid string) {
	r.locker.Lock()
	defer r.locker.Unlock()
	for _, client := range r.table {
		if client.HasQid(qid) {
			client.RmQid(qid)
		}
	}
}
func (r *Router) GetByQid(qid string) *Client {
	r.locker.Lock()
	defer r.locker.Unlock()
	for _, client := range r.table {
		if client.HasQid(qid) {
			return client
		}
	}
	return nil
}
