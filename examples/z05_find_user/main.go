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
	row := p.Find(sql)

	// 根据结构体查询
	user1 := User{}
	row.Scan(&user1.Id, &user1.Uuid, &user1.Name, &user1.Email, &user1.Password, &user1.CreatedAt)
	fmt.Println(user1.Id, user1.Uuid, user1.Name, user1.Email, user1.Password, user1.CreatedAt)

	// 根据普通字段查询
	var (
		id        int
		uuid      string
		name      string
		email     string
		password1 string
		create_at time.Time
	)
	sql = `
		SELECT id, uuid, name, email, password, created_at FROM users
		`
	row = p.Find(sql)
	row.Scan(&id, &uuid, &name, &email, &password1, &create_at)
	fmt.Println(id, uuid, name, email, password1, create_at)
}
