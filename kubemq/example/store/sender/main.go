package main

import (
	"context"
	"fmt"
	"log"
	"time"

	kubemq "github.com/kubemq-io/kubemq-go"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	client, err := kubemq.NewClient(ctx,
		kubemq.WithAddress("192.168.1.222", 50000),
		kubemq.WithClientId("test-command-client-id"),
		kubemq.WithTransportType(kubemq.TransportTypeGRPC))
	if err != nil {
		log.Fatal(err)
	}
	defer client.Close()
	channel := "hello-command"
	for i := 0; i < 1000; i++ {
		response, err := client.ES().
			SetId(fmt.Sprintf("some-command-id-%v", time.Now())).
			SetChannel(channel).
			SetMetadata("some-metadata").
			SetBody([]byte("hello kubemq - sending a command, please reply")).
			AddTag("seq", fmt.Sprintf("%d", i)).
			Send(ctx)
		if err != nil {
			log.Fatal(err)
		}
		log.Printf("Response Received:\nCommandID: %s\nExecutedAt:%v\n", response.Id, response.Sent)
	}

}
