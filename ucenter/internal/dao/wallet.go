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

func (m *MemberWalletDao) FindByIdAndCoinId(ctx context.Context, memId int64, coinId int64) (mw *model.MemberWallet, err error) {
	session := m.conn.Session(ctx)
	err = session.Model(&model.MemberWallet{}).
		Where("member_id=? and coin_id=?", memId, coinId).
		Take(&mw).Error
	if err == gorm.ErrRecordNotFound {
		return nil, nil
	}
	return
}

func (m *MemberWalletDao) FindByAddress(ctx context.Context, address string) (mw *model.MemberWallet, err error) {
	session := m.conn.Session(ctx)
	err = session.Model(&model.MemberWallet{}).Where("address=?", address).Take(&mw).Error
	if err == gorm.ErrRecordNotFound {
		return nil, nil
	}
	return
}

func (m *MemberWalletDao) FindAllAddress(ctx context.Context, coinName string) (list []string, err error) {
	session := m.conn.Session(ctx)
	err = session.Model(&model.MemberWallet{}).Where("coin_name=?", coinName).Select("address").Find(&list).Error
	return
}

func (m *MemberWalletDao) UpdateAddress(ctx context.Context, wallet *model.MemberWallet) error {
	updateSql := "update member_wallet set address=?,address_private_key=? where id=?"
	session := m.conn.Session(ctx)
	err := session.Model(&model.MemberWallet{}).Exec(updateSql, wallet.Address, wallet.AddressPrivateKey, wallet.Id).Error
	return err
}

func (m *MemberWalletDao) FindByMemberId(ctx context.Context, memId int64) (list []*model.MemberWallet, err error) {
	session := m.conn.Session(ctx)
	err = session.Model(&model.MemberWallet{}).Where("member_id=?", memId).Find(&list).Error
	return
}

func (m *MemberWalletDao) UpdateWallet(ctx context.Context, conn msdb.DbConn, id int64, balance float64, frozenBalance float64) error {
	gormConn := conn.(*gorms.GormConn)
	tx := gormConn.Tx(ctx)
	//Update
	updateSql := "update member_wallet set balance=?,frozen_balance=? where id=?"
	err := tx.Model(&model.MemberWallet{}).Exec(updateSql, balance, frozenBalance, id).Error
	return err
}

func (m *MemberWalletDao) UpdateFreeze(ctx context.Context, conn msdb.DbConn, memberId int64, symbol string, money float64) error {
	con := conn.(*gorms.GormConn)
	session := con.Tx(ctx)
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
