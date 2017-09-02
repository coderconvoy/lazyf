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

func TestFile(t *testing.T) {
	dt, err := ReadFile("test_data/t1.lz")
	if err != nil {
		t.Error(err)
	}
	if len(dt) != 4 {
		t.Errorf("Expected 4 chars, got: %d", len(dt))
	}

	w, ok := ByName(dt, "warrior")
	if !ok {
		t.Errorf("Expected warrior to exist")
	}
	at, ok := w.PString("at", "At")
	t.Logf(at)

}
