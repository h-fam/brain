package caliber

import (
	"context"

	"cloud.google.com/go/datastore"
)

type Caliber struct {
	Name     string
	Diameter int64 // micrometers
	URL      string
}

func (c *Caliber) Key() string {
	return c.Name
}

func Add(ctx context.Context, c *datastore.Client, caliber *Caliber) error {
	k := datastore.NameKey("Caliber", caliber.Key(), nil)
	m := datastore.NewInsert(k, caliber)
	if _, err := c.Mutate(ctx, m); err != nil {
		return err
	}
	return nil
}
