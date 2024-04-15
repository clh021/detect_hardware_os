package getinfo

import (
	"context"

	"github.com/zcalusic/sysinfo"
)

func Sys(ctx context.Context) (interface{}, error) {
	var si sysinfo.SysInfo
	si.GetSysInfo()
	return &si, nil
}
