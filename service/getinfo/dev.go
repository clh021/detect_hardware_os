package getinfo

import (
	"context"

	"github.com/clh021/detect_hardware_os/service/develop"
)

func Dev(ctx context.Context) (interface{}, error) {
	dev := develop.GetDevelopments()
	return &dev, nil
}
