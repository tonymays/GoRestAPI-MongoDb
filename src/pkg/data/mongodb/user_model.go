package mongodb

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"pkg"
)

type userModel struct {
	Mid			primitive.ObjectID	`bson:"_id,omitempty"`
	Id			string 				`bson:"id,omitempty"`
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
		Id:			rcvr.Id,
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
		Mid: 		rcvr.Mid.Hex( ),
		Id:			rcvr.Id,
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
