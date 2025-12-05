package main

import (
	"exc8/client"
	"exc8/server"
	"fmt"
	"time"
)

func main() {
	go func() {
		if err := server.StartGrpcServer(); err != nil {
			panic(err)
		}
	}()
	time.Sleep(1 * time.Second)
	c, err := client.NewGrpcClient()
	if err != nil {
		fmt.Println("Error creating client:", err)
		return
	}
	if err := c.Run(); err != nil {
		fmt.Println("Error running client:", err)
	}
	println("Orders complete!")
}