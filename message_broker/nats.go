package messagebroker

import (
	"log"

	"github.com/nats-io/nats.go"
)

var NatsConn *nats.Conn

func InitNATS() {
	nc, err := nats.Connect("nats://nats:4222")
	if err != nil {
		log.Fatalf("Failed to connect to NATS: %v", err)
	}
	NatsConn = nc
}

func CloseNATS() {
	if NatsConn != nil {
		NatsConn.Close()
	}
}
