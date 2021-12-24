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

	// 创建表格
	sql := `
		INSERT INTO users ( uuid, NAME, email, PASSWORD, created_at )
		VALUES
			( $1, $2, $3, $4, $5 ) RETURNING ID,
			uuid,
			created_at
		`
	result := p.Execute(sql, "aa1", "张大鹏", "zhangdapen1g@qq.com", "test", time.Now())
	fmt.Println(result)
}
