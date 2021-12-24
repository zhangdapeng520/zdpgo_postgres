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

	result := p.Find("users", []string{"name", "email"}, 2)
	fmt.Println(result)
	var name string
	var email string
	result.Scan(&name, &email)
	fmt.Println(name, email)
}
