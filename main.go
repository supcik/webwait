// SPDX-FileCopyrightText: 2025 Jacques Supcik <jacques.supcik@hefr.ch>
//
// SPDX-License-Identifier: MIT OR Apache-2.0

package main

import "github.com/supcik/webwait/cmd"

var version = "dev"

func main() {
	cmd.Execute(version)
}
