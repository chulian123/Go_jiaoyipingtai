package domain

import (
	"context"
	"errors"
	"github.com/zeromicro/go-zero/core/logx"
	"mscoin-common/msdb"
	"ucenter/internal/dao"
	"ucenter/internal/model"
	"ucenter/internal/repo"
)

type MemberDomain struct {
	MemberRepo repo.MemberRepo
}

func (d MemberDomain) FindByPhone(ctx context.Context, phone string) (*model.Member, error) {
	//涉及数据库的查询
	mem, err := d.MemberRepo.FindByPhone(ctx, phone)
	if err != nil {
		logx.Error(err)
		return nil, errors.New("数据库异常")
	}
	return mem, nil
}

func NewMemberDomain(db *msdb.MsDB) *MemberDomain {
	return &MemberDomain{
		MemberRepo: dao.NewMemberDao(db),
	}
}
