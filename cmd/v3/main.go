package main

import (
	"context"

	"github.com/clh021/detect_hardware_os/service/browser"
	"github.com/clh021/detect_hardware_os/service/cmd/version"
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