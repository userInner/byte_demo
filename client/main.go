package main

import (
	"context"
	"example/kitex_gen/api"
	"example/kitex_gen/api/echo"
	"github.com/cloudwego/kitex/client"
	"github.com/cloudwego/kitex/client/callopt"
	"log"
	"time"
)

func main() {
	c, err := echo.NewClient("example", client.WithHostPorts("0.0.0.0:8888"))
	if err != nil {
		log.Fatal(err)
	}
	req := &api.Request{
		Message: "Hello World",
		A:       1314,
		B:       520,
	}
	resp, err := c.Echo(
		context.Background(),
		req,
		callopt.WithRPCTimeout(3*time.Second))
	if err != nil {
		log.Fatal(err)
	}
	log.Println(resp)
}
