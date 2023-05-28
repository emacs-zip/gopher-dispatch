package policies

import (
	"gopher-dispatch/api/models"
	"gopher-dispatch/pkg/db"

	"github.com/google/uuid"
)

func GetPolicies() ([]models.Policy, error) {
	var policies []models.Policy
	if err := db.GetDB().Find(&policies).Error; err != nil {
		return nil, err
	}
	return policies, nil
}

func GetPolicy(id uuid.UUID) (models.Policy, error) {
	var policy models.Policy
	if err := db.GetDB().Where("id = ?", id).First(&policy).Error; err != nil {
		return models.Policy{}, err
	}
	return policy, nil
}

func CreatePolicy(newPolicy models.Policy) (models.Policy, error) {
	// Validate newPolicy data before saving to DB

	if err := db.GetDB().Create(&newPolicy).Error; err != nil {
		return models.Policy{}, err
	}
	return newPolicy, nil
}

func UpdatePolicy(id uuid.UUID, updatedPolicy models.Policy) (models.Policy, error) {
	var policy models.Policy
	if err := db.GetDB().Where("id = ?", id).First(&policy).Error; err != nil {
		return models.Policy{}, err
	}

	if err := db.GetDB().Model(&policy).Updates(updatedPolicy).Error; err != nil {
		return models.Policy{}, err
	}

	return policy, nil
}

func DeletePolicy(id uuid.UUID) error {
	var policy models.Policy
	if err := db.GetDB().Where("id = ?", id).First(&policy).Error; err != nil {
		return err
	}

	if err := db.GetDB().Delete(&policy).Error; err != nil {
		return err
	}

	return nil
}
