package getinfo

import (
	"context"

	"github.com/clh021/detect_hardware_os/service/browser"
)


func Browser(ctx context.Context) (interface{}, error) {
	bs := browser.GetBrowsers()
	return &bs, nil
}