package root

import (
	"encoding/json"
	"errors"
	"golang.org/x/crypto/bcrypt"
	"go.mongodb.org/mongo-driver/bson"
	"strings"
)

// ---- UserService ----
type UserService interface {
	CreateUser(u User) (User, error)
	FindUser(u User) ([]User, error)
	UpdateUser(f User, u User) (User, error)
}

// ---- User ----
type User struct {
	Id				string		`json:"_id,omitempty"`
	Userid			string		`json:"user_id,omitempty"`
	Username		string		`json:"username,omitempty"`
	Password		string		`json:"password,omitempty"`
	Firstname		string		`json:"first_name,omitempty"`
	Lastname		string		`json:"last_name,omitempty"`
	Address			string		`json:"address,omitempty"`
	City			string		`json:"city,omitempty"`
	State			string		`json:"state,omitempty"`
	Zip				string		`json:"zip,omitempty"`
	Country			string		`json:"country,omitempty"`
	Email			string		`json:"email,omitempty"`
	Phone			string		`json:"phone,omitempty"`
	Active			string		`json:"active,omitempty"`
	Created			string		`json:"created,omitempty"`
	Modified		string		`json:"modified,omitempty"`
}

// ---- UserToken ----
type UserToken struct {
	Userid			string		`json:"user_idid,omitempty"`
	Username		string		`json:"username,omitempty"`
	Email			string		`json:"email,omitempty"`
	RemoteAddr		string		`json:"remote_addr,omitempty"`
	ServiceCatalog	[]string	`json:"service_catalog,omitempty"`
}

// ---- User.Validate ----
func (rcvr *User) Validate(opCreate bool) error {
	// if opCreate is true ...
	if opCreate {
		// ... then, make sure we have what we need to create a user record
		if len(rcvr.Userid) == 0 {return errors.New("missing user id")}
		if len(rcvr.Username) == 0 {return errors.New("missing username")}
		if len(rcvr.Password) == 0 {return errors.New("missing password")}
		if len(rcvr.Firstname)== 0 {return errors.New("missing first name")}
		if len(rcvr.Lastname) == 0 {return errors.New("missing last name")}
		if len(rcvr.Address) == 0 {return errors.New("missing address")}
		if len(rcvr.City) == 0 {return errors.New("missing city")}
		if len(rcvr.State) == 0 {return errors.New("missing state")}
		if len(rcvr.Zip) == 0 {return errors.New("missing zip")}
		if len(rcvr.Country) == 0 {return errors.New("missing country")}
		if len(rcvr.Email) == 0 {return errors.New("missing email")}
		if len(rcvr.Phone) == 0 {return errors.New("missing phone")}
		if len(rcvr.Active) == 0 {return errors.New("missing active")}
	// otherwise, if opCreate is false ...
	} else {
		// ... then, we cannot update the following
		if len(rcvr.Userid) > 0 {return errors.New("updating user id not allowed")}
		if len(rcvr.Username) > 0 {return errors.New("updating username not allowed")}
		if len(rcvr.Password) > 0 {return errors.New("updating password not allowed")}
		if len(rcvr.Email) > 0 {return errors.New("updating email not allowed")}
	}

	// return nil if no errors
	return nil
}

// ---- User.HashPassword ----
func (rcvr *User) HashPassword(p string) (string, error) {
	byteP := []byte(p)
	hp, err := bcrypt.GenerateFromPassword(byteP, bcrypt.DefaultCost)
	if err != nil {
		return p, err
	}
	return string(hp), nil
}

// ---- User.ValidatePassword ----
func (rcvr *User) ValidatePassword(p string, hp string) bool {
	byteP := []byte(p)
	byteHp := []byte(hp)
	err := bcrypt.CompareHashAndPassword(byteHp, byteP)
	if err != nil {
		return false
	}
	return true
}

// ---- User.MakeBsonDQueryFilter ----
func (rcvr *User) MakeBsonDQueryFilter() bson.D {
	// create a nil bson.D struct
	filter := bson.D{}
	// marshall the user structure given into a json string
	jsonString, _ := json.Marshal(rcvr)
	// convert the json string to a string
	processString := string(jsonString)
	// if we do no have an empty structure string
	if string(processString) != "{}" {
		// convert the Json string to an array
		fieldArray := JsonStringToArray(string(processString))
		// walk the array to build out a bson.D query filter
		for _, elElement := range fieldArray {
			keys := strings.Split(elElement, ":")
			filter = append(filter, bson.E{keys[0], keys[1]})
		}
	}
	// return the filter
	return filter
}

// ---- User.MakeBsonDUpdateQueryFilter ----
func (rcvr *User) MakeBsonDUpdateQueryFilter() bson.D {
	inner := rcvr.MakeBsonDQueryFilter()
	outer := bson.D{{"$set", inner}}
	return outer
}