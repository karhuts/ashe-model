package models

type Mis struct {
	Id       *int64  `json:"id,omitempty" gorm:"id"`
	Uid      *int64  `json:"uid,omitempty" gorm:"uid"`
	Payload  *string `json:"payload,omitempty" gorm:"payload"`
	CreateAt *int64  `json:"create_at,omitempty" gorm:"create_at"`
	Type     *int    `json:"type,omitempty" gorm:"type"`
	Status   *int    `json:"status,omitempty" gorm:"status"`
	Operator *int    `json:"operator,omitempty" gorm:"operator"`

	User *Users `json:"user,omitempty" gorm:"foreignkey:uid"`
}

func (f *Mis) ToMap() map[string]interface{} {
	return map[string]interface{}{
		"id":         f.Id,
		"user":       f.User.JoinUserToMap(),
		"uid":        f.Uid,
		"payload":    f.Payload,
		"created_at": f.CreateAt,
		"status":     f.Status,
		"type":       f.Type,
		"operator":   f.Operator,
	}
}

type MisAdmin struct {
	Id          *int64  `json:"id,omitempty" gorm:"id"`
	Uid         *int64  `json:"uid,omitempty" gorm:"uid"`
	CreateAt    *int64  `json:"create_at,omitempty" gorm:"create_at"`
	LoginTime   *int64  `json:"login_time,omitempty" gorm:"login_time"`
	Status      *int    `json:"status,omitempty" gorm:"status"`
	IsSuper     *int    `json:"is_super,omitempty" gorm:"is_super"`
	Permissions *string `json:"permissions,omitempty" gorm:"permissions"`
	Nickname    *string `json:"nickname,omitempty" gorm:"nickname"`
	Token       *string `json:"token,omitempty" gorm:"token"`

	User *Users `json:"user,omitempty" gorm:"foreignkey:uid"`
}

func (a *MisAdmin) ToMap() map[string]interface{} {
	return map[string]interface{}{
		"id":          a.Id,
		"uid":         a.Uid,
		"create_at":   a.CreateAt,
		"login_time":  a.LoginTime,
		"is_super":    a.IsSuper,
		"status":      a.Status,
		"permissions": a.Permissions,
		"nickname":    a.Nickname,
		"token":       a.Token,
	}
}

type MisLogs struct {
	ID       *int64  `json:"id" gorm:"id"`
	Operator *int64  `json:"operator" gorm:"operator"`
	CreateAt *int64  `json:"create_at" gorm:"create_at"`
	Payload  *string `json:"payload" gorm:"payload"`
	Ip       *string `json:"ip" gorm:"ip"`

	User *Users `json:"user,omitempty" gorm:"foreignkey:operator"`
}

func (m *MisLogs) ToMap() map[string]interface{} {
	return map[string]interface{}{
		"id":        m.ID,
		"operator":  m.Operator,
		"create_at": m.CreateAt,
		"payload":   m.Payload,
		"users":     m.User.JoinUserToMap(),
	}
}
