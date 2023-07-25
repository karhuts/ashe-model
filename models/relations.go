package models

type Relations struct {
	Uid       *int64 `json:"uid" gorm:"uid,primary_key,not null;default:0"`
	TargetUid *int64 `json:"target_uid" gorm:"target_uid,primary_key;default:0"`
	Relation  *int   `json:"relation" gorm:"relation;default:0"`
	CreatedAT *int64 `json:"created_at" gorm:"created_at;default:0"`

	User   *Users `json:"user" gorm:"foreignkey:uid"`
	Target *Users `json:"target" gorm:"foreignkey:target_uid"`
}

func (f *Relations) ToMap() map[string]interface{} {
	return map[string]interface{}{
		"uid":        f.Uid,
		"target_uid": f.TargetUid,
		"relation":   f.Relation,
		"created_at": f.CreatedAT,

		"user":   f.User.JoinUserToMap(),
		"target": f.Target.JoinUserToMap(),
		}
}
