package models

type PhotosComments struct {
	Id            *int64  `json:"id,omitempty" gorm:"id;default:null"`
	Uid           *int64  `json:"uid,omitempty" gorm:"uid;default:null"`
	Pid           *int64  `json:"pid,omitempty" gorm:"pid;default:null"`
	CreatedAt     *int64  `json:"created_at,omitempty" gorm:"created_at;default:0"`
	Status        *int    `json:"status,omitempty" gorm:"status;default:0"`
	LastIpAddress *string `json:"last_ip_address,omitempty" gorm:"last_ip_address;default:null"`
	Text          *string `json:"text,omitempty" gorm:"text;default:null"`

	Photo *Photos `json:"photo" gorm:"foreignkey:pid"`
	User  *Users  `json:"user" gorm:"foreignkey:uid"`
}

func (u *PhotosComments) ToMap() map[string]interface{} {
	return map[string]interface{}{
		"id":              u.Id,
		"uid":             u.Uid,
		"pid":             u.Pid,
		"status":          u.Status,
		"created_at":      u.CreatedAt,
		"last_ip_address": u.LastIpAddress,
		"text":            u.Text,

		"photo": u.Photo.ToMap(),
		"user":  u.User.JoinUserToMap(),
	}
}
