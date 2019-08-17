package cases

import (
	"context"
	"fmt"

	"cloud.google.com/go/datastore"
	"hines-alloc/brain/loader/caliber"
	"hines-alloc/brain/loader/manufacturer"
)

type Case struct {
	Manufacturer string
	Caliber      string
	Primer       string
	Volume       int64 // millgrains
}

func (c *Case) Key() string {
	return fmt.Sprintf("%s/%s/%s", c.Manufacturer, c.Caliber, c.Primer)
}

func Add(ctx context.Context, c *datastore.Client, cs *Case) error {
	manu := &manufacturer.Manufacturer{}
	if err := c.Get(ctx, datastore.NameKey("Manufacturer", cs.Manufacturer, nil), manu); err != nil {
		return fmt.Errorf("manufacturer doesn't exist: %q", cs.Manufacturer)
	}

	cali := &caliber.Caliber{}
	if err := c.Get(ctx, datastore.NameKey("Caliber", cs.Caliber, nil), cali); err != nil {
		return fmt.Errorf("caliber doesn't exist: %q", cs.Caliber)
	}
	k := datastore.NameKey("Case", cs.Key(), nil)
	m := datastore.NewInsert(k, cs)
	if _, err := c.Mutate(ctx, m); err != nil {
		return err
	}
	return nil
}
