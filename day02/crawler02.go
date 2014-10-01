// crawler02.go
package main

import (
	"compress/gzip"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
)

func main() {
	client := &http.Client{}
	reqest, _ := http.NewRequest("GET", "http://www.baidu.com", nil)

	reqest.Header.Set("Accept", "text/html,application/xhtml+xml,application/xml;q=0.9,*/*;q=0.8")
	reqest.Header.Set("Accept-Charset", "GBK,utf-8;q=0.7,*;q=0.3")
	reqest.Header.Set("Accept-Encoding", "gzip,deflate,sdch")
	reqest.Header.Set("Accept-Language", "zh-CN,zh;q=0.8")
	reqest.Header.Set("Cache-Control", "max-age=0")
	reqest.Header.Set("Connection", "keep-alive")

	response, _ := client.Do(reqest)
	if response.StatusCode == 200 {
		var body string

		switch response.Header.Get("Content-Encoding") {
		case "gzip":
			reader, _ := gzip.NewReader(response.Body)
			for {
				buf := make([]byte, 1024)
				n, err := reader.Read(buf)

				if err != nil && err != io.EOF {
					panic(err)
				}

				if n == 0 {
					break
				}
				body += string(buf)
			}
		default:
			bodyByte, _ := ioutil.ReadAll(response.Body)
			body = string(bodyByte)
		}

		fmt.Println(body)
	}
}
