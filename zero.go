package generics

/*
	Zero just provides a concise way to return the zero value of the instantiated value of a type parameter.
	Mostly useful when returning an error from a generic function or method.
	It doesn't do any magic, but is just convenient shorthand for declaring a variable of the parametric type, and returning that variable.
*/
func Zero[T any]() T {
	var r T
	return r
}
