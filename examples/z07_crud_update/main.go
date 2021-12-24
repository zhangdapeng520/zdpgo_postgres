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

	// 添加数据
	sql := `
		UPDATE users SET name=$1 WHERE id=$2
		`
	result := p.Update(sql, "张大鹏2333", 2)
	fmt.Println(result)
}
