package browser

import (
	"fmt"
	"log"
	"regexp"
	"strings"
	"sync"
	"time"

	"github.com/clh021/detect_hardware_os/service/callsrv"
	"github.com/clh021/detect_hardware_os/service/common"
	"github.com/gogf/gf/v2/frame/g"
)

type ConfigWithLock struct {
	sync.Mutex
	Items []BrowserItem
}

func GetBrowsers() *[]BrowserItem {
	confWithLock := &ConfigWithLock{Items: getConf()}

	port, err := common.GetFreePort()
	if err != nil {
		log.Fatalln(err)
	}
	defaultUrl := fmt.Sprintf("http://127.0.0.1:%d?b=defaultbrowser", port)
	UserAgentServe(port, &confWithLock.Items)

	waitSecond := 6
	timeoutTimer := time.NewTimer(time.Duration(waitSecond) * time.Second)
	defer timeoutTimer.Stop()

	for i, _ := range confWithLock.Items {
		// getUserAgentVersion(port, &c)
		getVersion(&confWithLock.Items[i])
	}

	checkUserAgent := func() {
		for {
			select {
			case <-timeoutTimer.C:
				fmt.Printf("仍在获取浏览器信息，您可使用常用浏览器访问地址：%s 以完成浏览器信息采集。\n", defaultUrl)
				return
			default:
				time.Sleep(5 * time.Second)
			}
		}
	}
	checkUserAgent()
	fmt.Println("===================================================")
	g.Dump(confWithLock.Items)
	fmt.Println("===================================================")
	return &confWithLock.Items
}

func getVersion(b *BrowserItem) (e error) {
	out, e := callsrv.ExecGetSysInfoStdout(b.Bin, "--version")
	if e != nil {
		return
	}
	b.CmdVer = strings.TrimSpace(string(out))

	cmdVer := strings.TrimPrefix(b.CmdVer, "360Browser") // 针对360浏览器的数字调整

	b.Version, e = regVer(cmdVer, b.CmdReg)
	return
}

func regVer(str, reg string) (string, error) {
	pattern := regexp.MustCompile(reg)
	// 查找匹配项
	matches := pattern.FindStringSubmatch(str)
	if len(matches) < 2 {
		return "", fmt.Errorf("无法提取 版本号")
	}
	// 提取版本号（matches[1] 是第一个括号内的匹配内容）
	ver := matches[1]
	return ver, nil
}
