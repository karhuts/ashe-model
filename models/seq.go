package models

import "time"

// Seq  序列号
type Seq struct {
	ID         int64     `gorm:"column:id" db:"id" json:"id"`                            //  自增主键
	ObjectType int64     `gorm:"column:object_type" db:"object_type" json:"object_type"` //  对象类型,1:用户；2：群组
	ObjectId   int64     `gorm:"column:object_id" db:"object_id" json:"object_id"`       //  对象id
	Seq        int64     `gorm:"column:seq" db:"seq" json:"seq"`                         //  序列号
	CreateTime time.Time `gorm:"column:create_time" db:"create_time" json:"create_time"` //  创建时间
	UpdateTime time.Time `gorm:"column:update_time" db:"update_time" json:"update_time"` //  更新时间
}

func (Seq) TableName() string {
	return "seq"
}
