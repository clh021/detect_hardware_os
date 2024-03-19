package common

import (
	"context"
	"os/user"

	"github.com/gogf/gf/v2/os/glog"
)

func MustSuperUser(ctx context.Context) {
	current, err := user.Current()
	if err != nil {
		glog.Fatal(ctx, err)
	}

	if current.Uid != "0" {
		glog.Fatal(ctx, "requires superuser privilege")
	}
}