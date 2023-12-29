package pathology

import (
	"fmt"
	"golang.org/x/exp/maps"
	"os"
)

type AbsDir string

var absDirsPaths = map[AbsDir]string{}

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
	return nil
}

func GetAbsDir(ad AbsDir) (string, error) {
	if adp, ok := absDirsPaths[ad]; ok && adp != "" {
		return adp, nil
	}
	return "", fmt.Errorf("abs dir %s not set", ad)
}
