package pathology

import (
	"errors"
	"fmt"
	"github.com/boggydigital/wits"
	"os"
	"slices"
)

type AbsDir string

var (
	absDirsPaths    = map[AbsDir]string{}
	absDirsPathsSet = false
	absDirsKnown    []AbsDir
)

func SetAbsDirs(userDirectoriesFilename string, absDirs ...AbsDir) error {

	absDirsKnown = absDirs
	defaultDirs := getDefaultDirs(absDirs...)

	userDirs := make(map[string]string)

	if _, err := os.Stat(userDirectoriesFilename); err == nil {
		udFile, err := os.Open(userDirectoriesFilename)
		if err != nil {
			return err
		}

		userDirs, err = wits.ReadKeyValue(udFile)
		if err != nil {
			return err
		}
	}

	for _, adk := range absDirs {
		adp := userDirs[string(adk)]
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
