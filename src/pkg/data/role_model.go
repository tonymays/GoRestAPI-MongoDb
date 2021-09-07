package data

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"pkg"
)

// ---- RoleModel ----
type RoleModel struct {
	Id			primitive.ObjectID	`bson:"_id,omitempty"`
	RoleId		string	 			`bson:"role_id,omitempty"`
	Name		string 				`bson:"name,omitempty"`
	Active		string 				`bson:"active,omitempty"`
	Created		string 				`bson:"created,omitempty"`
	Modified	string 				`bson:"modified,omitempty"`
}

// ---- NewRoleModel ----
func NewRoleModel(rcvr root.Role) *RoleModel {
	return &RoleModel{
		RoleId:		rcvr.RoleId,
		Name:		rcvr.Name,
		Active:		rcvr.Active,
		Created:	rcvr.Created,
		Modified:	rcvr.Modified,
	}
}

// ---- RoleModel.ToRootRole ----
func (rcvr *RoleModel) ToRootRole() root.Role {
	return root.Role{
		RoleId:		rcvr.RoleId,
		Name:		rcvr.Name,
		Active:		rcvr.Active,
		Created:	rcvr.Created,
		Modified:	rcvr.Modified,
	}
}
