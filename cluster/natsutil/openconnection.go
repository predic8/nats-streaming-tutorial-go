package natsutil

import (
	"fmt"
	"github.com/nats-io/nats.go"
	"log"
)

func OpenNATSConnection(url string) (*nats.Conn, error) {
	nc, err := nats.Connect(url,
		nats.DisconnectErrHandler(func(conn *nats.Conn, err error) {
			fmt.Printf("Disconnect\n")
		}),
		nats.ReconnectHandler(func(conn *nats.Conn) {
			fmt.Printf("Reconnect\n")
			fmt.Printf("Server Liste: %s\n", conn.Opts.Servers)
			fmt.Printf("Verbunden mit %s\n\n", conn.ConnectedUrl())

		}),
		nats.ErrorHandler(func(conn *nats.Conn, subscription *nats.Subscription, err error) {
			fmt.Printf("Error %s!", err)
		}),
		nats.ClosedHandler(func(conn *nats.Conn) {
			fmt.Println("Closed!")
		}),
		nats.DiscoveredServersHandler(func(conn *nats.Conn) {
			fmt.Println("Server entdeckt!")
			fmt.Printf("Servers: %s\n", conn.Opts.Servers)
		}),
	)
	if err != nil {
		log.Println(err)
	}

	return nc, err
}