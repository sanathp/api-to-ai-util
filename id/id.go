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

func GetTimestamp() int64 {
	ctime := time.Now().UnixNano()
	//Converting nano timestamp to milli seconds
	return ctime / 1000000
}

func IsValidItemId(id string) (int64, bool) {
	if len(id) != IdLength {
		return 0, false
	}
	if value, err := strconv.Atoi(id); err == nil {
		return int64(value), true
	} else {
		return 0, false
	}
}

func IsValidItemIdInt64(id int64) bool {
	if len(strconv.FormatInt(id, 10)) != IdLength {
		return false
	}
	return true
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

//FTODO : each server should have an id ,
//take micro second value and appen server id at the end
//Make this more efficient
func GetNewItemIdAndTimestamp() (int64, int64) {
	ctime := time.Now().UnixNano()
	//Converting nano timestamp to milli seconds
	timestamp := ctime / 1000000
	newId := timestamp*100 + int64(rand.Intn(100))
	return newId, timestamp
}
