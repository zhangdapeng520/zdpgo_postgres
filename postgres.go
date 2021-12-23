package zdpgo_postgres

import (
	"database/sql"
	"fmt"

	"github.com/zhangdapeng520/zdpgo_log"
)

// PostgreSQL结构体
type Postgres struct {
	conn    *sql.DB        // 数据库连接对象
	logPath string         // 日志文件路径
	log     *zdpgo_log.Log // 日志对象
	debug   bool           // 是否为debug模式
}

// 设置debug模式
func (p *Postgres) SetDebug(debug bool) {
	p.debug = debug
}

// 设置日志文件路径
func (p *Postgres) SetLogPath(logPath string) {
	p.logPath = logPath
	p.log = zdpgo_log.New(logPath)
	p.log.SetDebug(p.debug)
}

// New 创建新的Postgres对象
// @param host 主机地址
// @param port 端口号
// @param user 用户名
// @param password 密码
// @param database 数据库名称
// @param logPath 日志存放路径
// @param debug 是否为debug模式
func New(host string, port int, user, password, database, logPath string, debug bool) *Postgres {
	// 创建postgres对象
	p := Postgres{}
	p.debug = debug

	// 创建日志对象
	p.logPath = logPath
	p.log = zdpgo_log.New(logPath)
	p.log.SetDebug(debug)

	// 创建数据库连接对象
	address := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, database)
	Db, err := sql.Open("postgres", address)
	if err != nil {
		p.log.Panic("连接PostgreSQL数据库失败：", err)
	}
	p.conn = Db

	// 返回初始化后的postgres对象
	return &p
}

// Default 按照默配置创建新的Postgres对象
func Default() *Postgres {
	host := "127.0.0.1"
	port := 5432
	user := "postgres"
	password := "postgres"
	database := "postgres"
	logPath := "zdpgo_postgres.log"
	return New(host, port, user, password, database, logPath, true)
}
