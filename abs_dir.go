package pathways

import (
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
		return "", ErrDefaultRootDirNotSet
	}
	if !absDirsPathsSet {
		return "", ErrAbsDirsPathsNotSet
	}
	if !slices.Contains(absDirsKnown, ad) {
		return "", errUnknownAbsDir(ad)
	}

	if adp, ok := absDirsPaths[ad]; ok && adp != "" {
		return adp, nil
	}

	return "", errAbsDirNotSet(ad)
}
