package errs

import (
	"reflect"
	"strings"
)

// Separator is used to separate error messages when calling Error on a list.
// Only package main should set Separator.  It should only be set in an init
// function defined in the main package.
var Separator = ", "

// An Error is a list of errors and implements the error interface.  An Error
// should never be declared directly, use a List and then it's Err
// method to return a proper error.
type Error struct {
	List
}

// List is the working representation of an Error, it does not implement
// the error interface.  Use the Add method to add errors to a List.
//
// Separator may optionally be set as the string to separate errors when
// displayed.  If not set, it defaults to the global Separator value.
type List struct {
	Separator string
	errors    []error
}

// Errors is implemented by error types that can return lists of errors.
type Errors interface {
	// Errors returns the list of errors associated with the recevier.  It
	// returns nil if there are no errors associated with the recevier.
	Errors() []error
}

// etype is used for detecting types that are simply slices of error.
var etype = reflect.TypeOf([]error{})

// Add adds all non-nil errs to the list of errors in e and returns true if errs
// contains a non-nil error.  If no non-nil errors are passed Add does nothing
// and returns false.  Add will never add a nil error to the List.  If err
// implementes the Errors interface or its underlying type is a slice of errors
// then e.Add is called on each individual error.
func (e *List) Add(errs ...error) bool {
	added := false
	for _, err := range errs {
		if err != nil {
			if el, ok := err.(Errors); ok {
				errs := el.Errors()
				if len(errs) > 0 {
					e.errors = append(e.errors, errs...)
					added = true
				}
				continue
			}
			if rv := reflect.ValueOf(err); rv.Type().AssignableTo(etype) {
				a := false
				n := rv.Len()
				for i := 0; i < n; i++ {
					errI, ok := rv.Index(i).Interface().(error)
					if ok {
						a = e.Add(errI) || a
					}
				}
				added = added || a
				continue
			}
			e.errors = append(e.errors, err)
			added = true
		}
	}
	return added
}

// Err returns e as an error of type Error if e has errors, or nil.
func (e List) Err() error {
	if len(e.errors) == 0 {
		return nil
	}
	return Error{e}
}

// Error implements the error interface.
func (e Error) Error() string {
	sep := e.Separator
	if sep == "" {
		sep = Separator
	}
	msgs := make([]string, len(e.errors))
	for x, err := range e.errors {
		msgs[x] = err.Error()
	}
	return strings.Join(msgs, sep)
}

// Errors returns the list of errors in e.
func (e Error) Errors() []error {
	return e.errors
}
