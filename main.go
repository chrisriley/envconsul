package main

import (
	"bytes"
	"fmt"
	"os"
)

// The git commit that was compiled. This will be filled in by the compiler.
var GitCommit string

const Name = "envconsul"
const Version = "0.5.1"
const VersionPrerelease = "dev"

func main() {
	cli := &CLI{outStream: os.Stdout, errStream: os.Stderr}
	os.Exit(cli.Run(os.Args))
}

// formattedVersion returns a formatted version string which includes the git
// commit and development information.
func formattedVersion() string {
	var versionString bytes.Buffer
	fmt.Fprintf(&versionString, "%s v%s", Name, Version)

	if VersionPrerelease != "" {
		fmt.Fprintf(&versionString, "-%s", VersionPrerelease)

		if GitCommit != "" {
			fmt.Fprintf(&versionString, " (%s)", GitCommit)
		}
	}
	return versionString.String()
}
