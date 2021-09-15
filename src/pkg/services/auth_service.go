package services

import (
	"context"
	"errors"
	"go.mongodb.org/mongo-driver/mongo"
	"pkg"
	"pkg/configuration"
	"pkg/data_models"
	"pkg/server"
	"strings"
	"time"
)

// ---- AuthService ----
type AuthService struct {
	config				configuration.Configuration
	dbClient			*mongo.Client
	userService			root.UserService
	blacklistCollection *mongo.Collection
	loginCollection		*mongo.Collection
}

// ---- NewAuthService ----
func NewAuthService(config configuration.Configuration, dbClient *mongo.Client, userService root.UserService) *AuthService {
	blc := dbClient.Database(config.DbName).Collection("blacklists")
	lc := dbClient.Database(config.DbName).Collection("logins")
	return &AuthService{config, dbClient, userService, blc, lc}
}

// ---- AuthService.StartSession ----
func (rcvr *AuthService) StartSession(payload root.AuthPayload) (root.UserToken, error) {
	// setup a nil user token struct
	var userToken root.UserToken

	// blacklist any token specified in the payload
	err := rcvr.blacklistAuthToken(payload)
	if err != nil {
		return root.UserToken{}, err
	}

	// setup nil user and users struct
	var user root.User

	// are we logging in with email or username?
	// I allowed both but you do not have to
	if strings.Contains(payload.Username, "@") {
		user.Email = payload.Username
	} else {
		user.Username = payload.Username
	}
	user.Active = "Yes"

	// setup login recording struct
	var login root.Login
	login.LoginId, _ = root.GenId()
	login.Username = payload.Username
	login.Success = "No"
	login.Created = root.GenTimestamp()
	login.Modified = login.Created

	// find the user from the userService
	users, err := rcvr.userService.FindUser(user)
	if err != nil {
		rcvr.recordLoginAttempt(login)
		return root.UserToken{}, errors.New("user not found")
	}

	// validate the password
	if !user.ValidatePassword(payload.Password, users[0].Password) {
		rcvr.recordLoginAttempt(login)
		return root.UserToken{}, errors.New("user not found")
	}

	// establish the userToken struct if we make this far
	user = root.User{}
	user.UserId = users[0].UserId
	userToken.UserId = users[0].UserId
	userToken.Username = users[0].Username
	userToken.Email	= users[0].Email
	userToken.RemoteAddr = payload.LoginIp
	userToken.ServiceCatalog, err = rcvr.userService.GetServiceCatalog(user)
	if err != nil {
		userToken.ServiceCatalog = []string{}
	}

	// finalize the login recording
	login.UserId = users[0].UserId
	login.Username = users[0].Username
	login.Email = users[0].Email
	login.Success = "Yes"
	rcvr.recordLoginAttempt(login)

	// so ... let them eat cake
	return userToken, nil
}

// ---- AuthService.KillSession ----
func (rcvr *AuthService) KillSession(a root.AuthPayload) error {
	return rcvr.blacklistAuthToken(a)
}

// ---- AuthService.CheckSession ----
func (rcvr *AuthService) CheckSession(payload root.AuthPayload) error {
	// lets decode the token for the vitals
	userToken := server.DecodeJWT(payload.AuthToken, rcvr.config)
	if userToken.Username == "" {
		return errors.New("invalid token")
	}

	// find that user and we are all good
	var u root.User
	u.UserId = userToken.UserId
	_, err := rcvr.userService.FindUser(u)
	if err != nil {
		return errors.New("invalid token")
	}

	// return no error
	return nil
}

// ---- AuthService.ChangePasword ----
func (rcvr *AuthService) ChangePassword(payload root.ChangePasswordPayload) error {
	if payload.Username == "" {
		return errors.New ("invalid credentials")
	} else 	if payload.Password == "" {
		return errors.New ("invalid credentials")
	} else 	if payload.NewPassword == "" {
		return errors.New ("invalid credentials")
	} else 	if payload.Password == payload.NewPassword {
		return errors.New ("new password cannot be the same as the current password")
	}
	var filterUser root.User
	var updateUser root.User
	var user root.User
	var users []root.User
	if strings.Contains(payload.Username, "@") { user.Email = payload.Username } else { user.Username = payload.Username }
	user.Active = "Yes"
	users, err := rcvr.userService.FindUser(user)
	if err != nil {
		return errors.New("invalid credentials")
	}
	if len(users) > 1 {
		return errors.New("critical error: change password failed")
	}
	filterUser.UserId = users[0].UserId
	if !user.ValidatePassword(payload.Password, users[0].Password) {
		return errors.New("invalid credentials")
	}
	hashedPassword, _ := user.HashPassword(payload.NewPassword)
	updateUser.Password = string(hashedPassword)
	_, err = rcvr.userService.UpdateUser(filterUser, updateUser)
	return err
}

// ---- AuthService.blacklistAuthToken
func (rcvr *AuthService) blacklistAuthToken(a root.AuthPayload) error {
	if a.AuthToken != "" {
		var b root.Blacklist
		b.Id, _ = root.GenId()
		b.AuthToken = a.AuthToken
		b.Created = root.GenTimestamp()
		m := data_models.NewBlacklistModel(b)
		ctx, cancel := context.WithTimeout(context.Background(), 30 * time.Second)
		defer cancel()
		_, err := rcvr.blacklistCollection.InsertOne(ctx, m)
		return err
	}
	return nil
}

// ---- AuthService.recordLoginAttemp ----
func (rcvr *AuthService) recordLoginAttempt(l root.Login) error {
	m := data_models.NewLoginModel(l)
	ctx, cancel := context.WithTimeout(context.Background(), 30 * time.Second)
	defer cancel()
	_, err := rcvr.loginCollection.InsertOne(ctx, m)
	return err
}
