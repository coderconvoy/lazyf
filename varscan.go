package lazyf

import (
	"os"
	"strings"
)

type VarBlock struct {
	S   string
	isV bool
}

func (vn VarBlock) String() string {
	if vn.isV {
		return "Var:" + vn.S
	}
	return "Str:" + vn.S
}

func VarScan(s, op, cl string) []VarBlock {
	mode := 0
	skip := 0
	cstring := ""
	res := []VarBlock{}
	for k, v := range s {
		if skip > k {
			continue
		}
		if mode == 0 && strings.HasPrefix(s[k:], op) {
			skip = k + len(op)
			mode = 1
			if cstring != "" {
				res = append(res, VarBlock{cstring, false})
			}
			cstring = ""
			continue
		}

		if mode == 1 && strings.HasPrefix(s[k:], cl) {
			skip = k + len(cl)
			mode = 0
			res = append(res, VarBlock{cstring, true})
			cstring = ""
			continue
		}
		cstring += string(v)
	}
	if cstring != "" {
		res = append(res, VarBlock{cstring, false})
	}
	return res
}

func VarReplace(vv []VarBlock, f func(string) string) string {
	res := ""
	for _, v := range vv {
		if !v.isV {
			res += v.S
			continue
		}
		res += f(v.S)
	}
	return res
}

func FVarReplace(s, op, cl string, f func(string) string) string {
	t := VarScan(s, op, cl)
	return VarReplace(t, f)
}

func EnvReplace(s string) string {
	return FVarReplace(s, "{", "}", os.Getenv)
}
