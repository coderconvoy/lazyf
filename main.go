package lazyf

import (
	"bufio"
	"io"
	"os"
	"strconv"
	"strings"
)

type LZ struct {
	Name  string
	Deets map[string]string
}

func ParseLZ(s string) *LZ {
	s = strings.TrimSpace(s)
	ss := strings.Split(s, ":")
	for k, v := range ss {
		ss[k] = strings.TrimSpace(v)
	}
	return NewLZ(ss[0], ss[1:]...)

}

func NewLZ(name string, ex ...string) *LZ {
	mp := make(map[string]string)
	for k, v := range ex {
		mp["ex"+strconv.Itoa(k)] = v
	}
	return &LZ{
		Name:  name,
		Deets: mp,
	}
}

func Read(r io.Reader) ([]LZ, error) {
	sc := bufio.NewScanner(r)
	res := []LZ{}

	for sc.Scan() {
		sc.Text()
		//TODO
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
