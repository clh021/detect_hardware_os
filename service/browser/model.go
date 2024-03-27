package browser

type BrowserItem struct {
	Title     string `json:"title"`
	Name      string `json:"name"`
	Desktop   string `json:"-"`
	IsDefault bool   `json:"is_default"`
	Version   string `json:"version"`
	CmdVer    string `json:"cmd_ver"`
	KernelVer string `json:"kernel_ver"`
	Bin       string `json:"-"`
	Agent     string `json:"agent"`
	KernelReg string `json:"-"`
	CmdReg    string `json:"-"`
}
