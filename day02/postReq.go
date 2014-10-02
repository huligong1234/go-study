package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

func main() {
	client := &http.Client{}
	reqest, _ := http.NewRequest("GET", "http://127.0.0.1/", nil)
	reqest.Header.Set("Accept", "text/html,application/xhtml+xml,application/xml;q=0.9,*/*;q=0.8")
	reqest.Header.Set("Accept-Charset", "GBK,utf-8;q=0.7,*;q=0.3")
	reqest.Header.Set("Accept-Encoding", "gzip,deflate,sdch")
	reqest.Header.Set("Accept-Language", "zh-CN,zh;q=0.8")
	reqest.Header.Set("Cache-Control", "max-age=0")
	reqest.Header.Set("Connection", "keep-alive")
	reqest.Header.Set("User-Agent", "chrome 100")

	response, _ := client.Do(reqest)
	if response.StatusCode == 200 {
		body, _ := ioutil.ReadAll(response.Body)
		bodystr := string(body)
		fmt.Println(bodystr)
	}

	/*========post request1===========*/
	v := url.Values{}
	v.Set("username", "jack")
	v.Add("password", "123456")
	requestPost, _ := http.NewRequest("POST", "http://127.0.0.1/", strings.NewReader(v.Encode()))
	requestPost.Header.Set("Content-Type", "application/x-www-form-urlencoded; param=value")
	response2, _ := client.Do(requestPost)
	if response2.StatusCode == 200 {
		body, _ := ioutil.ReadAll(response2.Body)
		bodystr := string(body)
		fmt.Println(bodystr)
	}

	/*========post request2===========*/
	response3, _ := http.PostForm("http://127.0.0.1/",
		url.Values{"code": {"admin"}, "pwd": {"123456"},
			"subject": {"study golang"}, "content": {"hello,golang "}})
	if response3.StatusCode == 200 {
		body, _ := ioutil.ReadAll(response3.Body)
		bodystr := string(body)
		fmt.Println(bodystr)
	}
}
