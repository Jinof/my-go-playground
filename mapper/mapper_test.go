package mapper

import (
	"testing"
)

func TestMapper(t *testing.T) {
	m := Mapper{}
	if m.Check(1, 1) {
		t.Error("bad check, n1, n2 should not in Mapper now")
	}
	if err := m.Add(1, 1, 1); err != nil {
		t.Error(err)
	}
	if !m.Check(1, 1) {
		t.Error("bad check, n1, n2 should in Mapper now")
	}
	t.Log(m)
}
