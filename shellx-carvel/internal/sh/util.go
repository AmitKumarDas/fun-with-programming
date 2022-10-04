package sh

import (
	"fmt"
	"os"
	"path"
	"strings"
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

// JoinPathsWithErrHandle returns the final path only if there
// was no error while joining these paths
func JoinPathsWithErrHandle(mErr *Error, paths ...string) string {
	if mErr.HasError() {
		return ""
	}
	final, err := JoinPaths(paths...)
	if err != nil {
		mErr.Add(err)
		return ""
	}
	return final
}

// IsEq verifies if the provided strings are equal. The string(s)
// can also be an env variable or consist of an env variable
// as a substring. ENV variable if any will have its value
// matched.
//
// Note: It will mostly return false for invalid env syntax or
// strings with special shell characters s.a $*, $#, etc.
// Refer os.ExpandEnv & os.Expand for more details
func IsEq(a, b string) bool {
	return strings.ToLower(os.ExpandEnv(a)) == strings.ToLower(os.ExpandEnv(b))
}

func IsNotEq(a, b string) bool {
	return !IsEq(a, b)
}
