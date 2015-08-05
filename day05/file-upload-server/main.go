package main
 
import (
    "fmt"
    "io"
	"crypto/md5"
	"strconv"
	"html/template"
	"time"
	"os"
    "net/http"
)
 

// 处理/upload 逻辑
func upload(w http.ResponseWriter, r *http.Request) {
    fmt.Println("method:", r.Method) //获取请求的方法
    if r.Method == "GET" {
		//fmt.Fprintln(w,"invalid request")
        crutime := time.Now().Unix()
        h := md5.New()
        io.WriteString(h, strconv.FormatInt(crutime, 10))
        token := fmt.Sprintf("%x", h.Sum(nil))
        t, _ := template.ParseFiles("upload.html")
        t.Execute(w, token)
    } else {
        r.ParseMultipartForm(32 << 20)
        file, handler, err := r.FormFile("uploadfile")
        if err != nil {
            fmt.Println(err)
            return
        }
        defer file.Close()
        //fmt.Fprintf(w, "%v", handler.Header)
        f, err := os.OpenFile("d:/"+handler.Filename, os.O_WRONLY|os.O_CREATE, 0666)
        if err != nil {
            fmt.Println(err)
            return
        }
        defer f.Close()
        io.Copy(f, file)
    }
}

func main() {
	http.HandleFunc("/upload", upload)
    //http.Handle("/", http.FileServer(http.Dir("./"))) //文件服务器
    http.ListenAndServe(":9090", nil)
}
