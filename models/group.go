package models

import "time"

// Group  群组
type Group struct {
	ID           int64     `gorm:"column:id" db:"id" json:"id"`                               //  自增主键
	Name         string    `gorm:"column:name" db:"name" json:"name"`                         //  群组名称
	AvatarUrl    string    `gorm:"column:avatar_url" db:"avatar_url" json:"avatar_url"`       //  群组头像
	Introduction string    `gorm:"column:introduction" db:"introduction" json:"introduction"` //  群组简介
	UserNum      int64     `gorm:"column:user_num" db:"user_num" json:"user_num"`             //  群组人数
	Extra        string    `gorm:"column:extra" db:"extra" json:"extra"`                      //  附加属性
	CreateTime   time.Time `gorm:"column:create_time" db:"create_time" json:"create_time"`    //  创建时间
	UpdateTime   time.Time `gorm:"column:update_time" db:"update_time" json:"update_time"`    //  更新时间
}

func (Group) TableName() string {
	return "group"
}
