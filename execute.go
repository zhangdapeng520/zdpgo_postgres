package zdpgo_postgres

// Delete 从数据库中删除用户
func (p *Postgres) Execute(sql string, args ...interface{}) (err error) {

	// 预处理SQL语句
	stmt, err := p.conn.Prepare(sql)
	if err != nil {
		p.log.Error("预处理SQL语句失败：", err)
		return err
	}
	defer stmt.Close()

	// 执行SQL语句
	_, err = stmt.Exec(args...)
	if err != nil {
		return err
	}
	return
}
