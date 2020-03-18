package uniq

import (
	"testing"
)

func TestUniq(t *testing.T) {

	has := make(map[string]bool)

	for i := 0; i < 100000; i++ {
		v := New()
		if has[v] {
			t.Fatal("uniq failed")
		}
		has[v] = true
	}
}
