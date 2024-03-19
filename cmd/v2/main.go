package main

import (
	"context"

	"github.com/clh021/detect_hardware_os/service/common"
	"github.com/clh021/detect_hardware_os/service/getinfo"
	"github.com/gogf/gf/v2/os/glog"
)

func main() {
	ctx := context.TODO()
	common.MustSuperUser(ctx)

	// sysinfo
	if e := common.PutJsonByFunc(ctx, "sys.json", getinfo.Sys); e != nil {
		glog.Error(ctx, e)
	}

	// devinfo
	if e := common.PutJsonByFunc(ctx, "dev.json", getinfo.Dev); e != nil {
		glog.Error(ctx, e)
	}

	// hardinfo
	if e := common.PutJsonByFunc(ctx, "hardinfo.json", getinfo.Hard); e != nil {
		glog.Error(ctx, e)
	}
}
