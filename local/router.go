package local

import (
	"fmt"
)

type Router struct {
	table map[string]*Client
}

func NewRouter() *Router {
	r := &Router{
		make(map[string]*Client),
	}
	return r
}

func (r *Router) NextCid() string {
	return fmt.Sprintf("%04d", len(r.table))
}

func (r *Router) Add(client *Client) {
	r.table[client.ID()] = client
}

func (r *Router) Rm(cid string) {
	delete(r.table, cid)
}

func (r *Router) Get(cid string) *Client {
	client, _ := r.table[cid]
	return client
}

func (r *Router) RmByQid(qid string) {
	for _, client := range r.table {
		if client.HasQid(qid) {
			client.RmQid(qid)
		}
	}
}
func (r *Router) GetByQid(qid string) *Client {
	for _, client := range r.table {
		if client.HasQid(qid) {
			return client
		}
	}
	return nil
}
