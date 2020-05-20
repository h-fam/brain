package main

import (
   "context"
   "github.com/kubemq-io/kubemq-go"
   "log"
   "time"
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
   response, err := client.C().
      SetId("some-command-id").
      SetChannel(channel).
      SetMetadata("some-metadata").
      SetBody([]byte("hello kubemq - sending a command, please reply")).
      SetTimeout(10 *time.Second).
      Send(ctx)
   if err != nil {
      log.Fatal(err)
   }
   log.Printf("Response Received:\nCommandID: %s\nExecutedAt:%s\n", response.CommandId, response.ExecutedAt)
}
