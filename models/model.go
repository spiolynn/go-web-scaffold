package models

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"go-web-scaffold/pkgs/logging"
	"go-web-scaffold/pkgs/setting"
)

/*
数据库入口
*/

var db *gorm.DB

type Model struct {
	ID         int `gorm:"primary_key" json:"id"`
	CreatedOn  int `json:"created_on"`
	ModifiedOn int `json:"modified_on"`
}

// 初始化
func _init() {

	logging.Info(" model init start ... ")

	var err error
	db, err = gorm.Open(setting.G_cfg_yaml.DB.Type, fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8&parseTime=True&loc=Local",
		setting.G_cfg_yaml.DB.User,
		setting.G_cfg_yaml.DB.Password,
		setting.G_cfg_yaml.DB.Host,
		setting.G_cfg_yaml.DB.Sid))

	if err != nil {
		logging.Fatal(fmt.Sprintf("Fail to open the DB: %v, with errors: %v", setting.G_cfg_yaml.DB.Host, err))
	} else {
		logging.Info("Connect to DB successful.")
	}

	// 数据库表头
	gorm.DefaultTableNameHandler = func(db *gorm.DB, defaultTableName string) string {
		return setting.G_cfg_yaml.DB.Tableprefix + defaultTableName
	}

	// 全局禁用表名复数
	// 如果设置为true,`User`的默认表名为`user`
	db.SingularTable(true)
	// 设置闲置的连接数
	db.DB().SetMaxIdleConns(setting.G_cfg_yaml.DB.Maxidleconn)
	// 设置最大打开的连接数，默认值为0表示不限制
	db.DB().SetMaxOpenConns(setting.G_cfg_yaml.DB.Maxopenconn)

	// 打印sql日志
	if setting.G_cfg_yaml.DB.DBdebug == "true" {
		db.LogMode(true)
	}

	logging.Info(" model init end ... ")
}

func OpenDB() {
	_init()
}

// 关闭数据库
func CloseDB() {
	defer db.Close()
}
