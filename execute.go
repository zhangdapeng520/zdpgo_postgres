package zdpgo_postgres

import (
	"context"
)

// Execute 执行增删改SQL语句
func (p *Postgres) Execute(sql string, args ...interface{}) int64 {
	p.log.Info("正在执行SQL语句：", sql)
	p.log.Info("参数：", args)

	// 执行SQL语句
	result, err := p.conn.Exec(context.Background(), sql, args...)
	if err != nil {
		p.log.Error("执行SQL语句失败：", err)
	}
	affected := result.RowsAffected()
	if affected == 0 {
		p.log.Warning("注意：本次执行未生效，没有影响数据库任何一行！！！")
	}
	return affected
}
