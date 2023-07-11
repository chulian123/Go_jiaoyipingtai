package dao

import (
	"context"
	"github.com/zeromicro/go-zero/core/logx"
	"gorm.io/gorm"
	"market/internal/model"
	"mscoin-common/msdb"
	"mscoin-common/msdb/gorms"
)

type ExchangeCoinDao struct {
	conn *gorms.GormConn
}

func (d *ExchangeCoinDao) FindByFindSymbol(ctx context.Context, symbol string) (*model.ExchangeCoin, error) {
	session := d.conn.Session(ctx)
	data := &model.ExchangeCoin{}
	err := session.Model(&model.ExchangeCoin{}).Where("symbol=?", symbol).Take(data).Error
	if err != nil && err == gorm.ErrRecordNotFound {
		logx.Info("数据不存在")
		return nil, nil
	}
	return data, err

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
