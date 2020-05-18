package component

import (
	"hfam/brain/testrunner/result"
)

type Component interface {
	Start() result.Result
	Stop() result.Result
}
