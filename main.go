package main

import (
	"fmt"
	"log"

	"github.com/AHMEDxHAGAG/RSSaggregator/internal/config"
)

func errPanic(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	conf, err := config.Read()
	errPanic(err)
	err = conf.SetUser("Ahmed Hagag")
	errPanic(err)
	conf, err = config.Read()
	errPanic(err)
	fmt.Println(conf.CurrentUserName)
	fmt.Println(conf.DBURL)
}
