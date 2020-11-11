package main

import (
	"fmt"
	"math/rand"
	"predic8.de/util"
)
import stan "github.com/nats-io/stan.go"

func main() {

	sc, err := stan.Connect("demo", "1")
	if err != nil {
		fmt.Printf("%s", err)
	}
	defer sc.Close()

	_, err = sc.Subscribe("foo", func(m *stan.Msg) {
		fmt.Printf("%s Redelivered: %t # Redeliveries: %d\n", string(m.Data), m.Redelivered, m.RedeliveryCount)
		if rand.Intn(3) == 2 {
			fmt.Println("Shit happened!")
			return
		}
		m.Ack()
	}, stan.SetManualAckMode())

	if err != nil {
		fmt.Printf("%s", err)
	}

	util.Block()

}
