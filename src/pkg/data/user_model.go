package data

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"pkg"
)

// ---- UserModel ----
type UserModel struct {
	Id			primitive.ObjectID	`bson:"_id,omitempty"`
	Userid		string	 			`bson:"user_id,omitempty"`
	Username	string 				`bson:"username,omitempty"`
	Password	string 				`bson:"password,omitempty"`
	Firstname	string 				`bson:"first_name,omitempty"`
	Lastname	string 				`bson:"last_name,omitempty"`
	Address		string 				`bson:"address,omitempty"`
	City		string 				`bson:"city,omitempty"`
	State		string 				`bson:"state,omitempty"`
	Zip			string 				`bson:"zip,omitempty"`
	Country		string 				`bson:"country,omitempty"`
	Email		string 				`bson:"email,omitempty"`
	Phone		string 				`bson:"phone,omitempty"`
	Active		string 				`bson:"active,omitempty"`
	Created		string 				`bson:"created,omitempty"`
	Modified	string 				`bson:"modified,omitempty"`
}

// ---- NewUserModel ----
func NewUserModel(rcvr root.User) *UserModel {
	return &UserModel{
		Userid:		rcvr.Userid,
		Username:	rcvr.Username,
		Password:	rcvr.Password,
		Firstname:	rcvr.Firstname,
		Lastname:	rcvr.Lastname,
		Address:	rcvr.Address,
		City:		rcvr.City,
		State:		rcvr.State,
		Zip:		rcvr.Zip,
		Country:	rcvr.Country,
		Email:		rcvr.Email,
		Phone:		rcvr.Phone,
		Active:		rcvr.Active,
		Created:	rcvr.Created,
		Modified:	rcvr.Modified,
	}
}

// ---- UserModel.ToRootUser ----
func (rcvr *UserModel) ToRootUser() root.User {
	return root.User{
		Userid:		rcvr.Userid,
		Username:	rcvr.Username,
		Password:	rcvr.Password,
		Firstname:	rcvr.Firstname,
		Lastname:	rcvr.Lastname,
		Address:	rcvr.Address,
		City:		rcvr.City,
		State:		rcvr.State,
		Zip:		rcvr.Zip,
		Country:	rcvr.Country,
		Email:		rcvr.Email,
		Phone:		rcvr.Phone,
		Active:		rcvr.Active,
		Created:	rcvr.Created,
		Modified:	rcvr.Modified,
	}
}
