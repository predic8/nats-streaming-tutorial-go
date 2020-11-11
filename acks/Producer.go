package main

import (
	"strconv"
	"time"

	stan "github.com/nats-io/stan.go"
)

func main() {

	sc, _ := stan.Connect("prod", "simple-pub")
	defer sc.Close()

	for i := 0; ; i++ {
		sc.Publish("foo", []byte("Botschaft "+strconv.Itoa(i)))
		time.Sleep(time.Second * 2)
	}

}
