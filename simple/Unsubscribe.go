package main

import (
	"fmt"
	"predic8.de/util"
)
import stan "github.com/nats-io/stan.go"

func main() {

	sc, _ := stan.Connect("prod", "sub-1")
	defer sc.Close()

	qs, _ := sc.QueueSubscribe("foo","team-a", func(m *stan.Msg) {
		fmt.Printf("Got: %s\n", string(m.Data))
		},
	)

	qs.Unsubscribe()

	util.Block()

}
