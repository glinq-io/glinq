package glinq

type Enumerable[T any] struct {
	eval func() []T
}

func (e Enumerable[T]) ToSlice() []T {
	return e.eval()
}

func NewEnumerable[T any](inputSlice []T) Enumerable[T] {
	return Enumerable[T]{
		eval: func() []T { return inputSlice },
	}
}