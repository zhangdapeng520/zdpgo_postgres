package zdpgo_postgres

import (
	"context"

	"github.com/jackc/pgx/v4"
)

// FindMany 从数据库中查询多条数据
func (p *Postgres) QueryMany(sql string, args ...interface{}) (rows pgx.Rows, err error) {
	p.log.Info("正在执行查询SQL语句：", sql)
	p.log.Info("参数：", args)

	// 执行SQL语句
	rows, err = p.conn.Query(context.Background(), sql, args...)
	if err != nil {
		return nil, err
	}
	return rows, err
}

// Find 从数据库中查询单条数据
func (p *Postgres) Query(sql string, args ...interface{}) pgx.Row {
	p.log.Info("正在执行查询SQL语句：", sql)
	p.log.Info("参数：", args)

	// 执行SQL语句
	row := p.conn.QueryRow(context.Background(), sql, args...)
	return row
}
