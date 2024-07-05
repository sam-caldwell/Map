package Map

import (
	"testing"
)

func TestNew(t *testing.T) {
	m := New[string, int]()
	(*m)["1"] = 1
}
