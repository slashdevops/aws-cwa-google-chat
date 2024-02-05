package version

import (
	"fmt"
)

var (
	// Version is the version as string.
	Version = "dev"

	// GitCommit is the git commit hash as string.
	GitCommit string

	// GitBranch is the git branch as string.
	GitBranch string

	// BuildDate is the build date as string.
	BuildDate string

	// BuildUser is the user who built the binary.
	BuildUser string
)

// GetVersionExtended returns a version string.
func GetVersionExtended() string {
	return fmt.Sprintf("version=%s\ngit=%s\nbranch=%s\ndate=%s\ngitUser=%s\n", Version, GitCommit, GitBranch, BuildDate, BuildUser)
}
