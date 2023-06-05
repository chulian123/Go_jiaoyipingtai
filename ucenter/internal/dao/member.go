package dao

import (
	"context"
	"gorm.io/gorm"
	"mscoin-common/msdb"
	"mscoin-common/msdb/gorms"
	"ucenter/internal/model"
)

type MemberDao struct {
	conn *gorms.GormConn
}

// Save  保存操作
//
//	func (m *MemberDao) Save(ctx context.Context, mem *model.Member) error {
//		//TODO implement me
//		session := m.conn.Session(ctx)
//		err := session.Save(mem).Error
//		return err
//	}
func (m *MemberDao) Save(ctx context.Context, mem *model.Member) error {
	session := m.conn.Session(ctx)
	err := session.Save(mem).Error
	return err
}

// FindByPhone 查询手机号码
func (m *MemberDao) FindByPhone(ctx context.Context, phone string) (mem *model.Member, err error) {
	//TODO implement me
	session := m.conn.Session(ctx)
	err = session.Model(&model.Member{}).Where("mobile_phone = ?", phone).Limit(1).Take(&mem).Error
	if err == gorm.ErrRecordNotFound {
		return nil, nil
	}
	return mem, err
}

// UpDateLoginCount 更新登录次数
func (m *MemberDao) UpDateLoginCount(ctx context.Context, id int64, step int) error {
	session := m.conn.Session(ctx)
	//实现+1
	//login_count =login_count+1
	err := session.Exec("update Member set login_count =login_count+? where id = ? ", step, id).Error
	return err
}

func NewMemberDao(db *msdb.MsDB) *MemberDao {
	return &MemberDao{
		conn: gorms.New(db.Conn),
	}
}
