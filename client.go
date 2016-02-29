package main

import (
	"fmt"
	"time"

	"gopkg.in/redis.v3"
)

func main() {
	client := redis.NewClient(&redis.Options{
		Addr:     "localhost:5112",
		Password: "",
		DB:       0,
	})

	psclient, err := client.Subscribe("mychannel")
	if err != nil {
		panic(err)
	}

	go func() {
		for i := 0; i < 10; i++ {
			client.Publish("mychannel", fmt.Sprintf("message %d", i))
			time.Sleep(350 * time.Millisecond)
		}
		time.Sleep(time.Second)
		psclient.Close()
	}()

	for {
		msg, err := psclient.ReceiveMessage()
		if err != nil {
			break
		}
		fmt.Println(msg)
	}

	fmt.Println("Closing client")
	client.Close()
}
