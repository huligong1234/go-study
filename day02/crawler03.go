// crawler03.go
// 下载百度图片
package main

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"github.com/bitly/go-simplejson"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"path/filepath"
)

func downloadImage(url string, to string) {
	client := &http.Client{}
	reqest, _ := http.NewRequest("GET", url, nil)
	response, err := client.Do(reqest)
	if err != nil {
		fmt.Println("[-] download failed(%d) : %s\n", response.StatusCode, url)
		return
	}
	content, _ := ioutil.ReadAll(response.Body)
	ioutil.WriteFile(to, content, 0755)
}

func md5Byte(s []byte) string {
	h := md5.New()
	h.Write(s)
	return hex.EncodeToString(h.Sum(nil))
}

func makeFname(basedir, s string) string {
	var extension = filepath.Ext(s)
	a := md5Byte([]byte(s)) + extension
	//c := "./" + basedir + "/" + a[:3] + "/" + a[3:6] + "/" + a[6:9] + "/"
	c := "./" + basedir + "/"
	_, err := os.Stat(c)
	if err != nil && !os.IsExist(err) {
		os.MkdirAll(c, 0755)
	}
	return c + a
}

func main() {
	var url = "http://image.baidu.com/channel/listjson?fr=channel&tag1=%E7%BE%8E%E5%A5%B3&tag2=%E5%85%A8%E9%83%A8&sorttype=0&pn=0&rn=30&ie=utf8&oe=utf-8"
	response, _ := http.Get(url)
	bodyByte, _ := ioutil.ReadAll(response.Body)
	//fmt.Println(string(bodyByte))
	json, err := simplejson.NewJson(bodyByte)
	if err != nil {
		log.Fatal("error:", err)
	}

	data, err := json.Get("data").Array()
	if err != nil {
		log.Fatal("error:", err)
	}
	for _, v := range data {
		vv := v.(map[string]interface{})
		var imgUrl = vv["download_url"].(string)
		var localFile = makeFname("tmp", imgUrl)
		downloadImage(imgUrl, localFile)
	}
}
