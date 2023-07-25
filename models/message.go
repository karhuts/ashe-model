package models

import "time"

// Message  消息
type Message struct {
	ID         int64     `gorm:"column:id" db:"id" json:"id"`                            //  自增主键
	UserId     int64     `gorm:"column:user_id" db:"user_id" json:"user_id"`             //  所属类型的id
	RequestId  int64     `gorm:"column:request_id" db:"request_id" json:"request_id"`    //  请求id
	Code       int64     `gorm:"column:code" db:"code" json:"code"`                      //  消息类型
	Content    string    `gorm:"column:content" db:"content" json:"content"`             //  消息内容
	Seq        int64     `gorm:"column:seq" db:"seq" json:"seq"`                         //  消息序列号
	SendTime   time.Time `gorm:"column:send_time" db:"send_time" json:"send_time"`       //  消息发送时间
	Status     int64     `gorm:"column:status" db:"status" json:"status"`                //  消息状态，0：未处理1：消息撤回
	CreateTime time.Time `gorm:"column:create_time" db:"create_time" json:"create_time"` //  创建时间
	UpdateTime time.Time `gorm:"column:update_time" db:"update_time" json:"update_time"` //  更新时间
}

func (Message) TableName() string {
	return "message"
}
