package lazyf

import (
	"flag"
	"os"
)

var flagloc string

type flagger struct {
	fset  *flag.FlagSet
	flist map[string]*string
	blist map[string]*bool
	args  []string
}

//defFlagger will normally be used, but for testing other flaggers may be
var defFlagger flagger = NewFlagger(flag.CommandLine, os.Args[1:])

func NewFlagger(fset *flag.FlagSet, args []string) flagger {
	return flagger{
		fset:  fset,
		args:  args,
		flist: make(map[string]*string),
		blist: make(map[string]*bool),
	}
}

//FlagString adds the flag, for setting the first section of the lzconfig
func FlagString(f, cname, info string) *string {
	return defFlagger.FlagString(f, cname, info)
}

func (ff flagger) FlagString(f, cname, info string) *string {
	s := ff.fset.String(f, "", info)

	ff.flist[cname] = s
	return s
}

//FlagBool adds a boolean flag. False is considered unset for adding at core
func FlagBool(f, cname, info string) *bool {
	return defFlagger.FlagBool(f, cname, info)
}

func (ff flagger) FlagBool(f, cname, info string) *bool {
	s := ff.fset.Bool(f, false, "info")
	ff.blist[cname] = s
	return s
}

//FlagLoad must be called after all FlagString methods and then never again
func FlagLoad(f string, deflocs ...string) []LZ {
	return defFlagger.FlagLoad(f, deflocs...)
}

func (ff flagger) FlagLoad(f string, deflocs ...string) []LZ {
	cloc := ff.fset.String(f, "", "Location of Configuration File")
	ff.fset.Parse(ff.args)

	if *cloc != "" {
		deflocs = []string{*cloc}
	}

	cfig, err := GetConfig(deflocs...)
	if err != nil {
		cfig = append(cfig, LZ{})
	}

	//Strings
	for k, v := range ff.flist {
		if *v != "" {
			cfig[0].Deets[k] = *v
			continue
		}
		*v = cfig[0].PStringD("", k)
	}

	//Bools are only false if not set in config, and not flagged
	for k, v := range ff.blist {
		if *v {
			cfig[0].Deets[k] = "T"
			continue
		}
		*v = cfig[0].PBoolD(false, k)
	}

	return cfig
}
