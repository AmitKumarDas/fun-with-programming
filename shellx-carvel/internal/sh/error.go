package sh

import (
	"fmt"
	"strings"
)

type InvalidEnvError struct {
	Context     string
	InvalidEnvs []string
}

func (e *InvalidEnvError) Error() string {
	msg := fmt.Sprintf("found invalid env(s) [%s]", strings.Join(e.InvalidEnvs, ", "))
	if e.Context == "" {
		return msg
	}
	return fmt.Sprintf("%s: %s", e.Context, msg)
}

type MultiError struct {
	Errors []error
}

func (mErr *MultiError) Error() string {
	if len(mErr.Errors) == 0 {
		return fmt.Sprintf("invalid use of %T", mErr)
	}
	if len(mErr.Errors) == 1 {
		return mErr.Errors[0].Error()
	}
	var msgs = make([]string, 0, len(mErr.Errors))
	for _, e := range mErr.Errors {
		msgs = append(msgs, e.Error())
	}
	return strings.Join(msgs, ", ")
}
