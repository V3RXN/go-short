package models

type Link struct {
	Id 				string `gorm:"primaryKey"`
	URL 			string 
}