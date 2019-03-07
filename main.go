package main

import (
	"fmt"
	"log"
	"time"

	"github.com/sf1/go-card/smartcard"
)

func main() {
	log.Println("card server v0.2")
	ctx, _ := smartcard.EstablishContext()

	defer ctx.Release()

	// handle error, if any
	fmt.Println("\nWaiting for card...")
	var reader *smartcard.Reader

	for {
		readers, _ := ctx.ListReaders()

		reader = readers[0]
		// handle error, if any
		if reader.IsCardPresent() {
			card, _ := reader.Connect()
			// handle error, if any
			log.Println(card.ATR().String())

			card.Disconnect()
		}

		for reader.IsCardPresent() {

			time.Sleep(100 * time.Microsecond)

		}
		for !reader.IsCardPresent() {
			time.Sleep(100 * time.Microsecond)

		}
	}

}
