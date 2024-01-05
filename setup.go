package pathology

import (
	"github.com/boggydigital/wits"
	"os"
)

func Setup(
	optionalDirsOverrideFilename string,
	defaultRootDir string,
	relToAbsDirs map[RelDir]AbsDir,
	absDirs ...AbsDir) error {
	SetDefaultRootDir(defaultRootDir)
	if err := SetAbsDirs(absDirs...); err != nil {
		return err
	}
	if _, err := os.Stat(optionalDirsOverrideFilename); err == nil {
		udFile, err := os.Open(optionalDirsOverrideFilename)
		if err != nil {
			return err
		}
		userDirs, err := wits.ReadKeyValue(udFile)
		if err != nil {
			return err
		}
		SetUserDirsOverrides(userDirs)
	}
	if relToAbsDirs != nil {
		SetRelToAbsDir(relToAbsDirs)
	}

	return nil
}
