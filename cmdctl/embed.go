package main

import _ "embed"

//go:embed completions/cmdctl.bash
var cmdctlScript string

//go:embed completions/uchar.bash
var ucharScript string

//go:embed completions/unbuffer.bash
var unbufferScript string

//go:embed completions/ver.bash
var verScript string
