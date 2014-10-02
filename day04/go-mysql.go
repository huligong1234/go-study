// go-mysql.go
package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"strconv" //这个是为了把int转换为string
)

func main() { //main函数

	db, err := sql.Open("mysql", "root:123456@tcp(localhost:3306)/test?charset=utf8")
	if err != nil {
		panic(err.Error())       //抛出异常
		fmt.Println(err.Error()) //仅仅是显示异常
	}
	defer db.Close() //只有在前面用了 panic 这时defer才能起作用
	rows, err := db.Query("select id,app_code,app_name from t_app")
	if err != nil {
		panic(err.Error())
		fmt.Println(err.Error())
		return
	}
	defer rows.Close()
	var id int
	var app_code, app_name string
	for rows.Next() {
		rerr := rows.Scan(&id, &app_code, &app_name)
		if rerr == nil {
			fmt.Println(strconv.Itoa(id) + "|" + app_code + "|" + app_name)
		}

	}
	//db.Close()
}
