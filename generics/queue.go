package generics

type Queue[T any] struct {
	values []T
}

func (s Queue[T]) Push(value T) {
	s.values = append(s.values, value)
}

func (s Queue[T]) IsEmpty() bool {
	return len(s.values) == 0
}

func (s Queue[T]) Pop() (T, bool) {
	if s.IsEmpty() {
		var zero T
		return zero, false
	}

	el := s.values[0]
	s.values = s.values[1:len(s.values)]
	return el, true
}
