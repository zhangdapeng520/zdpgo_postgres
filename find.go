package zdpgo_postgres

import "database/sql"

// FindMany 从数据库中查询多条数据
func (p *Postgres) FindMany(sql string, args ...interface{}) (rows *sql.Rows, err error) {
	// 执行SQL语句
	rows, err = p.conn.Query(sql, args...)
	if err != nil {
		return nil, err
	}
	return rows, err
}

// Find 从数据库中查询单条数据
func (p *Postgres) Find(sql string, args ...interface{}) *sql.Row {
	// 执行SQL语句
	row := p.conn.QueryRow(sql, args...)
	return row
}
