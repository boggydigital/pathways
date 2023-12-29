package pathology

var (
	defaultRootDirectory = "/var/lib/pathology"
	defaultRootDirSet    = false
)

func SetDefaultRootDir(drd string) {
	defaultRootDirectory = drd
	defaultRootDirSet = true
}

func getDefaultDirs(dirs ...AbsDir) map[AbsDir]string {
	defaultDirs := make(map[AbsDir]string, len(dirs))

	for _, d := range dirs {
		defaultDirs[d] = defaultRootDirectory + string(d)
	}

	return defaultDirs
}
