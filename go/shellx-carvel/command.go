package shellx_carvel

import "carvel.shellx.dev/internal/sh"

// Unix & generic commands as functions
var mkdir = sh.RunCmdStrict("mkdir", "-p")
var curl = sh.RunCmdStrict("curl")
var ls = sh.RunCmdStrict("ls")
var chmod = sh.RunCmdStrict("chmod")

// Docker CLI as function
var docker = sh.RunCmdStrict("docker")

// File creation as a function
var file = sh.File

func exists(file string) bool {
	return ls(file) == nil
}
