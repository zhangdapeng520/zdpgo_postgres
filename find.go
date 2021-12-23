package zdpgo_postgres

import "fmt"

func (p *Postgres) FindCountByColumn(table, columnName string, columnValue interface{}) (count int) {
	// 构建SQL语句
	sql := fmt.Sprintf("SELECT count(*) FROM %s where %s = $1", table, columnName)

	// 查询数据库
	rows, err := p.conn.Query(sql, columnValue)
	defer rows.Close()
	// rows, err := Db.Query("SELECT count(*) FROM posts where thread_id = $1", thread.Id)
	if err != nil {
		p.log.Error("查询数据库失败：", err)
	}

	// 取值
	for rows.Next() {
		if err = rows.Scan(&count); err != nil {
			p.log.Error("查询数据库失败：", err)
		}
	}

	return count
}
