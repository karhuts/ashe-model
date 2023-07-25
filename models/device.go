package models

import "time"

// Device  设备
const (
	DeviceOnLine  = 1 // 设备在线
	DeviceOffLine = 0 // 设备离线
)

type Device struct {
	ID            int64     `gorm:"column:id" db:"id" json:"id"`                                     //  自增主键
	UserId        int64     `gorm:"column:user_id" db:"user_id" json:"user_id"`                      //  账户id
	Type          int64     `gorm:"column:type" db:"type" json:"type"`                               //  设备类型,1:android；2：ios；3：windows; 4：macos；5：web
	Brand         string    `gorm:"column:brand" db:"brand" json:"brand"`                            //  手机厂商
	Model         string    `gorm:"column:model" db:"model" json:"model"`                            //  机型
	SystemVersion string    `gorm:"column:system_version" db:"system_version" json:"system_version"` //  系统版本
	SdkVersion    string    `gorm:"column:sdk_version" db:"sdk_version" json:"sdk_version"`          //  app版本
	Status        int64     `gorm:"column:status" db:"status" json:"status"`                         //  在线状态，0：离线；1：在线
	ConnAddr      string    `gorm:"column:conn_addr" db:"conn_addr" json:"conn_addr"`                //  连接层服务器地址
	ClientAddr    string    `gorm:"column:client_addr" db:"client_addr" json:"client_addr"`          //  客户端地址
	CreateTime    time.Time `gorm:"column:create_time" db:"create_time" json:"create_time"`          //  创建时间
	UpdateTime    time.Time `gorm:"column:update_time" db:"update_time" json:"update_time"`          //  更新时间
}

func (Device) TableName() string {
	return "device"
}
