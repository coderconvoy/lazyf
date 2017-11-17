package lazyf

import (
	"flag"
	"testing"
)

func pmatch(t *testing.T, exp, got, mess string) {
	if exp != got {
		t.Errorf(mess+": exp '%s', got '%s'", exp, got)
	}

}

func qerr(t *testing.T, err error, msg string, fargs ...string) {
	if err != nil {
		t.Errorf(msg+":%s", append(fargs, err.Error()))
	}
}

func Test_Flags(t *testing.T) {
	fg := flagger{
		args:  []string{"-g", "hello", "-s", "small", "-d", "MEGABYTES"},
		flist: make(map[string]*string),
		fset:  flag.NewFlagSet("test", flag.ContinueOnError),
	}
	gpp := fg.FlagString("g", "greeting", "A Greeting")
	app := fg.FlagString("a", "amount", "An Amount")
	dpp := fg.FlagString("d", "data", "Some Data")
	spp := fg.FlagString("s", "size", "A size")
	dt := fg.FlagLoad("cf", "test_data/flagtest1.lz")

	//greeting -- exists overwritten
	s, err := dt[0].PString("greeting")
	qerr(t, err, "No Greeting")

	if s != "hello" {
		t.Errorf("greeting expected '%s', got '%s'", "hello", s)
	}
	pmatch(t, *gpp, s, "Greeting Pointer")

	//amount -- exists !overwritten
	s, err = dt[0].PString("amount")
	qerr(t, err, "No Amount")

	if s != "twenty" {
		t.Errorf("amount expected 'twenty', got '%s'", s)
	}
	pmatch(t, *app, s, "Amount Pointer")

	//size -- !exists overwritten
	s, err = dt[0].PString("size")
	qerr(t, err, "No Size")
	if s != "small" {
		t.Errorf("size expected 'small', got '%s'", s)
	}
	pmatch(t, *spp, s, "Size Pointer")

	//slob -- !exists !overwritten
	s, err = dt[0].PString("slob")
	if err == nil {
		t.Errorf("Slob Assigened somehow")
	}

	//data -- Assigned before DATA
	s, err = dt[0].PString("data", "Data")
	qerr(t, err, "No data or Data")
	if s != "MEGABYTES" {
		t.Errorf("data expected 'MEGABYTES', got '%s'", s)
	}
	pmatch(t, *dpp, s, "Data Pointer")

}

func Test_FileFind(t *testing.T) {
	fg := flagger{
		args:  []string{"-cf", "test_data/flagtest2.lz"},
		flist: make(map[string]*string),
		fset:  flag.NewFlagSet("test", flag.ContinueOnError),
	}
	fg.FlagString("g", "greeting", "A Greeting")

	fg.FlagString("a", "amount", "An Amount")

	dt := fg.FlagLoad("cf", "test_data/flagtest1.lz")

	s, err := dt[0].PString("greeting")
	qerr(t, err, "No Greeting")
	if s != "Salut" {
		t.Errorf("greeting expected 'Salut', got '%s'", s)
	}
}
