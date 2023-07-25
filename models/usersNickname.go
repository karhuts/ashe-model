package models

// UsersNickname undefined
type UsersNickname struct {
	ID       *int64  `json:"id" gorm:"id"`
	Uid      *int64  `json:"uid" gorm:"uid"`
	CreateAt *int64  `json:"create_at" gorm:"create_at"`
	PrevName *string `json:"prev_name" gorm:"prev_name"`
	NewName  *string `json:"new_name" gorm:"new_name"`
}

func (u *UsersNickname) ToMap() map[string]interface{} {
	return map[string]interface{}{
		"id":        u.ID,
		"uid":       u.Uid,
		"create_at": u.CreateAt,
		"new_name":  u.NewName,
		"prev_name": u.PrevName,
	}
}
