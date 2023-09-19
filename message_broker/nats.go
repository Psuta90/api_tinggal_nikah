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

// SubscribeMessage adalah fungsi untuk menerima pesan dari NATS JetStream
func SubscribeMessage(subject string, callback func(*nats.Msg)) (*nats.Subscription, error) {
	sub, err := NatsConn.Subscribe(subject, callback)
	if err != nil {
		log.Printf("Failed to subscribe to subject %s: %v", subject, err)
		return nil, err
	}
	return sub, nil
}
