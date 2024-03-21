package browser

import (
	"fmt"
	"strings"

	"github.com/clh021/detect_hardware_os/service/callsrv"
)

type Conf [4]string

func GetBrowsers() []BrowserItem {
	scripting := []BrowserItem{}
	configs := [...]Conf{
		{"firefox", "Firefox", "firefox --version", `(\d+(\.\d+)*)`},
		{"python", "Python", "python --version 2>&1", `(\d+(\.\d+)*)`},
		{"qianxinbrowser", "奇安信浏览器", "userAgent|qaxbrowser-safe", `Chrome\/(\d+(\.\d+)*)( Safari|$)`},
	}
	for _, c := range configs {
		ver := ""
		if useAgentCmd, isUserAgent := strings.CutPrefix(c[2], "userAgent|"); isUserAgent {
			ver = getUserAgentVersion(useAgentCmd, c[3])
		} else {
			ver = getVersion(c[2], c[3])
		}
		scripting = append(scripting, BrowserItem{
			DisplayName: c[1],
			Name:        c[0],
			Version:     ver,
		})
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
