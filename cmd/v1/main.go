package main

import (
	"context"
	"os/user"

	"github.com/clh021/detect_hardware_os/service/common"
	"github.com/clh021/detect_hardware_os/service/develop"
	"github.com/zcalusic/sysinfo"

	"github.com/gogf/gf/v2/os/glog"
)

func main() {
	ctx := context.TODO()
	current, err := user.Current()
	if err != nil {
		glog.Fatal(ctx, err)
	}

	if current.Uid != "0" {
		glog.Fatal(ctx, "requires superuser privilege")
	}

	// sysinfo
	if e := PutSysInfo(ctx); e != nil {
		glog.Error(ctx, e)
	}

	// devinfo
	if e := PutDevInfo(ctx); e != nil {
		glog.Error(ctx, e)
	}
}

func PutDevInfo(ctx context.Context) error {
	dev := develop.GetDevelopments()
	return common.PutJsonByData(ctx, "devinfo.json", &dev)
}

func PutSysInfo(ctx context.Context) error {
	var si sysinfo.SysInfo
	si.GetSysInfo()

	return common.PutJsonByData(ctx, "sysinfo.json", &si)
}