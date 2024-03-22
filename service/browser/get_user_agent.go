package browser

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"os/exec"
)

var server *http.Server

func UserAgentServe(port int, userAgentMap *map[string]string) {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		userAgent := r.Header.Get("User-Agent")
		w.Header().Set("Content-Type", "text/html; charset=utf-8") // 设置响应类型为 HTML
		fmt.Fprintf(w,
			`<!DOCTYPE html><html lang="zh"><head><meta charset="UTF-8"><title>环境采集服务</title></head><body><h2>%s</h2></body></html>`,
			"这里是环境采集服务，目前采集已经完成，您可以关闭该页面。",
		)
		// log.Printf("Received User-Agent: %s", userAgent)

		queryValue := r.URL.Query().Get("b")
		if len(queryValue) > 0 {
			(*userAgentMap)[queryValue] = userAgent
		} else {
			log.Println("Debug:Query:", r.URL.Query())
			// log.Fatal("query parameter 'b' not found in request URL")
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

func getUserAgentVersion(port int, browserCommand, name string) (openUrl string, err error) {
	// 判断浏览器命令是否存在
	if _, err := exec.LookPath(browserCommand); err != nil {
		return "", fmt.Errorf("无法找到指定的浏览器程序 '%s': %w", browserCommand, err)
	}
	openUrl = fmt.Sprintf("http://127.0.0.1:%d?b=%s", port, name)
	// log.Printf("%s %s", browserCommand, openUrl)
	cmd := exec.Command(browserCommand, openUrl)
	cmd.Env = append(os.Environ(), "DISPLAY=:0")
	err = cmd.Start()
	return
}