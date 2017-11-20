package lazyf

import (
	"flag"
	"os"
)

var flagloc string

type stringflag struct {
	s   *string
	def string
}

type flagger struct {
	fset  *flag.FlagSet
	flist map[string]stringflag
	blist map[string]*bool
	args  []string
}

//defFlagger will normally be used, but for testing other flaggers may be
var defFlagger flagger = NewFlagger(flag.CommandLine, os.Args[1:])

func NewFlagger(fset *flag.FlagSet, args []string) flagger {
	return flagger{
		fset:  fset,
		args:  args,
		flist: make(map[string]stringflag),
		blist: make(map[string]*bool),
	}
}

//FlagString adds the flag, for setting the first section of the lzconfig
//def value will override pointer result if nothing is set, but will not change result of PStringD on the config
func FlagString(f, def, cname, info string) *string {
	return defFlagger.FlagString(f, def, cname, info)
}

func (ff flagger) FlagString(f, def, cname, info string) *string {
	p := ff.fset.String(f, "", info)

	ff.flist[cname] = stringflag{p, def}
	return p
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

//FlagLoad must be called after all FlagString methods and then never again returns []LZ for config, and selected Config Filename
func FlagLoad(f string, deflocs ...string) ([]LZ, string) {
	return defFlagger.FlagLoad(f, deflocs...)
}

func (ff flagger) FlagLoad(f string, deflocs ...string) ([]LZ, string) {
	cloc := ff.fset.String(f, "", "Location of Configuration File")
	ff.fset.Parse(ff.args)

	if *cloc != "" {
		deflocs = []string{*cloc}
	}

	cfig, fname, err := GetConfig(deflocs...)
	if err != nil {
		cfig = append(cfig, LZ{})
	}

	//Strings
	for k, v := range ff.flist {
		if *v.s != "" {
			cfig[0].Deets[k] = *v.s
			continue
		}
		fv, err := cfig[0].PString(k)
		if err == nil {
			*v.s = fv
			continue
		}
		*v.s = v.def
	}

	//Bools are only false if not set in config, and not flagged
	for k, v := range ff.blist {
		if *v {
			cfig[0].Deets[k] = "T"
			continue
		}
		*v = cfig[0].PBoolD(false, k)
	}

	return cfig, fname
}
