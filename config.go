package lazyf

import (
	"github.com/pkg/errors"
)

//GetConfig returns the config list, + selected filename, or error
func GetConfig(locs ...string) ([]LZ, string, error) {
	for _, v := range locs {
		lz, err := ReadFile(v)
		if err != nil {
			if _, ok := err.(interface {
				NErrs() int
			}); ok {
				return lz, v, err
			}
			continue
		}
		return lz, v, nil
	}

	return []LZ{}, "", errors.Errorf("Config not found")

}

func GetConfigN(n int, confLocs ...string) (LZ, string, error) {

	carr, fname, err := GetConfig(confLocs...)

	if err != nil {
		return LZ{}, fname, err
	}

	if len(carr) < n {
		return LZ{}, fname, errors.Errorf("No Entry in Config")
	}

	return carr[n], fname, nil

}
