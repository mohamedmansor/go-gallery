package models

import "time"

type Image struct {
	ID       int64  `gorm:"primary_key;auto_increment" json:"id"`
	Title    string `gorm:"size:200" json:"title"`
	ImageUrl string `json:"imgpath"`
	// ImageUrl	string        `bson:"imgUrl"`
	CreatedAt time.Time `json:"created_at,omitempty"`
	UpdatedAt time.Time `json:"updated_at,omitempty"`
}

// TableName method sets table name for Image model
func (image *Image) TableName() string {
	return "image"
}

//ResponseMap -> response map method of Image
func (image *Image) ResponseMap() map[string]interface{} {
	resp := make(map[string]interface{})
	resp["id"] = image.ID
	resp["title"] = image.Title
	resp["created_at"] = image.CreatedAt
	resp["updated_at"] = image.UpdatedAt
	return resp
}
