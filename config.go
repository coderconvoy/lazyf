package lazyf

import (
	"github.com/pkg/errors"
)

func GetConfig(locs ...string) ([]LZ, error) {
	for _, v := range locs {
		lz, err := ReadFile(v)
		if err != nil {
			if e, ok := err.(LineErr); ok {
				return lz, e
			}
			continue
		}
		return lz, nil
	}

	return []LZ{}, errors.Errorf("arr")

}
