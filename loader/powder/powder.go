package powder

import (
	"context"
	"fmt"

	"github.com/h-fam/brain/loader/manufacturer"

	"cloud.google.com/go/datastore"
)

type Powder struct {
	Manufacturer string
	Name         string
	URL          string
}

func (p *Powder) Key() string {
	return fmt.Sprintf("%s/%s", p.Manufacturer, p.Name)
}

func Add(ctx context.Context, c *datastore.Client, p *Powder) error {
	manu := &manufacturer.Manufacturer{}
	if err := c.Get(ctx, datastore.NameKey("Manufacturer", p.Manufacturer, nil), manu); err != nil {
		return fmt.Errorf("manufacturer doesn't exist: %q", p.Manufacturer)
	}
	k := datastore.NameKey("Powder", p.Key(), nil)
	m := datastore.NewInsert(k, p)
	if _, err := c.Mutate(ctx, m); err != nil {
		return err
	}
	return nil
}
