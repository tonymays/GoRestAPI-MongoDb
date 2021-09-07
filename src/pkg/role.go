package root

import (
	"errors"
)

// ---- RoleService ----
type RoleService interface {
	CreateRole(r Role) (Role, error)
	FindRole(r Role) ([]Role, error)
	UpdateRole(f Role, u Role) (Role, error)
}

// ---- Role ----
type Role struct {
	Id			string	`json:"_id,omitempty"`
	RoleId		string	`json:"role_id,omitempty"`
	Name		string	`json:"name,omitempty"`
	Active		string	`json:"active,omitempty"`
	Created		string	`json:"created,omitempty"`
	Modified	string	`json:"modified,omitempty"`
}

// ---- Role.Validate ----
func (rcvr *Role) Validate(opCreate bool) error {
	if opCreate {
		if len(rcvr.RoleId) == 0 {return errors.New("missing role id")}
		if len(rcvr.Name) == 0 {return errors.New("missing name")}
		if len(rcvr.Active) == 0 {return errors.New("missing active flag")}
	} else {
		if len(rcvr.RoleId) > 0 {return errors.New("update role id not permitted")}
	}
	return nil
}