package chouf

import (
	"crypto/sha1"
	"encoding/json"
	"fmt"
	"os"

	"github.com/4nth0/chouf/pkg/chouf/sender"
	"github.com/4nth0/chouf/pkg/chouf/sender/api"
)

type Sender interface {
	Send(sender.Message) error
}

type Client struct {
	service string
	scope   string
	sender  *api.Client
}

func New(service, user, key string) (*Client, error) {
	api, err := api.New(user, key)
	if err != nil {
		return nil, err
	}

	client := Client{
		service: service,
		sender:  api,
	}

	return &client, nil
}

func (c Client) WithScope(scope string) *Client {
	return &Client{
		service: c.service,
		scope:   scope,
		sender:  c.sender,
	}
}

func (c Client) Notify(level, message string, params ...interface{}) {
	hash := c.generateHash(c.service, c.scope, message, params)

	fmt.Println("hash: ", hash)

	host, _ := os.Hostname()
	m := sender.Message{
		Level:   level,
		Service: c.service,
		Scope:   c.scope,
		Message: message,
		Data:    params,
		Hash:    hash,
		Host:    host,
	}

	c.sender.Send(m)
}

func (c Client) NotifyInfo(message string, params ...interface{}) {
	c.Notify("info", message, params)
}

func (c Client) NotifyError(message string, params ...interface{}) {
	c.Notify("error", message, params)
}

func (c Client) generateHash(params ...interface{}) string {
	b, _ := json.Marshal(params)

	h := sha1.New()
	h.Write([]byte(b))

	return fmt.Sprintf("%x", h.Sum(nil))
}
