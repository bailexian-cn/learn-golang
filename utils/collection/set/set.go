package set

type empty struct{}
type t interface{}
type Set map[t]empty

func (s Set) has(item t) bool {
	_, exists := s[item]
	return exists
}

func (s Set) insert(item t) {
	s[item] = empty{}
}

func (s Set) delete(item t) {
	delete(s, item)
}

func (s Set) len() int {
	return len(s)
}
