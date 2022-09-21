package sh

import (
	"github.com/magefile/mage/sh"
	"testing"
)

func TestFile(t *testing.T) {
	requireNoErr(t, sh.Run("mkdir", "-p", "tmp"))
	requireErr(t, File("tmp/will-not-create.txt", "I $WILL_NOT_EXIST due to this unset env\n", 0644))
	requireNoErr(t, File("tmp/will-create.txt", "I WILL EXIST\n", 0644))
}

func TestIsEq(t *testing.T) {
	// mock env k:v pairs
	EnvTestIsEq := "${TEST_IS_EQ}"
	_ = MaybeSetEnv(EnvTestIsEq, "testing")
	defer func() {
		UnsetEnv(EnvTestIsEq)
	}()
	t.Run("match empty value given an unset env", func(t *testing.T) {
		requireTrue(t, IsEq("${DOES_NOT_EXIST}", ""))
	})
	t.Run("match actual value given a set env", func(t *testing.T) {
		requireTrue(t, IsEq(EnvTestIsEq, "testing"))
	})
	t.Run("match substring given a string containing an unset env", func(t *testing.T) {
		requireTrue(t, IsEq("Who am ${I}?", "Who am ?"))
	})
}
