Go Study
=====

### Install File
* http://golang.so/dl/go1.3.windows-amd64.msi

* http://golang.so/dl/go1.3.windows-386.msi

* http://golang.so/dl/go1.3.linux-amd64.tar.gz

* http://golang.so/dl/go1.3.linux-386.tar.gz


### CentOS6.5上golang环境配置


* 下载和解压go环境包

>>cd /usr/local/src/

>>wget -c http://golangtc.com/static/go/go1.4beta1.linux-amd64.tar.gz

>>tar zxvf go1.4beta1.linux-amd64.tar.gz -C /usr/local


* 设置系统环境变量

>>vi /etc/profile

export GOROOT=/usr/local/go

export GOBIN=$GOROOT/bin

export PATH=$PATH:$GOBIN

* 编译，使其生效

>>source /etc/profile  


* 验证，查看是否配置成功

>>go version  


-----
 
## IDE
* notepad++/sublime Text/.. + plugin
* goeclipse
* LiteIDE

-----
## Study Resources

http://studygolang.com/

### 入门资料

http://weekly.golang.org/ 推荐阅读该资料

http://tour.golang.org/#1

http://golang.org/doc/effective_go.html

http://golang.org/doc/go_spec.html

### 文档

http://golang.org/pkg/

内存模型 http://golang.org/doc/go_mem.html

Go 的Web开发介绍
http://golang.org/doc/codelab/wiki/

最后的Try it命令为
$ go build wiki.go


http://www.youtube.com/watch?v=-i0hat7pdpk
http://code.google.com/intl/zh-CN/appengine/docs/go/

### 其他资料
http://godashboard.appspot.com/project


参考书籍：《学习Go语言》、《Go语言编程》


### 资料：
1、Go语言半小时速成教程
http://www.vaikan.com/go/a-tour-of-go/#1

2、Go语言豆瓣小组
http://www.douban.com/group/topic/9766700/

3、Go语言(golang)开源项目大全
http://www.open-open.com/lib/view/open1396063913278.html

4、博客园上关于Go语言的主题
http://news.cnblogs.com/n/topic_278.htm

5、Go语言教程
http://www.yiibai.com/go/

6、golang中文社区
http://studygolang.com/

7、国内镜像

Go 指南国内镜像  http://tour.golangtc.com/welcome/1

Go 语言国内下载镜像   http://www.golangtc.com/download

Go 官方网站国内镜像  http://docs.studygolang.com/

8、Go Web开发框架

beego：http://beego.me/

revel：http://gorevel.cn/

9、Go名人

#### astaxie：https://github.com/astaxie 
* 《Go实战开发》 https://github.com/astaxie/Go-in-Action
* 《Go Web 编程》https://github.com/astaxie/build-web-application-with-golang

 
#### Unknwon：https://github.com/Unknwon
* 《Go入门指南》 https://github.com/Unknwon/the-way-to-go_ZH_CN
* 《Go编程基础》 https://github.com/Unknwon/go-fundamental-programming
* 《Go Web基础》 https://github.com/Unknwon/go-web-foundation
* 《Go名库讲解》 https://github.com/Unknwon/go-rock-libraries-showcases
