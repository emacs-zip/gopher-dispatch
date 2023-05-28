package tenants

import (
	"gopher-dispatch/api/models"
	"gopher-dispatch/pkg/db"

	"github.com/google/uuid"
)

func GetTenants() ([]models.Tenant, error) {
	var tenants []models.Tenant
	if err := db.GetDB().Find(&tenants).Error; err != nil {
		return nil, err
	}
	return tenants, nil
}

func GetTenant(id uuid.UUID) (models.Tenant, error) {
	var tenant models.Tenant
	if err := db.GetDB().Where("id = ?", id).First(&tenant).Error; err != nil {
		return models.Tenant{}, err
	}
	return tenant, nil
}

func CreateTenant(newTenant models.Tenant) (models.Tenant, error) {
	// Validate newTenant data before saving to DB

	if err := db.GetDB().Create(&newTenant).Error; err != nil {
		return models.Tenant{}, err
	}
	return newTenant, nil
}

func UpdateTenant(id uuid.UUID, updatedTenant models.Tenant) (models.Tenant, error) {
	var tenant models.Tenant
	if err := db.GetDB().Where("id = ?", id).First(&tenant).Error; err != nil {
		return models.Tenant{}, err
	}

	if err := db.GetDB().Model(&tenant).Updates(updatedTenant).Error; err != nil {
		return models.Tenant{}, err
	}

	return tenant, nil
}

func DeleteTenant(id uuid.UUID) error {
	var tenant models.Tenant
	if err := db.GetDB().Where("id = ?", id).First(&tenant).Error; err != nil {
		return err
	}

	if err := db.GetDB().Delete(&tenant).Error; err != nil {
		return err
	}

	return nil
}

func GetTenantUsers(tenantID uuid.UUID) ([]models.User, error) {
	var users []models.User
	if err := db.GetDB().Where("tenant_id = ?", tenantID).Find(&users).Error; err != nil {
		return nil, err
	}
	return users, nil
}

func GetTenantPolicies(tenantID uuid.UUID) ([]models.Policy, error) {
	var policies []models.Policy
	if err := db.GetDB().Where("tenant_id = ?", tenantID).Find(&policies).Error; err != nil {
		return nil, err
	}
	return policies, nil
}
