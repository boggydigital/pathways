package pathology

import (
	"errors"
	"fmt"
	"golang.org/x/exp/maps"
	"path/filepath"
	"slices"
)

type RelDir string

var (
	relToAbsDir    = map[RelDir]AbsDir{}
	relToAbsDirSet = false
	relDirsKnown   []RelDir
)

func SetRelToAbsDir(ra map[RelDir]AbsDir) {
	relDirsKnown = maps.Keys(ra)
	relToAbsDir = ra
	relToAbsDirSet = true
}

func GetAbsRelDir(rd RelDir) (string, error) {
	if !relToAbsDirSet {
		return "", errors.New("pathology rel to abs dir not set")
	}
	if !slices.Contains(relDirsKnown, rd) {
		return "", errors.New("unknown rel dir " + string(rd))
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

func GetRelDir(rd RelDir) (string, error) {
	if !slices.Contains(relDirsKnown, rd) {
		return "", errors.New("unknown rel dir " + string(rd))
	}
	return string(rd), nil
}
