package lazyf

import (
	"bufio"
	"io"
	"os"
	"strconv"
	"strings"

	"github.com/pkg/errors"
)

type LZ struct {
	Name  string
	Deets map[string]string
}

func ParseLZ(s string) LZ {
	s = strings.TrimSpace(s)
	ss := strings.Split(s, ":")
	for k, v := range ss {
		ss[k] = strings.TrimSpace(v)
	}
	return NewLZ(ss[0], ss[1:]...)

}

func NewLZ(name string, ex ...string) LZ {
	mp := make(map[string]string)
	for k, v := range ex {
		mp["ex"+strconv.Itoa(k)] = v
	}
	return LZ{
		Name:  name,
		Deets: mp,
	}
}

func Read(r io.Reader) ([]LZ, error) {
	sc := bufio.NewScanner(r)
	res := []LZ{}
	var curr LZ

	for sc.Scan() {
		t := sc.Text()
		tr := strings.TrimSpace(t)
		if len(tr) == 0 {
			continue
		}
		if tr[0] == '#' {
			continue
		}
		if tr[0] == t[0] {
			//New Entry
			curr = ParseLZ(tr)
			res = append(res, curr)
			continue
		}

		//Deets

		ss := strings.SplitN(tr, ":", 2)
		if len(ss) != 2 {
			return res, errors.New("No Colon in deets line")
		}
		curr.Deets[ss[0]] = ss[1]
	}

	return res, nil
}

func ReadFile(fname string) ([]LZ, error) {
	f, err := os.Open(fname)
	if err != nil {
		return []LZ{}, err
	}
	defer f.Close()
	return Read(f)
}

func (lz LZ) PString(ns ...string) (string, error) {
	for _, v := range ns {
		res, ok := lz.Deets[v]
		if ok {
			return res, nil
		}
	}
	return "", errors.New("Item not found")
}

func (lz LZ) PInt(ns ...string) (int, error) {
	s, err := lz.PString(ns...)
	if err != nil {
		return 0, err
	}
	conv, err := strconv.Atoi(s)
	if err != nil {
		return 0, errors.Wrap(err, "Could not convert Item")
	}

	return conv, nil
}

func (lz LZ) PBool(ns ...string) (bool, error) {
	s, err := lz.PString(ns...)
	if err != nil {
		return false, err
	}

	conv, err := strconv.ParseBool(s)
	if err != nil {
		return false, errors.Wrap(err, "Could not convert Item")
	}

	return conv, nil
}

func (lz LZ) PFloat(ns ...string) (float64, error) {
	s, err := lz.PString(ns...)
	if err != nil {
		return 0, err
	}

	conv, err := strconv.ParseFloat(s, 64)
	if err != nil {
		return 0, errors.Wrap(err, "Could not convert Item")
	}

	return conv, nil
}

func (lz LZ) PIntD(def int, ns ...string) int {
	r, err := lz.PInt(ns...)
	if err != nil {
		return def
	}
	return r
}
func (lz LZ) PBoolD(def bool, ns ...string) bool {
	r, err := lz.PBool(ns...)
	if err != nil {
		return def
	}
	return r
}
func (lz LZ) PFloatD(def float64, ns ...string) float64 {
	r, err := lz.PFloat(ns...)
	if err != nil {
		return def
	}
	return r
}

func ByName(ll []LZ, s string) (LZ, bool) {
	s = strings.ToLower(s)
	for _, v := range ll {
		if strings.ToLower(v.Name) == s {
			return v, true
		}
	}
	return LZ{}, false
}
