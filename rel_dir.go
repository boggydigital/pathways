package pathways

import (
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
		return "", ErrRelToAbsDirNotSet
	}
	if !slices.Contains(relDirsKnown, rd) {
		return "", errUnknownRelDir(rd)
	}

	if ad, ok := relToAbsDir[rd]; ok {

		adp, err := GetAbsDir(ad)
		if err != nil {
			return "", err
		}

		return filepath.Join(adp, string(rd)), nil
	} else {
		return "", errRelativityNotSet(rd)
	}
}

func GetRelDir(rd RelDir) (string, error) {
	if !slices.Contains(relDirsKnown, rd) {
		return "", errUnknownRelDir(rd)
	}
	return string(rd), nil
}
