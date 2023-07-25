package models

type Notes struct {
	Uid       *int64  `json:"uid,omitempty" gorm:"uid;default:null"`
	TargetUid *int64  `json:"target_uid,omitempty" gorm:"target_uid;default:null"`
	CreatedAt *int64  `json:"created_at,omitempty" gorm:"created_at;default:0"`
	Note      *string `json:"note,omitempty" gorm:"note;default:null"`

	User   *Users `json:"user" gorm:"foreignkey:uid"`
	Target *Users `json:"target" gorm:"foreignkey:target_uid"`
}

func (u *Notes) ToMap() map[string]interface{} {
	return map[string]interface{}{
		"uid":        u.Uid,
		"target_uid": u.TargetUid,
		"created_at": u.CreatedAt,
		"note":       u.Note,

		"user":   u.User.JoinUserToMap(),
		"target": u.Target.JoinUserToMap(),
	}
}
