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
