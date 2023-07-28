package domain

import (
	"context"
	"mscoin-common/msdb"
	"ucenter/internal/dao"
	"ucenter/internal/model"
	"ucenter/internal/repo"
)

type MemberTransactionDomain struct {
	memberTransactionRepo repo.MemberTransactionRepo
}

func (d *MemberTransactionDomain) FindTransaction(ctx context.Context, userid int64, pageNo int64, pageSize int64, symbol string, startTime string, endTime string, t string) ([]*model.MemberTransactionVo, int64, error) {
	//通过repo进行查询
	memberTransactions, total, err := d.memberTransactionRepo.FindTransaction(ctx, int(pageNo), int(pageSize), userid, startTime, symbol, endTime, t)
	if err != nil {
		return nil, total, err
	}
	volist := make([]*model.MemberTransactionVo, len(memberTransactions))
	for i, v := range memberTransactions {
		volist[i] = v.ToVo()
	}
	return volist, total, nil
}

func NewMemberTransactionDomain(db *msdb.MsDB) *MemberTransactionDomain {
	return &MemberTransactionDomain{
		memberTransactionRepo: dao.NewMemberTransactionDao(db),
	}
}
