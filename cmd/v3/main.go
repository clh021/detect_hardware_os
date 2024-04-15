package main

import (
	"context"
	"log"

	"github.com/clh021/detect_hardware_os/service/cmd/version"
	"github.com/clh021/detect_hardware_os/service/common"
	"github.com/clh021/detect_hardware_os/service/getinfo"
	"github.com/gogf/gf/v2/os/gcmd"
	"github.com/gogf/gf/v2/os/gctx"
	"github.com/gogf/gf/v2/os/glog"
)

var (
	Main = gcmd.Command{
		Name:        "detect_hardware_os",
		Usage:       "detect_hardware_os",
		Brief:       "scan some sysinfo",
		Description: "scan hardinfo, osinfo, development info, browser info",
		Func: func(ctx context.Context, parser *gcmd.Parser) (err error) {
			// common.MustSuperUser(ctx)
			log.Println("启动环境采集服务...")
			resultPath := "results"

			// browserinfo
			log.Println("采集浏览器环境信息...")
			if e := common.PutJsonByFunc(ctx, resultPath+"/browser.json", getinfo.Browser); e != nil {
				glog.Error(ctx, e)
			}

			// sysinfo
			log.Println("采集系统环境信息...")
			if e := common.PutJsonByFunc(ctx, resultPath+"/sys.json", getinfo.Sys); e != nil {
				glog.Error(ctx, e)
			}

			// devinfo
			log.Println("采集开发环境信息...")
			if e := common.PutJsonByFunc(ctx, resultPath+"/dev.json", getinfo.Dev); e != nil {
				glog.Error(ctx, e)
			}

			// hardinfo
			log.Println("采集硬件环境信息...")
			if e := common.PutJsonByFunc(ctx, resultPath+"/hardinfo.json", getinfo.Hard); e != nil {
				glog.Error(ctx, e)
			}
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
