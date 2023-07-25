package models

type Avatar struct {
	Url     string `json:"url" default:""`
	Type    string `json:"type" default:"image"`
	Index   int64  `json:"index" default:"0"`
	Audited int64  `json:"audited" default:"0"`
}

func (a *Avatar) ToMap() map[string]interface{} {
	return map[string]interface{}{
		"url":     a.Url,
		"type":    a.Type,
		"index":   a.Index,
		"audited": a.Audited,
	}
}

type Users struct {
	Id                *int64  `json:"id" gorm:"id;default:0"`
	Username          *string `json:"username" gorm:"username,default:null"`
	IsEmailConfirmed  *int    `json:"is_email_confirmed" gorm:"is_email_confirmed;default:0"`
	Avatar            *string `json:"avatar" gorm:"avatar;default:null"`
	LastSeenAt        *int64  `json:"last_seen_at" gorm:"last_seen_at;default:0"`
	SexualOrientation *int    `json:"sexual_orientation" gorm:"sexual_orientation;default:0"`
	Role              *int    `json:"role" gorm:"role;default:1"`
	PhotosCount       *int    `json:"photos_count" gorm:"photos_count;default:0"`
	Badge             *int    `json:"badge" gorm:"badge;default:0"`
	IsImmortal        *int    `json:"is_immortal" gorm:"is_immortal;default:0"`
	FansCount         *int    `json:"fans_count" gorm:"fans_count;default:0"`
	FollowedCount     *int    `json:"followed_count" gorm:"followed_count;default:0"`
	BlacklistCount    *int    `json:"blacklist_count" gorm:"blacklist_count;default:0"`
	EmoticonCount     *int    `json:"emoticon_count" gorm:"emoticon_count;default:0"`
	Birthday          *int64  `json:"birthday" gorm:"birthday;default:0"`
	IsLocked          *int    `json:"is_locked" gorm:"is_locked;default:0"`
	IsOnline          *int    `json:"is_online" gorm:"is_online;default:0"`
	IsAudited         *int    `json:"is_audited" gorm:"is_audited;default:0"`
	Height            *int    `json:"height" gorm:"height;default:0"`
	Weight            *int    `json:"weight" gorm:"weight;default:0"`
	Ethnicity         *int    `json:"ethnicity" gorm:"ethnicity;default:0"`
	Email             *string `json:"email" gorm:"email;default:null"`
	CitySettled       *string `json:"city_settled" gorm:"city_settled;default:null"`
	Hometown          *string `json:"hometown" gorm:"hometown;default:null"`
	Bio               *string `json:"bio" gorm:"bio;default:null"`
	Password          *string `json:"password" gorm:"password;default:null"`
	RegIpAddress      *string `json:"reg_ip_address" gorm:"reg_ip_address;default:127.0.0.1"`
	RegVersionName    *string `json:"reg_version_name" gorm:"reg_version_name;default:1.0.0"`
	JoinedAt          *int64  `json:"joined_at" gorm:"joined_at;default:0"`
	Education         *int64  `json:"education" gorm:"education;default:0"`
	Delisting         *int64  `json:"delisting" gorm:"delisting;default:0"`
	Progress          *int64  `json:"progress" gorm:"progress;default:0"`
	About             *string `json:"about" gorm:"about;default:"`
	Hobbies           *string `json:"hobbies" gorm:"hobbies;default:"`
	Emotional         *string `json:"emotional" gorm:"emotional;default:"`
	Dream             *string `json:"dream" gorm:"dream;default:"`
	Marital           *int64  `json:"marital" gorm:"marital;default:0"`
	Salary            *int64  `json:"salary" gorm:"salary;default:0"`
	Job               *int64  `json:"job" gorm:"marital;job:0"`
	AccessTokens      *AccessTokens  `json:"access_tokens" gorm:"foreignkey:UserId"`
}

func (u *Users) ToMap() map[string]interface{} {
	return map[string]interface{}{
		"role":               u.Role,
		"dream":              u.Dream,
		"about":              u.About,
		"emotional":          u.Emotional,
		"hobbies":            u.Hobbies,
		"uid":                u.Id,
		"username":           u.Username,
		"email":              u.Email,
		"password":           u.Password,
		"reg_ip_address":     u.RegIpAddress,
		"reg_version_name":   u.RegVersionName,
		"joined_at":          u.JoinedAt,
		"is_email_confirmed": u.IsEmailConfirmed,
		"avatar":             u.Avatar,
		"last_seen_at":       u.LastSeenAt,
		"sexual_orientation": u.SexualOrientation,
		"photos_count":       u.PhotosCount,
		"badge":              u.Badge,
		"is_immortal":        u.IsImmortal,
		"fans_count":         u.FansCount,
		"followed_count":     u.FollowedCount,
		"blacklist_count":    u.BlacklistCount,
		"emoticon_count":     u.EmoticonCount,
		"birthday":           u.Birthday,
		"is_locked":          u.IsLocked,
		"is_audited":         u.IsAudited,
		"height":             u.Height,
		"weight":             u.Weight,
		"ethnicity":          u.Ethnicity,
		"city_settled":       u.CitySettled,
		"hometown":           u.Hometown,
		"bio":                u.Bio,
		"is_online":          u.IsOnline,
		"education":          u.Education,
		"delisting":          u.Delisting,
		"progress":           u.Progress,
		"marital":            u.Marital,
		"salary":             u.Salary,
		"job":                u.Job,

		"access_tokens": u.AccessTokens.JoinUserToMap(),
	}
}

func (u *Users) JoinUserToMap() map[string]interface{} {
	return map[string]interface{}{
		"uid":      u.Id,
		"username": u.Username,
		//"email":              u.Email,
		//"password":           u.Password,
		//"reg_ip_address":     u.RegIpAddress,
		//"reg_version_name":   u.RegVersionName,
		"joined_at": u.JoinedAt,
		//"is_email_confirmed": u.IsEmailConfirmed,
		"avatar":             u.Avatar,
		"last_seen_at":       u.LastSeenAt,
		"sexual_orientation": u.SexualOrientation,
		"photos_count":       u.PhotosCount,
		"badge":              u.Badge,
		//"is_immortal":        u.IsImmortal,
		"fans_count":      u.FansCount,
		"followed_count":  u.FollowedCount,
		"blacklist_count": u.BlacklistCount,
		"emoticon_count":  u.EmoticonCount,
		"birthday":        u.Birthday,
		//"is_locked":          u.IsLocked,
		//"is_audited":         u.IsAudited,
		"height":       u.Height,
		"weight":       u.Weight,
		"ethnicity":    u.Ethnicity,
		"city_settled": u.CitySettled,
		"hometown":     u.Hometown,
		"bio":          u.Bio,
		"is_online":    u.IsOnline,

		"access_tokens": u.AccessTokens.JoinUserToMap(),
	}
}


