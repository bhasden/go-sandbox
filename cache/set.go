package cache

type Set[E comparable] map[E]struct{}

func (set Set[E]) Add(elements ...E) {
	for _, element := range elements {
		set[element] = struct{}{}
	}
}

func (set Set[E]) Union(anotherSet Set[E]) Set[E] {
	unionSet := New[E]()
	unionSet.Add(set.Members()...)
	unionSet.Add(anotherSet.Members()...)
	return unionSet
}

func (set Set[E]) Members() []E {
	members := make([]E, 0, len(set))
	for key := range set {
		members = append(members, key)
	}
	return members
}

func New[E comparable]() Set[E] {
	return Set[E]{}
}
