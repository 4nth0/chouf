package api

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"

	"github.com/4nth0/chouf/pkg/chouf/sender"
)

type Message struct {
	Level   string      `json:"level"`
	Service string      `json:"service"`
	Scope   string      `json:"scope"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
	Hash    string      `json:"hash"`
	Host    string      `json:"host"`
}

type Client struct {
	user   string
	key    string
	client *http.Client
}

func New(user, key string) (*Client, error) {
	if user == "" {
		return nil, errors.New("No User supplied")
	}
	if key == "" {
		return nil, errors.New("No API Key supplied")
	}

	httpClient := &http.Client{
		Timeout: time.Second * 10,
	}

	client := Client{
		user:   user,
		key:    key,
		client: httpClient,
	}

	return &client, nil
}

func (c Client) Send(message sender.Message) error {
	jsonStr, _ := json.Marshal(message)

	req, _ := http.NewRequest("POST", "http://127.0.0.1:8787/message", bytes.NewBuffer(jsonStr))
	req.Header.Set("Content-Type", "application/json")

	resp, err := c.client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	fmt.Println("response Status:", resp.Status)
	fmt.Println("response Headers:", resp.Header)
	body, _ := ioutil.ReadAll(resp.Body)
	fmt.Println("response Body:", string(body))

	return nil
}
