package main

//导入依赖包
//使用【import _ 包路径】只是引用该包，仅仅是为了调用init()函数，所以无法通过包名来调用包中的其他函数
import (
	_ "github.com/go-sql-driver/mysql"
	"html/template"
	"log"
	"net/http"
)

//用const定义常量
const (
	username = "root"
	password = "123456"
	dbName   = "test"
)

// 定义结构体，类似Java中的POJO，或Django中的models.py
type App struct {
	ID      string
	AppCode string
	AppName string
}

func (app *App) toString() string {
	return "appId:" + app.ID + ";appCode:" + app.AppCode + ";appName:" + app.AppName
}

/*
golang中的interface类似于java的interface、PHP的interface或C++的纯虚基类。
接口就是一个协议，规定了一组成员。
*/

// 简要封装定义页面渲染函数，函数定义用func关键字
func render(w http.ResponseWriter, tplName string, context map[string]interface{}) {
	tpl, err := template.ParseFiles(tplName)
	/*
		golang的nil在概念上和其它语言的null、None、nil、NULL一样，都指代零值或空值。
		nil是预先说明的标识符，也即通常意义上的关键字。
		在golang中，nil只能赋值给指针、channel、func、interface、map或slice类型的变量。
		如果未遵循这个规则，则会引发panic。对此官方有明确的说明：http://pkg.golang.org/pkg/builtin/#Type
	*/
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	tpl.Execute(w, context)
	return
}

//定义请求处理器,类似Struts中的Action或SpringMVC中的Controller类,或Django中的views.py
func indexHandler(w http.ResponseWriter, r *http.Request) {
	//创建数据连接对象
	db, err := getDB(username, password, dbName)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	//数据表查询
	rows, err := db.Query("SELECT id, app_code,app_name FROM t_app")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	/*
		关键字defer 允许我们推迟到函数返回之前（或任意位置执行 return 语句之后）一刻才执行某个语句或函数
		defer的用法类似于面向对象编程语言 Java 和 C# 的 finally 语句块，它一般用于释放某些已分配的资源。
	*/
	defer rows.Close()

	//定义变量
	var id, app_code, app_name string

	/*
		Golang中可以使用指针，并提供了两种内存分配机制：
		new：分配长度为0的空白内存，返回类型T*。
		make：仅用于 切片、map、chan消息管道，返回类型T而不是指针。

		map 是引用类型，可以使用如下声明：
		var map1 map[keytype]valuetype
		var map1 map[string]int

		在声明的时候不需要知道 map 的长度，map 是可以动态增长的。

		未初始化的 map 的值是 nil。
		key 可以是任意可以用 == 或者 != 操作符比较的类型，比如 string、int、float。
		所以数组、切片和结构体不能作为 key，但是指针和接口类型可以。

		如果要用结构体作为 key 可以提供 Key() 和 Hash() 方法，
		这样可以通过结构体的域计算出唯一的数字或者字符串的 key。

		value 可以是任意类型的；通过使用空接口类型(interface{})，我们可以存储任意值，但是使用这种类型作为值时需要先做一次类型断言。

		不要使用 new，永远用 make 来构造 map
	*/

	/*
		定义变量locals，类型为map,类似Java中 Map<String,Object> locals = new HashMap<String,Object>();
		这里期望作用类似Django里面的locals(),或SpringMVC中Controller里面的Model
	*/
	locals := make(map[string]interface{})

	//定义数组，用于存放app
	apps := []App{}

	//迭代结果集
	for rows.Next() {
		//将行数据保存到结构体App中
		err = rows.Scan(&id, &app_code, &app_name)
		if err == nil {
			//log.Println(id, app_code, app_name)
			apps = append(apps, App{id, app_code, app_name})
		}
	}

	//for..range 类似Java中的foreach
	for pos, app := range apps {
		//log.Println(pos, app)
		log.Println(pos, app.toString())
	}

	locals["apps"] = apps

	//调用渲染函数
	render(w, "index.html", locals)
	return
}

func main() {

	//定义请求映射，类似Django中的urls.py,或SpringMVC中的HandlerMapping作用
	http.HandleFunc("/", indexHandler)

	//开启监听9080端口HTTP请求
	err := http.ListenAndServe(":9080", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err.Error())
	}
}
