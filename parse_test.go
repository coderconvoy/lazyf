package lazyf

import "testing"

func Test_Parse(t *testing.T) {
	ts := []struct {
		s    string
		elen int
	}{
		{"go:", 1},
		{"gop", 0},
		{"pop:44:2", 2},
	}

	for k, v := range ts {
		lz := ParseLZ(v.s)
		if len(lz.Deets) != v.elen {
			t.Errorf("with %d, expected len %d, got %d", k, v.elen, len(lz.Deets))
		}

	}

}
