package generics

/*	Each iterates over its arguments ... i.e. it returns a chan containing the arguments.
	It's just a convenient way to write
		for e := range Each(1,2,3,4) {
*/
func Each[T any](elements ...T) chan T {
	c := make(chan T, len(elements))
	go func() {
		defer close(c)
		for _, e := range elements {
			c <- e
		}
	}()
	return c
}
