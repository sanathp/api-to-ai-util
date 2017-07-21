package authUtil

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sanathp/api-to-ai-util/constants"
	"github.com/sanathp/api-to-ai-util/logger"
)

func GetUserIdFromContext(c *gin.Context) *UserId {
	//return &UserId{"1", 1, 1}

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
}
