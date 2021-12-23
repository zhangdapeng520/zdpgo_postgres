package zdpgo_postgres

import "database/sql"

// Delete 从数据库中删除用户
func (p *Postgres) Add(sql string, args ...interface{}) (stmt *sql.Stmt, err error) {

	// 预处理SQL语句
	stmt, err = p.conn.Prepare(sql)
	if err != nil {
		p.log.Error("预处理SQL语句失败：", err)
		return nil, err
	}
	defer stmt.Close()

	// 执行SQL语句
	_, err = stmt.Exec(args...)
	if err != nil {
		return nil, err
	}
	return stmt, err
}
