package main

import (
	"context"
	"log"
	"time"

	"github.com/kubemq-io/kubemq-go"
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
	errCh := make(chan error)
	commandsCh, err := client.SubscribeToCommands(ctx, channel, "", errCh)
	if err != nil {
		log.Fatal(err)
	}
	for {
		log.Println("omg")
		select {
		case err := <-errCh:
			log.Fatal(err)
			return
		case command, more := <-commandsCh:
			if !more {
				log.Println("Command Received , done")
				return
			}
			log.Printf("Command Received:\nId %s\nChannel: %s\nMetadata: %s\nBody: %s\n", command.Id, command.Channel, command.Metadata, command.Body)
			err := client.R().
				SetRequestId(command.Id).
				SetResponseTo(command.ResponseTo).
				SetExecutedAt(time.Now()).
				Send(ctx)
			if err != nil {
				log.Fatal(err)
			}
		case <-ctx.Done():
			return
		}
	}

}
