package services

import (
	"context"
	"errors"
	"fmt"
	"go.mongodb.org/mongo-driver/mongo"
	"pkg"
	"pkg/configuration"
	"pkg/data_models"
	"time"
)

// ---- PermissionService ----
type PermissionService struct {
	config					configuration.Configuration
	dbClient				*mongo.Client
	permissionsCollection	*mongo.Collection
}

// ---- NewPermissionService ----
func NewPermissionService(config configuration.Configuration, dbClient *mongo.Client) *PermissionService {
	pc := dbClient.Database(config.DbName).Collection("permissions")
	return &PermissionService{config, dbClient, pc}
}

// ---- PermissionService.CreatePermission ----
func (rcvr *PermissionService) CreatePermission(p root.Permission) (root.Permission, error) {
	// does the email address exists
	var f root.Permission
	f.Tag = p.Tag
	_, err := rcvr.FindPermission(f)
	if err == nil {
		return root.Permission{}, errors.New("tag taken")
	}

	// establish active flag or fail is passed
	if len(p.Active) == 0 {
		p.Active = "Yes"
	} else {
		return root.Permission{}, errors.New("setting the active flag manually is not authorized")
	}

	// establish the user id or fail if passed
	if len(p.PermissionId) == 0 {
		id, err := root.GenId()
		if err != nil {
			return root.Permission{}, err
		}
		p.PermissionId = id
	} else {
		return root.Permission{}, errors.New("setting permission_id manually is not authorized")
	}

	// update the record timestamps
	p.Created = root.GenTimestamp()
	p.Modified = p.Created

	// validate the user and fail if we do not ave what we are looking for
	err = p.Validate(true)
	if err != nil {
		return root.Permission{}, err
	}

	// add the record or return err on insert error
	ctx, cancel := context.WithTimeout(context.Background(), 30 * time.Second)
	defer cancel()
	m := data_models.NewPermissionModel(p)
	_, err = rcvr.permissionsCollection.InsertOne(ctx, m)
	if err != nil {
		return root.Permission{}, err
	}

	// return what was given
	return p, nil
}

// ---- PermissionService.FindPermission ----
func (rcvr *PermissionService) FindPermission(p root.Permission) ([]root.Permission, error) {
	// establish a nil slice
	var permissions []root.Permission

	// make the filter and query the users collection
	ctx, cancel := context.WithTimeout(context.Background(), 30 * time.Second)
	defer cancel()
	filter := root.MakeBsonDQueryFilter(p)
	fmt.Println(filter)
	count := 0
	cursor, err := rcvr.permissionsCollection.Find(ctx, filter)
	if err != nil {
		return permissions, err
	}
	defer cursor.Close(ctx)

	// walk the cursor returned
	for cursor.Next(ctx) {
		var permission = data_models.NewPermissionModel(root.Permission{})
		cursor.Decode(&permission)
		permissions = append(permissions, permission.ToRootPermission())
		count++
	}

	// no permissions found then toss back an error
	if count == 0 {
		return permissions, errors.New("no permissions found")
	}

	// otherwise, return the permissions
	return permissions, nil
}

// ---- PermissionService.UpdatePermission
func (rcvr *PermissionService) UpdatePermission(f root.Permission, u root.Permission) (root.Permission, error) {
	// find the permission
	_, err := rcvr.FindPermission(f)

	// if the permission exists ...
	if err == nil {
		// ... then update the permission
		ctx, cancel := context.WithTimeout(context.Background(), 30 * time.Second)
		defer cancel()
		u.Modified = root.GenTimestamp()
		filter := root.MakeBsonDQueryFilter(f)
		update := root.MakeBsonDUpdateQueryFilter(u)
		_, err := rcvr.permissionsCollection.UpdateMany(ctx, filter, update)
		if err != nil {
			return root.Permission{}, err
		}
		return u, nil
	// otherwise, toss back an error
	} else {
		return root.Permission{}, errors.New("permission not found")
	}
}