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

func handleMessage(ctx context.Context, msg *pubsub.Message) {
	if ctx.Err() != nil {
		return
	}
	fmt.Println(msg)
}
