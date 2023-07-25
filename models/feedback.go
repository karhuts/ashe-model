package models



type Feedback struct {
	Id          *int64  `json:"id,omitempty" gorm:"id"`
	Uid         *int64  `json:"uid,omitempty" gorm:"uid"`
	Contents    *string `json:"contents,omitempty" gorm:"contents"`
	CreatedAt   *int64  `json:"created_at,omitempty" gorm:"created_at"`
	App         *int    `json:"app,omitempty" gorm:"app"`
	LastIp      *string `json:"last_ip,omitempty" gorm:"last_ip"`
	Platform    *string `json:"platform,omitempty" gorm:"platform"`
	Version     *string `json:"version,omitempty" gorm:"version"`
	Status      *int    `json:"status,omitempty" gorm:"status"`
	Attachments *string `json:"attachments,omitempty" gorm:"attachments"`

	User *Users `json:"user,omitempty" gorm:"foreignkey:uid"`
}

func (f *Feedback) ToMap() map[string]interface{} {
	return map[string]interface{}{
		"id":          f.Id,
		"user":        f.User.JoinUserToMap(),
		"uid":         f.Uid,
		"contents":    f.Contents,
		"created_at":  f.CreatedAt,
		"app":         f.App,
		"last_ip":     f.LastIp,
		"platform":    f.Platform,
		"version":     f.Version,
		"status":      f.Status,
		"attachments": f.Attachments,
		}
}