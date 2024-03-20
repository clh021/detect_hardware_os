package browser

import (
	"fmt"
	"log"
	"net/http"
	"os/exec"
	"sync"
	"time"

	"github.com/clh021/detect_hardware_os/service/version"
)

var userAgentReceived = false
var shutdownServerOnce sync.Once // 用于保证只关闭一次服务器
var server *http.Server

func Main() {
	fmt.Println("启动环境采集服务...")

	version.ShowVersionInfo()

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		userAgent := r.Header.Get("User-Agent")

		w.Header().Set("Content-Type", "text/html; charset=utf-8") // 设置响应类型为 HTML
		fmt.Fprintf(w, `
			<!DOCTYPE html>
			<html lang="en">
			<head>
				<meta charset="UTF-8">
				<title>环境采集服务</title>
			</head>
			<body>
				<h2>%s</h2>
			</body>
			</html>
		`, "这里是环境采集服务，目前采集已经完成，您可以关闭该页面。")
		log.Printf("Received User-Agent: %s", userAgent)

		// 设置标志并尝试关闭服务器
		userAgentReceived = true
		shutdownServer()
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

	cmd := exec.Command("qaxbrowser-safe", "http://localhost:8080")
	err := cmd.Start()
	if err != nil {
		log.Printf("无法启动浏览器: %v", err)
	} else {
		log.Println("已启动浏览器访问服务地址...")
	}

	// 等待浏览器发送请求或者在超时后退出程序
	timeout := time.After(10 * time.Second) // 可以根据实际情况调整超时时间
	for !userAgentReceived {
		select {
		case <-timeout:
			log.Println("等待浏览器请求超时，正在强制关闭服务器...")
			shutdownServer()
			return
		default:
			time.Sleep(100 * time.Millisecond) // 延迟一小段时间再检查
		}
	}
}

func shutdownServer() {
	shutdownServerOnce.Do(func() {
		log.Println("开始关闭服务器...")
		// ctx, cancel := context.WithTimeout(context.Background(), 25*time.Second)
		// defer cancel()

		// err := server.Shutdown(ctx)
		// if err != nil {
		// 	log.Fatalf("服务器关闭失败: %v", err)
		// }
		// log.Println("服务器已成功关闭.")
	})
}
