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

func UserAgentServe(userAgentMap *map[string]string) {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		userAgent := r.Header.Get("User-Agent")
		w.Header().Set("Content-Type", "text/html; charset=utf-8") // 设置响应类型为 HTML
		fmt.Fprintf(w,
			`<!DOCTYPE html><html lang="zh"><head><meta charset="UTF-8"><title>环境采集服务</title></head><body><h2>%s</h2></body></html>`,
			"这里是环境采集服务，目前采集已经完成，您可以关闭该页面。",
		)
		log.Printf("Received User-Agent: %s", userAgent)

		queryValue := r.URL.Query().Get("b")
		if len(queryValue) > 0 {
			(*userAgentMap)[queryValue] = userAgent
		} else {
			g.Dump(r.URL.Query())
			// log.Fatal("query parameter 'b' not found in request URL")
		}
	})
	server = &http.Server{
		Addr: ":8080",
	}
	go func() {
		err := server.ListenAndServe()
		if err != nil && err != http.ErrServerClosed {
			log.Fatal(err)
		}
	}()
}

func getUserAgentVersion(browserCommand, reg string) error {
	log.Println("browserCommand:", browserCommand)
	log.Println("reg:", reg)
	log.Println("启动环境采集服务...")

	cmd := exec.Command(browserCommand, fmt.Sprintf("http://127.0.0.1:8080?b=%s", browserCommand))
	cmd.Env = append(os.Environ(), "DISPLAY=:0")
	return cmd.Start()
}