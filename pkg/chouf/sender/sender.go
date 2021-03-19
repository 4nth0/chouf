package sender

type Message struct {
	Level   string      `json:"level"`
	Service string      `json:"service"`
	Scope   string      `json:"scope"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
	Hash    string      `json:"hash"`
	Host    string      `json:"host"`
}
