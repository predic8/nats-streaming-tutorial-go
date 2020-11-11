package main

import (
	"predic8.de/util"
)
import stan "github.com/nats-io/stan.go"

func main() {

	sc, _ := stan.Connect("prod", "sub-1")
	defer sc.Close()

	qs, _ := sc.QueueSubscribe("foo","team-a", func(m *stan.Msg) {})

	qs.Unsubscribe()

	util.Block()

}
