package stack

type Stack(type T) []T

func (s Stack(T)) Pop() T {
	return s[len(s)-1]
}

func(s *Stack(T)) Push(value T) {
	*s = append(*s, value)
}

