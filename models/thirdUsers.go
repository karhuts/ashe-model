package models


type ThirdUsers struct {
	OpenId   string `json:"open_id,omitempty" gorm:"open_id"`
	Uid      int64  `json:"uid,omitempty" gorm:"uid"`
	CreateAt int    `json:"create_at,omitempty" gorm:"create_at"`
	Type     string `json:"type,omitempty" gorm:"type"`
	OpenName string `json:"open_name,omitempty" gorm:"open_name"`
}

func (t *ThirdUsers) ToMap() map[string]interface{} {
	return map[string]interface{}{
		"open_id":   t.OpenId,
		"uid":       t.Uid,
		"create_at": t.CreateAt,
		"type":      t.Type,
		"open_name": t.OpenName,
		}
}

func (*ThirdUsers) TableName() string {
	return "third_users"
}
