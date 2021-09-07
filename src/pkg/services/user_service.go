package services

import (
	"context"
	"errors"
	"go.mongodb.org/mongo-driver/mongo"
	"pkg"
	"pkg/configuration"
	"pkg/data"
	"time"
)

// ---- UserService ----
type UserService struct {
	config			configuration.Configuration
	dbClient		*mongo.Client
	usersCollection	*mongo.Collection
}

// ---- NewUserService ----
func NewUserService(config configuration.Configuration, dbClient *mongo.Client) *UserService {
	uc := dbClient.Database(config.DbName).Collection("users")
	return &UserService{config, dbClient, uc}
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
	if len(u.Userid) == 0 {
		id, err := root.GenId()
		if err != nil {
			return root.User{}, err
		}
		u.Userid = id
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
	m := data.NewUserModel(u)
	_, err = rcvr.usersCollection.InsertOne(ctx, m)
	if err != nil {
		return root.User{}, err
	}

	// return what was given
	return u, nil
}

// ---- UserService.FindUser ----
func (rcvr *UserService) FindUser(u root.User) ([]root.User, error) {
	var users []root.User
	ctx, cancel := context.WithTimeout(context.Background(), 30 * time.Second)
	defer cancel()
	filter := u.MakeBsonDQueryFilter()
	count := 0
	cursor, err := rcvr.usersCollection.Find(ctx, filter)
	if err != nil {
		return users, err
	}
	defer cursor.Close(ctx)
	for cursor.Next(ctx) {
		var user = data.NewUserModel(root.User{})
		cursor.Decode(&user)
		users = append(users, user.ToRootUser())
		count++
	}
	if count == 0 {
		return users, errors.New("no users found")
	}
	return users, nil
}

// ---- UserService.UpdateUser
func (rcvr *UserService) UpdateUser(f root.User, u root.User) (root.User, error) {
	_, err := rcvr.FindUser(f)
	if err == nil {
		ctx, cancel := context.WithTimeout(context.Background(), 30 * time.Second)
		defer cancel()
		u.Modified = root.GenTimestamp()
		filter := f.MakeBsonDQueryFilter()
		update := u.MakeBsonDUpdateQueryFilter()
		_, err := rcvr.usersCollection.UpdateMany(ctx, filter, update)
		if err != nil {
			return root.User{}, err
		}
		return root.User{}, nil
	} else {
		return u, errors.New("user not found")
	}
}
