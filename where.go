package glinq

func (e Enumerable[T]) Where(predicate func(T) bool) Enumerable[T] {
	return Enumerable[T]{
		eval: func() []T {
			var result []T
			res := e.eval()
			for i := 0; i < len(res); i++ {
				if predicate(res[i]) {
					result = append(result, res[i])
				}
			}
			return result
		},
	}
}