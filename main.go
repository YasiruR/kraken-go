package main

import (
	"github.com/YasiruR/feeder/feeder"
	"log"
)

func main() {
	f := feeder.NewFeeder(``)
	if err := f.AssetInfo(); err != nil {
		log.Fatalln(err)
	}

	if err := f.Subscribe(`XBT/USD`); err != nil {
		log.Fatalln(err)
	}
}
