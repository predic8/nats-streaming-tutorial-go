package main

import (
	"strconv"
	"time"

	stan "github.com/nats-io/stan.go"
)

func main() {

	sc, _ := stan.Connect("prod", "simple-pub")
	defer sc.Close()

	for i := 1; ; i++ {
		sc.Publish("bestellungen", []byte("Bestellung "+strconv.Itoa(i)))
		time.Sleep(2 * time.Second)
	}

}
