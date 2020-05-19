package main

import (
	"context"
	"fmt"
	"hfam/brain/base/go/log"

	"cloud.google.com/go/pubsub"
)

const (
	subID = "build-brain"
)

func main() {
	ctx := context.Background()
	c, err := pubsub.NewClient(ctx, "hines-alloc")
	if err != nil {
		log.Errorf("failed to create client: %v", err)
		return
	}
	subIter := c.Subscriptions(ctx)
	for {
		s, err := subIter.Next()
		if err != nil {
			log.Errorf("iterator stop: %v", err)
			break
		}
		fmt.Println(s.String())
	}
	s := c.Subscription(subID)
	e, err := s.Exists(ctx)
	if err != nil {
		log.Errorf("Failed to check if subscription exists: %v", err)
		return
	}
	if !e {
		var err error
		s, err = c.CreateSubscription(ctx, subID, pubsub.SubscriptionConfig{})
		if err != nil {
			log.Errorf("Failed to create subscription: %v", err)
			return
		}
	}
	for {
		if err := s.Receive(ctx, handleMessage); err != nil {
			log.Errorf("Failed to recieve message: %v", err)
			break
		}
	}
}

/*
{
  "name": "projects/hines-alloc/repos/brain",
  "url": "https://source.developers.google.com/p/hines-alloc/r/brain",
  "eventTime": "2020-05-18T22:50:48.531823Z",
  "refUpdateEvent": {
    "email": "marcus.hines@gmail.com",
    "refUpdates": {
      "refs/heads/master": {
        "refName": "refs/heads/master",
        "updateType": "UPDATE_FAST_FORWARD",
        "oldId": "a27ea64b41dbf5ee515227a7ae3c1151fe6336e0",
        "newId": "b8763d451e09907d974d9f9ff983a2d8da08565c"
      }
    }
  }
}
*/
func handleMessage(ctx context.Context, msg *pubsub.Message) {
	if ctx.Err() != nil {
		return
	}
	fmt.Println(string(msg.Data))
}
