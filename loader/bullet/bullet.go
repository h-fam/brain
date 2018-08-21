package bullet

import (
	"context"
	"fmt"

	"cloud.google.com/go/datastore"
	"github.com/marcushines/brain/loader/caliber"
	"github.com/marcushines/brain/loader/manufacturer"
)

type Bullet struct {
	Manufacturer string
	Name         string
	Caliber      string
	Weight       int // in milligrains
	Shape        string
}

func (b *Bullet) Key() string {
	return fmt.Sprintf("%s/%s/%s/%s/%d", b.Manufacturer, b.Name, b.Caliber, b.Shape, b.Weight)
}

func Add(ctx context.Context, c *datastore.Client, b *Bullet) error {
	manu := &manufacturer.Manufacturer{}
	if err := c.Get(ctx, datastore.NameKey("Manufacturer", b.Manufacturer, nil), manu); err != nil {
		return fmt.Errorf("manufacturer doesn't exist: %q", b.Manufacturer)
	}

	cali := &caliber.Caliber{}
	if err := c.Get(ctx, datastore.NameKey("Caliber", b.Caliber, nil), cali); err != nil {
		return fmt.Errorf("caliber doesn't exist: %q", b.Caliber)
	}
	k := datastore.NameKey("Bullet", b.Key(), nil)
	m := datastore.NewInsert(k, b)
	if _, err := c.Mutate(ctx, m); err != nil {
		return err
	}
	return nil
}
