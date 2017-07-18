package id

import (
	"math/rand"
	"strconv"
	"time"

	"github.com/satori/go.uuid"
)

const (
	IdLength        = 15
	TimestampLength = 13
)

func GetNewUUID() string {
	return uuid.NewV4().String()
}

func GetNewClientSecret() string {
	//FTODO: implement a better secret value
	secretValue := uuid.NewV4().String()
	return secretValue
}

func GetTimestamp() int {
	ctime := int(time.Now().UnixNano())
	//Converting nano timestamp to milli seconds
	return ctime / 1000000
}

func IsValidItemId(id string) (int, bool) {
	if len(id) != IdLength {
		return 0, false
	}
	if value, err := strconv.Atoi(id); err == nil {
		return value, true
	} else {
		return 0, false
	}
}

func IsValidStringTimestamp(timestamp string) (int64, bool) {
	if len(timestamp) != TimestampLength {
		return 0, false
	}
	if value, err := strconv.Atoi(timestamp); err == nil {
		return int64(value), true
	} else {
		return 0, false
	}
}

func GetNewUUIDAndTimestamp() (string, int) {
	return GetNewUUID(), GetTimestamp()
}

//FTODO : each server should have an id ,
//take micro second value and appen server id at the end
//Make this more efficient
func GetNewItemIdAndTimestamp() (int, int) {
	ctime := int(time.Now().UnixNano())
	//Converting nano timestamp to milli seconds
	timestamp := ctime / 1000000
	newId := timestamp*100 + rand.Intn(100)
	return newId, timestamp
}

func GenerateUserAppKey(userId string) (string, int) {
	return userId + "|" + GetNewUUID(), GetTimestamp()
}

func GetUserAppKeyAndTimestamp(userId string, appId string) (string, int) {
	return userId + "|" + appId, GetTimestamp()
}

func GetUserAppKey(userId string, appId string) string {
	return userId + "|" + appId
}
