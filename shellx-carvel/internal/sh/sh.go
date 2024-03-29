package sh

// Borrowed from https://github.com/magefile/mage/tree/master/sh

import (
	"bytes"
	"fmt"
	"io"
	"log"
	"os"
	"os/exec"
	"strings"
)

const (
	CMDLogLevelNone  = "0"
	CMDLogLevelInfo  = "1"
	CMDLogLevelDebug = "2"
)

const EnvCMDLogLevel = "${CMD_LOG_LEVEL}"

func init() {
	MaybeSetEnv(EnvCMDLogLevel, CMDLogLevelNone)
}

func IsDebug() bool {
	return IsEq(EnvCMDLogLevel, CMDLogLevelDebug)
}

func IsInfo() bool {
	return IsEq(EnvCMDLogLevel, CMDLogLevelInfo)
}

// RunCmd returns a function that will call Run with the given command. This is
// useful for creating command aliases to make your scripts easier to read, like
// this:
//
// In a helper file somewhere
// var g0 = sh.RunCmd("go")  // go is a keyword :(
//
// Somewhere in your main code
//
//	if err := g0("install", "github.com/gohugo/hugo"); err != nil {
//	 return err
//	}
func RunCmd(cmd string, args ...string) func(args ...string) error {
	return func(args2 ...string) error {
		return Run(cmd, append(args, args2...)...)
	}
}

// OutCmd is like RunCmd except the command returns the output of the
// command.
func OutCmd(cmd string, args ...string) func(args ...string) (string, error) {
	return func(args2 ...string) (string, error) {
		return Output(cmd, append(args, args2...)...)
	}
}

// Run is like RunWith, but doesn't specify any environment variables.
func Run(cmd string, args ...string) error {
	return RunWith(nil, cmd, args...)
}

// RunWith runs the given command with the environment variables for
// the command being run.
//
// Note: Environment variables should be in the format name=value.
func RunWith(env map[string]string, cmd string, args ...string) error {
	_, err := OutputWith(env, cmd, args...)
	return err
}

// RunV is like Run, but always sends the command's stdout to os.Stdout.
func RunV(cmd string, args ...string) error {
	_, err := Exec(nil, os.Stdout, os.Stderr, cmd, args...)
	return err
}

// RunE accepts error if any and execute the provided
// command iff there was no previous error(s).
//
// This function might be helpful to avoiding error handling
// after each invocation. This enables handling the error
// only once at the end of all invocations.
//
//	func workflow() error {
//		var e Error
//		RunE(&e, "kubectl", "get", "po", "-n", "kube-system", "pod-a")
//		RunE(&e, "kubectl", "describe", "po", "-n", "kube-system", "pod-a")
//		RunE(&e, "kubectl", "get", "svc", "-n", "kube-system", "svc-a")
//		return e.ErrOrNil()
//	}
func RunE(shErr *Error, cmd string, args ...string) {
	if shErr.HasError() {
		return
	}
	_, err := OutputWith(nil, cmd, args...)
	_ = shErr.Add(err)
}

// RunWithV is like RunWith, but always sends the command's stdout to os.Stdout.
func RunWithV(env map[string]string, cmd string, args ...string) error {
	_, err := Exec(env, os.Stdout, os.Stderr, cmd, args...)
	return err
}

// Output runs the command and returns the text from stdout.
func Output(cmd string, args ...string) (string, error) {
	return OutputWith(nil, cmd, args...)
}

// OutputWith returns what is written to stdout. Its error handling
// suits debuggability.
func OutputWith(env map[string]string, cmd string, args ...string) (string, error) {
	bufOut := &bytes.Buffer{}
	bufErr := &bytes.Buffer{}
	if _, err := Exec(env, bufOut, bufErr, cmd, args...); err != nil {
		msg1 := strings.TrimSuffix(bufErr.String(), "\n")
		msg2 := strings.TrimSuffix(bufOut.String(), "\n")
		newErr := fmt.Errorf("%w : %s : %s", err, msg1, msg2)
		return "", newErr
	}
	return strings.TrimSuffix(bufOut.String(), "\n"), nil
}

