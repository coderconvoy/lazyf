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
		t.FailNow()
	}
	at, err := w.PString("at", "At")
	if err != nil {
		t.Errorf("Expected Warrior to have attack")
	}
	if at != "2h+6" {
		t.Errorf("Expected Warrior to have 2h+6 : got %d", at)
	}

	_, err = w.PString("hat")
	if err == nil {
		t.Errorf("Expected Warrior not to have hat")
	}

	h := w.PIntD(5, "H")
	if h != 3 {
		t.Errorf("Expected Warrior Health to be 3, got %d", h)
	}

	h = w.PIntD(5, "Hat")
	if h != 5 {
		t.Errorf("Expected Warrior Had to be 5, got %d", h)
	}

	h = w.PIntD(5, "At")
	if h != 5 {
		t.Errorf("Attack should not be parseable int")
	}
}
