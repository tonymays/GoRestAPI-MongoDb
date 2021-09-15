package data_models

import (
	"pkg"
)

// ---- LoginModel ----
type LoginModel struct {
	LoginId		string 	`bson:"login_id,omitempty"`
	UserId		string 	`bson:"user_id,omitempty"`
	Username 	string	`bson:"username,omitempty"`
	Email 		string 	`bson:"email,omitempty"`
	Success		string 	`bson:"success,omitempty"`
	Created 	string 	`bson:"created,omitempty"`
	Modified	string 	`bson:"modified,omitempty"`
}

// ---- NewLoginModel ----
func NewLoginModel(rcvr root.Login) *LoginModel {
	return &LoginModel{
		LoginId:	rcvr.LoginId,
		UserId:		rcvr.UserId,
		Username:	rcvr.Username,
		Email:		rcvr.Email,
		Success:	rcvr.Success,
		Created:	rcvr.Created,
		Modified:	rcvr.Modified,
	}
}

// ---- LoginModel.ToRootLogin ----
func  (rcvr *LoginModel) ToRootLogin() root.Login {
	return root.Login{
		LoginId:	rcvr.LoginId,
		UserId:		rcvr.UserId,
		Username:	rcvr.Username,
		Email:		rcvr.Email,
		Created:	rcvr.Created,
		Success:	rcvr.Success,
		Modified:	rcvr.Modified,
	}
}

// ---- BlackListModel ----
type BlacklistModel struct {
	Id 			string	`bson:"id,omitempty"`
	AuthToken	string	`bson:"auth_token,omitempty"`
	Created		string	`bson:"created,omitempty"`
}

// ---- NewBlackListModel ----
func NewBlacklistModel(rcvr root.Blacklist) *BlacklistModel {
	return &BlacklistModel{
		Id:			rcvr.Id,
		AuthToken:	rcvr.AuthToken,
		Created:	rcvr.Created,
	}
}

// ---- BlacklistModel.ToRootBlacklist ----
func (rcvr *BlacklistModel) ToRootBlacklist() root.Blacklist {
	return root.Blacklist{
		Id:			rcvr.Id,
		AuthToken:	rcvr.AuthToken,
		Created: 	rcvr.Created,
	}
}
