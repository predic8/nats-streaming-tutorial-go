package main

import (
	"fmt"
	"predic8.de/util"
)
import stan "github.com/nats-io/stan.go"

func main() {

	sc, _ := stan.Connect("prod", "sub-3")
	defer sc.Close()

	sc.Subscribe("foo4",  func(m *stan.Msg) {
		fmt.Printf("Got: %s\n", string(m.Data))
		},
		stan.DeliverAllAvailable(),
	)

	util.Block()

}
