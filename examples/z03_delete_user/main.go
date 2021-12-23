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
		delete from users where id = $1
		`
	err := p.Execute(sql, 1)
	if err != nil {
		fmt.Println("删除用户失败：", err)
	}
	fmt.Println("删除用户成功。")
}
