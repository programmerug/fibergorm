// 文件名：entities/dog.go
package entities

import "gorm.io/gorm"

type Dog struct {
	gorm.Model
	Name      string `json:"name"`
	Age       int    `json:"age"`
	Breed     string `json:"breed"`
	IsGoodBoy bool   `json:"isGoodBoy" gorm:"default:true"`
}
