package root

// ---- AuthService ----
type AuthService interface {
	StartSession(payload AuthPayload) (UserToken, error)
	KillSession(auth AuthPayload) error
	CheckSession(payload AuthPayload) error
	ChangePassword(payload ChangePasswordPayload) error
}

// ---- AuthPayload ----
type AuthPayload struct {
	UserId		string `json:"user_id,omitempty"`
	Username	string `json:"username,omitempty"`
	Password	string `json:"password,omitempty"`
	AuthToken	string `json:"auth_token,omitempty"`
	LoginIp		string `json:"login_ip,omitempty"`
}

// ---- ChangePasswordPayload ----
type ChangePasswordPayload struct {
	Username 	string `json:"username,omitempty"`
	Password 	string `json:"password,omitempty"`
	NewPassword string `json:"new_password,omitempty"`
}

// ---- Login ----
type Login struct {
	LoginId		string 	`json:"login_id,omitempty"`
	UserId		string 	`json:"user_id,omitempty"`
	Username 	string	`json:"username,omitempty"`
	Email 		string 	`json:"email,omitempty"`
	Success		string 	`json:"success,omitempty"`
	Created 	string 	`json:"created,omitempty"`
	Modified	string 	`json:"modified,omitempty"`
}

// ---- Blacklist ----
type Blacklist struct {
	Id	 				string `json:"id,omitempty"`
	AuthToken        	string `json:"auth_token,omitempty"`
	Created 			string `json:"created,omitempty"`
}