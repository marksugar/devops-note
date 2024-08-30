package main

import (
	"fmt"
	"net"
	"net/http"
	"runtime"
	"sync"
	"time"
)

// 检测HTTP接口是否可用并计算耗时
func checkHTTP(url string, timeout time.Duration) (bool, time.Duration) {
	start := time.Now() // 开始计时

	client := http.Client{
		Timeout: timeout,
	}
	resp, err := client.Get(url)
	elapsed := time.Since(start) // 计算耗时

	if err != nil {
		fmt.Printf("接口 %s 无法访问: %v，耗时: %v\n", url, err, elapsed)
		return false, elapsed
	}
	defer resp.Body.Close()

	if resp.StatusCode == http.StatusOK {
		fmt.Printf("接口 %s 可用 (HTTP 200 OK)，耗时: %v\n", url, elapsed)
		return true, elapsed
	} else {
		fmt.Printf("接口 %s 返回状态码 %d，耗时: %v\n", url, resp.StatusCode, elapsed)
		return false, elapsed
	}
}

// 检测IP和端口的连通性并计算耗时
func checkIPPort(address string, timeout time.Duration, protocol string, wg *sync.WaitGroup, resultChan chan<- Ipaddre, semaphore chan struct{}) (bool, time.Duration) {
	defer wg.Done()
	defer func() { <-semaphore }() // 释放信号量
	start := time.Now()            // 开始计时

	conn, err := net.DialTimeout(protocol, address, timeout)
	elapsed := time.Since(start) // 计算耗时

	if err != nil {
		resultChan <- Ipaddre{IpAddr: address, TimeOut: timeout, Result: fmt.Sprintf("IP:Port %s 无法连接: %v", address, err), TimeWasted: elapsed}
		return false, elapsed
	}
	defer conn.Close()
	resultChan <- Ipaddre{IpAddr: address, TimeOut: timeout, Result: "连接成功", TimeWasted: elapsed}
	return true, elapsed
}

type Ipaddre struct {
	Http       string        `json:"http"`
	IpAddr     string        `json:"ipaddr"`
	Protocol   string        `json:"protocol"`
	TimeOut    time.Duration `json:"timeout"`
	TimeWasted time.Duration `json:"timewasted"`
	Result     string        `json:"result"`
}

func main() {

	// timeout :=   // 设置超时时间
	http := Ipaddre{
		Http:    "http://example.com",
		TimeOut: 5 * time.Second,
	}
	checkHTTP(http.Http, http.TimeOut)

	// IP和端口连通性检测
	var ip []Ipaddre
	ip = append(ip,
		Ipaddre{IpAddr: "172.25.110.31:22992", TimeOut: 5 * time.Second, Protocol: "tcp"},
		Ipaddre{IpAddr: "172.25.110.31:58080", TimeOut: 5 * time.Second, Protocol: "tcp"})

	fmt.Println(GetAvailableCPUs())
	ExplorationIpaddr(ip, GetAvailableCPUs())

}

func ExplorationIpaddr(ipaddre []Ipaddre, concurrencyLimit int) {
	var wg sync.WaitGroup
	resultChan := make(chan Ipaddre, len(ipaddre))

	semaphore := make(chan struct{}, concurrencyLimit) // 信号量，限制并发数量

	for _, ipport := range ipaddre {
		wg.Add(1)
		semaphore <- struct{}{} // 获取信号量
		go checkIPPort(ipport.IpAddr, ipport.TimeOut, ipport.Protocol, &wg, resultChan, semaphore)
	}
	wg.Wait()
	close(resultChan)
	for result := range resultChan {
		fmt.Println(result)
	}
}

// 自动获取一半的CPU数量
func GetAvailableCPUs() int {
	if runtime.NumCPU() == 1 {
		return 1
	}
	return runtime.NumCPU() / 2
}
