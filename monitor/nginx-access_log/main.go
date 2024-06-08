package main

import (
	"bufio"
	"fmt"
	"os"
	"time"

	"github.com/satyrius/gonx"
)

// ### nginx log_format
// log_format upstream2  '[$proxy_add_x_forwarded_for]-[$geoip2_data_country_code]-[$geoip2_data_province_name]-[$geoip2_data_city]'
// ' $remote_user [$time_local] "$request" $http_host'
// ' [$body_bytes_sent] $request_body "$http_referer" "$http_user_agent" [$ssl_protocol] [$ssl_cipher]'
// ' [$request_time] [$status] [$upstream_status] [$upstream_response_time] [$upstream_addr]';

func main() {
	// 指定需要监控的Nginx日志文件路径
	filePath := "access.log"

	// 打开日志文件
	file, err := os.Open(filePath)
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	// 获取文件的初始大小
	fileInfo, err := file.Stat()
	if err != nil {
		fmt.Printf("Error getting file info: %v\n", err)
		return
	}
	initialSize := fileInfo.Size()

	// 定义日志格式（根据你的Nginx日志格式进行调整）
	format := `[$proxy_add_x_forwarded_for]-[$geoip2_data_country_code]-[$geoip2_data_province_name]-[$geoip2_data_city] $remote_user [$time_local] "$request" $http_host [$body_bytes_sent] $request_body "$http_referer" "$http_user_agent" [$ssl_protocol] [$ssl_cipher] [$request_time] [$status] [$upstream_status] [$upstream_response_time] [$upstream_addr]`

	// 创建日志解析器
	parser := gonx.NewParser(format)

	// 实时监控日志文件的变化
	for {
		fileInfo, err := file.Stat()
		if err != nil {
			fmt.Printf("Error getting file info: %v\n", err)
			continue
		}

		// 如果文件大小变化，读取新内容
		if fileInfo.Size() > initialSize {
			if _, err := file.Seek(initialSize, 0); err != nil {
				fmt.Printf("Error seeking file: %v\n", err)
				continue
			}

			// 逐行读取新内容
			scanner := bufio.NewScanner(file)
			for scanner.Scan() {
				line := scanner.Text()
				// 解析日志条目
				entry, err := parser.ParseString(line)
				if err != nil {
					fmt.Printf("Error parsing log entry: %v\n", err)
					continue
				}
				fmt.Println("entry:", entry)
				// 获取状态码并进行过滤
				status, err := entry.Field("status")
				if err != nil {
					fmt.Printf("Error getting status field: %v\n", err)
					continue
				}

				upstream_status, err := entry.Field("upstream_status")
				if err != nil {
					fmt.Printf("Error getting upstream_status field: %v\n", err)
					continue
				}
				request_time, err := entry.Field("request_time")
				if err != nil {
					fmt.Printf("Error getting request_time field: %v\n", err)
					continue
				}
				upstream_response_time, err := entry.Field("upstream_response_time")
				if err != nil {
					fmt.Printf("Error getting upstream_response_time field: %v\n", err)
					continue
				}

				if status == "404" || upstream_status == "404" || upstream_response_time > "5" || request_time > "5" {
					fmt.Println(line)
				}
			}

			if err := scanner.Err(); err != nil {
				fmt.Printf("Error reading log file: %v\n", err)
			}

			// 更新已读取的文件大小
			initialSize = fileInfo.Size()
		}

		// 每隔一段时间检查一次文件变化
		time.Sleep(1 * time.Second)
	}
}
