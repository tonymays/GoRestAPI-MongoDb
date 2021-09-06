package data

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"pkg"
)

type userModel struct {
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

func newUserModel(rcvr root.User) *userModel {
	return &userModel{
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

func (rcvr *userModel) toRootUser() root.User {
	return root.User{
		Id: 		rcvr.Id.Hex(),
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
