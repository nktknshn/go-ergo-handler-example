package set

type Set[T comparable] struct {
	items map[T]struct{}
}

func NewFromSlice[T comparable](items []T) *Set[T] {
	set := New[T]()
	for _, item := range items {
		set.Add(item)
	}
	return set
}

func New[T comparable](items ...T) *Set[T] {
	set := &Set[T]{items: make(map[T]struct{})}
	for _, item := range items {
		set.Add(item)
	}
	return set
}

func (s *Set[T]) Add(items ...T) {
	for _, item := range items {
		s.items[item] = struct{}{}
	}
}

func (s *Set[T]) Has(item T) bool {
	_, ok := s.items[item]
	return ok
}

func (s *Set[T]) Size() int {
	return len(s.items)
}

func (s *Set[T]) Remove(item T) {
	delete(s.items, item)
}

func (s *Set[T]) Clear() {
	s.items = make(map[T]struct{})
}

func (s *Set[T]) Values() []T {
	values := make([]T, 0, len(s.items))
	for item := range s.items {
		values = append(values, item)
	}
	return values
}

func (s *Set[T]) IsEmpty() bool {
	return len(s.items) == 0
}

func (s *Set[T]) Union(other *Set[T]) *Set[T] {
	result := New[T]()
	for item := range s.items {
		result.Add(item)
	}

	for item := range other.items {
		result.Add(item)
	}

	return result
}

func (s *Set[T]) Intersection(other *Set[T]) *Set[T] {
	result := New[T]()
	for item := range s.items {
		if other.Has(item) {
			result.Add(item)
		}
	}
	return result
}

// Difference возвращает множество, содержащее элементы, которые есть в текущем множестве, но нет в другом множестве
func (s *Set[T]) Difference(other *Set[T]) *Set[T] {
	result := New[T]()
	for item := range s.items {
		if !other.Has(item) {
			result.Add(item)
		}
	}
	return result
}
