package main

import (
	"fmt"

	"github.com/zhangdapeng520/zdpgo_postgres"
)

func main() {
	host := "192.168.18.101"
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
	flag := p.Execute(sql)
	fmt.Println(flag)
}
