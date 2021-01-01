package key

// String is a keyable string.
type String string

// String returns the string representation of String key.
func (s String) String() string {
	return string(s)
}

// Equal compares s to dst for equality.
func (s String) Equal(dst interface{}) bool {
	sd, ok := dst.(string)
	if !ok {
		return false
	}
	if string(s) == sd {
		return true
	}
	return false
}

// NewString returns a new String key.
func NewString(s string) String {
	return String(s)
}
