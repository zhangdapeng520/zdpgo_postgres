package main

import (
	"fmt"
	"time"

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
		INSERT INTO users ( uuid, NAME, email, PASSWORD, created_at )
		VALUES
			( $1, $2, $3, $4, $5 ) RETURNING ID,
			uuid,
			created_at
		`
	err := p.Execute(sql, "aa1", "张大鹏", "zhangdapen1g@qq.com", "test", time.Now())
	if err != nil {
		fmt.Println("创建用户失败：", err)
	}
	fmt.Println("创建用户成功。")
}
