package src

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func Run(c *gin.Context) {
	resp, err := monitor()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"data": "",
			"err":  err,
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"data": resp,
		})
	}
}

func monitor() (*LinuxResp, error) {
	linux := NewLinux()
	memoryInfo, err := linux.GetMemoryInfo()
	if err != nil {
		return nil, err
	}
	cpuInfo, err := linux.GetCpuInfo()
	if err != nil {
		return nil, err
	}

	diskInfo, err := linux.GetDiskInfo()
	if err != nil {
		return nil, err
	}
	if err != nil {
		return nil, err
	}
	netInfo, err := linux.GetNetInfo()
	if err != nil {
		return nil, err
	}
	proInfo, err := linux.GetProcessInfo()
	if err != nil {
		return nil, err
	}

	resp := &LinuxResp{
		MemoryInfo:  memoryInfo,
		CpuInfo:     cpuInfo,
		DiskInfo:    diskInfo,
		NetInfo:     netInfo,
		ProcessInfo: proInfo,
	}

	return resp, nil
}

type Response struct {
	Value float64 `json:"value"`
}
