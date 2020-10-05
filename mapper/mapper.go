package mapper

import "errors"

type Mapper map[int]map[int]int

func (m *Mapper) Add(n1, n2, n3 int) error {
	if m.Check(n1, n2) {
		return errors.New("value exists")
	}
	(*m)[n1] = map[int]int{n2: n3}
	return nil
}

func (m *Mapper) Check(n1, n2 int) bool {
	_, ok := (*m)[n1]
	if !ok {
		return false
	}
	_, ok = (*m)[n1][n2]
	return ok
}
