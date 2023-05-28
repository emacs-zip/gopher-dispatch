package users

import (
	"gopher-dispatch/api/models"
	"gopher-dispatch/pkg/db"

	"github.com/google/uuid"
)

func GetUsers() ([]models.User, error) {
	// Implementation to fetch users from database
	// Assume db is your database instance
	var users []models.User
	if err := db.GetDB().Find(&users).Error; err != nil {
		return nil, err
	}
	return users, nil
}

func GetUser(id uuid.UUID) (models.User, error) {
	var user models.User
	if err := db.GetDB().Where("id = ?", id).First(&user).Error; err != nil {
		return models.User{}, err
	}
	return user, nil
}

func CreateUser(newUser models.User) (models.User, error) {
	// Validate newUser data before saving to DB

	if err := db.GetDB().Create(&newUser).Error; err != nil {
		return models.User{}, err
	}
	return newUser, nil
}

func UpdateUser(id uuid.UUID, updatedUser models.User) (models.User, error) {
	// Validate updatedUser data before saving to DB

	var user models.User
	if err := db.GetDB().Where("id = ?", id).First(&user).Error; err != nil {
		return models.User{}, err
	}

	if err := db.GetDB().Model(&user).Updates(updatedUser).Error; err != nil {
		return models.User{}, err
	}

	return user, nil
}

func DeleteUser(id uuid.UUID) error {
	var user models.User
	if err := db.GetDB().Where("id = ?", id).First(&user).Error; err != nil {
		return err
	}

	if err := db.GetDB().Delete(&user).Error; err != nil {
		return err
	}

	return nil
}

func GetUserAttributes(userId uuid.UUID) ([]models.UserAttribute, error) {
	var attributes []models.UserAttribute
	if err := db.GetDB().Where("user_id = ?", userId).Find(&attributes).Error; err != nil {
		return nil, err
	}
	return attributes, nil
}

func GetUserAttribute(userId, attributeId uuid.UUID) (models.UserAttribute, error) {
	var attribute models.UserAttribute
	if err := db.GetDB().Where("user_id = ? AND id = ?", userId, attributeId).First(&attribute).Error; err != nil {
		return models.UserAttribute{}, err
	}
	return attribute, nil
}

func AddUserAttribute(userId uuid.UUID, newAttribute models.UserAttribute) (models.UserAttribute, error) {
	newAttribute.UserID = userId
	if err := db.GetDB().Create(&newAttribute).Error; err != nil {
		return models.UserAttribute{}, err
	}
	return newAttribute, nil
}

func UpdateUserAttribute(userId, attributeId uuid.UUID, updatedAttribute models.UserAttribute) (models.UserAttribute, error) {
	var attribute models.UserAttribute
	if err := db.GetDB().Where("user_id = ? AND id = ?", userId, attributeId).First(&attribute).Error; err != nil {
		return models.UserAttribute{}, err
	}

	if err := db.GetDB().Model(&attribute).Updates(updatedAttribute).Error; err != nil {
		return models.UserAttribute{}, err
	}

	return attribute, nil
}

func DeleteUserAttribute(userId, attributeId uuid.UUID) error {
	var attribute models.UserAttribute
	if err := db.GetDB().Where("user_id = ? AND id = ?", userId, attributeId).First(&attribute).Error; err != nil {
		return err
	}

	if err := db.GetDB().Delete(&attribute).Error; err != nil {
		return err
	}

	return nil
}
