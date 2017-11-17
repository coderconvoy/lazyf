package lazyf

import (
	"flag"
	"os"
)

var flagloc string

type flagger struct {
	fset  *flag.FlagSet
	flist map[string]*string
	args  []string
}

//defFlagger will normally be used, but for testing other flaggers may be
var defFlagger flagger = flagger{
	fset:  flag.CommandLine,
	flist: make(map[string]*string),
	args:  os.Args[1:],
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

	for k, v := range ff.flist {
		if *v != "" {
			cfig[0].Deets[k] = *v
			continue
		}
		*v = cfig[0].PStringD("", k)
	}
	return cfig
}
