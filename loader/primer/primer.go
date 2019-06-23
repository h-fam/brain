package primer

import (
	"context"
	"fmt"

	"cloud.google.com/go/datastore"
	"source.cloud.google.com/hines-alloc/brain/loader/manufacturer"
)

type Primer struct {
	Manufacturer string
	Name         string
	Size         string
	Type         string
	URL          string
}

func (p *Primer) Key() string {
	return fmt.Sprintf("%s/%s", p.Manufacturer, p.Name)
}

func Add(ctx context.Context, c *datastore.Client, p *Primer) error {
	manu := &manufacturer.Manufacturer{}
	if err := c.Get(ctx, datastore.NameKey("Manufacturer", p.Manufacturer, nil), manu); err != nil {
		return fmt.Errorf("manufacturer doesn't exist: %q", p.Manufacturer)
	}
	k := datastore.NameKey("Primer", p.Key(), nil)
	m := datastore.NewInsert(k, p)
	if _, err := c.Mutate(ctx, m); err != nil {
		return err
	}
	return nil
}
