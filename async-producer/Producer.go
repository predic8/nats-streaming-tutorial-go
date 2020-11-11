package main

import (
	"fmt"
	"strconv"
	"time"

	stan "github.com/nats-io/stan.go"
)

func main() {

	sc, _ := stan.Connect("prod", "simple-pub")

	t1 := time.Now().UnixNano()


	for i := 1; i <= 1000; i++ {
		sc.Publish("besellungen", []byte("Bestellung "+strconv.Itoa(i)))
	}

	t2 := time.Now().UnixNano()


	sc.Close()

	fmt.Printf("%f Nachrichten/Sekunde", 1000.0/(float64(t2-t1)/1e9))

}
