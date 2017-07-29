package authUtil

import (
	"net/http"

	"github.com/didip/tollbooth"
	"github.com/didip/tollbooth/config"
	"github.com/gin-gonic/gin"
	"github.com/sanathp/api-to-ai-util/constants"
	"github.com/sanathp/api-to-ai-util/logger"
	"github.com/sanathp/api-to-ai-util/vErr"
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
			return
		} else {
			logger.Error("UserIdAuthMiddleware", authErr)
			c.JSON(http.StatusUnauthorized, gin.H{})
			c.Abort()
			return
		}
	}
}

func UserIdAuthCookieMiddleware(userLimiter *config.Limiter) gin.HandlerFunc {

	return func(c *gin.Context) {
		userIdObj, authErr := getUserIdFromCookie(c)

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
			return
		} else {
			logger.Error("UserIdAuthMiddleware", authErr)

			//FTODO: use common constnt here for login path
			c.Redirect(http.StatusTemporaryRedirect, "/login")
			c.Abort()
			return
		}
	}
}

func getUserIdFromAuthHeader(c *gin.Context) (UserId, vErr.Error) {
	authHeader := c.Request.Header.Get("Authorization")

	tokenStr, err := GetBearerToken(authHeader)

	if err != nil {
		return UserId{}, err
	}

	return ValidateAuthToken(tokenStr)
}

func getUserIdFromCookie(c *gin.Context) (UserId, vErr.Error) {
	cookie, cookieErr := c.Request.Cookie(constants.AuthTokenCookieKey)

	if cookieErr != nil {
		return UserId{}, vErr.SendError(cookieErr)
	}

	return ValidateAuthToken(cookie.Value)
}
