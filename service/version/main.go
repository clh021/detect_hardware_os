package version

import (
	"context"
	"fmt"

	"github.com/gogf/gf/v2/os/gcmd"
)

var (
    // GitTime is the timestamp of the latest commit.
    GitTime = ""
    // BuildTime is the timestamp when this binary was built.
    BuildTime = ""
    // GitHash represents the short hash of the current commit.
    GitHash = ""
    // GitCount is the total number of commits in the repository.
    GitCount = ""
)

func ShowVersionInfo() {
	fmt.Println("Version Build Info:")
	fmt.Println("\tBuild Time: ", BuildTime)
	fmt.Println("\tBuild  Ver: ", fmt.Sprintf("v0.0.%s", GitCount))
	fmt.Println("\tGit   Time: ", GitTime)
	fmt.Println("\tGit Commit: ", GitHash)
}

var (
	CmdVer = gcmd.Command{
		Name:        "version",
		Usage:       "version",
		Brief:       "show version info",
		Description: "show version info",
		Func: func(ctx context.Context, parser *gcmd.Parser) (err error) {
			ShowVersionInfo()
			return
		},
	}
)