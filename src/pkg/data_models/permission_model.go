package data_models

import (
	"pkg"
)

// ---- PermissionModel ----
type PermissionModel struct {
	PermissionId	string	`bson:"permission_id,omitempty"`
	Tag				string 	`bson:"tag,omitempty"`
	Active			string 	`bson:"active,omitempty"`
	Created			string 	`bson:"created,omitempty"`
	Modified		string 	`bson:"modified,omitempty"`
}

// ---- NewPermissionModel ----
func NewPermissionModel(rcvr root.Permission) *PermissionModel {
	return &PermissionModel{
		PermissionId:		rcvr.PermissionId,
		Tag:				rcvr.Tag,
		Active:				rcvr.Active,
		Created:			rcvr.Created,
		Modified:			rcvr.Modified,
	}
}

// ---- PermissionModel.ToRootPermission ----
func (rcvr *PermissionModel) ToRootPermission() root.Permission {
	return root.Permission{
		PermissionId:		rcvr.PermissionId,
		Tag:				rcvr.Tag,
		Active:				rcvr.Active,
		Created:			rcvr.Created,
		Modified:			rcvr.Modified,
	}
}
