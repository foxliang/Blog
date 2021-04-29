package config

import (
	"fmt"
	"github.com/BurntSushi/toml"
	"go-mysql/mysql"
	"sync"
)

//func flags() {
//	var mysqlAddr string
//
//	flag.StringVar(&mysqlAddr, "mysql", "", "user:password@tcp(host)/database")
//	fmt.Println("mysqlAddr", mysqlAddr)
//	flag.Parse()
//	mysql.SetCfg(&mysql.Config{
//		Addr:         mysqlAddr,
//		MaxOpenConns: 8,
//		MaxIdleConns: 2,
//	})
//}

var (
	once sync.Once
	// Cfg 配置文件
	Cfg  Config
	path = "/var/config/config.toml"
)

// Config 配置文件类型
type Config struct {
	MysqlAddr string `toml:"mysql-addr"`
}

// InitCfg 初始化配置文件
func InitCfg() {

	once.Do(func() {
		if _, err := toml.DecodeFile(path, &Cfg); err != nil {
			fmt.Println("cannot parse config file err", err)
		}
	})

	fmt.Println("Mysql", Cfg.MysqlAddr)

	mysql.SetCfg(&mysql.Config{
		Addr:         Cfg.MysqlAddr,
		MaxOpenConns: 8,
		MaxIdleConns: 2,
	})
}
