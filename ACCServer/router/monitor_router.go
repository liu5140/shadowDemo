package router

import (
	"fmt"
	"net/http"
	"os"
	"runtime"
	ip2 "shadowDemo/zframework/utils/ip"
	time2 "shadowDemo/zframework/utils/time"

	"strconv"
	"time"

	"shadowDemo/ACCServer/router/vo"

	"github.com/gin-gonic/gin"
	"github.com/shirou/gopsutil/cpu"
	"github.com/shirou/gopsutil/disk"
	"github.com/shirou/gopsutil/host"
	"github.com/shirou/gopsutil/load"
	"github.com/shirou/gopsutil/mem"
)

var StartTime = time.Now()

func monitorRouter(r *gin.RouterGroup) {
	// swagger:route GET /server server monitorServer
	//
	// 获取服务器信息;
	//
	//     Consumes:
	//     - application/json
	//
	//     Produces:
	//     - application/json
	//
	//     Schemes: http
	//
	//     Security:
	//       api_key:
	//       oauth: read, write
	//
	//     Responses:
	//       200: monitorResponse
	//       default: genericError
	r.GET("/server", monitorServer)
}

func monitorServer(c *gin.Context) {
	cpuNum := runtime.NumCPU() //核心数

	var cpuUsed float64 = 0  //用户使用率
	var cpuAvg5 float64 = 0  //CPU负载5
	var cpuAvg15 float64 = 0 //当前空闲率

	cpuInfo, err := cpu.Percent(time.Duration(time.Second), false)
	if err == nil {
		cpuUsed, _ = strconv.ParseFloat(fmt.Sprintf("%.2f", cpuInfo[0]), 64)
	}

	loadInfo, err := load.Avg()
	if err == nil {
		cpuAvg5, _ = strconv.ParseFloat(fmt.Sprintf("%.2f", loadInfo.Load5), 64)
		cpuAvg15, _ = strconv.ParseFloat(fmt.Sprintf("%.2f", loadInfo.Load5), 64)
	}

	var memTotal uint64 = 0  //总内存
	var memUsed uint64 = 0   //总内存  := 0 //已用内存
	var memFree uint64 = 0   //剩余内存
	var memUsage float64 = 0 //使用率

	v, err := mem.VirtualMemory()
	if err == nil {
		memTotal = v.Total / 1024 / 1024
		memFree = v.Free / 1024 / 1024
		memUsed = v.Used / 1024 / 1024
		memUsage, _ = strconv.ParseFloat(fmt.Sprintf("%.2f", v.UsedPercent), 64)
	}

	var goTotal uint64 = 0  //go分配的总内存数
	var goUsed uint64 = 0   //go使用的内存数
	var goFree uint64 = 0   //go剩余的内存数
	var goUsage float64 = 0 //使用率

	var gomem runtime.MemStats
	runtime.ReadMemStats(&gomem)
	goUsed = gomem.Sys / 1024 / 1024

	sysComputerIp := "" //服务器IP

	ip, err := ip2.GetLocalIP()
	if err == nil {
		sysComputerIp = ip
	}

	sysComputerName := "" //服务器名称
	sysOsName := ""       //操作系统
	sysOsArch := ""       //系统架构

	sysInfo, err := host.Info()

	if err == nil {
		sysComputerName = sysInfo.Hostname
		sysOsName = sysInfo.OS
		sysOsArch = sysInfo.KernelArch
	}

	goName := "GoLang"             //语言环境
	goVersion := runtime.Version() //版本
	goStartTime := StartTime       //启动时间

	goRunTime := time2.GetHourDiffer(StartTime.String(), time.Now().String()) //运行时长
	goHome := runtime.GOROOT()                                                //安装路径
	goUserDir := ""                                                           //项目路径

	curDir, err := os.Getwd()

	if err == nil {
		goUserDir = curDir
	}

	//服务器磁盘信息
	disklist := make([]vo.UsageStat, 0)
	diskInfo, err := disk.Partitions(true) //所有分区
	if err == nil {
		for _, p := range diskInfo {
			diskDetail, err := disk.Usage(p.Mountpoint)
			if err == nil {
				diskVo := vo.UsageStat{}
				diskVo.UsedPercent, _ = strconv.ParseFloat(fmt.Sprintf("%.2f", diskDetail.UsedPercent), 64)
				diskVo.Total = diskDetail.Total / 1024 / 1024
				diskVo.Used = diskDetail.Used / 1024 / 1024
				diskVo.Free = diskDetail.Free / 1024 / 1024
				disklist = append(disklist, diskVo)
			}
		}
	}

	c.JSON(http.StatusOK, vo.MonitorResponse{
		CpuNum:          cpuNum,
		CpuUsed:         cpuUsed,
		CpuAvg5:         cpuAvg5,
		CpuAvg15:        cpuAvg15,
		MemTotal:        memTotal,
		GoTotal:         goTotal,
		MemUsed:         memUsed,
		GoUsed:          goUsed,
		MemFree:         memFree,
		GoFree:          goFree,
		MemUsage:        memUsage,
		GoUsage:         goUsage,
		SysComputerName: sysComputerName,
		SysOsName:       sysOsName,
		SysComputerIp:   sysComputerIp,
		SysOsArch:       sysOsArch,
		GoName:          goName,
		GoVersion:       goVersion,
		GoStartTime:     goStartTime,
		GoRunTime:       goRunTime,
		GoHome:          goHome,
		GoUserDir:       goUserDir,
		Disklist:        disklist,
	})
}
