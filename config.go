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

func GetConfigN(n int, confLocs ...string) (LZ, error) {

	carr, err := GetConfig(confLocs...)

	if err != nil {
		return LZ{}, err
	}

	if len(carr) < n {
		return LZ{}, errors.Errorf("No Entry in Config")
	}

	return carr[n], nil

}
