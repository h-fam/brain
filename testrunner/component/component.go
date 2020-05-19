package component

import (
	"context"
	"hfam/brain/testrunner/result"
)

type ComponentGetter interface {
	Get(n string) (Component, error)
}

type Component interface {
	Start(context.Context) result.Result
	Stop(context.Context) result.Result
}
