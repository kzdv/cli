package version

import (
	_ "embed"
	"fmt"
)

//go:generate bash gen_data.sh
//go:embed version.txt
var Version string

//go:embed build_time.txt
var BuildTime string

//go:embed git_hash.txt
var GitCommit string

//go:embed go_version.txt
var GoVersion string

func FriendlyVersion() string {
	return fmt.Sprintf("%s-%s (built: %s [GoVersion %s])", Version, GitCommit, BuildTime, GoVersion)
}
