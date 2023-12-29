package pathology

import (
	"errors"
	"fmt"
	"golang.org/x/exp/maps"
	"os"
)

type AbsDir string

var (
	absDirsPaths    = map[AbsDir]string{}
	absDirsPathsSet = false
)

func SetAbsDirs(kv map[AbsDir]string) error {

	defaultDirs := GetDefaultDirs(maps.Keys(kv)...)

	for adk, adp := range kv {
		if adp != defaultDirs[string(adk)] {
			// make sure directory exists and is accessible to current process
			if _, err := os.Stat(adp); err != nil {
				return err
			}
		}
		absDirsPaths[adk] = adp
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

	if adp, ok := absDirsPaths[ad]; ok && adp != "" {
		return adp, nil
	}
	return "", fmt.Errorf("abs dir %s not set", ad)
}
