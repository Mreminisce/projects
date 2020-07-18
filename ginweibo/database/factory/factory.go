package factory

import (
	"ginweibo/database"
)

// 清空表
func DropAndCreateTable(table interface{}) {
	database.DB.DropTable(table)
	database.DB.CreateTable(table)
}
