package vo

import (
	"time"
)

// 登录这类返回值
// swagger:response monitorResponse
type MonitorResponse struct {
	//核心数
	CpuNum int
	//CPU使用率
	CpuUsed float64
	//Load Avg 5
	CpuAvg5 float64
	//Load Avg 15
	CpuAvg15 float64
	//总内存
	MemTotal uint64
	//go 内存
	GoTotal uint64
	//已用内存
	MemUsed uint64
	// go用的内存
	GoUsed uint64
	//剩余内存
	MemFree uint64
	//go 内存
	GoFree uint64
	//使用率
	MemUsage float64
	//
	GoUsage float64
	//服务器名称
	SysComputerName string
	//操作系统
	SysOsName string
	//服务器IP
	SysComputerIp string
	//系统架构
	SysOsArch string
	//语言环境
	GoName string
	//版本
	GoVersion string
	//启动时间
	GoStartTime time.Time
	//运行时长
	GoRunTime int64
	//安装路径
	GoHome string
	//项目路径
	GoUserDir string
	//磁盘列表
	Disklist []UsageStat
}

type UsageStat struct {
	Path              string
	Fstype            string
	Total             uint64
	Free              uint64
	Used              uint64
	UsedPercent       float64
	InodesTotal       uint64
	InodesUsed        uint64
	InodesFree        uint64
	InodesUsedPercent float64
}
