package browser

import (
	"fmt"
	"strings"

	"github.com/clh021/detect_hardware_os/service/callsrv"
)

type devConf [4]string

func GetScripting() []DevItem {
	scripting := []DevItem{}
	configs := [...]devConf{
		{"firefox", "Firefox", "firefox --version", `(\d+\.\d+\.\d+)`},
		{"python", "Python", "python --version", `(\d+\.\d+\.\d+)`},
	}
	for _, c := range configs {
		scripting = append(scripting, DevItem{
			DisplayName: c[1],
			Name:        c[0],
			Version:     getDevVersion(c[2], c[3]),
		})
	}
	return scripting
}

func getDevVersion(bin, grepArg string) string {
	cmd := fmt.Sprintf("%s | grep -P \"%s\" -m 1 -o", bin, grepArg)
	// log.Println(cmd)
	out, _ := callsrv.ExecGetSysInfoStdout("bash", "-c", cmd)
	// if err != nil {
	// 	log.Println(strings.ReplaceAll(err.Error(), "\r", ""))
	// }
	return strings.Split(string(out), "\n")[0]
}
