package browser

import (
	"fmt"
	"regexp"
)

func ExtractChromeVersion(userAgent string) (string, error) {
	// 正则表达式匹配任意数量的 . 分隔的数字序列，直到 " Safari" 或字符串结束
	pattern := regexp.MustCompile(`Chrome\/(\d+(\.\d+)*)( Safari|$)`)

	// 查找匹配项
	matches := pattern.FindStringSubmatch(userAgent)
	if len(matches) < 2 {
		return "", fmt.Errorf("无法从 User-Agent 中提取 Chrome 版本号")
	}

	// g.Dump(matches)
	// 提取版本号（matches[1] 是第一个括号内的匹配内容）
	chromeVersion := matches[1]
	return chromeVersion, nil
}

func TestExtractChromeVersion() {
	userAgent := "User-Agent: Mozilla/5.0 (X11; Linux aarch64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/90.0.4430.246 Safari/537.36 Qaxbrowser"

	version, err := ExtractChromeVersion(userAgent)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("Chrome 版本号为: %s\n", version)
	}
}
