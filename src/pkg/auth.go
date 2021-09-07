package root

// ---- AuthService ----
type AuthService interface {
}

// ---- AuthPayload ----
type AuthPayload struct {
	Id					string `json:"id,omitempty"`
	Username         	string `json:"username,omitempty"`
	Password         	string `json:"password,omitempty"`
	LoginIp				string `json:"login_ip,omitempty"`
}

// ---- ChangePasswordPayload ----
type ChangePasswordPayload struct {
	Username 			string `json:"username,omitempty"`
	Password 			string `json:"password,omitempty"`
	NewPassword 		string `json:"new_password,omitempty"`
}

// ---- Login ----
type Login struct {
	Id 					string 	`json:"_id,omitempty"`
	Username 			string	`json:"username,omitempty"`
	Email 				string 	`json:"email,omitempty"`
	Success				string 	`json:"success,omitempty"`
}

// ---- Blacklist ----
type Blacklist struct {
	Id	 				string `json:"id,omitempty"`
	AuthToken        	string `json:"auth_token,omitempty"`
	Created 			string `json:"created,omitempty"`
	Modified     		string `json:"modified,omitempty"`
}