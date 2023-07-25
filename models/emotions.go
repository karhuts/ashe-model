package models


type Emoticons struct {
	Id            *int64  `json:"id,omitempty" gorm:"id;default:null;"`
	Uid           *int64  `json:"uid,omitempty" gorm:"uid;default:null;"`
	CreatedAt     *int64  `json:"created_at,omitempty" gorm:"created_at;default:0;"`
	Type          *string `json:"type,omitempty" gorm:"type;default:png"`
	Path          *string `json:"path,omitempty" gorm:"path;default:null;"`
	Status        *int    `json:"status,omitempty" gorm:"status;default:0"`
	LastIpAddress *string `json:"last_ip_address,omitempty" gorm:"last_ip_address;default:null;"`

	User *Users `json:"user" gorm:"foreignkey:uid"`
}

func (u *Emoticons) ToMap() map[string]interface{} {
	return map[string]interface{}{
		"id":              u.Id,
		"uid":             u.Uid,
		"created_at":      u.CreatedAt,
		"type":            u.Type,
		"path":            u.Path,
		"status":          u.Status,
		"last_ip_address": u.LastIpAddress,

		"user": u.User.JoinUserToMap(),
		}
}