package lazyf

import "flag"

var flagloc string

var flagList map[string]*string = make(map[string]*string)

//FlagString adds the flag, for setting the first section of the lzconfig
func FlagString(f, cname, info string) *string {
	s := flag.String(f, "", info)
	flagList[cname] = s

}

//FlagLoad must be called after all FlagString methods and then never again
func FlagLoad(f string, deflocs ...string) []LZ {
	cloc := flag.String(f, "", "Location of Configuration File")
	flag.Parse()

	if *cloc != "" {
		deflocs = []string{*cloc}
	}

	cfig, err := GetConfig(deflocs)

	for k, v := range flagList {
		if *v != "" {
			cfig[0].Deets[k] = *v
		}
	}
	return cfig

}
