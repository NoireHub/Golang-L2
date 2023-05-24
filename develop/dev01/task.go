package main

import (
	"fmt"
	"log"
	"os"

	"github.com/beevik/ntp"
)

func CurrentTime() {
	time, err := ntp.Time("0.beevik-ntp.pool.ntp.org")
	if err != nil {
		fmt.Fprint(os.Stderr, err)
		log.Fatal()
	}

	log.Print(time)
}

func main() {
	CurrentTime()
}
