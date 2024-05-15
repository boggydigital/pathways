package pathways

import "errors"

var (
	ErrDefaultRootDirNotSet = errors.New("pathways: default root dir not set")
	ErrAbsDirsPathsNotSet   = errors.New("pathways: abs dirs paths not set")
	ErrRelToAbsDirNotSet    = errors.New("pathways: rel to abs dir not set")
)

func errAbsDirNotSet(d AbsDir) error {
	return errors.New("pathways: abs dir " + string(d) + " not set")
}

func errUnknownAbsDir(d AbsDir) error {
	return errors.New("pathways: unknown abs dir " + string(d))
}

func errUnknownRelDir(d RelDir) error {
	return errors.New("pathways: unknown rel dir " + string(d))
}

func errRelativityNotSet(d RelDir) error {
	return errors.New("pathways: " + string(d) + "dir relativity not set")
}
