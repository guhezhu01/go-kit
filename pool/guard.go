package pool

import (
	"fmt"
	"github.com/shirou/gopsutil/cpu"
	"github.com/shirou/gopsutil/mem"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"
)

var cpuPercent float64
var memPercent float64

// 获取cpu使用情况
func getCpuPercent() float64 {
	percent, _ := cpu.Percent(time.Second, false)
	return percent[0]
}

// 获取内存使用情况
func getMemPercent() float64 {
	memInfo, _ := mem.VirtualMemory()
	return memInfo.UsedPercent
}

func init() {
	// 监视内存和cpu使用情况
	go func() {
		quit := make(chan os.Signal)
		signal.Notify(quit, syscall.SIGINT)
		for {
			select {
			case <-quit:
			default:
				cpuPercent = getCpuPercent()
				memPercent = getMemPercent()
				checkCpuUsed(cpuPercent)
				checkMemUsed(memPercent)
				time.Sleep(100)
			}

		}

	}()
}

func checkCpuUsed(maxCpuUsed float64) bool {
	if maxCpuUsed < cpuPercent {
		str := fmt.Sprintf("Cpu exceeds maximum value.Cpu is already %f used", cpuPercent)
		log.Println(str)
		return false
	}
	return true

}

func checkMemUsed(maxMemUsed float64) bool {
	if maxMemUsed < memPercent {
		str := fmt.Sprintf("Memory exceeds maximum value.Memory is already %f used", memPercent)
		log.Println(str)
		return false
	}
	return true
}
