package key

// Key is a comparable interface for use in map keys or any place for providing equality.
type Key interface {
	Equal(interface{}) bool
	String() string
}

// DuplicateError is a type error for dealing with Key collisions.
type DuplicateError string

func (d DuplicateError) Error() string {
	return string(d)
}

// Slice is a slice of keys. The KeySlice is order dependent for equality.
type Slice []Key

// Equal compares equal element in the slice for equality.
func (k Slice) Equal(v interface{}) bool {
	t, ok := v.(Slice)
	if !ok {
		return false
	}
	if len(k) != len(t) {
		return false
	}
	for i := range t {
		if !k[i].Equal(k[i]) {
			return false
		}
	}
	return true
}
