# detect_hardware_os
detect hardware info and os info

# Usage
```bash
go get github.com/clh021/detect_hardware_os@v0.0.3
```

# Collection List

## Model
```golang
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
type DevItem struct {
	DisplayName string `json:"displayName"`
	Name        string `json:"name"`
	Version     string `json:"version"`
}
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

si.getMetaInfo()

// DMI info
si.getProductInfo()
type Product struct {
	Name    string    `json:"name,omitempty"`
	Vendor  string    `json:"vendor,omitempty"`
	Version string    `json:"version,omitempty"`
	Serial  string    `json:"serial,omitempty"`
	UUID    uuid.UUID `json:"uuid,omitempty"`
	SKU     string    `json:"sku,omitempty"`
}
si.getBoardInfo()
type Board struct {
	Name     string `json:"name,omitempty"`
	Vendor   string `json:"vendor,omitempty"`
	Version  string `json:"version,omitempty"`
	Serial   string `json:"serial,omitempty"`
	AssetTag string `json:"assettag,omitempty"`
}
si.getChassisInfo()
type Chassis struct {
	Type     uint   `json:"type,omitempty"`
	Vendor   string `json:"vendor,omitempty"`
	Version  string `json:"version,omitempty"`
	Serial   string `json:"serial,omitempty"`
	AssetTag string `json:"assettag,omitempty"`
}
si.getBIOSInfo()
type BIOS struct {
	Vendor  string `json:"vendor,omitempty"`
	Version string `json:"version,omitempty"`
	Date    string `json:"date,omitempty"`
}

// SMBIOS info
si.getMemoryInfo()
type Memory struct {
	Type  string `json:"type,omitempty"`
	Speed uint   `json:"speed,omitempty"` // RAM data rate in MT/s
	Size  uint   `json:"size,omitempty"`  // RAM size in MB
}

// Node info
si.getNodeInfo() // depends on BIOS info
type Node struct {
	Hostname   string `json:"hostname,omitempty"`
	MachineID  string `json:"machineid,omitempty"`
	Hypervisor string `json:"hypervisor,omitempty"`
	Timezone   string `json:"timezone,omitempty"`
}

// Hardware info
si.getCPUInfo() // depends on Node info
type CPU struct {
	Vendor  string `json:"vendor,omitempty"`
	Model   string `json:"model,omitempty"`
	Speed   uint   `json:"speed,omitempty"`   // CPU clock rate in MHz
	Cache   uint   `json:"cache,omitempty"`   // CPU cache size in KB
	Cpus    uint   `json:"cpus,omitempty"`    // number of physical CPUs
	Cores   uint   `json:"cores,omitempty"`   // number of physical CPU cores
	Threads uint   `json:"threads,omitempty"` // number of logical (HT) CPU cores
}
si.getStorageInfo()
type StorageDevice struct {
	Name   string `json:"name,omitempty"`
	Driver string `json:"driver,omitempty"`
	Vendor string `json:"vendor,omitempty"`
	Model  string `json:"model,omitempty"`
	Serial string `json:"serial,omitempty"`
	Size   uint   `json:"size,omitempty"` // device size in GB
}
si.getNetworkInfo()
type NetworkDevice struct {
	Name       string `json:"name,omitempty"`
	Driver     string `json:"driver,omitempty"`
	MACAddress string `json:"macaddress,omitempty"`
	Port       string `json:"port,omitempty"`
	Speed      uint   `json:"speed,omitempty"` // device max supported speed in Mbps
}

// Software info
si.getOSInfo()
type OS struct {
	Name         string `json:"name,omitempty"`
	Vendor       string `json:"vendor,omitempty"`
	Version      string `json:"version,omitempty"`
	Release      string `json:"release,omitempty"`
	Architecture string `json:"architecture,omitempty"`
}
si.getKernelInfo()
type Kernel struct {
	Release      string `json:"release,omitempty"`
	Version      string `json:"version,omitempty"`
	Architecture string `json:"architecture,omitempty"`
}


```

## Table

### 浏览器环境信息

支持采集目标系统环境中安装的几乎所有浏览器，采集浏览器字段信息如下:

浏览器名称,是否为默认浏览器,浏览器版本,内核版本,浏览器UserAgent

### 系统环境信息

#### DMI 信息

Product信息: 名称,厂商,版本,序列号,UUID,SKU
Board信息: 名称,厂商,版本,序列号,资产标签
Chassis信息: 类型,厂商,版本,序列号,资产标签
BIOS信息: 厂商,版本,日期

#### SMBIOS 信息

网络设备信息: 名称,驱动,MAC地址,端口,最大传输速率
操作系统信息: 名称,厂商,版本,发行版本,架构
系统内核信息: 版本,发行版本,架构

### 开发环境信息

支持以下开发语言环境信息采集:

gcc, clang, dmd, gbc3, java, csharp_old, csharp, vala, haskell, pascal, go, rust

采集字段信息如下:

当前语言环境软件是否安装，当前安装版本号

### 硬件环境信息

CPU信息: 厂商,型号,速度,缓存大小,物理CPU,物理CPU核心,逻辑CPU核心
Memory信息: 类型,速度,大小
存储设备信息: 名称,驱动,厂商,型号,序列号,大小
Node信息: 主机名,机器ID,虚拟化类型,时区
