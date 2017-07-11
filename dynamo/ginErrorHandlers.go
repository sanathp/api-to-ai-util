package dynamo

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sanathp/api-to-ai-util/vErr"
)

func GetItemErrorHandler(c *gin.Context, err vErr.Error) {
	if err.Type() == vErr.ItemNotFoundType {
		c.JSON(http.StatusBadRequest, vErr.BadRequestErr.Json())
		return
	} else {
		c.JSON(http.StatusInternalServerError, vErr.InternalServerErr.Json())
		return
	}
}

func GetListErrorHandler(c *gin.Context, err vErr.Error) {
	c.JSON(http.StatusInternalServerError, vErr.InternalServerErr.Json())
	return
}

func PutItemErrorHandler(c *gin.Context, err vErr.Error) {
	if err.Type() == vErr.InvalidInputDataType {
		c.JSON(http.StatusBadRequest, err.Json())
	} else {
		c.JSON(http.StatusInternalServerError, vErr.InternalServerErr.Json())
		return
	}
}

func BulkInsertErrorHandler(c *gin.Context, err vErr.Error) {
	//FTODO: find out possible errors of batch write input and handle all those here
	if err.Type() == vErr.InvalidInputDataType || err.Type() == vErr.BadRequestType {
		c.JSON(http.StatusBadRequest, err.Json())
	} else {
		c.JSON(http.StatusInternalServerError, vErr.InternalServerErr.Json())
		return
	}
}

func UpdateItemErrorHandler(c *gin.Context, err vErr.Error) {
	if err.Type() == vErr.ItemNotFoundType {
		c.JSON(http.StatusBadRequest, err.Json())
		return
	} else if err.Type() == vErr.InvalidInputDataType {
		c.JSON(http.StatusBadRequest, err.Json())
		return
	} else {
		c.JSON(http.StatusInternalServerError, vErr.InternalServerErr.Json())
		return
	}
}

func TagsInsertErrorHandler(c *gin.Context, err vErr.Error) {
	//FTODO: write error handler here when you complete InsertTags operations.go TODOs
	if err.Type() == vErr.ItemNotFoundType {
		c.JSON(http.StatusBadRequest, err.Json())
		return
	} else if err.Type() == vErr.InvalidInputDataType {
		c.JSON(http.StatusBadRequest, err.Json())
		return
	} else {
		c.JSON(http.StatusInternalServerError, vErr.InternalServerErr.Json())
		return
	}
}

func DeleteItemErrorHandler(c *gin.Context, err vErr.Error) {
	if err.Type() == vErr.ItemNotFoundType {
		c.JSON(http.StatusBadRequest, err.Json())
		return
	} else {
		c.JSON(http.StatusInternalServerError, vErr.InternalServerErr.Json())
		return
	}
}
