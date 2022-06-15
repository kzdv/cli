package version

import "fmt"

var (
	Version   = "dev"
	GitCommit = "HEAD"
	BuildTime = "unknown"
	GoVersion = ""
)

func FriendlyVersion() string {
	return fmt.Sprintf("%s-%s (built: %s [GoVersion %s])", Version, GitCommit, BuildTime, GoVersion)
}
