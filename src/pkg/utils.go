package root

import (
	"github.com/gofrs/uuid"
	"github.com/lithammer/shortuuid"
	"strings"
	"time"
)

func GenId() (string, error) {
	id, err := uuid.NewV4()
	if err != nil {
		return "", err
	}
	return id.String(), nil
}

func GenShortId() string {
	return shortuuid.New()
}

func GenTimeStamp() string {
	return time.Now().UTC().String()
}

func JsonStringToArray(jsonString string) []string {
	replacer := strings.NewReplacer("\",\"", "|", "{", "", "}", "", "\"", "")
	newString := replacer.Replace(jsonString)
	return strings.Split(newString, "|")
}
