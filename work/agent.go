package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/shirou/gopsutil/cpu"
	"github.com/shirou/gopsutil/mem"
	"net/http"
	"time"
)

type Metric struct {
	Name  string      `json:"name"`
	Value interface{} `json:"value"`
}

func main() {
	// 设置监控服务器地址和端口
	serverURL := "http://localhost:8080/agent"

	for {
		// 采集 CPU 使用率
		cpuUsage, err := cpu.Percent(time.Second, false)
		if err != nil {
			fmt.Println("Failed to get CPU usage:", err)
			continue
		}
		// 采集内存使用情况
		memInfo, err := mem.VirtualMemory()
		if err != nil {
			fmt.Println("Failed to get memory usage:", err)
			continue
		}
		// 构建采集数据
		metrics := []Metric{
			{Name: "cpu_usage", Value: cpuUsage[0]},
			{Name: "mem_usage", Value: memInfo.UsedPercent},
		}
		data, err := json.Marshal(metrics)
		if err != nil {
			fmt.Println("Failed to marshal data:", err)
			continue
		}
		// 发送采集数据到监控服务器
		resp, err := http.Post(serverURL, "application/json", bytes.NewReader(data))
		if err != nil {
			fmt.Println("Failed to send data to server:", err)
			continue
		}
		if resp.StatusCode != http.StatusOK {
			fmt.Println("Server returned an error:", resp.StatusCode)
			continue
		}
		fmt.Println("Data sent to server successfully.")
		// 等待一段时间后进行下一次采集
		time.Sleep(time.Second * 10)
	}
}
