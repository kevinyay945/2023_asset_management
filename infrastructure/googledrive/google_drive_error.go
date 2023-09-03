package googledrive

import "fmt"

type GoogleDriveError struct {
	funcName string
	error    error
}

func newGoogleDriveError(funcName string, err error) error {
	if err == nil {
		return nil
	}
	driveError := GoogleDriveError{
		funcName: funcName,
		error:    fmt.Errorf("%w", err),
	}
	return &driveError
}

func (g GoogleDriveError) Error() string {
	return fmt.Sprintf("infrastructure.googledrive.%s: %s", g.funcName, g.error.Error())
}
