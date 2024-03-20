package version

import (
	"fmt"
	"os"
)

const (
    // GitTime is the timestamp of the latest commit.
    GitTime = "2024-03-19 22:20:48"
    // BuildTime is the timestamp when this binary was built.
    BuildTime = "2024-03-20 17:42:30"
    // GitHash represents the short hash of the current commit.
    GitHash = "6e75aa8a3ae9030966937e90053d8b6a68d79076"
    // GitCount is the total number of commits in the repository.
    GitCount = 7
)

func ShowVersionInfo()  {
	fmt.Println("Version Build Info:")
	fmt.Println("\tBuild Time: ", BuildTime)
	fmt.Println("\tBuild  Ver: ", fmt.Sprintf("v0.0.%d", GitCount))
	fmt.Println("\tGit   Time: ", GitTime)
	fmt.Println("\tGit Commit: ", GitHash)
	os.Exit(0)
}
