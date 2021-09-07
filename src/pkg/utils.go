package root

import (
	"encoding/json"
	"github.com/gofrs/uuid"
	"github.com/lithammer/shortuuid"
	"go.mongodb.org/mongo-driver/bson"
	"strings"
	"time"

)

// ---- GenId ----
func GenId() (string, error) {
	id, err := uuid.NewV4()
	if err != nil {
		return "", err
	}
	return id.String(), nil
}

// ---- GenShortId ----
func GenShortId() string {
	return shortuuid.New()
}

// ---- GenTimestamp ----
func GenTimestamp() string {
	return time.Now().UTC().String()
}

// ---- JsonStringToArray ----
func JsonStringToArray(jsonString string) []string {
	replacer := strings.NewReplacer("\",\"", "|", "{", "", "}", "", "\"", "")
	newString := replacer.Replace(jsonString)
	return strings.Split(newString, "|")
}

// ---- MakeBsonDQueryFilter ----
func MakeBsonDQueryFilter(i interface{}) bson.D {
	// create a nil bson.D struct
	filter := bson.D{}
	// marshall the user structure given into a json string
	jsonString, _ := json.Marshal(i)
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

// ---- MakeBsonDUpdateQueryFilter ----
func MakeBsonDUpdateQueryFilter(i interface{}) bson.D {
	inner := MakeBsonDQueryFilter(i)
	outer := bson.D{{"$set", inner}}
	return outer
}
