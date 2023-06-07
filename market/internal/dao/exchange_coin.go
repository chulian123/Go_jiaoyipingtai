package dao

import (
	"context"
	"market/internal/model"
	"mscoin-common/msdb"
	"mscoin-common/msdb/gorms"
)

type ExchangeCoinDao struct {
	conn *gorms.GormConn
}

func (d *ExchangeCoinDao) FindVisible(ctx context.Context) (list []*model.ExchangeCoin, err error) {
	session := d.conn.Session(ctx)
	err = session.Model(&model.ExchangeCoin{}).Where("visible=?", 1).Find(&list).Error
	return
}

func NewExchangeCoinDao(db *msdb.MsDB) *ExchangeCoinDao {
	return &ExchangeCoinDao{
		conn: gorms.New(db.Conn),
	}
}
