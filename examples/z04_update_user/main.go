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
		update users set name = $2, email = $3 where id = $1
		`
	err := p.Execute(sql, 2, "主宰者", "xxx@qq.com")
	if err != nil {
		fmt.Println("更新用户失败：", err)
	} else {
		fmt.Println("更新用户成功。")
	}
}
