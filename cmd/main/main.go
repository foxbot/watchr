package main

import (
	"flag"
	"log"
	"os"
	"os/signal"

	"github.com/foxbot/watchr"
)

var (
	hostFlag = flag.String("host", "localhost:3000", "-host=localhost:0")
)

func main() {
	println("watchr")
	println("\tfoxbot.me")
	println("")

	w := watchr.NewWatchr(*hostFlag)
	w.Run()
	errors := w.Errors()

	sc := make(chan os.Signal)
	signal.Notify(sc, os.Interrupt, os.Kill)

	for {
		select {
		case e := <-errors:
			log.Println(e)
		case <-sc:
			goto die
		}
	}
die:
	println("bye")
}
