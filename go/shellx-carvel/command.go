package shellx_carvel

import "carvel.shellx.dev/internal/sh"

// Carvel binaries / CLIs as functions
var kbld = sh.RunCmdStrict(binPathCarvel + "/kbld")
var whichKbld = sh.RunCmdStrict("ls", binPathCarvel+"/kbld")
var imgpkg = sh.RunCmdStrict(binPathCarvel + "/imgpkg")
var whichImgpkg = sh.RunCmdStrict("ls", binPathCarvel+"/imgpkg")
var ytt = sh.RunCmdStrict(binPathCarvel + "/ytt")
var whichYtt = sh.RunCmdStrict("ls", binPathCarvel+"/ytt")

// KIND CLI as function
var kind = sh.RunCmdStrict(binPathKind + "/kind")
var whichKind = sh.RunCmdStrict("ls", binPathKind+"/kind")

// Docker CLI as function
var docker = sh.RunCmdStrict("docker")

// Unix & generic commands as functions
var mkdir = sh.RunCmdStrict("mkdir", "-p")
var curl = sh.RunCmdStrict("curl")
var ls = sh.RunCmdStrict("ls")
var chmod = sh.RunCmdStrict("chmod")

// File creation as a function
var file = sh.File
