package services

import (
	messagebroker "api_tinggal_nikah/message_broker"
	"fmt"

	"github.com/labstack/echo/v4"
	"github.com/nats-io/nats.go"
)

func SubscribePaymentStatusServices(c echo.Context, order_id string) error {

	subjectNats := "payment.status." + order_id

	// Create a channel for this client
	clientCh := make(chan string)
	defer close(clientCh)

	// Subscribe to NATS subject
	sub, err := messagebroker.NatsConn.Subscribe(subjectNats, func(msg *nats.Msg) {
		// Send message to the client channel
		clientCh <- string(msg.Data)
	})
	if err != nil {
		return err
	}
	defer sub.Unsubscribe()

	// Send a welcome message to the client
	fmt.Fprintf(c.Response().Writer, "data: Welcome!\n\n")
	c.Response().Flush()

	// Continuously send messages to the client
	for {
		select {
		case message := <-clientCh:
			// Send the message to the client
			fmt.Fprintf(c.Response().Writer, "data: %s\n\n", message)
			c.Response().Flush()
		case <-c.Request().Context().Done():
			return nil // Client disconnected
		}
	}

}
