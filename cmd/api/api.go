package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"sync"
	"time"

	"github.com/4nth0/chouf/pkg/chouf/sender"
	log "github.com/sirupsen/logrus"
)

var store = NewStore()

type Entry struct {
	CreatedAt time.Time       `json:"createAt"`
	UpdatedAt time.Time       `json:"updatedAt"`
	Counter   int64           `json:"counter"`
	Message   *sender.Message `json:"message"`
}

type Storage struct {
	entries map[string]*Entry
	mu      *sync.Mutex
}

type Rule struct {
	Limit int64
}

var rules = map[string]map[string]*Rule{
	"wizads": {
		"*": &Rule{
			Limit: 10,
		},
	},
}

func SearcheRule(service, scope string) *Rule {
	if rules[service] == nil {
		return nil
	}
	if rules[service][scope] != nil {
		return rules[service][scope]
	}
	if rules[service]["*"] != nil {
		return rules[service]["*"]
	}
	return nil
}

func checkForAlert(entry *Entry) {

	rule := SearcheRule(entry.Message.Service, entry.Message.Scope)

	if rule == nil {
		return
	}

	if rule.Limit == entry.Counter {
		fmt.Println("")
		fmt.Println("--")
		fmt.Println("-- Alert --")
		fmt.Println("-- Service", entry.Message.Service)
		fmt.Println("-- Scope", entry.Message.Scope)
		fmt.Println("-- Message", entry.Message.Message)
		fmt.Println("-- Data", entry.Message.Data)
		fmt.Println("--")
		fmt.Println("")
	}
}

func NewStore() *Storage {
	return &Storage{
		entries: make(map[string]*Entry),
		mu:      &sync.Mutex{},
	}
}

func (s Storage) All() []*Entry {
	output := make([]*Entry, len(s.entries))
	index := 0

	for _, entry := range s.entries {
		output[index] = entry
		index++
	}

	return output
}

func (s *Storage) Upsert(message sender.Message) *Entry {
	s.mu.Lock()
	defer s.mu.Unlock()

	if s.entries[message.Hash] == nil {
		s.entries[message.Hash] = &Entry{
			Counter:   1,
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
			Message:   &message,
		}
	} else {
		s.entries[message.Hash].Counter++
		s.entries[message.Hash].UpdatedAt = time.Now()
	}

	return s.entries[message.Hash]
}

func main() {

	server := http.NewServeMux()

	server.HandleFunc("/message", func(w http.ResponseWriter, req *http.Request) {
		switch req.Method {
		case http.MethodGet:
			entries := store.All()

			b, _ := json.Marshal(entries)
			w.Header().Add("content-type", "application/json")
			w.Write(b)

		case http.MethodPost:
			decoder := json.NewDecoder(req.Body)
			var message sender.Message
			err := decoder.Decode(&message)
			if err != nil {
				panic(err)
			}
			inserted := store.Upsert(message)
			checkForAlert(inserted)
		}
	})

	err := http.ListenAndServe(":8787", server)
	if err != nil {
		log.WithFields(
			log.Fields{
				"err":  err,
				"port": "8787",
			}).Error("Unable to start server listening.")
	}
}