// Exec executes the command, piping its stderr to mage's stderr and
// piping its stdout to the given writer. If the command fails, it will return
// an error that, if returned from a target or mg.Deps call, will cause mage to
// exit with the same code as the command failed with.  Env is a list of
// environment variables to set when running the command, these override the
// current environment variables set (which are also passed to the command). cmd
// and envs may include references to environment variables in $FOO format, in
// which case these will be expanded before the command is run.
//
// Ran reports if the command ran (rather than was not found or not executable).
// Code reports the exit code the command returned if it ran. If err == nil, ran
// is always true and code is always 0.
func Exec(env map[string]string, stdout, stderr io.Writer, cmd string, args ...string) (ran bool, err error) {
	var invalidEnvs = make([]string, 0, len(args)+1) // size includes all envs & cmd
	// a strict expand function that adds to error if
	// there was no expansion
	mustExpand := func(s string) string {
		s2, ok := env[s]
		if ok {
			return s2
		}
		val := os.Getenv(s)
		if val == "" {
			invalidEnvs = append(invalidEnvs, s)
		}
		return val
	}
	cmd = os.Expand(cmd, mustExpand)
	for i := range args {
		args[i] = os.Expand(args[i], mustExpand)
	}
	if len(invalidEnvs) != 0 {
		return false, &InvalidEnvError{
			Context:     fmt.Sprintf(`failed to run "%s %s"`, cmd, strings.Join(args, " ")),
			InvalidEnvs: invalidEnvs,
		}
	}
	ran, code, err := run(env, stdout, stderr, cmd, args...)
	if err == nil {
		return true, nil
	}
	if ran {
		return ran, fmt.Errorf(`running "%s %s" failed with exit code %d`, cmd, strings.Join(args, " "), code)
	}
	return ran, fmt.Errorf(`failed to run "%s %s: %v"`, cmd, strings.Join(args, " "), err)
}

func run(env map[string]string, stdout, stderr io.Writer, cmd string, args ...string) (ran bool, code int, err error) {
	c := exec.Command(cmd, args...)
	c.Env = os.Environ()
	for k, v := range env {
		c.Env = append(c.Env, k+"="+v)
	}
	c.Stderr = stderr
	c.Stdout = stdout
	c.Stdin = os.Stdin

	var quoted []string
	for i := range args {
		quoted = append(quoted, fmt.Sprintf("%q", args[i]))
	}
	err = c.Run()
	ran = CmdRan(err)
	code = ExitStatus(err)
	// When logging is required irrespective of successful or unsuccessful runs
	if IsInfo() {
		log.Println("exec:", cmd, strings.Join(quoted, " "), "status:", ran, code)
	}
	return ran, code, err
}

// CmdRan examines the error to determine if it was generated as a result of a
// command running via os/exec.Command.  If the error is nil, or the command ran
// (even if it exited with a non-zero exit code), CmdRan reports true.  If the
// error is an unrecognized type, or it is an error from exec.Command that says
// the command failed to run (usually due to the command not existing or not
// being executable), it reports false.
func CmdRan(err error) bool {
	if err == nil {
		return true
	}
	ee, ok := err.(*exec.ExitError)
	if ok {
		return ee.Exited()
	}
	return false
}

type exitStatus interface {
	ExitStatus() int
}

// ExitStatus returns the exit status of the error if it is an exec.ExitError
// or if it implements ExitStatus() int.
// 0 if it is nil or 1 if it is a different error.
func ExitStatus(err error) int {
	if err == nil {
		return 0
	}
	if e, ok := err.(exitStatus); ok {
		return e.ExitStatus()
	}
	if e, ok := err.(*exec.ExitError); ok {
		if ex, ok := e.Sys().(exitStatus); ok {
			return ex.ExitStatus()
		}
	}
	return 1
}
