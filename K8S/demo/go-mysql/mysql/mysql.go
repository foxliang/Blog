package mysql

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"sync"
	"xorm.io/xorm"
)

var (
	currentConfig Config
)

// GetCfg get current config
func GetCfg() *Config {
	return &currentConfig
}

// SetCfg set current config
func SetCfg(cfg *Config) {
	currentConfig = *cfg
}

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
	//fmt.Println("GetCfg().Addr", GetCfg().Addr)
	_dbOnce.Do(func() {
		//Addr = "root:123456@tcp(192.168.79.2:31306)/test"
		//engine, err := xorm.NewEngine("mysql", "root:123456@tcp(192.168.79.2:31306)/test")
		engine, err := xorm.NewEngine("mysql", GetCfg().Addr)
		if err != nil {
			fmt.Println("mysql connection err", err)
		}

		engine.SetMaxOpenConns(GetCfg().MaxOpenConns)
		engine.SetMaxIdleConns(GetCfg().MaxIdleConns)
		//engine.SetLogger(newDbLogger())
		//engine.ShowSQL(true)

		_db = engine

		fmt.Println("mysql init successed")
	})

	return _db
}
