package pathology

var (
	defaultRootDirectory = "/var/lib/pathology"
	defaultRootDirSet    = false
)

func SetDefaultRootDir(drd string) {
	defaultRootDirectory = drd
	defaultRootDirSet = true
}

func GetDefaultDirs(dirs ...AbsDir) map[string]string {
	defaultDirs := make(map[string]string, len(dirs))

	for _, d := range dirs {
		defaultDirs[string(d)] = defaultRootDirectory + string(d)
	}

	return defaultDirs
}
