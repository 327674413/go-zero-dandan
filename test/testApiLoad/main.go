package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"sync"
	"time"
)

func main() {
	url := "http://rf.xing-ao.cn/api/ProjectRiffle/getCompanyList"
	concurrency := 100 // 并发数

	var wg sync.WaitGroup
	wg.Add(concurrency)

	start := time.Now()

	var successCount, failureCount int

	for i := 0; i < concurrency; i++ {
		go func() {
			defer wg.Done()

			client := &http.Client{}
			req, err := http.NewRequest("POST", url, nil)
			if err != nil {
				fmt.Println("Error creating HTTP request:", err)
				failureCount++
				return
			}

			resp, err := client.Do(req)
			if err != nil {
				fmt.Println("Error sending HTTP request:", err)
				failureCount++
				return
			}

			defer resp.Body.Close()
			_, err = ioutil.ReadAll(resp.Body)
			if err != nil {
				fmt.Println("Error reading HTTP response:", err)
				failureCount++
				return
			}

			// 假设根据响应的状态码判断请求是否成功
			if resp.StatusCode == http.StatusOK {
				successCount++
			} else {
				failureCount++
			}
		}()
	}

	wg.Wait()

	elapsed := time.Since(start)
	fmt.Println("总请求数:", concurrency)
	fmt.Println("请求总时长:", elapsed)
	fmt.Println("平均每秒处理请求数:", float64(concurrency)/elapsed.Seconds())
	fmt.Println("成功请求数:", successCount)
	fmt.Println("失败请求数", failureCount)
	fmt.Println("并发成功请求率:", float64(successCount)/float64(concurrency)*100, "%")
}
