package msdb

import "gorm.io/gorm"

// 事物的处理
type DbConn interface {
	Begin()
	Rollback()
	Commit()
}

type MsDB struct {
	Conn *gorm.DB
}
