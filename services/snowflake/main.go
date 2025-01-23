package main

import (
	"log"

	"github.com/vhall1/shorturl/services/snowflake/service"
)

func main() {
	// TODO: dinamically generate machine IDs
	snowflake, err := service.NewSnowflake(int64(0))
	if err != nil {
		log.Fatal(err)
	}

	ch := make(chan error)

	go func() {
		s := NewHttpServer(snowflake)
		ch <- s.Start()
	}()

	go func() {
		s := NewGrpcServer(snowflake)
		ch <- s.Start()
	}()

	err = <-ch
	if err != nil {
		log.Fatal(err)
	}
}
