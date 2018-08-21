package manufacturer

import (
	"context"

	"cloud.google.com/go/datastore"
)

type Manufacturer struct {
	Name string
	URL  string
}

func (m *Manufacturer) Key() string {
	return m.Name
}

func Add(ctx context.Context, c *datastore.Client, manu *Manufacturer) error {
	k := datastore.NameKey("Manufacturer", manu.Key(), nil)
	m := datastore.NewInsert(k, manu)
	if _, err := c.Mutate(ctx, m); err != nil {
		return err
	}
	return nil
}
