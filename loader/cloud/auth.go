package cloud

import (
	"context"
	"sync"

	"cloud.google.com/go/datastore"
	"google.golang.org/api/option"

	log "github.com/golang/glog"
)

var (
	jsonPath = "c:\\Users\\marcus\\Downloads\\hines-alloc-13bac0ee250f.json"
	mu       sync.Mutex
	dsClient *datastore.Client
	AuthFile string
)

func GetDSClient(ctx context.Context) *datastore.Client {
	mu.Lock()
	defer mu.Unlock()
	if dsClient == nil {
		// Create a datastore client. In a typical application, you would create
		// a single client which is reused for every datastore operation.
		var err error
		dsClient, err = datastore.NewClient(ctx, "hines-alloc", option.WithCredentialsFile(jsonPath))
		if err != nil {
			log.Errorf("failed to create datastore client: %v", err)
			return nil
		}
	}
	return dsClient
}
