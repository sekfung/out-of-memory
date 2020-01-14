package models

import (
	"github.com/jinzhu/gorm"
	"math"
	"outofmemory/errors"
)

type UserRole struct {
	BaseModel
	Uid    uint32 `gorm:"not null;index:uid" json:"uid"`
	RoleId uint32 `gorm:"not null;index:role_id" json:"role_id"`
}

func GetUserRole(uid uint32) (uint32, error) {
	var userRole UserRole
	err := db.Select("role_id").Where("uid = ?", uid).Find(&userRole).Error
	if err != nil {
		switch err {
		case gorm.ErrRecordNotFound:
			return math.MaxUint32, errors.ErrUserRoleNotExist
		default:
			return math.MaxUint32, errors.ErrGetUserRoleFailed
		}
	}
	return userRole.RoleId, nil
}

func AddUserRole(adminId, uid, roleId uint32) (uint32, error) {
	err := checkUserRole(adminId, roleId)
	if err != nil {
		return math.MaxUint32, err
	}
	userRoleId, err := GetUserRole(uid)
	if err != nil && err != errors.ErrUserRoleNotExist {
		return math.MaxUint32, err
	}
	if userRoleId  == roleId {
		return roleId, errors.ErrUserRoleIsExist
	}
	userRole := UserRole{
		Uid:uid,
		RoleId:roleId,
	}
	err = db.Create(&userRole).Error
	if err != nil {
		return math.MaxUint32, errors.ErrCreateUserRoleFailed
	}
	return roleId, err
}

func UpdateUserRole(adminId, uid, roleId uint32) (uint32, error) {
	err := checkUserRole(adminId, roleId)
	if err != nil {
		return math.MaxUint32, err
	}
	userRoleId, err := GetUserRole(uid)
	if err != nil {
		return userRoleId, err
	}
	err = db.Table("user_role").Where("uid = ?", uid).Update("role_id", roleId).Error
	if err != nil {
		return math.MaxUint32, errors.ErrUpdateUserRoleFailed
	}
	return roleId, err
}

func DeleteUserRole(adminId, uid, roleId uint32) (uint32, error) {
	err := checkUserRole(adminId, roleId)
	if err != nil {
		return math.MaxUint32, err
	}
	userRoleId, err := GetUserRole(uid)
	if err != nil {
		return userRoleId, err
	}
	err = db.Where("role_id = ? AND uid = ?", roleId, uid).Delete(&UserRole{}).Error
	if err != nil {
		return math.MaxUint32, errors.ErrDeleteUserRoleFailed
	}
	return roleId, err
}

func checkUserRole(adminId, roleId uint32) error {
	adminRoleId, err := GetUserRole(adminId)
	if err != nil {
		return err
	}
	// only system admin can create/update/delete user role
	if adminRoleId > SystemAdmin {
		return errors.ErrForBidden
	}
	if roleId > Member || roleId < SystemAdmin {
		return errors.ErrInvalidParam
	}
	return err
}