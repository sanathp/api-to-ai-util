package authUtil

import (
	"strconv"
	"strings"
	"time"

	"github.com/pborman/uuid"
	"github.com/sanathp/api-to-ai-util/vErr"
	"github.com/sanathp/api-to-ai-util/vcrypt"
)

func GenerateAuthToken(userId string, validUntil int64) (string, error) {
	created := time.Now().Unix()
	expiresAt := created + validUntil
	randomText := uuid.New()

	tokenString := userId + "}|{" + strconv.FormatInt(created, 10) + "}|{" + strconv.FormatInt(expiresAt, 10) + "}|{" + randomText

	authToken, err := vcrypt.HexaAesEncrypt(tokenString)

	return authToken, err
}

func ValidateAuthToken(token string) (UserId, vErr.Error) {

	tokenString, err := vcrypt.HexaAesDecrypt(token)

	if err != nil {
		return UserId{}, vErr.ErrInvalidToken
	}

	tokensArray := strings.Split(tokenString, "}|{")

	if len(tokensArray) != 4 {
		return UserId{}, vErr.ErrInvalidToken
	}

	created, err2 := strconv.Atoi(tokensArray[1])

	if err2 != nil {
		return UserId{}, vErr.ErrInvalidToken
	}

	expires, err2 := strconv.Atoi(tokensArray[2])

	if err2 != nil {
		return UserId{}, vErr.ErrInvalidToken
	}

	if expires < int(time.Now().Unix()) {
		return UserId{}, vErr.ErrInvalidToken
	}

	userId := UserId{
		tokensArray[0],
		created,
		expires,
	}

	return userId, nil
}

func GetBearerToken(authHeader string) (string, vErr.Error) {
	if strings.HasPrefix(authHeader, BearerAuthScheme) {
		token := strings.TrimPrefix(authHeader, BearerAuthScheme+" ")
		if len(token) > BearerTokenMinLength {
			return token, nil
		} else {
			return "", vErr.ErrInvalidToken
		}
	} else {
		return "", vErr.ErrInvalidToken
	}
}
