package develop

import (
	"fmt"
	"strings"

	"github.com/clh021/detect_hardware_os/service/callsrv"
)

type developments struct {
	Compilers []DevItem `json:"compilers"`
	Scripting []DevItem `json:"scripting"`
	Tools     []DevItem `json:"tools"`
}

type devConf [4]string

func GetDevelopments() *developments {
	return &developments{
		Compilers: GetCompilers(),
		Scripting: GetScripting(),
		Tools:     GetTools(),
	}
}

func GetCompilers() []DevItem {
	compilers := []DevItem{}
	configs := [...]devConf{
		{"gcc", "C (GCC)", "gcc --version 2>&1", `(\d+\.\d+\.\d+)`},
		{"clang", "C (Clang)", "clang -v", `(\d+\.\d+)`},
		{"dmd", "D (dmd)", "dmd --help", `(\d+\.\d+)`},
		{"gbc3", "Gambas3 (gbc3)", "gbc3 --version", `(\d+\.\d+\.\d+)`},
		{"java", "Java", "javac -version", `(\d+\.\d+\.\d+)`},
		{"csharp_old", "CSharp (Mono, old)", "mcs --version", `(\d+\.\d+\.\d+\.\d+)`},
		{"csharp", "CSharp (Mono)", "gmcs --version", `(\d+\.\d+\.\d+\.\d+)`},
		{"vala", "Vala", "valac --version", `(\d+\.\d+\.\d+)`},
		{"haskell", "Haskell (GHC)", "ghc -v", `(\d+\.\d+\.\d+)`},
		{"pascal", "FreePascal", "fpc -iV", `(\d+\.\d+\.?\d*)`},
		{"go", "Go", "go version", `(\d+\.\d+\.?\d* )`},
		{"rust", "Rust", "rustc --version", `(\d+\.\d+\.?\d* )`},
	}
	for _, c := range configs {
		compilers = append(compilers, DevItem{
			DisplayName: c[1],
			Name:        c[0],
			Version:     getDevVersion(c[2], c[3]),
		})
	}
	return compilers
}

func GetScripting() []DevItem {
	scripting := []DevItem{}
	configs := [...]devConf{
		{"gbr3", "Gambas3 (gbr3)", "gbr3 --version", `(\d+\.\d+\.\d+)`},
		{"python", "Python", "python --version 2>&1", `(\d+\.\d+\.\d+)`},
		{"python2", "Python2", "python2 --version 2>&1", `(\d+\.\d+\.\d+)`},
		{"python3", "Python3", "python3 --version", `(\d+\.\d+\.\d+)`},
		{"perl", "Perl", "perl --version", `(\d+\.\d+\.\d+)`},
		{"perl6", "Perl6 (VM)", "perl6 --version", `(\d+\.\d+\.\d+)`},
		//{"perl6_vm", "Perl6", "perl6 --version", `((?<=implementing Perl )\w*\.\w*)`},
		{"php", "PHP", "php --version", `(\d+\.\d+\.\S+)`},
		{"ruby", "Ruby", "ruby --version", `(\d+\.\d+\.\d+)`},
		{"bash", "Bash", "bash --version", `(\d+\.\d+\.\S+)`},
	}
	for _, c := range configs {
		scripting = append(scripting, DevItem{
			DisplayName: c[1],
			Name:        c[0],
			Version:     getDevVersion(c[2], c[3]),
		})
	}
	return scripting
}

func getDevVersion(bin, grepArg string) string {
	cmd := fmt.Sprintf("%s | grep -P \"%s\" -m 1 -o", bin, grepArg)
	// log.Println(cmd)
	out, _ := callsrv.ExecGetSysInfoStdout("bash", "-c", cmd)
	// if err != nil {
	// 	log.Println(strings.ReplaceAll(err.Error(), "\r", ""))
	// }
	return strings.Split(string(out), "\n")[0]
}

func GetTools() []DevItem {
	tools := []DevItem{}
	configs := [...]devConf{
		{"make", "make", "make --version", `(\d+\.\d+)`},
		{"gdb", "GDB", "gdb --version", "(?<=^GNU gdb ).*"},
		{"strace", "strace", "strace -V", `(\d+\.\d+\.?\d*)`},
		{"valgrind", "valgrind", "valgrind --version", `(\d+\.\d+\.\S+)`},
		{"qmake", "QMake", "qmake --version", `(\d+\.\S+)`},
		{"cmake", "CMake", "cmake --version", `(\d+\.\d+\.?\d*)`},
		{"gambas3", "Gambas3 IDE", "gambas3 --version", `(\d+\.\d+\.\d+)`},
	}
	for _, c := range configs {
		tools = append(tools, DevItem{
			DisplayName: c[1],
			Name:        c[0],
			Version:     getDevVersion(c[2], c[3]),
		})
	}
	return tools
}
