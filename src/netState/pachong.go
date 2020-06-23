package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {
	// 生成client客户端
	client := &http.Client{}
	// 生成Request对象
	req, err := http.NewRequest("Get", "https://www.baidu.com", nil)
	if err != nil {
		fmt.Println(err)
	}
	// 添加Header
	req.Header.Add("User-Agent", "Mozilla/5.0 (Windows NT 10.0; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/69.0.3497.100 Safari/537.36")
	// 发起请求
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
	}
	// 设定关闭响应体
	defer resp.Body.Close()
	// 读取响应体
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(body))
}
