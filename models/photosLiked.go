package models

type PhotosLiked struct {
	Id            *int64  `json:"id,omitempty" gorm:"id;default:null"`
	Uid           *int64  `json:"uid,omitempty" gorm:"uid;default:null"`
	Pid           *int64  `json:"pid,omitempty" gorm:"pid;default:null"`
	CreatedAt     *int64  `json:"created_at,omitempty" gorm:"created_at;default:null"`
	LastIpAddress *string `json:"last_ip_address,omitempty" gorm:"last_ip_address;default:null"`

	Photo *Photos `json:"photo" gorm:"foreignkey:pid"`
	User  *Users  `json:"user" gorm:"foreignkey:uid"`
}

func (u *PhotosLiked) ToMap() map[string]interface{} {
	return map[string]interface{}{
		"id":              u.Id,
		"uid":             u.Uid,
		"pid":             u.Pid,
		"created_at":      u.CreatedAt,
		"last_ip_address": u.LastIpAddress,

		"photo": u.Photo.ToMap(),
		"user":  u.User.JoinUserToMap(),
	}
}
