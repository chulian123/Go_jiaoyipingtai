package domain

import (
	"context"
	"errors"
	"github.com/zeromicro/go-zero/core/logx"
	"mscoin-common/msdb"
	"mscoin-common/tools"
	"ucenter/internal/dao"
	"ucenter/internal/model"
	"ucenter/internal/repo"
)

type MemberDomain struct {
	MemberRepo repo.MemberRepo
}

// FindByPhone 查找手机号
func (d MemberDomain) FindByPhone(ctx context.Context, phone string) (*model.Member, error) {
	//涉及数据库的查询
	mem, err := d.MemberRepo.FindByPhone(ctx, phone)
	if err != nil {
		logx.Error(err)
		return nil, errors.New("数据库异常")
	}
	return mem, nil
}

// Register 注册  promotion字段是邀请码
func (d *MemberDomain) Register(ctx context.Context, phone string, password string, username string, country string, partner string, promotion string) error {
	mem := model.NewMember()
	//对密码进行md5加密 加盐 md5加密不安全(容易被破解
	//member表字段比较多，而且所有字段不为null，很多字段要填写默认值 写一个工具类 通过反射来填充默认值即可
	_ = tools.Default(mem)
	salt, pwd := tools.Encode(password, nil) //给password字段放到md5工具类加密
	mem.Username = username
	mem.Country = country
	mem.Password = pwd
	mem.MobilePhone = phone
	mem.FillSuperPartner(partner)
	mem.PromotionCode = promotion
	mem.MemberLevelId = model.GENERAL
	mem.Salt = salt
	//调用数据库操作 存入数据库
	err := d.MemberRepo.Save(ctx, mem)
	if err != nil {
		logx.Info("数据库异常")
		return errors.New("数据库异常")
	}
	return nil
}

func NewMemberDomain(db *msdb.MsDB) *MemberDomain {
	return &MemberDomain{
		MemberRepo: dao.NewMemberDao(db),
	}
}
