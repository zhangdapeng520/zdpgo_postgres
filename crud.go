package zdpgo_postgres

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/jackc/pgx/v4"
)

// 增删改查的SQL接口
type CrudSql interface {
	Add(sql string, args ...interface{}) int64                         // 添加数据
	Delete(table string, id int) bool                                  // 根据ID删除数据
	DeleteIds(table string, ids ...int) bool                           // 根据ID列表删除数据
	Update(sql string, args ...interface{}) bool                       // 根据ID修改数据
	Find(table string, columns []string, id int) pgx.Row               // 根据ID查询数据
	FindIds(table string, columns []string, ids []int) pgx.Rows        // 根据ID列表查询数据
	FindPages(table string, columns []string, page, size int) pgx.Rows // 根据分页查询数据
}

// 添加数据
func (p *Postgres) Add(sql string, args ...interface{}) (affected int64) {
	affected = p.Execute(sql, args...)
	return
}

// 根据ID删除数据
func (p *Postgres) Delete(table string, id int) (flag bool) {
	s := fmt.Sprintf("DELETE FROM %s WHERE id = %d;", table, id)
	ret := p.Execute(s)
	flag = ret > 0
	return
}

// 根据ID列表删除数据
func (p *Postgres) DeleteIds(table string, ids ...int) (flag bool) {
	// 整理ID列表
	var ids_ []string
	for _, v := range ids {
		vs := strconv.Itoa(v)
		ids_ = append(ids_, vs)
	}
	idsStr := strings.Join(ids_, ",")

	// 执行SQL语句
	s := fmt.Sprintf("DELETE FROM %s WHERE id IN (%s);",
		table, idsStr)
	p.log.Info("执行批量删除的SQL语句：", s)
	ret := p.Execute(s)
	flag = ret > 0
	return
}

// 根据ID修改数据
func (p *Postgres) Update(sql string, args ...interface{}) (flag bool) {
	ret := p.Execute(sql, args...)
	flag = ret > 0
	return
}

// 根据ID查询数据
func (p *Postgres) Find(table string, columns []string, id int) (row pgx.Row) {
	columnsStr := strings.Join(columns, ",")
	s := fmt.Sprintf("SELECT %s FROM %s WHERE id = $1;", columnsStr, table)
	row = p.Query(s, id)
	return
}

// 根据ID列表查询数据
func (p *Postgres) FindIds(table string, columns []string, ids []int) (rows pgx.Rows) {
	// 整理ID列表
	var ids_ []string
	for _, v := range ids {
		vs := strconv.Itoa(v)
		ids_ = append(ids_, vs)
	}
	idsStr := strings.Join(ids_, ",")

	// 整理字段列表
	columnsStr := strings.Join(columns, ",")

	// 执行SQL语句
	s := fmt.Sprintf("SELECT %s FROM %s WHERE id IN (%s);",
		columnsStr, table, idsStr)

	// 执行查询
	rows, err := p.QueryMany(s)
	if err != nil {
		p.log.Error("根据ID列表查询多条数据失败：", err)
		return nil
	}
	return
}

// 根据分页查询数据
func (p *Postgres) FindPages(table string, columns []string, page, size int) (rows pgx.Rows) {
	// 整理字段列表
	columnsStr := strings.Join(columns, ",")

	// 计算偏移量
	offset := (page - 1) * size
	s := fmt.Sprintf("SELECT %s FROM %s ORDER BY id DESC LIMIT %d OFFSET %d;", columnsStr, table, size, offset)

	// 执行查询
	rows, err := p.QueryMany(s)
	if err != nil {
		p.log.Error("根据ID列表查询多条数据失败：", err)
		return nil
	}
	return
}
