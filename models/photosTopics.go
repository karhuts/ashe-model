package models

type PhotosTopics struct {
	Id          *int64  `gorm:"column:id;primary_key;AUTO_INCREMENT;NOT NULL" json:"id,omitempty"`
	Name        *string `gorm:"column:name;default:;NOT NULL" json:"name,omitempty"`
	ViewCount   *int64  `gorm:"column:view_count;default:0;NOT NULL" json:"viewCount,omitempty"`
	PhotosCount *int64  `gorm:"column:photos_count;default:0;NOT NULL" json:"photosCount,omitempty"`
	Description *string `gorm:"column:description;NOT NULL" json:"description,omitempty"`
	CreateTime  *int64  `gorm:"column:create_time;default:0;NOT NULL" json:"createTime,omitempty"`
	Type        *int64  `gorm:"column:type;default:0;NOT NULL" json:"type,omitempty"`
}

func (p *PhotosTopics) ToMap() map[string]interface{} {
	return map[string]interface{}{
		"id":           p.Id,
		"name":         p.Name,
		"view_count":   p.ViewCount,
		"photos_count": p.PhotosCount,
		"description":  p.Description,
		"creat_time":   p.CreateTime,
		"type":         p.Type,
		}
}

func (p *PhotosTopics) TableName() string {
	return "photos_topics"
}
