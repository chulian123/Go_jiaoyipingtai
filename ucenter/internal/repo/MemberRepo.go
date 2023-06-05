package repo

import (
	"context"
	"ucenter/internal/model"
)

//接口 对应Member的操作

type MemberRepo interface {
	FindByPhone(ctx context.Context, phone string) (*model.Member, error)
	Save(ctx context.Context, mem *model.Member) error
	UpDateLoginCount(ctx context.Context, id int64, step int) error
}
