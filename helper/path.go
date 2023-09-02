package helper

import (
	"os"
	"os/user"
	"path/filepath"
)

// OsExecutable get current exe path
var OsExecutable = os.Executable

var GetCurrentExecutableFolderPath = func() (string, error) {
	executable, err := OsExecutable()
	if err != nil {
		return "", err
	}
	return filepath.Dir(executable), nil
}

var GetCurrentUserPath = func() (string, error) {
	myself, err := user.Current()
	if err != nil {
		return "", err
	}
	homedir := myself.HomeDir
	return homedir, nil
}
