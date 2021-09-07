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

// ---- RoleService ----
type RoleService struct {
	config			configuration.Configuration
	dbClient		*mongo.Client
	rolesCollection	*mongo.Collection
}

// ---- NewRoleService ----
func NewRoleService(config configuration.Configuration, dbClient *mongo.Client) *RoleService {
	rc := dbClient.Database(config.DbName).Collection("roles")
	return &RoleService{config, dbClient, rc}
}

// ---- RoleService.CreateRole ----
func (rcvr *RoleService) CreateRole(r root.Role) (root.Role, error) {
	// does the role name exists
	var f root.Role
	f.Name = r.Name
	_, err := rcvr.FindRole(f)
	if err == nil {
		return root.Role{}, errors.New("role exists")
	}

	// establish active flag or fail is passed
	if len(r.Active) == 0 {
		r.Active = "Yes"
	} else {
		return root.Role{}, errors.New("setting the active flag manually is not authorized")
	}

	// establish the role id or fail if passed
	if len(r.RoleId) == 0 {
		id, err := root.GenId()
		if err != nil {
			return root.Role{}, err
		}
		r.RoleId = id
	} else {
		return root.Role{}, errors.New("setting role_id manually is not authorized")
	}

	// validate the user and fail if we do not ave what we are looking for
	err = r.Validate(true)
	if err != nil {
		return root.Role{}, err
	}

	// update the record timestamps
	r.Created = root.GenTimestamp()
	r.Modified = r.Created

	// add the record or return err on insert error
	ctx, cancel := context.WithTimeout(context.Background(), 30 * time.Second)
	defer cancel()
	m := data.NewRoleModel(r)
	_, err = rcvr.rolesCollection.InsertOne(ctx, m)
	if err != nil {
		return root.Role{}, err
	}

	// return what was given
	return r, nil
}

// ---- RoleService.FindRole ----
func (rcvr *RoleService) FindRole(r root.Role) ([]root.Role, error) {
	// establish a nil slice
	var roles []root.Role

	// make the filter and query the roles collection
	ctx, cancel := context.WithTimeout(context.Background(), 30 * time.Second)
	defer cancel()
	filter := root.MakeBsonDQueryFilter(r)
	count := 0
	cursor, err := rcvr.rolesCollection.Find(ctx, filter)
	if err != nil {
		return roles, err
	}
	defer cursor.Close(ctx)

	// walk the cursor returned
	for cursor.Next(ctx) {
		var role = data.NewRoleModel(root.Role{})
		cursor.Decode(&role)
		roles = append(roles, role.ToRootRole())
		count++
	}

	// no users found then toss back an error
	if count == 0 {
		return roles, errors.New("no roles found")
	}

	// otherwise, return the roles
	return roles, nil
}

// ---- RoleService.UpdateRole
func (rcvr *RoleService) UpdateRole(f root.Role, u root.Role) (root.Role, error) {
	// find the role
	_, err := rcvr.FindRole(f)

	// if the role exists ...
	if err == nil {
		// ... then update the role
		ctx, cancel := context.WithTimeout(context.Background(), 30 * time.Second)
		defer cancel()
		u.Modified = root.GenTimestamp()
		filter := root.MakeBsonDQueryFilter(f)
		update := root.MakeBsonDUpdateQueryFilter(u)
		_, err := rcvr.rolesCollection.UpdateMany(ctx, filter, update)
		if err != nil {
			return root.Role{}, err
		}
		return u, nil
	// otherwise, toss back an error
	} else {
		return root.Role{}, errors.New("role not found")
	}
}