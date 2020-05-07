package keyable

type Key interface {
	Equal(interface{}) bool
}
