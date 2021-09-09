package data_models

import (
	"pkg"
)

// ---- RoleModel ----
type RolePermissionModel struct {
	RolePermissionId	string	`bson:"role_permission_id,omitempty"`
	RoleId				string 	`bson:"role_id,omitempty"`
	PermissionId		string 	`bson:"permission_id,omitempty"`
	Created				string 	`bson:"created,omitempty"`
}

// ---- NewRolePermissionModel ----
func NewRolePermissionModel(rcvr root.RolePermission) *RolePermissionModel {
	return &RolePermissionModel{
		RolePermissionId:	rcvr.RolePermissionId,
		RoleId:				rcvr.RoleId,
		PermissionId:		rcvr.PermissionId,
		Created:			rcvr.Created,
	}
}

// ---- RolePermissionModel.ToRootRolePermission ----
func (rcvr *RolePermissionModel) ToRootRolePermission() root.RolePermission {
	return root.RolePermission{
		RolePermissionId:	rcvr.RolePermissionId,
		RoleId:				rcvr.RoleId,
		PermissionId:		rcvr.PermissionId,
		Created:			rcvr.Created,
	}
}
