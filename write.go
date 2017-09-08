package lazyf

import (
	"fmt"
	"io"
	"os"
	"sort"
	"strconv"
)

func (l LZ) WriteOut(w io.Writer) error {
	//First line
	l1 := l.Name
	i := 0
	for {
		nx, ok := l.Deets["ex"+strconv.Itoa(i)]
		if !ok {
			break
		}
		l1 += ":" + nx
		i++
	}
	_, err := fmt.Fprintln(w, l1)
	if err != nil {
		return err
	}

	//Get sorted keys for sorted output
	klist := []string{}
	for k, _ := range l.Deets {
		klist = append(klist, k)
	}
	sort.Strings(klist)
	for _, kv := range klist {
		v, _ := l.Deets[kv]
		_, err = fmt.Fprintf(w, "\t%s:%s\n\n", kv, v)
		if err != nil {
			return err
		}
	}

	return nil
}

func Write(ll []LZ, w io.Writer) error {
	for _, v := range ll {
		err := v.WriteOut(w)
		if err != nil {
			return err
		}
	}
	return nil
}

func WriteFile(ll []LZ, fname string) error {
	f, err := os.OpenFile(fname, os.O_WRONLY|os.O_CREATE, 0777)
	if err != nil {
		return err
	}

	defer f.Close()

	return Write(ll, f)

}
