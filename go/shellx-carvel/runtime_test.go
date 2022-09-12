package shellx_carvel

import (
	"os"
	"testing"
)

func TestFilterInvalidEnvs(t *testing.T) {
	var tmpEnvKey = "TESTING_FILTER_INVALID_ENVS_FUNC"
	os.Setenv(tmpEnvKey, "")
	defer func() {
		os.Unsetenv(tmpEnvKey)
	}()
	var scenarios = []struct {
		name  string
		envs  []string
		isErr bool
	}{
		{
			name:  "given all valid envs",
			envs:  []string{EnvGOOS, EnvGOARCH, EnvVersion, EnvBinPathCarvel, EnvRegistryName},
			isErr: false,
		},
		{
			name:  "given all valid envs & one env with empty value",
			envs:  []string{EnvGOOS, EnvGOARCH, EnvVersion, EnvBinPathCarvel, EnvRegistryName, "$" + tmpEnvKey},
			isErr: false,
		},
		{
			name:  "given all envs that can not be looked up",
			envs:  []string{"$NOOP", "$DONT_EXIST", "${NO}"},
			isErr: true,
		},
		{
			name:  "given all envs with invalid syntax",
			envs:  []string{"${}", "${"},
			isErr: true,
		},
		{
			name:  "given one valid & then one invalid env",
			envs:  []string{EnvRegistryPort, "${}"},
			isErr: true,
		},
		{
			name:  "given many valid envs & then one invalid env",
			envs:  []string{EnvRegistryPort, EnvVersion, EnvAppBundleName, "${}"},
			isErr: true,
		},
		{
			name:  "given one invalid env & then many valid envs",
			envs:  []string{"${", EnvRegistryPort, EnvVersion, EnvAppBundleName},
			isErr: true,
		},
	}

	for _, s := range scenarios {
		t.Run(s.name, func(t *testing.T) {
			got := verifyArgs(s.envs...)
			if s.isErr {
				requireErr(t, got)
			} else {
				requireNoErr(t, got)
			}
		})
	}
}

func TestFile(t *testing.T) {
	requireErr(t, file("tmp/will-not-create.txt", "I $WILL_NOT_EXIST due to this unset env\n", 0644))
	requireNoErr(t, file("tmp/will-create.txt", "I WILL EXIST\n", 0644))
}
