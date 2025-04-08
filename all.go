package glinq

func (e Enumerable[T]) All(predicate func(T) bool) bool {
	for _, v := range e.eval() {
		if !predicate(v) {
			return false
		}
	}
	return true
}