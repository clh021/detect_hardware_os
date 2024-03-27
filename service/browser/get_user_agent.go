package browser

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"os/exec"
	"strings"
)

var server *http.Server

func UserAgentServe(port int, Conf *[]BrowserItem, agentGetCnt *int) {
	nameIdxMap := make(map[string]int)
	conf := *Conf
	for i, c := range conf {
		nameIdxMap[c.Name] = i
	}

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/html; charset=utf-8") // 设置响应类型为 HTML
		fmt.Fprintf(w,
			`<!DOCTYPE html><html lang="zh"><head><meta charset="UTF-8"><title>环境采集服务</title></head><body><h2>%s</h2></body></html>`,
			"这里是环境采集服务，目前采集已经完成，您可以关闭该页面。",
		)

		bTag:= r.URL.Query().Get("b")
		if len(bTag) > 0 {
			userAgent := r.Header.Get("User-Agent")
			if nameIdxMap[bTag] > 0 {
				*agentGetCnt++
				log.Println(strAlign(bTag, 20), userAgent)
				conf[nameIdxMap[bTag]].Agent = userAgent
				conf[nameIdxMap[bTag]].KernelVer, _ = regVer(userAgent, conf[nameIdxMap[bTag]].KernelReg)
			}
		}
	})
	server = &http.Server{
		Addr: fmt.Sprintf(":%d", port),
	}
	go func() {
		err := server.ListenAndServe()
		if err != nil && err != http.ErrServerClosed {
			log.Fatal(err)
		}
	}()
}

func strAlign(str string, strMaxlen int) string {
	sLen := len(str)
	if sLen < strMaxlen {
		str += strings.Repeat(" ", strMaxlen-sLen)
	}
	return str
}

func sendUserAgentRequest(port int, b *BrowserItem) (openUrl string, err error) {
	// 判断浏览器命令是否存在
	if _, err := exec.LookPath(b.Bin); err != nil {
		return "", fmt.Errorf("无法找到指定的浏览器程序 '%s': %w", b.Bin, err)
	}
	openUrl = fmt.Sprintf("http://127.0.0.1:%d?b=%s", port, b.Name)
	// log.Printf("openUrl: %s %s", b.Bin, openUrl)
	cmd := exec.Command(b.Bin, openUrl)
	cmd.Env = append(os.Environ(), "DISPLAY=:0")
	err = cmd.Start()
	return
}
