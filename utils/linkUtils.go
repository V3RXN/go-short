package utils

import (
	"github.com/google/uuid"
	"github.com/v3rxn/go-short/database"
	"github.com/v3rxn/go-short/models"
)

func GenerateUniqueId() string {
	var exists bool
	var id string
	for {
		id = generateTrimmedUUID()
		database.DB.Model(&models.Link{}).Where("id = ?", id).Find(&exists)
		if !exists{
			break
		}
	}
	return id
}

func generateTrimmedUUID() string {
	uuid := uuid.NewString()
	trimmedId := uuid[:6]
	return trimmedId
}