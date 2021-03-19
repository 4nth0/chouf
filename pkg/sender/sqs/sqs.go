package sqs

import (
	"encoding/json"
	"fmt"

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

type SQSSenderClient struct{}

func SQSSender() *SQSSenderClient {
	return &SQSSenderClient{}
}

func (s SQSSenderClient) Send(message sender.Message) error {
	b, _ := json.Marshal(message)

	fmt.Println("b: ", string(b))

	return nil
}
