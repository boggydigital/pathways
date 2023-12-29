package pathology

import (
	"fmt"
	"path/filepath"
)

type RelDir string

var relToAbsDir = map[RelDir]AbsDir{}

func SetRelToAbsDir(ra map[RelDir]AbsDir) {
	relToAbsDir = ra
}

func GetAbsRelDir(rd RelDir) (string, error) {
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
