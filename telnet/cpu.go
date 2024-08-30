package main

import (
	"fmt"
	"log"
	"runtime"

	"github.com/shirou/gopsutil/cpu"
)

// GetCPUInfo 获取当前运行主机的CPU信息
func GetCPUInfo() ([]cpu.InfoStat, error) {
	// 使用gopsutil库来获取CPU信息
	return cpu.Info()
}

// GetCPUUsage 获取当前CPU使用率
func GetCPUUsage() ([]float64, error) {
	// 获取CPU使用率
	return cpu.Percent(0, true)
}

func main() {
	// 获取CPU信息
	info, err := GetCPUInfo()
	if err != nil {
		log.Fatalf("获取CPU信息失败: %v", err)
	}

	fmt.Printf("CPU信息: %v\n", info)

	// 获取CPU使用率
	usage, err := GetCPUUsage()
	if err != nil {
		log.Fatalf("获取CPU使用率失败: %v", err)
	}

	fmt.Printf("CPU使用率: %v%%\n", usage)

	// 打印当前系统类型
	fmt.Printf("当前操作系统: %v\n", runtime.GOOS)
}
