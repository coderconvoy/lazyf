package lazyf

import (
	"github.com/pkg/errors"
)

func GetConfig(locs ...string) ([]LZ, error) {
	for _, v := range locs {
		lz, err := ReadFile(v)
		if err != nil {
			if _, ok := err.(interface {
				NErrs() int
			}); ok {
				return lz, err
			}
			continue
		}
		return lz, nil
	}

	return []LZ{}, errors.Errorf("Config not found")

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
