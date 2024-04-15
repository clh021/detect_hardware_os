package browser

func getConf() (b []BrowserItem) {
	b = append(b, BrowserItem{
		Title:     "Firefox浏览器",
		Name:      "firefox",
		Desktop:   "firefox.desktop",
		IsDefault: false,
		Bin:       "firefox",                 // Exec=firefox %u
		KernelReg: `Firefox\/(\d+(\.\d+)*)$`, // "firefox2":        "Mozilla/5.0 (X11; Linux x86_64; rv:52.0) Gecko/20100101 Firefox/52.0",
		CmdReg:    `(\d+(\.\d+)*)`,
	}, BrowserItem{
		Title:     "Firefox浏览器ESR",
		Name:      "firefox-esr",
		Desktop:   "firefox-esr.desktop",
		IsDefault: false,
		Bin:       "firefox-esr",             // Exec=firefox-esr %u
		KernelReg: `Firefox\/(\d+(\.\d+)*)$`, // "firefox":         "Mozilla/5.0 (X11; Ubuntu; Linux x86_64; rv:78.0) Gecko/20100101 Firefox/78.0",
		CmdReg:    `(\d+(\.\d+)*)`,
	}, BrowserItem{
		Title:     "Chrome浏览器",
		Name:      "chrome",
		Desktop:   "google-chrome.desktop",
		IsDefault: false,
		Bin:       "google-chrome-stable",             // Exec=/usr/bin/google-chrome-stable %U
		KernelReg: `Chrome\/(\d+(\.\d+)*)( Safari|$)`, // "Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/116.0.0.0 Safari/537.36",
		CmdReg:    `(\d+(\.\d+)*)`,
	}, BrowserItem{
		Title:     "奇安信浏览器",
		Name:      "qaxbrowser",
		Desktop:   "qaxbrowser-safe.desktop",
		IsDefault: false,
		Bin:       "qaxbrowser-safe-stable",           // Exec=/usr/bin/qaxbrowser-safe-stable %U
		KernelReg: `Chrome\/(\d+(\.\d+)*)( Safari|$)`, // "qaxbrowser-safe": "Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/102.0.5005.200 Safari/537.36 Qaxbrowser",
		CmdReg:    `(\d+(\.\d+)*)`,
	}, BrowserItem{
		Title:     "奇安信先锋浏览器",
		Name:      "qaxbrowser-pioneer",
		Desktop:   "qaxbrowser-pioneer.desktop",
		IsDefault: false,
		Bin:       "qaxbrowser-pioneer-stable",        // Exec=/usr/bin/qaxbrowser-pioneer-stable %U
		KernelReg: `Chrome\/(\d+(\.\d+)*)( Safari|$)`, // "qianxinxianfeng": "Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/110.0.0.0 Safari/537.36 Qaxbrowser",
		CmdReg:    `(\d+(\.\d+)*)`,
	}, BrowserItem{
		Title:     "360安全浏览器",
		Name:      "360browser",
		Desktop:   "browser360-cn.desktop",
		IsDefault: false,
		Bin:       "browser360-cn-stable",             // Exec=/usr/bin/browser360-cn-stable
		KernelReg: `Chrome\/(\d+(\.\d+)*)( Safari|$)`, // "browser360-cn-stable":      "Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/108.0.5359.125 Safari/537.36",
		CmdReg:    `(\d+(\.\d+)*)`,
	}, BrowserItem{
		Title:     "Opera浏览器",
		Name:      "opera",
		Desktop:   "opera.desktop",
		IsDefault: false,
		Bin:       "opera",                            // Exec=opera %U
		KernelReg: `Chrome\/(\d+(\.\d+)*)( Safari|$)`, // "opera":           "Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/106.0.0.0 Safari/537.36 OPR/92.0.0.0",
		CmdReg:    `(\d+(\.\d+)*)`,
	}, BrowserItem{
		Title:     "龙芯浏览器",
		Name:      "lbrowser",
		Desktop:   "lbrowser.desktop",
		IsDefault: false,
		Bin:       "lbrowser-stable",                  // Exec=/usr/bin/lbrowser %U
		KernelReg: `Chrome\/(\d+(\.\d+)*)( Safari|$)`, // "longxinbrowser":  "Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/114.0.5735.349 Safari/537.36",
		CmdReg:    `(\d+(\.\d+)*)`,
	}, BrowserItem{
		Title:     "Edge浏览器",
		Name:      "edge",
		Desktop:   "microsoft-edge-beta.desktop",
		IsDefault: false,
		Bin:       "microsoft-edge-beta",              // Exec=/usr/bin/microsoft-edge-beta %U
		KernelReg: `Chrome\/(\d+(\.\d+)*)( Safari|$)`, // "edge":            "Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/107.0.0.0 Safari/537.36 Edg/107.0.1418.16",
		CmdReg:    `(\d+(\.\d+)*)`,
	}, BrowserItem{
		Title:     "Brave浏览器",
		Name:      "bravebrowser",
		Desktop:   "brave-browser.desktop",
		IsDefault: false,
		Bin:       "brave-browser-stable",             // Exec=/usr/bin/brave-browser-stable
		KernelReg: `Chrome\/(\d+(\.\d+)*)( Safari|$)`, // "brave":           "Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/92.0.4515.131 Safari/537.36",
		CmdReg:    `(\d+(\.\d+)*)`,
		// }, BrowserItem{
		// 	Title:      "Modori浏览器",
		// 	Name:       "modori",
		// 	VersionCmd: "midori",                 // Exec=midori %U
		// 	KernelReg:        `Chrome\/(\d+(\.\d+)*)( Safari|$)`, // "midori":          "Mozilla/5.0 (X11; Ubuntu; Linux x86_64) AppleWebKit/605.1.15 (KHTML, like Gecko) Version/13.0 Safari/605.1.15 Midori/6",
		//  CmdReg:    `(\d+(\.\d+)*)`,
	}, BrowserItem{
		Title:     "Vivaldi浏览器",
		Name:      "vivaldi",
		Desktop:   "vivaldi-stable.desktop",
		IsDefault: false,
		Bin:       "vivaldi-stable",                   // Exec=/usr/bin/vivaldi-stable %U
		KernelReg: `Chrome\/(\d+(\.\d+)*)( Safari|$)`, // "vivaldi":         "Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/104.0.5112.126 Safari/537.36",
		CmdReg:    `(\d+(\.\d+)*)`,
	}, BrowserItem{
		Title:     "红莲花安全浏览器",
		Name:      "honglianhua",
		Desktop:   "htbrowser.desktop",
		IsDefault: false,
		Bin:       "/opt/apps/htbrowser/htbrowser.sh", // Exec=/opt/apps/htbrowser/htbrowser.sh %U
		KernelReg: `Chrome\/(\d+(\.\d+)*)( Safari|$)`, // "honglianhua":     "Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/92.0.4515.159 Safari/537.36 htbrowser",
		CmdReg:    `(\d+(\.\d+)*)`,
	},
	)
	return
}
