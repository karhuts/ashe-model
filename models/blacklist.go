package models

type Blacklist struct {
	Uid       *int64 `json:"uid" gorm:"uid,primary_key,not null;default:0"`
	TargetUid *int64 `json:"target_uid" gorm:"target_uid,primary_key;default:0"`
	CreatedAT *int64 `json:"created_at" gorm:"created_at;default:0"`

	User   *Users `json:"user" gorm:"foreignkey:uid"`
	Target *Users `json:"target" gorm:"foreignkey:target_uid"`
}

func (b *Blacklist) ToMap() map[string]interface{} {
	return map[string]interface{}{
		"uid":        b.Uid,
		"target_uid": b.TargetUid,
		"created_at": b.CreatedAT,

		"user":   b.User.JoinUserToMap(),
		"target": b.Target.JoinUserToMap(),
		}
}
