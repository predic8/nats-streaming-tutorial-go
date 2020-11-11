package main

import (
	"fmt"

	"predic8.de/util"

	stan "github.com/nats-io/stan.go"
)

func main() {

	sc, _ := stan.Connect("prod", "1")
	defer sc.Close()

	sc.Subscribe("foo", func(m *stan.Msg) {
		fmt.Printf("Speichere: %s\n", string(m.Data))
	})

	util.Block()
}
