package services

import (
	"context"
	"errors"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"pkg"
	"pkg/configuration"
	"pkg/data_models"
	"time"
)

// ---- UserService ----
type UserService struct {
	config					configuration.Configuration
	dbClient				*mongo.Client
	roleService				root.RoleService
	rolePermissionService	root.RolePermissionService
	usersCollection			*mongo.Collection
	userRolesCollection		*mongo.Collection
}

// ---- NewUserService ----
func NewUserService(config configuration.Configuration, dbClient *mongo.Client, roleService root.RoleService, rolePermissionService root.RolePermissionService) *UserService {
	uc := dbClient.Database(config.DbName).Collection("users")
	urc := dbClient.Database(config.DbName).Collection("user_roles")
	return &UserService{config, dbClient, roleService, rolePermissionService, uc, urc}
}

// ---- UserService.CreateUser ----
func (rcvr *UserService) CreateUser(u root.User) (root.User, error) {
	// does the email address exists
	var f root.User
	f.Email = u.Email
	_, err := rcvr.FindUser(f)
	if err == nil {
		return root.User{}, errors.New("email address taken")
	}

	// establish active flag or fail is passed
	if len(u.Active) == 0 {
		u.Active = "Yes"
	} else {
		return root.User{}, errors.New("setting the active flag manually is not authorized")
	}

	// establish the user id or fail if passed
	if len(u.UserId) == 0 {
		id, err := root.GenId()
		if err != nil {
			return root.User{}, err
		}
		u.UserId = id
	} else {
		return root.User{}, errors.New("setting user_id manually is not authorized")
	}

	// validate the user and fail if we do not ave what we are looking for
	err = u.Validate(true)
	if err != nil {
		return root.User{}, err
	}

	// has the password given
	hp, err := u.HashPassword(u.Password)
	if err != nil {
		return root.User{}, err
	}
	u.Password = hp

	// update the record timestamps
	u.Created = root.GenTimestamp()
	u.Modified = u.Created

	// add the record or return err on insert error
	ctx, cancel := context.WithTimeout(context.Background(), 30 * time.Second)
	defer cancel()
	m := data_models.NewUserModel(u)
	_, err = rcvr.usersCollection.InsertOne(ctx, m)
	if err != nil {
		return root.User{}, err
	}

	// return what was given
	return u, nil
}

// ---- UserService.FindUser ----
func (rcvr *UserService) FindUser(u root.User) ([]root.User, error) {
	// establish a nil slice
	var users []root.User

	// make the filter and query the users collection
	ctx, cancel := context.WithTimeout(context.Background(), 30 * time.Second)
	defer cancel()
	filter := root.MakeBsonDQueryFilter(u)
	count := 0
	cursor, err := rcvr.usersCollection.Find(ctx, filter)
	if err != nil {
		return users, err
	}
	defer cursor.Close(ctx)

	// walk the cursor returned
	for cursor.Next(ctx) {
		var user = data_models.NewUserModel(root.User{})
		cursor.Decode(&user)
		users = append(users, user.ToRootUser())
		count++
	}

	// no users found then toss back an error
	if count == 0 {
		return users, errors.New("no users found")
	}

	// otherwise, return the users
	return users, nil
}

// ---- UserService.UpdateUser ----
func (rcvr *UserService) UpdateUser(f root.User, u root.User) (root.User, error) {
	// find the user
	_, err := rcvr.FindUser(f)

	// if the user exists ...
	if err == nil {
		// ... then update the user
		ctx, cancel := context.WithTimeout(context.Background(), 30 * time.Second)
		defer cancel()
		u.Modified = root.GenTimestamp()
		filter := root.MakeBsonDQueryFilter(f)
		update := root.MakeBsonDUpdateQueryFilter(u)
		_, err := rcvr.usersCollection.UpdateMany(ctx, filter, update)
		if err != nil {
			return root.User{}, err
		}
		return u, nil
	// otherwise, toss back an error
	} else {
		return root.User{}, errors.New("user not found")
	}
}

// ---- UserService.AssignUserRole ----
func (rcvr *UserService) AssignUserRole(userRole root.UserRole) error {
	// step 1: validate the assignment
	err := userRole.Validate(true)
	if err != nil {
		return err
	}

	// step 2: find the user
	var user root.User
	user.UserId = userRole.UserId
	_, err = rcvr.FindUser(user)
	if err != nil {
		return errors.New("user not found")
	}

	// step 3: find the role
	var role root.Role
	role.RoleId = userRole.RoleId
	_, err = rcvr.roleService.FindRole(role)
	if err != nil {
		return errors.New("role not found")
	}

	// step 4: add the document
	ctx, cancel := context.WithTimeout(context.Background(), 30 * time.Second)
	defer cancel()
	m := data_models.NewUserRoleModel(userRole)
	_, err = rcvr.userRolesCollection.InsertOne(ctx, m)

	// step 5: return any errors
	return err
}

