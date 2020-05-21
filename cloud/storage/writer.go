package main

import (
	"context"

	"cloud.google.com/go/storage"
	"github.com/h-fam/brain/base/go/log"
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
