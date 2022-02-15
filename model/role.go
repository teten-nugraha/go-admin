package model

type Role struct {
	Id         uint         `json:"id" gorm:"primary_key, AUTO_INCREMENT"`
	Name       string       `json:"name"`
	Permission []Permission `json:"permissions" gorm:"many2many:role_permission"`
}
