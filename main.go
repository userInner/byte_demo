package main

import (
	api "example/kitex_gen/api/echo"
	"log"
)

func main() {
	svr := api.NewServer(new(EchoImpl))
	err := svr.Run()

	if err != nil {
		log.Println(err.Error())
	}
}
