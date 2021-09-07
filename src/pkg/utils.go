package root

import (
	"github.com/gofrs/uuid"
	"github.com/lithammer/shortuuid"
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
