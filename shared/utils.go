package shared

// Set lets you store unique values.
type Set map[int]struct{}

// Put inserts a value into the set.
func (s Set) Put(v int) {
	s[v] = struct{}{}
}

// Has returns whether a value exists in a set.
func (s Set) Has(v int) bool {
	_, ok := s[v]
	return ok
}

// Delete removes a value from the set.
func (s Set) Delete(v int) {
	delete(s, v)
}
