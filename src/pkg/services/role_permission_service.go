package services


import (
	"context"
	"errors"
	"go.mongodb.org/mongo-driver/mongo"
	"pkg"
	"pkg/configuration"
	"pkg/data_models"
	"time"
)

// ---- RolePermissionService ----
type RolePermissionService struct {
	config						configuration.Configuration
	dbClient					*mongo.Client
	rolePermissionsCollection	*mongo.Collection
	roleService					root.RoleService
	permissionService			root.PermissionService
}

// ---- NewRolePermissionService ----
func NewRolePermissionService(config configuration.Configuration, dbClient *mongo.Client, roleService root.RoleService, permissionService root.PermissionService) *RolePermissionService {
	rpc := dbClient.Database(config.DbName).Collection("role_permissions")
	return &RolePermissionService{config, dbClient, rpc, roleService, permissionService}
}

// ---- RolePermissionService.SetRolePermissions ----
func (rcvr *RolePermissionService) SetRolePermissions(role root.Role, tags []string) error {
	// step 1: verify that we have a good role
	_, err := rcvr.roleService.FindRole(role)
	if err != nil {
		return errors.New("role " + role.RoleId + " not found")
	}

	// step 2: verify permission tags given and fail on the first bad one
	var permissions []root.Permission
	var permission root.Permission
	for _, elPermission := range tags {
		permission.Tag = elPermission
		permission, err := rcvr.permissionService.FindPermission(permission)
		if err != nil {
			return errors.New("permission tag `" + elPermission + "` not found")
		}
		permissions = append(permissions, permission[0])
	}

	// step 3: delete the current role permissions
	err = rcvr.DeleteRolePermissions(role)
	if err != nil {
		return err
	}

	// step 4: write out the given permissions for the specified role
	for _, elPermission := range permissions {
		var rolePermission root.RolePermission
		rolePermission.RolePermissionId, _ = root.GenId()
		rolePermission.RoleId = role.RoleId
		rolePermission.PermissionId = elPermission.PermissionId
		rolePermission.Created = root.GenTimestamp()
		_, err := rcvr.InsertRolePermission(rolePermission)
		if err != nil {
			return err
		}
	}

	return nil
}

// ---- RolePermissionService.InsertRolePermissions ----
func (rcvr *RolePermissionService) InsertRolePermission(rolePermission root.RolePermission) (root.RolePermission, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 30 * time.Second)
	defer cancel()
	rolePermissionModel := data_models.NewRolePermissionModel(rolePermission)
	_, err := rcvr.rolePermissionsCollection.InsertOne(ctx, rolePermissionModel)
	return rolePermission, err
}

// ---- RolePermissionService.DeleteRolePermissions ----
func (rcvr *RolePermissionService) DeleteRolePermissions(role root.Role) error {
	var rolePermission root.RolePermission
	rolePermission.RoleId = role.RoleId
	filter := root.MakeBsonDQueryFilter(rolePermission)
	ctx, cancel := context.WithTimeout(context.Background(), 30 * time.Second)
	defer cancel()
	_, err := rcvr.rolePermissionsCollection.DeleteMany(ctx, filter)
	return err
}

// ---- RolePermissionService.FindRolePermissions ----
func (rcvr *RolePermissionService) FindRolePermission(role root.Role) ([]root.RolePermissionPayload, error) {
	// step 1: does the role exist
	_, err := rcvr.roleService.FindRole(role)
	if err != nil {
		return []root.RolePermissionPayload{}, errors.New("role `" + role.RoleId + "` does not exist")
	}

	// step 2: get role permissions
	var rolePermission root.RolePermission
	rolePermission.RoleId = role.RoleId
	ctx, cancel := context.WithTimeout(context.Background(), 30 * time.Second)
	filter := root.MakeBsonDQueryFilter(rolePermission)
	defer cancel()
	cursor, err := rcvr.rolePermissionsCollection.Find(ctx, filter)
	if err != nil {
		return []root.RolePermissionPayload{}, err
	}
	defer cursor.Close(ctx)

	var rolePermissions []root.RolePermissionPayload
	for cursor.Next(ctx) {
		var rolePermission = data_models.NewRolePermissionModel(root.RolePermission{})
		cursor.Decode(&rolePermission)
		var permission root.Permission
		var permissions []root.Permission
		permission.PermissionId = rolePermission.PermissionId
		permissions, err := rcvr.permissionService.FindPermission(permission)
		if err != nil {
			return []root.RolePermissionPayload{}, nil
		}
		var payload root.RolePermissionPayload
		payload.RolePermissionId = rolePermission.RolePermissionId
		payload.RoleId = rolePermission.RoleId
		payload.PermissionId = rolePermission.PermissionId
		payload.Tag = permissions[0].Tag
		payload.Created = rolePermission.Created
		rolePermissions = append(rolePermissions, payload)
	}


	return rolePermissions, nil
}
