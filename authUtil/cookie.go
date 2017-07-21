package authUtil

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sanathp/api-to-ai-util/constants"
)

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
