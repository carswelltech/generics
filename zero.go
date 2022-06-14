package generics

// Zero just returns the zero value of its instantiated type.
func Zero[T any]() T {
	var r T
	return r
}
