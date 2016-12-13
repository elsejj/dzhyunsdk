package local

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/elsejj/dzhyunsdk/dzhyun"
	"github.com/golang/protobuf/proto"
	"github.com/gorilla/websocket"
	"log"
	"net/url"
)

type Client struct {
	id    string
	conn  *websocket.Conn
	qids  map[string]string
	reqsn int
}

func NewClient(id string, conn *websocket.Conn) *Client {
	c := &Client{
		id:    id,
		conn:  conn,
		qids:  make(map[string]string),
		reqsn: 0,
	}
	return c
}

func (c *Client) ID() string {
	return c.id
}

func (c *Client) NewQid() string {
	qid := fmt.Sprintf("%v_%d", c.id, c.reqsn)
	c.reqsn++
	return qid
}

func (c *Client) AddUserQid(qid, uqid string) {
	c.qids[qid] = uqid
}

func (c *Client) ReBuildQS(buf []byte) string {
	p := bytes.IndexByte(buf, '?')
	path := string(buf[0:p])
	var qs url.Values
	var err error
	if p > 0 {
		qs, err = url.ParseQuery(string(buf[p+1:]))
		if err != nil {
			log.Println("user request is ")
		}
	} else {
		qs = make(url.Values)
	}

	qid := qs.Get("qid")
	sub := qs.Get("sub") == "1"
	qs.Set("output", "pb")

	newQid := c.NewQid()
	qs.Set("qid", newQid)
	c.AddUserQid(newQid, qid)

	if sub {
	}

	path = path + "?" + qs.Encode()

	return path
}

func (c *Client) HasQid(qid string) bool {
	_, ok := c.qids[qid]
	return ok
}

func (c *Client) GetOriginQid(qid string) string {
	origin, _ := c.qids[qid]
	return origin
}

func (c *Client) RmQid(qid string) {
	delete(c.qids, qid)
}

func (c *Client) Send(ua *dzhyun.UAResponse) {

	fmt.Println("send to", c.ID(), ua.Qid)
	var err error
	ua.Qid = c.GetOriginQid(ua.Qid)

	var msg dzhyun.MSG

	err = proto.Unmarshal(ua.Data, &msg)
	if err != nil {
		log.Println("\t", "result is not a msg", err)
		return
	}

	msgData, err := json.MarshalIndent(msg, "", "  ")
	if err != nil {
		log.Println("client write result faield:", err)
	}

	data := fmt.Sprintf(`{
  "Qid": "%s",
  "Counter": %d,
  "Err": %d,
  "Data": %s
}`, ua.Qid, ua.Counter, ua.Err, msgData)


	var mm proto.Message

	

	c.conn.WriteMessage(websocket.TextMessage, []byte(data))
}
