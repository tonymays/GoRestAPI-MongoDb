package root

import (
	"encoding/json"
	"errors"
	"golang.org/x/crypto/bcrypt"
	"go.mongodb.org/mongo-driver/bson"
	"strings"
)

type UserService interface {
}

type User struct {
	Id				string		`json:"id,omitempty"`
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

type UserToken struct {
	Id				string		`json:"id,omitempty"`
	Username		string		`json:"username,omitempty"`
	Email			string		`json:"email,omitempty"`
	RemoteAddr		string		`json:"remote_addr,omitempty"`
	ServiceCatalog	[]string	`json:"service_catalog,omitempty"`
}

func (rcvr *User) Validate(opCreate bool) error {
	if opCreate {
		if len(rcvr.Id) == 0 {return errors.New("missing id")}
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
	} else {
		if len(rcvr.Password) > 0 {return errors.New("setting password not allowed")}
	}
	return nil
}

func (rcvr *User) HashPassword(p string) (string, error) {
	byteP := []byte(p)
	hp, err := bcrypt.GenerateFromPassword(byteP, bcrypt.DefaultCost)
	if err != nil {
		return p, err
	}
	return string(hp), nil
}

func (rcvr *User) ValidatePassword(p string, hp string) bool {
	byteP := []byte(p)
	byteHp := []byte(hp)
	err := bcrypt.CompareHashAndPassword(byteHp, byteP)
	if err != nil {
		return false
	}
	return true
}

func (rcvr *User) MakeBsonDQueryFilter() bson.D {
	filter := bson.D{}
	jsonString, _ := json.Marshal(rcvr)
	processString := string (jsonString)
	if string(processString) != "{}" {
		fieldArray := JsonStringToArray(string (processString))
		for _, elElement := range fieldArray {
			keys := strings.Split(elElement, ":")
			filter = append (filter, bson.E{keys [ 0 ], keys[ 1 ] })
		}
	}
	return filter
}

func (rcvr *User) MakeBsonDUpdateQueryFilter() bson.D {
	inner := rcvr.MakeBsonDQueryFilter()
	outer := bson.D{{"$set", inner } }
	return outer
}
