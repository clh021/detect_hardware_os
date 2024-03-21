package browser

import (
	"fmt"
	"log"
	"strings"
	"time"

	"github.com/clh021/detect_hardware_os/service/callsrv"
	"github.com/clh021/detect_hardware_os/service/common"
)

type Conf [4]string

func GetBrowsers() []BrowserItem {
	Conf := getConf()
	port, err := common.GetFreePort()
	if err != nil {
		log.Fatalln(err)
	}
	userAgentMap := make(map[string]string)
	UserAgentServe(port, &userAgentMap)
	userAgentLen := 0
	for _, c := range Conf {
		if useAgentCmd, isUserAgent := strings.CutPrefix(c.VersionCmd, "userAgent|"); isUserAgent {
			userAgentLen++
			url, err := getUserAgentVersion(port, useAgentCmd, c.Name)
			if err != nil {
				log.Println("系统如法自动获取浏览器版本，请手动协助采集，打开浏览器，访问地址:", url)
				log.Fatalln(err)
			}
		} else {
			c.Version = getVersion(c.VersionCmd, c.Reg)
		}
	}
	bItem := []BrowserItem{}
	for {
		time.Sleep(2 * time.Second)
		// g.Dump(userAgentMap)
		if len(userAgentMap) >= userAgentLen {
			for _, v := range Conf {
				if strings.HasPrefix(v.VersionCmd, "userAgent|") {
					if userAgentMap[v.Name] != "" {
						v.Agent = userAgentMap[v.Name]
						v.Version, err = ExtractChromeVersion(v.Agent)
						if err != nil {
							fmt.Println(err)
						}
						bItem = append(bItem, v)
					}
				} else {
					bItem = append(bItem, v)
				}
			}
		  break
		}
	}
	return bItem
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
