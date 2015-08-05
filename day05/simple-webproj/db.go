// db.go
package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"log"
)

func getDB(username, password, dbName string) (*sql.DB, error) {
	dataSourceName := fmt.Sprintf("%s:%s@tcp(localhost:3306)/%s?charset=utf8", username, password, dbName)
	db, err := sql.Open("mysql", dataSourceName)
	if err != nil {
		log.Println(err.Error()) //仅显示异常
		return nil, err
	}

	//SetMaxOpenConns用于设置最大打开的连接数，默认值为0表示不限制。
	db.SetMaxOpenConns(2000)

	//SetMaxIdleConns用于设置闲置的连接数。
	db.SetMaxIdleConns(1000)
	db.Ping()

	return db, nil
}
