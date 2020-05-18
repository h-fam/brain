package main

import (
	"context"

	"cloud.google.com/go/storage"
	log "github.com/golang/glog"
)

func main() {
	ctx := context.Background()
	client, err := storage.NewClient(ctx)
	if err != nil {
		log.Errorf("failed to create client: %v", err)
		return
	}
	bh := client.Bucket("mine")
	bh.Create(ctx, "foo", nil)

}
