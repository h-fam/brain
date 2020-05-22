package user

import (
	"context"

	"cloud.google.com/go/datastore"
)

type User struct {
	Email string
}

func Add(ctx context.Context, c *datastore.Client, u *User) error {
	k := datastore.NameKey("User", u.Email, nil)
	m := datastore.NewInsert(k, u)
	if _, err := c.Mutate(ctx, m); err != nil {
		return err
	}
	return nil
}
