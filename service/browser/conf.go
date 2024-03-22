package browser

func getConf() (b []BrowserItem) {
	b = append(b, BrowserItem{
		Title:      "默认浏览器",
		Name:       "defaultbrowser",
		VersionCmd: "",
		Reg:        `(\d+(\.\d+)*)`,
	}, BrowserItem{
		Title:      "Firefox浏览器",
		Name:       "firefox",
		VersionCmd: "firefox --version", // Exec=firefox %u
		Reg:        `(\d+(\.\d+)*)`,     // "firefox2":        "Mozilla/5.0 (X11; Linux x86_64; rv:52.0) Gecko/20100101 Firefox/52.0",
	}, BrowserItem{
		Title:      "Firefox浏览器ESR",
		Name:       "firefox-esr",
		VersionCmd: "firefox-esr --version", // Exec=firefox-esr %u
		Reg:        `(\d+(\.\d+)*)`,         // "firefox":         "Mozilla/5.0 (X11; Ubuntu; Linux x86_64; rv:78.0) Gecko/20100101 Firefox/78.0",
	}, BrowserItem{
		Title:      "Chrome浏览器",
		Name:       "chrome",
		VersionCmd: "google-chrome-stable --version", // Exec=/usr/bin/google-chrome-stable %U
		Reg:        `(\d+(\.\d+)*)`,
	}, BrowserItem{
		Title:      "奇安信浏览器",
		Name:       "qaxbrowser",
		VersionCmd: "userAgent|qaxbrowser-safe-stable", // Exec=/usr/bin/qaxbrowser-safe-stable %U
		Reg:        `Chrome\/(\d+(\.\d+)*)( Safari|$)`, // "qaxbrowser-safe": "Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/102.0.5005.200 Safari/537.36 Qaxbrowser",
	}, BrowserItem{
		Title:      "奇安信先锋浏览器",
		Name:       "qaxbrowser-pioneer",
		VersionCmd: "userAgent|qaxbrowser-pioneer-stable", // Exec=/usr/bin/qaxbrowser-pioneer-stable %U
		Reg:        `Chrome\/(\d+(\.\d+)*)( Safari|$)`,    // "qianxinxianfeng": "Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/110.0.0.0 Safari/537.36 Qaxbrowser",
	}, BrowserItem{
		Title:      "360安全浏览器",
		Name:       "360browser",
		VersionCmd: "userAgent|browser360-cn-stable",          // Exec=/usr/bin/browser360-cn-stable
		Reg:        `Chrome\/(\d+(\.\d+)*)( Safari|$)`, // "browser360-cn-stable":      "Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/108.0.5359.125 Safari/537.36",
	}, BrowserItem{
		Title:      "Opera浏览器",
		Name:       "opera",
		VersionCmd: "userAgent|opera",                  // Exec=opera %U
		Reg:        `Chrome\/(\d+(\.\d+)*)( Safari|$)`, // "opera":           "Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/106.0.0.0 Safari/537.36 OPR/92.0.0.0",
	}, BrowserItem{
		Title:      "龙芯浏览器",
		Name:       "lbrowser",
		VersionCmd: "userAgent|lbrowser",               // Exec=/usr/bin/lbrowser %U
		Reg:        `Chrome\/(\d+(\.\d+)*)( Safari|$)`, // "longxinbrowser":  "Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/114.0.5735.349 Safari/537.36",
	}, BrowserItem{
		Title:      "Edge浏览器",
		Name:       "edge",
		VersionCmd: "userAgent|microsoft-edge-beta",    // Exec=/usr/bin/microsoft-edge-beta %U
		Reg:        `Chrome\/(\d+(\.\d+)*)( Safari|$)`, // "edge":            "Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/107.0.0.0 Safari/537.36 Edg/107.0.1418.16",
	}, BrowserItem{
		Title:      "Brave浏览器",
		Name:       "bravebrowser",
		VersionCmd: "userAgent|brave-browser-stable",   // Exec=/usr/bin/brave-browser-stable
		Reg:        `Chrome\/(\d+(\.\d+)*)( Safari|$)`, // "brave":           "Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/92.0.4515.131 Safari/537.36",
	// }, BrowserItem{
	// 	Title:      "Modori浏览器",
	// 	Name:       "modori",
	// 	VersionCmd: "userAgent|midori",                 // Exec=midori %U
	// 	Reg:        `Chrome\/(\d+(\.\d+)*)( Safari|$)`, // "midori":          "Mozilla/5.0 (X11; Ubuntu; Linux x86_64) AppleWebKit/605.1.15 (KHTML, like Gecko) Version/13.0 Safari/605.1.15 Midori/6",
	}, BrowserItem{
		Title:      "Vivaldi浏览器",
		Name:       "vivaldi",
		VersionCmd: "userAgent|vivalid-stable",         // Exec=/usr/bin/vivaldi-stable %U
		Reg:        `Chrome\/(\d+(\.\d+)*)( Safari|$)`, // "vivaldi":         "Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/104.0.5112.126 Safari/537.36",
	}, BrowserItem{
		Title:      "红莲花浏览器",
		Name:       "honglianhua",
		VersionCmd: "userAgent|/opt/apps/htbrowser/htbrowser.sh", // Exec=/opt/apps/htbrowser/htbrowser.sh %U
		Reg:        `Chrome\/(\d+(\.\d+)*)( Safari|$)`,           // "honglianhua":     "Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/92.0.4515.159 Safari/537.36 htbrowser",
	},
	)
	return
}
