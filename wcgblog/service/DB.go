package service

import (
	"log"
	"time"
	"wcgblog/model"
	"wcgblog/system"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var db *gorm.DB
var err error

func ConnectDB() {
	db, err = gorm.Open("mysql", system.GetConfiguration().DSN)
	if err != nil {
		log.Panicln("连接数据库失败: " + err.Error())
	}
	// 最大空闲数
	db.DB().SetMaxIdleConns(10)
	// 最大连接数
	db.DB().SetMaxOpenConns(50)
	// 最长连接时间
	db.DB().SetConnMaxLifetime(5 * time.Minute)
	// 开启日志
	db.LogMode(true)
	// 创建数据表的时候 当新增数据表或第一次运行需要打开注释
	CreateTable()
}

// 从数据库断开
func DisconnectDB() {
	if err := db.Close(); nil != err {
		log.Println("数据库断开失败：" + err.Error())
	}
}

// 创建表，db.HasTable 判断表是否存在，db.set 设置表的属性
func CreateTable() {
	//创建评论数据表
	if !db.HasTable(&model.Comment{}) {
		if err := db.Set("gorm:table_options", "ENGINE=InnoDB DEFAULT CHARSET=utf8 comment '评论数据表'").CreateTable(&model.Comment{}).Error; err != nil {
			panic(err)
		}
	}
	// 创建用户订阅表
	if !db.HasTable(&model.Link{}) {
		if err := db.Set("gorm:table_options", "ENGINE=InnoDB DEFAULT CHARSET=utf8 comment '用户订阅数据表'").CreateTable(&model.Link{}).Error; err != nil {
			panic(err)
		}
	}
	// 分页 数据表
	if !db.HasTable(&model.Page{}) {
		if err := db.Set("gorm:table_options", "ENGINE=InnoDB DEFAULT CHARSET=utf8 comment '分页'").CreateTable(&model.Page{}).Error; err != nil {
			panic(err)
		}
	}
	//  标签 帖子 中间表
	if !db.HasTable(&model.PostTag{}) {
		if err := db.Set("gorm:table_options", "ENGINE=InnoDB DEFAULT CHARSET=utf8 comment ' 标签 帖子 中间表'").CreateTable(&model.PostTag{}).Error; err != nil {
			panic(err)
		}
	}
	// 标签 数据表
	if !db.HasTable(&model.Tag{}) {
		if err := db.Set("gorm:table_options", "ENGINE=InnoDB DEFAULT CHARSET=utf8 comment '标签 数据表'").CreateTable(&model.Tag{}).Error; err != nil {
			panic(err)
		}
	}
	// 用户 数据表
	if !db.HasTable(&model.User{}) {
		if err := db.Set("gorm:table_options", "ENGINE=InnoDB DEFAULT CHARSET=utf8 comment '用户 数据表'").CreateTable(&model.User{}).Error; err != nil {
			panic(err)
		}
	}
	//  帖子 数据表
	if !db.HasTable(&model.Post{}) {
		if err := db.Set("gorm:table_options", "ENGINE=InnoDB DEFAULT CHARSET=utf8 comment '帖子'").CreateTable(&model.Post{}).Error; err != nil {
			panic(err)
		}
	}
}
