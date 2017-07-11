package auth

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/didip/tollbooth"
	"github.com/didip/tollbooth/config"

	"github.com/gin-gonic/gin"
	"github.com/sanathp/api-to-ai-util/constants"
	"github.com/sanathp/api-to-ai-util/logger"
	"github.com/sanathp/api-to-ai-util/vErr"
)

type (
	UserId struct {
		Id         string `json:"userId"`
		SourceId   string `json:"sourceId"`
		DirectedId string `json:"directedId"`
	}

	UserData struct {
		Id         string `json:"userId"`
		SourceId   string `json:"sourceId"`
		DirectedId string `json:"directedId"`
	}
)

var AuthEndPoint string

const (
	UserIdPath           = "/userid"
	Authorization        = "Authorization"
	BearerAuthScheme     = "Bearer"
	BearerTokenMinLength = 20
)

func UserIdAuthMiddleware(userLimiter *config.Limiter) gin.HandlerFunc {
	return func(c *gin.Context) {
		userIdObj, authErr := getUserIdFromAuthHeader(c)

		if authErr == nil {
			httpError := tollbooth.LimitByKeys(userLimiter, []string{userIdObj.Id})
			if httpError != nil {
				c.String(httpError.StatusCode, httpError.Message)
				c.Abort()
				return
			}

			c.Set(constants.UserObjectKey, userIdObj)
			c.Next()
			return

		} else if authErr.Type() == vErr.InternalServerErrorType {
			logger.AuthInternalServerError("UserIdAuthMiddleware", authErr)
			c.JSON(http.StatusInternalServerError, gin.H{})
			c.Abort()
		} else {
			logger.Error("UserIdAuthMiddleware", authErr)
			c.JSON(http.StatusUnauthorized, gin.H{})
			c.Abort()
		}
	}
}

func UserIdAuthCookieMiddleware(c *gin.Context) {
	userIdObj, authErr := getUserIdFromCookie(c)

	if authErr == nil {
		c.Set(constants.UserObjectKey, userIdObj)
		c.Next()
		return
	} else if authErr.Type() == vErr.InternalServerErrorType {
		logger.AuthInternalServerError("UserIdAuthMiddleware", authErr)
		c.JSON(http.StatusInternalServerError, gin.H{})
		c.Abort()
	} else {
		//FTODO: use common constnt here for login path
		c.Redirect(http.StatusTemporaryRedirect, "/login")
		c.Abort()
	}
}

func getUserIdFromAuthHeader(c *gin.Context) (UserId, vErr.Error) {
	authHeader := c.Request.Header.Get("Authorization")

	tokenStr, err := GetBearerToken(authHeader)
	fmt.Sprint(tokenStr)

	if err != nil {
		return UserId{}, err
	}

	//TODO: implement mechanism to validte token
	//userId, err := validateTokenAndGetUserId(tokenStr)

	return UserId{}, nil
}

func getUserIdFromCookie(c *gin.Context) (UserId, vErr.Error) {
	cookie, cookieErr := c.Request.Cookie(constants.AuthTokenCookieKey)

	fmt.Sprint(cookie)

	if cookieErr != nil {
		return UserId{}, vErr.SendError(cookieErr)
	}
	//TODO: implement mechanism to validte token
	//userId, err := validateTokenAndGetUserId(tokenStr)

	return UserId{}, nil
}

func SetAuthCookie(ginWriter gin.ResponseWriter, value string, age int) {

	cookie := &http.Cookie{
		Name:   constants.AuthTokenCookieKey,
		Value:  value,
		Path:   "/",
		MaxAge: age,
	}
	http.SetCookie(ginWriter, cookie)
}

func ClearAuthCookie(ginWriter gin.ResponseWriter) {
	cookie := &http.Cookie{
		Name:   constants.AuthTokenCookieKey,
		Value:  "",
		Path:   "/",
		MaxAge: -1,
	}

	http.SetCookie(ginWriter, cookie)
}

func GetUserIdFromContext(c *gin.Context) *UserId {
	return &UserId{"1", "1", "1"}
	//TODO: acitvate below logic after implemetation of auth hard coding userId for now
	/*
		userObject, exists := c.Get(constants.UserObjectKey)

		if !exists {
			logger.NoUserIdInContextErr()
			c.JSON(http.StatusUnauthorized, gin.H{})
			return nil
		}
		userIdObj, ok := userObject.(UserId)
		if ok {
			return &userIdObj
		} else {
			logger.NoUserIdInContextErr()
			c.JSON(http.StatusUnauthorized, gin.H{})
			return nil
		}
	*/
}

func GetBearerToken(authHeader string) (string, vErr.Error) {
	if strings.HasPrefix(authHeader, BearerAuthScheme) {
		token := strings.TrimPrefix(authHeader, BearerAuthScheme+" ")
		if len(token) > BearerTokenMinLength {
			return token, nil
		} else {
			return "", vErr.ErrInvalidBearerToken
		}
	} else {
		return "", vErr.ErrNoBearerToken
	}
}
