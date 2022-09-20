package sh

import (
	"fmt"
	"os"
	"path"
)

func File(name, data string, perm os.FileMode) error {
	expanded, err := ExpandStrictAll(name, data)
	if err != nil {
		return err
	}
	return os.WriteFile(expanded[0], []byte(expanded[1]), perm)
}

func JoinPaths(paths ...string) (string, error) {
	if len(paths) == 0 {
		return "", fmt.Errorf("no paths given")
	}
	expanded, err := ExpandStrictAll(paths...)
	if err != nil {
		return "", err
	}
	return path.Join(expanded...), nil
}
