package domain

import "fmt"

type GoogleCloudFileStoreError struct {
	funcName string
	error    error
}

func newGoogleCloudFileStoreError(funcName string, err error) error {
	if err == nil {
		return nil
	}
	storeError := GoogleCloudFileStoreError{
		funcName: funcName,
		error:    err,
	}
	return storeError
}

func (g GoogleCloudFileStoreError) Error() string {
	return fmt.Sprintf("domain.GoogleCloudFileStoreError.%s: %s", g.funcName, g.error.Error())
}
