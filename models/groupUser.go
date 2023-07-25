package models

import "time"

// GroupUser  群组成员
type GroupUser struct {
	ID         int64     `gorm:"column:id" db:"id" json:"id"`                            //  自增主键
	GroupId    int64     `gorm:"column:group_id" db:"group_id" json:"group_id"`          //  组id
	UserId     int64     `gorm:"column:user_id" db:"user_id" json:"user_id"`             //  用户id
	MemberType int64     `gorm:"column:member_type" db:"member_type" json:"member_type"` //  成员类型，1：管理员；2：普通成员
	Remarks    string    `gorm:"column:remarks" db:"remarks" json:"remarks"`             //  备注
	Extra      string    `gorm:"column:extra" db:"extra" json:"extra"`                   //  附加属性
	Status     int64     `gorm:"column:status" db:"status" json:"status"`                //  状态
	CreateTime time.Time `gorm:"column:create_time" db:"create_time" json:"create_time"` //  创建时间
	UpdateTime time.Time `gorm:"column:update_time" db:"update_time" json:"update_time"` //  更新时间
}

func (GroupUser) TableName() string {
	return "group_user"
}
