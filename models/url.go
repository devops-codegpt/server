package models

// Url stores all api paths of the current application
type Url struct {
	Model
	Method   string `gorm:"comment:'request method'" json:"method"`
	Path     string `gorm:"comment:'url path'" json:"path"`
	Category string `gorm:"comment:'category'" json:"category"`
	Desc     string `gorm:"comment:'description'" json:"desc"`
	Creator  string `gorm:"comment:'creator'" json:"creator"`
}
