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
	errors []error
}

func (mErr *MultiError) Add(err error) error {
	if mErr == nil {
		return nil
	}
	mErr.errors = append(mErr.errors, err)
	return mErr
}

func (mErr *MultiError) AddAll(errs []error) error {
	if mErr == nil {
		return nil
	}
	for _, e := range errs {
		mErr.Add(e)
	}
	return mErr
}

func (mErr *MultiError) HasError() bool {
	if mErr == nil {
		return false
	}
	return len(mErr.errors) > 0
}

func (mErr *MultiError) Error() string {
	if mErr == nil {
		return ""
	}
	if len(mErr.errors) == 0 {
		return "no errors found"
	}
	if len(mErr.errors) == 1 {
		return mErr.errors[0].Error()
	}
	var msgs = make([]string, 0, len(mErr.errors))
	for _, e := range mErr.errors {
		msgs = append(msgs, e.Error())
	}
	return strings.Join(msgs, ", ")
}

func (mErr *MultiError) ErrOrNil() error {
	if mErr == nil || len(mErr.errors) == 0 {
		return nil
	}
	return mErr
}
