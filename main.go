package main

import (
	"time"

	"github.com/4nth0/chouf/pkg/chouf"
)

func main() {
	c, _ := chouf.New("wizads", "tf1_pub", "7967869769866678")

	for index := 0; index < 10; index++ {

		time.Sleep(10 * time.Millisecond)

		c.WithScope("Move ID Matcher").NotifyInfo("expired program")
		c.WithScope("Move ID Matcher").NotifyInfo("expired program")
		c.WithScope("Move ID Matcher").NotifyInfo("expired program")
		c.WithScope("Move ID Matcher").NotifyInfo("expired program")
		c.WithScope("Move ID Matcher").NotifyInfo("expired program")

		C3PO := c.WithScope("C3PO")

		C3PO.NotifyError("invalid bitrate", "987687957986", "1700")
		C3PO.NotifyError("invalid bitrate", "978668767868", "2000")
	}
}
