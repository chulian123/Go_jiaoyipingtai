package domain

import (
	"context"
	"mscoin-common/msdb"
	"ucenter/internal/dao"
	"ucenter/internal/model"
	"ucenter/internal/repo"
)

type MemberAddressDomain struct {
	memberAddressRepo repo.MemberAddressRepo
}

func (d *MemberAddressDomain) FindAddressByCoinId(ctx context.Context, userid int64, coinId int64) ([]*model.MemberAddress, error) {
	return d.memberAddressRepo.FindByMemIdAndCoinId(ctx, userid, coinId)
}

func NewMemberAddressDomain(db *msdb.MsDB) *MemberAddressDomain {
	return &MemberAddressDomain{
		memberAddressRepo: dao.NewMemberAddressDao(db),
	}
}
