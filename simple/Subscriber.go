package main

import (
	"fmt"

	"predic8.de/util"

	stan "github.com/nats-io/stan.go"
)

func main() {

	sc, _ := stan.Connect("prod", "sub-1")
	defer sc.Close()

	sc.Subscribe("bestellungen", func(m *stan.Msg) {
		fmt.Printf("Got: %s\n", string(m.Data))
	})

	util.Block()

}
