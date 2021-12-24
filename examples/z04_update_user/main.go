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
		update users set name = $2, email = $3 where id = $1
		`
	result := p.Execute(sql, 1, "主宰者", "xxx@qq.com")
	fmt.Println(result)
}
