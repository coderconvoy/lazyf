package lazyf

import (
	"testing"

	"github.com/pkg/errors"
)

func compareBlockSlice(a, b []VarBlock) error {
	if len(a) != len(b) {
		return errors.Errorf("lengths not equal: %d,%d", len(a), len(b))
	}

	for k, v := range a {
		if v != b[k] {
			return errors.Errorf("element %d not equal: %s,%s", k, v, b[k])
		}
	}
	return nil
}

func basicReplacer(s string) string {
	switch s {
	case "home":
		return "/root/"
	case "a":
		return "AAA"

	}
	return "--nf--"
}

func Test_VScan(t *testing.T) {
	td := []struct {
		s    string
		o, c string
		r    []VarBlock
	}{
		{"hello {name}, good to meet you", "{", "}",
			[]VarBlock{{"hello ", false}, {"name", true}, {", good to meet you", false}}},
		{"a{$b}{c}{$d}}", "{$", "}",
			[]VarBlock{{"a", false}, {"b", true}, {"{c}", false}, {"d", true}, {"}", false}}},
	}

	for k, v := range td {
		r := VarScan(v.s, v.o, v.c)
		err := compareBlockSlice(v.r, r)
		if err != nil {
			t.Error(k, err)
		}

	}
}

func Test_VReplace(t *testing.T) {
	td := []struct {
		in, out string
	}{
		{"hello{home}is a good place", "hello/root/is a good place"},
		{"h{{hom}is", "h--nf--is"},
	}
	for k, v := range td {
		r := FVarReplace(v.in, "{", "}", basicReplacer)
		if r != v.out {
			t.Errorf("Line %d not equal:%s:%s", k, v.out, r)
		}
	}
}
