package main

import (
	"fmt"
	"time"

	"github.com/zhangdapeng520/zdpgo_postgres"
)

type User struct {
	Id        int
	Uuid      string
	Name      string
	Email     string
	Password  string
	CreatedAt time.Time
}

func main() {
	host := "192.168.18.101"
	port := 5432
	user := "postgres"
	password := "postgres"
	database := "test"

	p := zdpgo_postgres.New(host, port, user, password, database, "zdp_postgres.log", true)

	sql := `
		SELECT id, uuid, name, email, password, created_at FROM users
		`
	rows, err := p.QueryMany(sql)
	defer rows.Close()
	if err != nil {
		fmt.Println("查找用户失败：", err)
	} else {
		fmt.Println("查找用户成功。")
	}
	fmt.Println(rows)

	// 根据结构体查询
	var users []User
	for rows.Next() {
		user1 := User{}
		if err = rows.Scan(&user1.Id, &user1.Uuid, &user1.Name, &user1.Email, &user1.Password, &user1.CreatedAt); err != nil {
			fmt.Println(err)
			return
		}
		fmt.Println(&user1.Id, &user1.Uuid, &user1.Name, &user1.Email, &user1.Password, &user1.CreatedAt)
		users = append(users, user1)
	}
	fmt.Println(users)
}
