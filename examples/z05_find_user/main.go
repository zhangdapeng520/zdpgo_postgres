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
	host := "127.0.0.1"
	port := 5432
	user := "postgres"
	password := "postgres"
	database := "test"

	p := zdpgo_postgres.New(host, port, user, password, database, "zdp_postgres.log", true)

	sql := `
		SELECT id, uuid, name, email, password, created_at FROM users
		`
	row := p.Find(sql)
	user1 := User{}
	row.Scan(&user1.Id, &user1.Uuid, &user1.Name, &user1.Email, &user1.Password, &user1.CreatedAt)
	fmt.Println(user1.Id, user1.Uuid, user1.Name, user1.Email, user1.Password, user1.CreatedAt)
}
