package browser

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"os/exec"

	"github.com/gogf/gf/v2/frame/g"
)

var server *http.Server

func UserAgentServe(port int, Conf *[]BrowserItem) {
	nameIdxMap := make(map[string]int)
	for i, c := range *Conf {
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
			log.Println("-------------------------------")
			userAgent := r.Header.Get("User-Agent")
			g.Dump(userAgent)
			g.Dump(bTag)
			g.Dump(nameIdxMap[bTag])

			// (*userAgentMap)[queryValue] = userAgent
			// } else {
			// 	log.Println("-------------------------------")
			// 	g.Dump(r)
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

func getUserAgentVersion(port int, b *BrowserItem) (openUrl string, err error) {
	// 判断浏览器命令是否存在
	if _, err := exec.LookPath(b.Bin); err != nil {
		return "", fmt.Errorf("无法找到指定的浏览器程序 '%s': %w", b.Bin, err)
	}
	openUrl = fmt.Sprintf("http://127.0.0.1:%d?b=%s", port, b.Name)
	log.Printf("openUrl: %s %s", b.Bin, openUrl)
	cmd := exec.Command(b.Bin, openUrl)
	cmd.Env = append(os.Environ(), "DISPLAY=:0")
	err = cmd.Start()
	return
}
