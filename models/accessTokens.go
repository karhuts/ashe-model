package models

type AccessTokens struct {
	UserId           *int64   `json:"user_id" gorm:"user_id;default:0"`
	Token            *string  `json:"token" gorm:"token;default:null"`
	LastActivityAt   *int64   `json:"last_activity_at" gorm:"last_activity_at;default:0"`
	LastGeoLongitude *float64 `json:"last_geo_longitude" gorm:"last_geo_longitude;default:0.00"`
	LastGeoLatitude  *float64 `json:"last_geo_latitude" gorm:"last_geo_latitude;default:0.00"`
	CreatedAt        *int64   `json:"created_at" gorm:"created_at;default:0"`
	Type             *string  `json:"type" gorm:"type;default:jwt"`
	LastIpAddress    *string  `json:"last_ip_address" gorm:"last_ip_address;default:127.0.0.1"`
	LastUserAgent    *string  `json:"last_user_agent" gorm:"last_user_agent;default:null"`
}

func (a *AccessTokens) ToMap() map[string]interface{} {
	return map[string]interface{}{
		"user_id":            a.UserId,
		"token":              a.Token,
		"last_activity_at":   a.LastActivityAt,
		"last_geo_longitude": a.LastGeoLongitude,
		"last_geo_latitude":  a.LastGeoLatitude,
		"created_at":         a.CreatedAt,
		"type":               a.Type,
		"last_ip_address":    a.LastIpAddress,
		"last_user_agent":    a.LastUserAgent,
	}
}

func (a *AccessTokens) JoinUserToMap() map[string]interface{} {
	return map[string]interface{}{
		"token":              a.Token,
		"last_activity_at":   a.LastActivityAt,
		"last_geo_longitude": a.LastGeoLongitude,
		"last_geo_latitude":  a.LastGeoLatitude,
		"created_at":         a.CreatedAt,
		"type":               a.Type,
		"last_ip_address":    a.LastIpAddress,
		"last_user_agent":    a.LastUserAgent,
	}
}
