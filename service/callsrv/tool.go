package callsrv

import (
	"errors"
	"fmt"
	"path/filepath"
	"strconv"

	"github.com/clh021/detect_hardware_os/service/common"
	"github.com/gogf/gf/v2/os/gfile"
)

func ExecTool(execSet []string, workDir string) (int, error) {
	TTYBin := filepath.Join(common.GetProgramPath(), "tooltty")
	if !gfile.Exists(TTYBin) {
		return 0, errors.New("error: Can't find ToolTTY bin")
	}
	fmt.Printf("ToolExec: %+v \n ", execSet)
	port, err := common.GetFreePort()
	if err != nil {
		return 0, err
	} else {
		execSet = append([]string{"-p", strconv.Itoa(port)}, execSet...)
	}
	go ExecWithSIGINT(TTYBin, workDir, execSet...)
	return port, nil
}
