package dao

import (
	"context"
	"gorm.io/gorm"
	"mscoin-common/msdb"
	"mscoin-common/msdb/gorms"
	"ucenter/internal/model"
)

type MemberWalletDao struct {
	conn *gorms.GormConn
}

func (m *MemberWalletDao) UpdateFreeze(ctx context.Context, conn msdb.DbConn, memberId int64, symbol string, money float64) error {
	m.conn = conn.(*gorms.GormConn)
	session := m.conn.Tx(ctx)
	sql := "update member_wallet set balance=balance-?, frozen_balance=frozen_balance+? where member_id=? and coin_name=?"
	err := session.Model(&model.MemberWallet{}).Exec(sql, money, money, memberId, symbol).Error
	return err
}

func (m *MemberWalletDao) Save(ctx context.Context, mw *model.MemberWallet) error {
	session := m.conn.Session(ctx)
	err := session.Save(&mw).Error
	return err
}

func (m *MemberWalletDao) FindByIdAndCoinName(ctx context.Context, memId int64, coinName string) (mw *model.MemberWallet, err error) {
	session := m.conn.Session(ctx)
	err = session.Model(&model.MemberWallet{}).
		Where("member_id=? and coin_name=?", memId, coinName).
		Take(&mw).Error
	if err == gorm.ErrRecordNotFound {
		return nil, nil
	}
	return
}

func NewMemberWalletDao(db *msdb.MsDB) *MemberWalletDao {
	return &MemberWalletDao{
		conn: gorms.New(db.Conn),
	}
}
