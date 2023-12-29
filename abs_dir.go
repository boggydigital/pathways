package pathology

import (
	"errors"
	"fmt"
	"golang.org/x/exp/maps"
	"os"
	"slices"
)

type AbsDir string

var (
	absDirsPaths    = map[AbsDir]string{}
	absDirsPathsSet = false
	absDirsKnown    []AbsDir
)

func SetAbsDirs(kv map[AbsDir]string) error {

	absDirsKnown = maps.Keys(kv)
	defaultDirs := GetDefaultDirs(absDirsKnown...)

	for adk, adp := range kv {
		if adp == "" {
			adp = defaultDirs[adk]
		}
		absDirsPaths[adk] = adp
		if _, err := os.Stat(adp); err != nil {
			return err
		}
	}

	absDirsPathsSet = true
	return nil
}

func GetAbsDir(ad AbsDir) (string, error) {

	if !defaultRootDirSet {
		return "", errors.New("pathology default root dir not set")
	}
	if !absDirsPathsSet {
		return "", errors.New("pathology abs dirs paths not set")
	}
	if !slices.Contains(absDirsKnown, ad) {
		return "", errors.New("unknown abs dir " + string(ad))
	}

	if adp, ok := absDirsPaths[ad]; ok && adp != "" {
		return adp, nil
	}
	return "", fmt.Errorf("abs dir %s not set", ad)
}
