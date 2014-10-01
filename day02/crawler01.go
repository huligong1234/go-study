// crawlerBaiduImage.go
package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {
	response, _ := http.Get("http://www.baidu.com")
	defer response.Body.Close()
	body, _ := ioutil.ReadAll(response.Body)
	fmt.Println(string(body))
}
