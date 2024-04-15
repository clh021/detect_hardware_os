package common

import (
	"context"
	"os"
	"os/user"

	"github.com/gogf/gf/v2/os/glog"
)

func MustSuperUser(ctx context.Context) {
	current, err := user.Current()
	if err != nil {
		glog.Fatal(ctx, err)
	}

	if current.Uid != "0" {
		glog.Print(ctx, "requires superuser privilege")
		os.Exit(0)
	}
}
