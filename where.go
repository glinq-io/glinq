package glinq

func (e Enumerable[T]) Where(predicate func(T) bool) Enumerable[T] {
	return Enumerable[T]{
		eval: func() []T {
			var result []T
			for _, v := range e.eval() {
				if predicate(v) {
					result = append(result, v)
				}
			}
			return result
		},
	}
}