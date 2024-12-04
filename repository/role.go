package repository

import (
	"mainbe/model"
	"gorm.io/gorm"
)

func GetAllRole(db *gorm.DB) ([]model.Role, error) {
	var roles []model.Role

	if err := db.Find(&roles).Error; err != nil {
		return nil, err
	}

	return roles, nil
}

func GetRoleById(db *gorm.DB, idRole int) (model.Role, error) {
	var role model.Role

	if err := db.First(&role, idRole).First(&role).Error; err != nil {
		return role, err
	}

	return role, nil
}

func InsertRole(db *gorm.DB, role *model.Role) error {
	if err := db.Create(&role).Error; err != nil {
		return err
	}

	return nil
}

func UpdateRole(db *gorm.DB, role *model.Role) error {
	if err := db.Save(&role).Error; err != nil {
		return err
	}

	return nil
}

func DeleteRole(db *gorm.DB, role *model.Role) error {
	if err := db.Delete(&role).Error; err != nil {
		return err
	}

	return nil
}