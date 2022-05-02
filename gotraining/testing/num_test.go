package num

import (
	"testing"
)

func BenchmarkSample(b *testing.B) {
	for i := 0; i < b.N; i++ {
		add(2, 3)
	}
}
func Test_add(t *testing.T) {
	tc := []struct {
		i, j, exp int
	}{
		{3, 1, 4},
		{9, 1, 10},
		{3, 8, 11},
		{3, 14, 17},
		{3, 14, 17},
		{3, 14, 17},
	}
	// inp := 3
	// exp := 6
	// act := add(inp)

	// if act != exp {
	// 	t.Errorf("Test failed")
	// }
	for _, te := range tc {
		if add(te.i, te.j) != te.exp {
			t.Errorf(("Test failed"))
		}
	}
}
