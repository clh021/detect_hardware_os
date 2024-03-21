package browser

type BrowserItem struct {
	Title      string `json:"title"`
	Name       string `json:"name"`
	Version    string `json:"version"`
	VersionCmd string `json:"-"`
	Agent      string `json:"agent"`
	Reg        string `json:"-"`
}
