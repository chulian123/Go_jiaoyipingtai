package dao

import (
	"context"
	"mscoin-common/msdb"
	"mscoin-common/msdb/gorms"
	"mscoin-common/tools"
	"ucenter/internal/model"
)

type MemberTransactionDao struct {
	conn *gorms.GormConn
}

func (d *MemberTransactionDao) FindTransaction(ctx context.Context, pageNo int, pageSize int, memberId int64, startTime string, endTime string, symbol string, transactionType string) (list []*model.MemberTransaction, total int64, err error) {
	session := d.conn.Session(ctx)
	db := session.Model(&model.MemberTransaction{}).Where("member_id=?", memberId)
	if transactionType != "" {
		db.Where("type=?", tools.ToInt64(transactionType))
	}
	if startTime != "" && endTime != "" {
		sTime := tools.ToMill(startTime)
		eTime := tools.ToMill(endTime)
		db.Where("create_time >= ? and create_time <= ?", sTime, eTime)
	}
	if symbol != "" {
		db.Where("symbol=?", symbol)
	}
	offset := (pageNo - 1) * pageSize
	db.Count(&total)
	db.Order("create_time desc").Offset(offset).Limit(pageSize)
	err = db.Find(&list).Error
	return
}
func NewMemberTransactionDao(db *msdb.MsDB) *MemberTransactionDao {
	return &MemberTransactionDao{
		conn: gorms.New(db.Conn),
	}
}
