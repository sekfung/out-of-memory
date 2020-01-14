package models

import (
	"math"
	"outofmemory/errors"
	"outofmemory/models/test"
	"outofmemory/utils"
	"testing"
)

const adminID = 4107665161

func TestAddUserRole(t *testing.T) {
	test.SetupForTest()
	type addUserRole struct {
		adminId uint32
		uid     uint32
		roleId  uint32
	}
	var uid = utils.GenerateUID()
	tests := []struct {
		name    string
		args    addUserRole
		want    uint32
		wantErr error
	}{
		{
			"User Role Is Exist",
			addUserRole{
				adminId: adminID,
				uid:     adminID,
				roleId:  1,
			},
			1,
			errors.ErrUserRoleIsExist,
		},
		{
			"User Role Not Exist",
			addUserRole{
				adminId: utils.GenerateUID(),
				uid:     uid,
				roleId:  3,
			},
			math.MaxUint32,
			errors.ErrUserRoleNotExist,
		},
		{
			"Add User Role Successful",
			addUserRole{
				adminId: adminID,
				uid:     uid,
				roleId:  3,
			},
			3,
			nil,
		},
		{
			"Add User Role Forbidden",
			addUserRole{
				adminId: uid,
				uid:     utils.GenerateUID(),
				roleId:  1,
			},
			math.MaxUint32,
			errors.ErrForBidden,
		},
		{
			"Invalid User Role Id",
			addUserRole{
				adminId: adminID,
				uid:     uid,
				roleId:  math.MaxUint32,
			},
			math.MaxUint32,
			errors.ErrInvalidParam,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := AddUserRole(tt.args.adminId, tt.args.uid, tt.args.roleId)
			if err != nil && err != tt.wantErr {
				t.Errorf("AddUserRole() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("AddUserRole() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUpdateUserRole(t *testing.T) {
	test.SetupForTest()
	type updateUserRole struct {
		adminId uint32
		uid     uint32
		roleId  uint32
	}
	var uid = utils.GenerateUID()
	AddUserRole(4107665161, uid, 3)
	tests := []struct {
		name    string
		args    updateUserRole
		want    uint32
		wantErr error
	}{{
		"User Role Is Exist",
		updateUserRole{
			adminId: adminID,
			uid:     adminID,
			roleId:  1,
		},
		1,
		errors.ErrUserRoleIsExist,
	},
		{
			"User Role Not Exist",
			updateUserRole{
				adminId: utils.GenerateUID(),
				uid:     uid,
				roleId:  3,
			},
			math.MaxUint32,
			errors.ErrUserRoleNotExist,
		},
		{
			"User Role Not Exist",
			updateUserRole{
				adminId: adminID,
				uid:     utils.GenerateUID(),
				roleId:  3,
			},
			math.MaxUint32,
			errors.ErrUserRoleNotExist,
		},
		{
			"Update User Role Forbidden",
			updateUserRole{
				adminId: uid,
				uid:     utils.GenerateUID(),
				roleId:  1,
			},
			math.MaxUint32,
			errors.ErrForBidden,
		},
		{
			"Update User Role Successful",
			updateUserRole{
				adminId: adminID,
				uid:     uid,
				roleId:  1,
			},
			1,
			nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := UpdateUserRole(tt.args.adminId, tt.args.uid, tt.args.roleId)
			if err != nil && err != tt.wantErr {
				t.Errorf("UpdateUserRole() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("UpdateUserRole() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDeleteUserRole(t *testing.T) {
	test.SetupForTest()
	type deleteUserRole struct {
		adminId uint32
		uid     uint32
		roleId  uint32
	}
	var uid = utils.GenerateUID()
	AddUserRole(4107665161, uid, 3)
	tests := []struct {
		name    string
		args    deleteUserRole
		want    uint32
		wantErr error
	}{
		{
			"User Role Not Exist",
			deleteUserRole{
				adminId: utils.GenerateUID(),
				uid:     utils.GenerateUID(),
				roleId:  3,
			},
			math.MaxUint32,
			errors.ErrUserRoleNotExist,
		},
		{
			"User Role Not Exist",
			deleteUserRole{
				adminId: adminID,
				uid:     utils.GenerateUID(),
				roleId:  3,
			},
			math.MaxUint32,
			errors.ErrUserRoleNotExist,
		},
		{
			"Delete User Role Forbidden",
			deleteUserRole{
				adminId: uid,
				uid:     adminID,
				roleId:  1,
			},
			math.MaxUint32,
			errors.ErrForBidden,
		},
		{
			"Delete User Role Successful",
			deleteUserRole{
				adminId: adminID,
				uid:     uid,
				roleId:  1,
			},
			1,
			nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := DeleteUserRole(tt.args.adminId, tt.args.uid, tt.args.roleId)
			if err != nil && err != tt.wantErr {
				t.Errorf("DeleteUserRole() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("DeleteUserRole() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func BenchmarkAddUserRole(b *testing.B) {
	for i := 0; i < b.N; i++ {
		var uid = utils.GenerateUID()
		AddUserRole(adminID, uid, 3)
	}
}

func BenchmarkGetUserRole(b *testing.B) {
	for i := 0; i < b.N; i++ {
		GetUserRole(adminID)
	}
}

func BenchmarkUpdateUserRole(b *testing.B) {
	for i := 0; i < b.N; i++ {
		UpdateUserRole(adminID, adminID, 1)
	}
}

