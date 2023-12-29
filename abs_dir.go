package pathology

import (
	"errors"
	"fmt"
	"slices"
)

type AbsDir string

var (
	absDirsPaths    = map[AbsDir]string{}
	absDirsPathsSet = false
	absDirsKnown    []AbsDir
)

func SetAbsDirs(absDirs ...AbsDir) error {
	absDirsKnown = absDirs
	absDirsPaths = getDefaultDirs(absDirs...)
	absDirsPathsSet = true
	return nil
}

func SetUserDirsOverrides(userDirs map[string]string) {
	for absDir, absPath := range userDirs {
		absDirsPaths[AbsDir(absDir)] = absPath
	}
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
