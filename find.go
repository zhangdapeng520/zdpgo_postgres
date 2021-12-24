package zdpgo_postgres

import (
	"context"

	"github.com/jackc/pgx/v4"
)

// FindMany 从数据库中查询多条数据
func (p *Postgres) FindMany(sql string, args ...interface{}) (rows pgx.Rows, err error) {
	// 执行SQL语句
	rows, err = p.conn.Query(context.Background(), sql, args...)
	if err != nil {
		return nil, err
	}
	return rows, err
}

// Find 从数据库中查询单条数据
func (p *Postgres) Find(sql string, args ...interface{}) pgx.Row {
	// 执行SQL语句
	row := p.conn.QueryRow(context.Background(), sql, args...)
	return row
}
