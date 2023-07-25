package models

import (
	"encoding/json"

	"github.com/gookit/goutil/strutil"
)

type Photos struct {
	Id            *int64  `json:"id,omitempty" gorm:"id;default:null"`
	Uid           *int64  `json:"uid,omitempty" gorm:"uid;default:null"`
	CreatedAt     *int64  `json:"created_at,omitempty" gorm:"created_at;default:null"`
	CommentCount  *int64  `json:"comment_count,omitempty" gorm:"comment_count;default:0"`
	ViewCount     *int64  `json:"view_count,omitempty" gorm:"view_count;default:0"`
	LikedCount    *int64  `json:"liked_count,omitempty" gorm:"liked_count;default:0"`
	Status        *int    `json:"status,omitempty" gorm:"status;default:0"`
	Type          *string `json:"type,omitempty" gorm:"type;default:png"`
	LastIpAddress *string `json:"last_ip_address,omitempty" gorm:"last_ip_address;default:null"`
	Path          *string `json:"path,omitempty" gorm:"path;default:null"`
	Description   *string `json:"description,omitempty" gorm:"description;default:null"`
	Public        *int    `json:"public,omitempty" gorm:"public;default:0"`
	Geo           *string `json:"geo,omitempty" gorm:"geo;default:0"`
	TopicId       *int64  `json:"topic_id,omitempty" gorm:"topic_id;default:0"`

	Topic *PhotosTopics `json:"topic,omitempty" gorm:"foreignkey:topic_id"`
	User  *Users        `json:"user,omitempty" gorm:"foreignkey:uid;"`
}

func (u *Photos) ToMap() map[string]interface{} {
	p := make(map[string]string)
	data := strutil.MustString(*u.Path)
	if data != "" {
		err := json.Unmarshal([]byte(data), &p)
		if err != nil {
			return nil
		}
	}
	var user map[string]interface{}
	if u.User != nil {
		user = u.User.JoinUserToMap()
	}
	var topic map[string]interface{}
	if u.Topic != nil {
		topic = u.Topic.ToMap()
	}

	return map[string]interface{}{
		"id":              u.Id,
		"uid":             u.Uid,
		"created_at":      u.CreatedAt,
		"comment_count":   u.CommentCount,
		"view_count":      u.ViewCount,
		"liked_count":     u.LikedCount,
		"status":          u.Status,
		"type":            u.Type,
		"last_ip_address": u.LastIpAddress,
		"path":            p,
		"description":     u.Description,
		"public":          u.Public,
		"photo_count":     len(p),
		"geo":             u.Geo,

		"user":  user,
		"topic": topic,
	}
}
