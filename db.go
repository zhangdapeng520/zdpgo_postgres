package zdpgo_postgres

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

// 获取DB连接
func GetDB(host string, port int, user, password, database string) (*sql.DB, error) {
	address := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, database)
	Db, err := sql.Open("postgres", address)
	return Db, err
}
