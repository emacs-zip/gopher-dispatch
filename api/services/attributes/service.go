package attributes

import (
	"gopher-dispatch/api/models"
	"gopher-dispatch/pkg/db"

	"github.com/google/uuid"
)

func GetAttributes() ([]models.Attribute, error) {
	var attributes []models.Attribute
	if err := db.GetDB().Find(&attributes).Error; err != nil {
		return nil, err
	}
	return attributes, nil
}

func GetAttribute(id uuid.UUID) (models.Attribute, error) {
	var attribute models.Attribute
	if err := db.GetDB().Where("id = ?", id).First(&attribute).Error; err != nil {
		return models.Attribute{}, err
	}
	return attribute, nil
}

func CreateAttribute(newAttribute models.Attribute) (models.Attribute, error) {
	// Validate newAttribute data before saving to DB

	if err := db.GetDB().Create(&newAttribute).Error; err != nil {
		return models.Attribute{}, err
	}
	return newAttribute, nil
}

func UpdateAttribute(id uuid.UUID, updatedAttribute models.Attribute) (models.Attribute, error) {
	var attribute models.Attribute
	if err := db.GetDB().Where("id = ?", id).First(&attribute).Error; err != nil {
		return models.Attribute{}, err
	}

	if err := db.GetDB().Model(&attribute).Updates(updatedAttribute).Error; err != nil {
		return models.Attribute{}, err
	}

	return attribute, nil
}

func DeleteAttribute(id uuid.UUID) error {
	var attribute models.Attribute
	if err := db.GetDB().Where("id = ?", id).First(&attribute).Error; err != nil {
		return err
	}

	if err := db.GetDB().Delete(&attribute).Error; err != nil {
		return err
	}

	return nil
}
