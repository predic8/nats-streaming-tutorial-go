package main

import (
	"fmt"
	"strconv"
)
import stan "github.com/nats-io/stan.go"

func main() {

	fmt.Print("producer")

	sc, _ := stan.Connect("prod", "simple-pub")

	for i := 1;i <= 10; i++ {
		sc.Publish("bestellungen", []byte("Nr. " + strconv.Itoa(i)))
	}

	sc.Close()

}
