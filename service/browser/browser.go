package browser

import (
	"fmt"
	"strings"
	"time"

	"github.com/clh021/detect_hardware_os/service/callsrv"
	"github.com/gogf/gf/v2/frame/g"
)

type Conf [4]string

func GetBrowsers() []BrowserItem {
	scripting := []BrowserItem{}
	configs := [...]Conf{
		{"firefox", "Firefox浏览器", "firefox --version", `(\d+(\.\d+)*)`},
		{"chrome", "Chrome浏览器", "google-chrome --version", `(\d+(\.\d+)*)`},
		{"python", "Python", "python --version 2>&1", `(\d+(\.\d+)*)`},
		{"qianxinbrowser", "奇安信浏览器", "userAgent|qaxbrowser-safe", `Chrome\/(\d+(\.\d+)*)( Safari|$)`},
		// {"360browser", "360安全浏览器", "userAgent|browser360-cn", `Chrome\/(\d+(\.\d+)*)( Safari|$)`},
		// Exec=/usr/bin/brave-browser-stable
		// Exec=/usr/bin/browser360-cn-stable
		// Exec=firefox-esr %u
		// Exec=firefox %u
		// Exec=/usr/bin/google-chrome-stable %U
		// Exec=/opt/apps/htbrowser/htbrowser.sh %U
		// Exec=/usr/bin/lbrowser %U
		// Exec=/usr/bin/microsoft-edge-beta %U
		// Exec=opera %U
		// Exec=midori %U
		// Exec=/usr/bin/qaxbrowser-pioneer-stable %U
		// Exec=/usr/bin/qaxbrowser-safe-stable %U
		// Exec=/usr/bin/vivaldi-stable %U
		// "firefox":         "Mozilla/5.0 (X11; Ubuntu; Linux x86_64; rv:78.0) Gecko/20100101 Firefox/78.0",
		// "qaxbrowser-safe": "Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/102.0.5005.200 Safari/537.36 Qaxbrowser",
		// "edge":            "Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/107.0.0.0 Safari/537.36 Edg/107.0.1418.16",
		// "browser360-cn-stable":      "Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/108.0.5359.125 Safari/537.36",
		// "honglianhua":     "Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/92.0.4515.159 Safari/537.36 htbrowser",
		// "opera":           "Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/106.0.0.0 Safari/537.36 OPR/92.0.0.0",
		// "brave":           "Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/92.0.4515.131 Safari/537.36",
		// "firefox2":        "Mozilla/5.0 (X11; Linux x86_64; rv:52.0) Gecko/20100101 Firefox/52.0",
		// "midori":          "Mozilla/5.0 (X11; Ubuntu; Linux x86_64) AppleWebKit/605.1.15 (KHTML, like Gecko) Version/13.0 Safari/605.1.15 Midori/6",
		// "vivaldi":         "Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/104.0.5112.126 Safari/537.36",
		// "qianxinxianfeng": "Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/110.0.0.0 Safari/537.36 Qaxbrowser",
		// "longxinbrowser":  "Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/114.0.5735.349 Safari/537.36",
	}
	userAgentMap := make(map[string]string)
	UserAgentServe(&userAgentMap)
	for _, c := range configs {
		ver := ""
		if useAgentCmd, isUserAgent := strings.CutPrefix(c[2], "userAgent|"); isUserAgent {
			getUserAgentVersion(useAgentCmd, c[3])
		} else {
			ver = getVersion(c[2], c[3])
		}
		scripting = append(scripting, BrowserItem{
			DisplayName: c[1],
			Name:        c[0],
			Version:     ver,
		})
	}
	for {
		// 使用time.Sleep让程序暂停2秒
		time.Sleep(2 * time.Second)
		g.Dump(userAgentMap)
		// 如果myVar达到某个条件时你想停止打印，可以添加一个退出循环的条件
		if len(userAgentMap) > 20 {
		  break
		}
	}
	return scripting
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
