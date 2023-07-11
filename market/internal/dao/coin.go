package dao

import (
	"context"
	"gorm.io/gorm"
	"market/internal/model"
	"mscoin-common/msdb"
	"mscoin-common/msdb/gorms"
)

type CoinDao struct {
	conn *gorms.GormConn
}

func (d *CoinDao) FindByUnit(ctx context.Context, unit string) (*model.Coin, error) {
	session := d.conn.Session(ctx)
	coin := &model.Coin{}
	err := session.Model(&model.Coin{}).Where("unit=?", unit).Take(coin).Error
	if err != nil && err == gorm.ErrRecordNotFound {
		return nil, nil
	}
	return coin, err
}

func NewCoinDao(db *msdb.MsDB) *CoinDao {
	return &CoinDao{
		conn: gorms.New(db.Conn),
	}
}
