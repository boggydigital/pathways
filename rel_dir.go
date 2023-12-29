package pathology

import (
	"errors"
	"fmt"
	"path/filepath"
)

type RelDir string

var (
	relToAbsDir    = map[RelDir]AbsDir{}
	relToAbsDirSet = false
)

func SetRelToAbsDir(ra map[RelDir]AbsDir) {
	relToAbsDir = ra
	relToAbsDirSet = true
}

func GetAbsRelDir(rd RelDir) (string, error) {
	if !relToAbsDirSet {
		return "", errors.New("pathology rel to abs dir not set")
	}

	if ad, ok := relToAbsDir[rd]; ok {

		adp, err := GetAbsDir(ad)
		if err != nil {
			return "", err
		}

		return filepath.Join(adp, string(rd)), nil
	} else {
		return "", fmt.Errorf("%s dir relativity not set", rd)
	}
}
