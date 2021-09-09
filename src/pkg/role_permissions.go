package root

import (
	"errors"
)

// ---- RolePermissionService ----
type RolePermissionService interface {
	SetRolePermissions(role Role, tags []string) error
	InsertRolePermission(rolePermission RolePermission) (RolePermission, error)
	DeleteRolePermissions(role Role) error
	FindRolePermission(role Role) ([]RolePermissionPayload, error)
}

// ---- RolePermission ----
type RolePermission struct {
	RolePermissionId	string	`json:"role_permission_id,omitempty"`
	RoleId				string	`json:"role_id,omitempty"`
	PermissionId		string	`json:"permission_id,omitempty"`
	Created				string	`json:"created,omitempty"`
}

// ---- RolePermissionPaylpad ----
type RolePermissionPayload struct {
	RolePermissionId	string	`json:"role_permission_id,omitempty"`
	RoleId				string	`json:"role_id,omitempty"`
	PermissionId		string	`json:"permission_id,omitempty"`
	Tag					string	`json:"tag,omitempty"`
	Created				string	`json:"created,omitempty"`
}

// ---- RolePermission.Validate ----
func (rcvr *RolePermission) Validate(opCreate bool) error {
	if opCreate {
		if len(rcvr.RolePermissionId) == 0 {return errors.New("missing role_permission_id")}
		if len(rcvr.RoleId) == 0 {return errors.New("missing role_id")}
		if len(rcvr.PermissionId) == 0 {return errors.New("missing permission_id")}
		if len(rcvr.Created) == 0 {return errors.New("missing created timestamp")}
	} else {
		return errors.New("updating a role permission is not allowed")
	}
	return nil
}