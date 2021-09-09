package data_models

import (
	"pkg"
)

// ---- UserRoleModel ----
type UserRoleModel struct {
	UserRoleId	string	`bson:"user_role_id,omitempty"`
	UserId		string	`bson:"user_id,omitempty"`
	RoleId		string	`bson:"role_id,omitempty"`
	Active		string 	`bson:"active,omitempty"`
	Created		string 	`bson:"created,omitempty"`
	Modified	string 	`bson:"modified,omitempty"`
}

// ---- NewUserRoleModel ----
func NewUserRoleModel(rcvr root.UserRole) *UserRoleModel {
	return &UserRoleModel{
		UserRoleId:	rcvr.UserRoleId,
		UserId:		rcvr.UserId,
		RoleId:		rcvr.RoleId,
		Active:		rcvr.Active,
		Created:	rcvr.Created,
		Modified:	rcvr.Modified,
	}
}

// ---- UserRoleModel.ToRootUserRole ----
func (rcvr *UserRoleModel) ToRootUserRole() root.UserRole {
	return root.UserRole{
		UserRoleId:	rcvr.UserRoleId,
		UserId:		rcvr.UserId,
		RoleId:		rcvr.RoleId,
		Active:		rcvr.Active,
		Created:	rcvr.Created,
		Modified:	rcvr.Modified,
	}
}
