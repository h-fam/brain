package main

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/h-fam/brain/base/go/log"

	"cloud.google.com/go/pubsub"
)

const (
	subID   = "build-brain"
	topicID = "projects/hines-alloc/topic/brain-build"
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
	if e, err := s.Exists(ctx); err != nil || !e {
		t := c.Topic(topicID)
		if ok, err := t.Exists(ctx); err != nil || !ok {
			log.Errorf("failed to get topic %q: %v", topicID, err)
			return
		}
		subCfg := pubsub.SubscriptionConfig{
			Topic:               t,
			RetainAckedMessages: false,
			Labels: map[string]string{
				"ci": "update",
			},
		}
		s, err = c.CreateSubscription(ctx, subID, subCfg)
		if err != nil {
			log.Errorf("failed to create subscription: %v", err)
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
type Event struct {
	Email      string
	RefUpdates map[string]Update
	UpdateType string
}

type Update struct {
	RefName    string
	UpdateType string
	OldID      string
	NewID      string
}

type Message struct {
	Name           string `json:"name"`
	URL            string `json:"url"`
	EventTime      string `json:"eventTime"`
	RefUpdateEvent Event  `json:"refUpdateEvent"`
}

func handleMessage(ctx context.Context, msg *pubsub.Message) {
	if ctx.Err() != nil {
		return
	}
	m := &Message{}
	if err := json.Unmarshal(msg.Data, m); err != nil {
		log.Errorf("failed to unmarshal message %+v: %v", msg, err)
	}
	fmt.Printf("%+v\n", m)
}
