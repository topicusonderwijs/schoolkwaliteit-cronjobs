package utils

import (
	"log/slog"
)

// DeferLogger creates a deferred error logging function.
//
// DeferLogger takes a function `f` that returns an error, and it returns a new function that, when deferred, will execute `f`, check for an error, and log it if an error is present.
//
// Example:
//
//	defer DeferLogger(func() error {
//	    // Your code that might return an error
//	    return someOperation()
//	})()
//
// This function is useful for deferring error logging and is commonly used with defer statements in Go.
func DeferLogger(f func() error) func() {
	return func() {
		err := f()
		if err != nil {
			slog.Error("Error in defer function", ErrAttr(err))
		}
	}
}

func ErrAttr(err error) slog.Attr {
	return slog.Any("error", err)
}
