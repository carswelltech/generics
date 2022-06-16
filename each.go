package generics

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
