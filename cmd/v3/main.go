package main

import (
	"context"

	"github.com/clh021/detect_hardware_os/service/browser"
	"github.com/clh021/detect_hardware_os/service/cmd/version"
	"github.com/clh021/detect_hardware_os/service/common"
	"github.com/gogf/gf/v2/os/gcmd"
	"github.com/gogf/gf/v2/os/gctx"
)

var (
	Main = gcmd.Command{
		Name:        "detect_hardware_os",
		Usage:       "detect_hardware_os",
		Brief:       "scan some sysinfo",
		Description: "scan hardinfo, osinfo, development info, browser info",
		Func: func(ctx context.Context, parser *gcmd.Parser) (err error) {
			common.MustSuperUser(ctx)

			// // sysinfo
			// if e := common.PutJsonByFunc(ctx, "sys.json", getinfo.Sys); e != nil {
			// 	glog.Error(ctx, e)
			// }

			// // devinfo
			// if e := common.PutJsonByFunc(ctx, "dev.json", getinfo.Dev); e != nil {
			// 	glog.Error(ctx, e)
			// }

			// // hardinfo
			// if e := common.PutJsonByFunc(ctx, "hardinfo.json", getinfo.Hard); e != nil {
			// 	glog.Error(ctx, e)
			// }

			// browserinfo
			// browser.TestExtractChromeVersion()
			browser.Main()
			return nil
		},
	}
)

func main() {
	err := Main.AddCommand(&version.CmdVer)
	if err != nil {
		panic(err)
	}
	Main.Run(gctx.GetInitCtx())
}
