package main

import (
	"fmt"
	"time"

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
		INSERT INTO users ( uuid, NAME, email, PASSWORD, created_at )
		VALUES
			( $1, $2, $3, $4, $5 )
		`
	result := p.Add(sql, "aa12", "张大鹏2", "zhangdapen12g@qq.com", "test", time.Now())
	fmt.Println(result)
}
