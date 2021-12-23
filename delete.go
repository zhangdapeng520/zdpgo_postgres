package zdpgo_postgres

import "fmt"

// Delete 从数据库中删除用户
func (p *Postgres) Delete(column string, columnValue interface{}) (err error) {
	statement := fmt.Sprintf("delete from users where %s = $1", column)
	stmt, err := p.conn.Prepare(statement)
	if err != nil {
		return
	}
	defer stmt.Close()
	_, err = stmt.Exec(columnValue)
	return
}
