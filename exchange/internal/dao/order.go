package dao

import (
	"context"
	"exchange/internal/model"
	"gorm.io/gorm"
	"mscoin-common/msdb"
	"mscoin-common/msdb/gorms"
)

type ExchangeOrderDao struct {
	conn *gorms.GormConn
}

func (e *ExchangeOrderDao) UpdateOrderComplete(
	ctx context.Context,
	orderId string,
	tradedAmount float64,
	turnover float64,
	status int) error {
	session := e.conn.Session(ctx)
	updateSql := "update exchange_order set traded_amount=?,turnover=?,status=? where order_id=? and status=?"
	err := session.Model(&model.ExchangeOrder{}).Exec(updateSql, tradedAmount, turnover, status, orderId, model.Trading).Error
	return err
}

func (e *ExchangeOrderDao) FindOrderListBySymbol(ctx context.Context, symbol string, status int) (list []*model.ExchangeOrder, err error) {
	session := e.conn.Session(ctx)
	err = session.Model(&model.ExchangeOrder{}).
		Where("symbol=? and status=?", symbol, status).Find(&list).Error
	return
}

func (e *ExchangeOrderDao) UpdateOrderStatusTrading(ctx context.Context, orderId string) error {
	session := e.conn.Session(ctx)
	err := session.Model(&model.ExchangeOrder{}).
		Where("order_id=?", orderId).Update("status", model.Trading).Error
	return err
}

func (e *ExchangeOrderDao) UpdateStatusCancel(ctx context.Context, orderId string) error {
	session := e.conn.Session(ctx)
	err := session.Model(&model.ExchangeOrder{}).
		Where("order_id=?", orderId).Update("status", model.Canceled).Error
	return err
}

func (e *ExchangeOrderDao) FindOrderByOrderId(ctx context.Context, orderId string) (order *model.ExchangeOrder, err error) {
	session := e.conn.Session(ctx)
	err = session.Model(&model.ExchangeOrder{}).
		Where("order_id=?", orderId).
		Take(&order).Error
	if err != nil && err == gorm.ErrRecordNotFound {
		return nil, nil
	}
	return
}

func (e *ExchangeOrderDao) Save(ctx context.Context, conn msdb.DbConn, order *model.ExchangeOrder) error {
	e.conn = conn.(*gorms.GormConn)
	tx := e.conn.Tx(ctx)
	err := tx.Save(&order).Error
	return err
}

func (e *ExchangeOrderDao) FindCurrentTradingCount(ctx context.Context, id int64, symbol string, direction int) (total int64, err error) {
	session := e.conn.Session(ctx)
	err = session.Model(&model.ExchangeOrder{}).
		Where("symbol=? and member_id=? and direction=? and status=?", symbol, id, direction, model.Trading).
		Count(&total).Error
	return
}

func (e *ExchangeOrderDao) FindOrderHistory(ctx context.Context, symbol string, page int64, size int64, memberId int64) (list []*model.ExchangeOrder, total int64, err error) {
	session := e.conn.Session(ctx)
	err = session.Model(&model.ExchangeOrder{}).
		Where("symbol=? and member_id=?", symbol, memberId).
		Limit(int(size)).
		Offset(int((page - 1) * size)).Find(&list).Error
	err = session.Model(&model.ExchangeOrder{}).
		Where("symbol=? and member_id=?", symbol, memberId).
		Count(&total).Error
	return
}

func (e *ExchangeOrderDao) FindOrderCurrent(ctx context.Context, symbol string, page int64, size int64, memberId int64) (list []*model.ExchangeOrder, total int64, err error) {
	session := e.conn.Session(ctx)
	err = session.Model(&model.ExchangeOrder{}).
		Where("symbol=? and member_id=? and status=?", symbol, memberId, model.Trading).
		Limit(int(size)).
		Offset(int((page - 1) * size)).Find(&list).Error
	err = session.Model(&model.ExchangeOrder{}).
		Where("symbol=? and member_id=? and status=?", symbol, memberId, model.Trading).
		Count(&total).Error
	return
}

func NewExchangeOrderDao(db *msdb.MsDB) *ExchangeOrderDao {
	return &ExchangeOrderDao{
		conn: gorms.New(db.Conn),
	}
}
