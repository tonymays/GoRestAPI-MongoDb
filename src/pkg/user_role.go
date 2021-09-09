package root

import (
	"errors"
)

// --- UserRole ----
type UserRole struct {
	UserRoleId	string	`json:"user_role_id,omitempty"`
	UserId		string	`json:"user_id,omitempty"`
	RoleId		string	`json:"role_id,omitempty"`
	Active		string	`json:"active,omitempty"`
	Created		string	`json:"created,omitempty"`
	Modified	string	`json:"modified,omitempty"`
}

// ---- UserRoles ----
type UserRoles struct {
	UserRoleId	string	`json:"user_role_id,omitempty"`
	UserId		string	`json:"user_id,omitempty"`
	RoleId		string	`json:"role_id,omitempty"`
	Username	string	`json:"user_name,omitempty"`
	Rolename	string	`json:"role_name,omitempty"`
	Active		string	`json:"active,omitempty"`
	Created		string	`json:"created,omitempty"`
	Modified	string	`json:"modified,omitempty"`
}

// ---- UserRole.Validate ----
func (rcvr *UserRole) Validate(opCreate bool) error {
	if opCreate {
		if len(rcvr.UserRoleId) == 0 {return errors.New("missing user_role_id")}
		if len(rcvr.UserId) == 0 {return errors.New("missing user_id")}
		if len(rcvr.RoleId) == 0 {return errors.New("missing role_id")}
		if len(rcvr.Active) == 0 {return errors.New("missing active flag")}
		if len(rcvr.Created) == 0 {return errors.New("missing created timestamp")}
		if len(rcvr.Modified) == 0 {return errors.New("missing modified timestamp")}
	} else {
		if len(rcvr.UserRoleId) > 0 {return errors.New("updating user_role_id not allowed")}
		if len(rcvr.UserId) > 0 {return errors.New("updating user_id not allowed")}
		if len(rcvr.RoleId) > 0 {return errors.New("updating role_id not allowed")}
		if len(rcvr.Active) == 0 {return errors.New("missing active flag")}
		if len(rcvr.Created) > 0 {return errors.New("updating created timestamp not allowed")}
		if len(rcvr.Modified) == 0 {return errors.New("missing modified timestamp")}
	}
	return nil
}