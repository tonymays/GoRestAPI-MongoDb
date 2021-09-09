package root

import (
	"errors"
)

type PermissionService interface {
	CreatePermission(p Permission) (Permission, error)
	FindPermission(p Permission) ([]Permission, error)
	UpdatePermission(f Permission, u Permission) (Permission, error)
}

type Permission struct {
	PermissionId	string	`json:"permission_id,omitempty"`
	Tag				string	`json:"tag,omitempty"`
	Active			string	`json:"active,omitempty"`
	Created			string	`json:"created,omitempty"`
	Modified		string	`json:"modified,omitempty"`
}

func (rcvr *Permission) Validate(opCreate bool) error {
	if opCreate {
		if len(rcvr.PermissionId) == 0 {return errors.New("missing permission id")}
		if len(rcvr.Tag) == 0 {return errors.New("missing tag")}
		if len(rcvr.Active) == 0 {return errors.New("missing active flag")}
		if len(rcvr.Created) == 0 {return errors.New("missing created timestamp")}
		if len(rcvr.Modified) == 0 {return errors.New("missing modified timestamp")}
	} else {
		if len(rcvr.PermissionId) > 0 {return errors.New("updating permission id not allowed")}
		if len(rcvr.Created) > 0 {return errors.New("updating created timestamp not allowed")}
	}
	return nil
}
