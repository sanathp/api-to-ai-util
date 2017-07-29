package id

import (
	"fmt"
	"math/rand"
	"strconv"
	"time"

	"github.com/satori/go.uuid"
)

const (
	IdLength        = 15
	TimestampLength = 13
)

var stringMap = map[string]string{
	"0": "a",
	"1": "b",
	"2": "c",
	"3": "d",
	"4": "e",
	"5": "f",
	"6": "g",
	"7": "h",
	"8": "i",
	"9": "j",
}

func GetStringId() string {
	stringId := ""
	timestamp := strconv.Itoa(int(time.Now().UnixNano()))
	for _, elem := range timestamp {
		stringId = stringId + stringMap[string(elem)]
	}
	fmt.Println(stringId)
	//Worst case, but this should not happen
	if len(stringId) == 0 {
		stringId = GetNewUUID()
	}
	return stringId
}

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