// ---- UserService.FindUserRole ----
func (rcvr *UserService) FindUserRole(userRole root.UserRole) ([]root.UserRoles, error) {
	// establish a nil slice
	var userRoles []root.UserRoles

	// make the filter and query the userRoles collection
	ctx, cancel := context.WithTimeout(context.Background(), 30 * time.Second)
	defer cancel()
	filter := root.MakeBsonDQueryFilter(userRole)
	count := 0
	cursor, err := rcvr.userRolesCollection.Find(ctx, filter)
	if err != nil {
		return userRoles, err
	}
	defer cursor.Close(ctx)

	// walk the cursor returned
	for cursor.Next(ctx) {
		var userRole = data_models.NewUserRoleModel(root.UserRole{})
		cursor.Decode(&userRole)
		var payload root.UserRoles
		payload.UserRoleId = userRole.UserRoleId
		payload.UserId = userRole.UserId
		payload.RoleId = userRole.RoleId
		payload.Active = userRole.Active
		payload.Created = userRole.Created
		payload.Modified = userRole.Modified

		// grab the user name
		var user root.User
		user.UserId = userRole.UserId
		users, err := rcvr.FindUser(user)
		if err != nil {
			return []root.UserRoles{}, nil
		}
		payload.Username = users[0].Username

		// grab the role name
		var role root.Role
		role.RoleId = userRole.RoleId
		roles, err := rcvr.roleService.FindRole(role)
		if err != nil {
			return []root.UserRoles{}, nil
		}
		payload.Rolename = roles[0].Name

		// append a new userRole
		userRoles = append(userRoles, payload)
		count++
	}

	// no userRoles found then toss back an error
	if count == 0 {
		return userRoles, errors.New("no user roles found")
	}

	// otherwise, return the users
	return userRoles, nil
}

// ---- UserService.ActivateUserRole ----
func (rcvr *UserService) ActivateUserRole(f root.UserRole, u root.UserRole) error {
	_, err := rcvr.FindUserRole(f)
	if err == nil {
		// ... then update the userrole
		ctx, cancel := context.WithTimeout(context.Background(), 30 * time.Second)
		defer cancel()
		filter := root.MakeBsonDQueryFilter(f)
		update := root.MakeBsonDUpdateQueryFilter(u)
		_, err := rcvr.userRolesCollection.UpdateMany(ctx, filter, update)
		if err != nil {
			return err
		}
	// otherwise, toss back an error
	} else {
		return errors.New("user roles not found")
	}
	return nil
}

// ---- UserService.GetServiceCatalog ----
func (rcvr *UserService) GetServiceCatalog(u root.User) ([]string, error) {
	var serviceCatalog []string
	_, err := rcvr.FindUser(u)
	if err != nil {
		return []string{}, errors.New("user not found")
	}
	var userRole root.UserRole
	userRole.UserId = u.UserId
	userRoles, err := rcvr.FindUserRole(userRole)
	if err != nil {
		return []string{}, err
	}
	for _, elUserRole := range userRoles {
		var role root.Role
		role.RoleId = elUserRole.RoleId
		rolePermissions, _ := rcvr.rolePermissionService.FindRolePermission(role)
		for _, elRolePermission := range rolePermissions {
			serviceCatalog = append(serviceCatalog, elRolePermission.Tag)
		}
	}
	serviceCatalog = rcvr.removeDuplicateStrings(serviceCatalog)
	return serviceCatalog, nil
}

// ---- UserService.CountUsers ----
func (rcvr *UserService) CountUser() int64 {
	ctx, cancel := context.WithTimeout(context.Background(), 30 * time.Second)
	defer cancel()
	count, _ := rcvr.usersCollection.CountDocuments(ctx, bson.D{})
	return count
}

// ---- UserService.removeDuplicateStrings ----
func (rcvr *UserService) removeDuplicateStrings(strList []string) []string {
	keys := make(map[string]bool)
	list := []string{}
	for _, entry := range strList {
		if _, value := keys[entry]; !value {
			keys[entry] = true
			list = append(list, entry)
		}
	}
	return list
}



