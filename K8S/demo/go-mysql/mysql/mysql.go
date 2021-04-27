package mysql

import (
	"fmt"
	"sync"
	"xorm.io/xorm"

	_ "github.com/go-sql-driver/mysql"
)

var (
	_dbOnce sync.Once
	_db     *xorm.Engine
)

// Config global config
type Config struct {
	Addr         string `json:"addr"`
	MaxOpenConns int    `json:"max_open_conns"`
	MaxIdleConns int    `json:"max_idle_conns"`
	ShowLog      bool   `json:"show_log"`
}

// GetDB get db instance from config
func GetDB() *xorm.Engine {
	_dbOnce.Do(func() {
		engine, err := xorm.NewEngine("mysql", "root:123456@tcp(192.168.79.2:31306)/test")
		if err != nil {
			fmt.Println("mysql connection err", err)
		}

		engine.SetMaxOpenConns(8)
		engine.SetMaxIdleConns(2)
		//engine.SetLogger(newDbLogger())
		//engine.ShowSQL(true)

		_db = engine

		fmt.Println("mysql init successed")
	})

	return _db
}
