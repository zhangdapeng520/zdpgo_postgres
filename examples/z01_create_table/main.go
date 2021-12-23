package main

import (
	"fmt"

	"github.com/zhangdapeng520/zdpgo_postgres"
)

func main() {
	host := "127.0.0.1"
	port := 5432
	user := "postgres"
	password := "postgres"
	database := "test"

	p := zdpgo_postgres.New(host, port, user, password, database, "zdp_postgres.log", true)

	// 创建表格
	sql := `
		create table if not exists users (
		  id         serial primary key,
		  uuid       varchar(64) not null unique,
		  name       varchar(255),
		  email      varchar(255) not null unique,
		  password   varchar(255) not null,
		  created_at timestamp not null
		);
		`
	err := p.Execute(sql)
	if err != nil {
		fmt.Println("创建表格失败：", err)
	}
	fmt.Println("创建表格成功。")
}
