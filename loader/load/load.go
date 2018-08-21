package load

import (
	"context"
	"fmt"

	"cloud.google.com/go/datastore"
	"github.com/marcushines/brain/loader/bullet"
	"github.com/marcushines/brain/loader/cases"
	"github.com/marcushines/brain/loader/powder"
)

type Load struct {
	Caliber string
	Case    *cases.Case
	Primer  *primer.Primer
	Bullet  *bullet.Bullet
	Powder  *powder.Powder
	Charge  int64 // millgrains
	COAL    int64 // micrometers
}

func (l *Load) Key() string {
	return fmt.Sprintf("%s~%s~%s~%s~%s~%s~%s", l.Caliber, l.Case.Key(), l.Primer.Key(), l.Bullet.Key(), l.Powder.Key(), l.Charge, l.COAL)
}

func Add(ctx context.Context, c *datastore.Client, l *Load) error {
	k := datastore.NameKey("Case", l.Key(), nil)
	m := datastore.NewInsert(k, l)
	if _, err := c.Mutate(ctx, m); err != nil {
		return err
	}
	return nil
}
