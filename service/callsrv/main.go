package callsrv

import (
	"errors"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"syscall"

	"github.com/clh021/detect_hardware_os/service/common"

	"github.com/gogf/gf/v2/os/gfile"
)

// 直接编写命令就可以拿到执行结果
// func ExecCmd(cmd string) (string, error) {
// 	return "", nil
// }

func Exec(pathBin string, args ...string) error {
	return ExecWithWorkDir(pathBin, common.GetProgramPath(), args...)
}

func ExecGetCmdStdoutWithEnv(appendEnv []string, pathBin string, args ...string) ([]byte, error) {
	// env := os.Environ()
	cmd := exec.Command(pathBin, args...)
	// argArr := []string{"-c"}
	// argArr = append(argArr, pathBin)
	// argArr = append(argArr, args...)
	// cmd := exec.Command("bash", argArr...)
	// cmd.SysProcAttr = &syscall.SysProcAttr{
	// 	Pdeathsig: syscall.SIGINT, //如果主进程退出，则将 SIGINT 发送给子进程
	// }
	// cmd.Stdout = os.Stdout
	// cmd.Stderr = os.Stderr
	// cmd.Env = append(env, appendEnv...)
	cmd.Stdin = os.Stdin
	return cmd.Output()
}

func ExecGetSysInfoStdout(pathBin string, args ...string) ([]byte, error) {
	emptyEnv := make([]string, 0)
	return ExecGetCmdStdoutWithEnv(emptyEnv, pathBin, args...)
}

func ExecWithWorkDir(pathBin string, workDir string, args ...string) error {
	if !gfile.Exists(pathBin) {
		return errors.New("error: Can't find " + pathBin)
	}

	env := os.Environ()
	cmd := exec.Command(pathBin, args...)
	cmd.Dir = workDir
	// cmd.SysProcAttr = &syscall.SysProcAttr{
	// 	Pdeathsig: syscall.SIGINT, //如果主进程退出，则将 SIGINT 发送给子进程
	// }
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.Env = env
	cmd.Stdin = os.Stdin
	return cmd.Run()
}

func ExecGetStdout(pathBin string, args ...string) ([]byte, error) {
	if !gfile.Exists(pathBin) {
		return nil, errors.New("error: Can't find " + pathBin)
	}

	env := os.Environ()
	cmd := exec.Command(pathBin, args...)
	// cmd.SysProcAttr = &syscall.SysProcAttr{
	// 	Pdeathsig: syscall.SIGINT, //如果主进程退出，则将 SIGINT 发送给子进程
	// }
	// cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.Env = env
	cmd.Stdin = os.Stdin
	output, err := cmd.Output()
	if err != nil {
		return nil, err
	}
	return output, cmd.Run()
}

func ExecWithSIGINT(pathBin, workDir string, args ...string) error {
	if !gfile.Exists(pathBin) {
		return errors.New("error: Can't find ToolTTY bin")
	}
	env := os.Environ()
	cmd := exec.Command(pathBin, args...)
	cmd.Dir = workDir
	cmd.SysProcAttr = &syscall.SysProcAttr{
		Pdeathsig: syscall.SIGINT, //如果主进程退出，则将 SIGINT 发送给子进程
	}
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.Env = env
	cmd.Stdin = os.Stdin
	return cmd.Run()
}

func ExecScript(scriptFile string, args ...string) ([]byte, error) {
	if !gfile.Exists(scriptFile) {
		return nil, fmt.Errorf("error: Can`t find script file %s", scriptFile)
	}
	cmd := exec.Command(scriptFile, args...)
	return cmd.Output()
}

func WriteCSV(name string) error {
	bin := filepath.Join(common.GetProgramPath(), "csv")
	return Exec(bin, "-f", name)
}

func GetReleaseOutput() ([]byte, error) {
	return exec.Command("sh", "-c", "cat /etc/*release").Output()
}

