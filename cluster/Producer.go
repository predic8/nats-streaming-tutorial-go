package main

import (
	"cluster/natsutil"
	"fmt"
	"github.com/nats-io/stan.go"
	"log"
	"strconv"
	"time"
)

func main() {

	connected := false
	i := 1
	var sc stan.Conn = nil

	for {

		for {

			time.Sleep(time.Second * 2)

			if !connected {

				fmt.Println("Baue Verbindung auf!")

				nc, err := natsutil.OpenNATSConnection("nats://127.0.0.1:4221,nats://127.0.0.1:4222,nats://127.0.0.1:4223")

				sc, err = stan.Connect("prod", "pub1", stan.NatsConn(nc), stan.SetConnectionLostHandler(func(conn stan.Conn, err error) {
					fmt.Println("Verbindung verloren")
					connected = false
				}))
				if err != nil {
					log.Println(err)
					continue
				}

				connected = true
			}

			fmt.Printf("Sende Nachricht %d\n", i)
			err := sc.Publish("foo", []byte("Botschaft "+strconv.Itoa(i)))
			if err != nil {
				log.Println("Fehler beim Senden: ", err)
				sc.Close()
				connected = false
			} else {
				break
			}
		}
		i++

	}

}


