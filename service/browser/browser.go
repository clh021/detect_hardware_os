package browser

import (
	"fmt"
	"log"
	"strings"
	"sync"
	"time"

	"github.com/clh021/detect_hardware_os/service/callsrv"
	"github.com/clh021/detect_hardware_os/service/common"
)

type Conf [4]string

func GetBrowsers() *[]BrowserItem {
	Conf := getConf()
	port, err := common.GetFreePort()
	if err != nil {
		log.Fatalln(err)
	}
	defaultUrl := fmt.Sprintf("http://127.0.0.1:%d?b=defaultbrowser", port)
	userAgentMap := make(map[string]string)
	UserAgentServe(port, &userAgentMap)
	userAgentLen := 0

	bItem := []BrowserItem{}
	waitSecond := 5
	timeoutTimer := time.NewTimer(time.Duration(waitSecond) * time.Second)
	defer timeoutTimer.Stop()

	for _, c := range Conf {
		if useAgentCmd, isUserAgent := strings.CutPrefix(c.VersionCmd, "userAgent|"); isUserAgent {
			go getUserAgentVersion(port, useAgentCmd, c.Name)
			userAgentLen++
		} else {
			c.Version = getVersion(c.VersionCmd, c.Reg)
		}
	}

	var wg sync.WaitGroup
	wg.Add(userAgentLen)

	checkUserAgent := func() {
		for {
			select {
			case <-timeoutTimer.C:
				fmt.Printf("仍在获取浏览器信息，您可使用常用浏览器访问地址：%s 以完成浏览器信息采集。\n", defaultUrl)
			default:
				if len(userAgentMap) >= userAgentLen || userAgentMap["defaultbrowser"] != "" {
					for _, v := range Conf {
						if userAgentMap[v.Name] != "" {
							v.Agent = userAgentMap[v.Name]
							v.Version, _ = ExtractChromeVersion(v.Agent)
							bItem = append(bItem, v)
						} else {
							bItem = append(bItem, v)
						}
					}
					wg.Done()
					return
				}
				time.Sleep(2 * time.Second)
			}
		}
	}

	checkUserAgent()
	return &bItem
}

func getVersion(bin, grepArg string) string {
	cmd := fmt.Sprintf("%s | grep -P \"%s\" -m 1 -o", bin, grepArg)
	// log.Println(cmd)
	out, _ := callsrv.ExecGetSysInfoStdout("bash", "-c", cmd)
	// if err != nil {
	// 	log.Println(strings.ReplaceAll(err.Error(), "\r", ""))
	// }
	return strings.Split(string(out), "\n")[0]
}
