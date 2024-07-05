package Map

import (
	"fmt"
	"testing"
)

func TestPurge(t *testing.T) {

	expectedSize := 100

	lhs := map[string]int{}

	for i := 0; i < expectedSize; i++ {
		key := fmt.Sprintf("%d", i)
		lhs[key] = i
	}

	Purge[string, int](&lhs)

	if sz := len(lhs); sz != 0 {
		t.Fatalf("lhs is not the expected size (sz: %d)", 0)
	}

	for i := 0; i < expectedSize; i++ {
		key := fmt.Sprintf("%d", i)
		if _, ok := lhs[key]; ok {
			t.Fatalf("key ('%s') should have been deleted", key)
		}
	}

}
