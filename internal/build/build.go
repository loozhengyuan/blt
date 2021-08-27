// Package build provides build information.
package build

import (
	"runtime"

	"github.com/loozhengyuan/grench/build"
)

// These variables are set dynamically at build time.
var (
	// Version is the current version of the application.
	Version = "v0.0.0"

	// CommitHash is the commit hash at build time.
	CommitHash = "dev"

	// Timestamp is the timestamp at build time.
	Timestamp = "1970-01-01T00:00:00Z"
)

// Info returns the build information of the application.
func Info(name string) build.Info {
	return build.Info{
		App:       name,
		System:    runtime.GOOS,
		Arch:      runtime.GOARCH,
		Version:   Version,
		Commit:    CommitHash,
		Timestamp: Timestamp,
	}
}
