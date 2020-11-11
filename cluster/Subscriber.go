package main

import (
	"cluster/natsutil"
	"fmt"
	"github.com/nats-io/stan.go"
	"log"
	"time"
)

func main() {

	connected := false
	var sc stan.Conn = nil

	for {

		if !connected {

			fmt.Println("Baue Verbindung auf!")

			nc, err := natsutil.OpenNATSConnection("nats://127.0.0.1:4221,nats://127.0.0.1:4222,nats://127.0.0.1:4223")

			sc, err = stan.Connect("prod", "1", stan.NatsConn(nc), stan.SetConnectionLostHandler(func(conn stan.Conn, err error) {
				fmt.Println("SetConnectionLostHandler")
				connected = false
			}))
			if err != nil {
				log.Print("%s", err)
				continue
			}

			_, err = sc.Subscribe("foo", func(m *stan.Msg) {
				fmt.Printf("Got: %s\n", string(m.Data))
			},
				stan.DurableName("my-durable"),
			)
			if err != nil {
				fmt.Printf("Subscribe Fehler: %s\n", err)
				continue
			}

			connected = true
		}
		time.Sleep(time.Second)
	}

}
