package models


type Settings struct {
	Id        *int64  `json:"id,omitempty" gorm:"id"`
	Uid       *int64  `json:"uid,omitempty" gorm:"uid"`
	Field     *string `json:"field,omitempty" gorm:"field"`
	Value     *string `json:"value,omitempty" gorm:"value"`
	CreatedAt *int64  `json:"created_at,omitempty" gorm:"created_at"`

	User *Users `json:"user" gorm:"foreignkey:uid"`
}

func (s *Settings) ToMap() map[string]interface{} {
	return map[string]interface{}{
		"uid":        s.Uid,
		"field":      s.Field,
		"value":      s.Value,
		"created_at": s.CreatedAt,

		"user": s.User.JoinUserToMap(),
		}
}
