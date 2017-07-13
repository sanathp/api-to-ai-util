package logger

import (
	"os"
	"time"

	"github.com/Sirupsen/logrus"
	"github.com/Sirupsen/logrus/formatters/logstash"
	"github.com/gin-gonic/gin"
	"github.com/sanathp/api-to-ai-util/vErr"
)

const (
	DEVELOPMENT = "development"
	PRODUCTION  = "production"
)

func Initialize(mode string, appName string) {

	if mode == DEVELOPMENT {
		//in development mode
		logrus.SetFormatter(&logstash.LogstashFormatter{Type: appName})
		logrus.SetOutput(os.Stderr)
		logrus.SetLevel(logrus.DebugLevel)
	} else if mode == PRODUCTION {
		//in production mode
		logrus.SetFormatter(&logstash.LogstashFormatter{Type: appName})
		logrus.SetOutput(os.Stderr)
		logrus.SetLevel(logrus.DebugLevel)
		//FTODO: save production logs somewhere,but not required for now i guess
		/*
			logrus.SetOutput(&lumberjack.Logger{
				Filename:   "/home/sanath/voblet.log",
				MaxSize:    100, // megabytes
				MaxBackups: 2,
				MaxAge:     28, //days
			})*/
	} else {
		panic("Invalid Mode for logger")
	}
}

func CustomLogger() gin.HandlerFunc {
	return func(c *gin.Context) {
		t := time.Now()
		path := c.Request.URL.Path

		// before request
		c.Next()

		// after request
		//latency in milliseconds
		latency := time.Since(t) / 1000000
		// access the status we are sending
		status := c.Writer.Status()
		clientIP := c.ClientIP()
		method := c.Request.Method

		logrus.WithFields(logrus.Fields{
			"request_type": method,
			"client_ip":    clientIP,
			"path":         path,
			"latency":      latency,
			"status":       status,
		}).Info("request")

	}
}

func Debug(args ...interface{}) {
	logrus.Debug(args)
}

func Error(args ...interface{}) {
	logrus.Error(args)
}

func Warn(args ...interface{}) {
	logrus.Warn(args)
}

func NoUserIdInContextErr() {
	logrus.WithFields(logrus.Fields{
		"error_type": vErr.NoUserIdInContextType,
	}).Error(vErr.NoUserIdInContextType)
}

func NoUserDataInContextErr() {
	logrus.WithFields(logrus.Fields{
		"error_type": vErr.NoUserDateInContextType,
	}).Error(vErr.NoUserDateInContextType)
}

func BadRequestError(tableName string, operationType string, methodName string, err error) {
	errorStr := "nil"
	if err != nil {
		errorStr = err.Error()
	}
	logrus.WithFields(logrus.Fields{
		"error_type":     vErr.BadRequestType,
		"table_name":     tableName,
		"method_name":    methodName,
		"operation_type": operationType,
		"error":          errorStr,
	}).Error(vErr.BadRequestType)
}

func DbError(tableName string, methodName string, operationType string, err error) {
	errorStr := "nil"
	if err != nil {
		errorStr = err.Error()
	}
	logrus.WithFields(logrus.Fields{
		"error_type":     vErr.DatabaseErrorType,
		"table_name":     tableName,
		"method_name":    methodName,
		"operation_type": operationType,
		"error":          errorStr,
	}).Error(vErr.DatabaseErrorType)
}

func OAuthError(methodName string, err error) {
	errorStr := "nil"
	if err != nil {
		errorStr = err.Error()
	}
	logrus.WithFields(logrus.Fields{
		"error_type":  vErr.OAuthErrorType,
		"method_name": methodName,
		"error":       errorStr,
	}).Error(vErr.OAuthErrorType)
}

func ErrorOfType(errorType string, methodName string, err error) {
	errorStr := "nil"
	if err != nil {
		errorStr = err.Error()
	}
	logrus.WithFields(logrus.Fields{
		"error_type":  errorType,
		"method_name": methodName,
		"error":       errorStr,
	}).Error(vErr.OAuthErrorType)
}

func GCMError(methodName string, err error) {
	errorStr := "nil"
	if err != nil {
		errorStr = err.Error()
	}
	logrus.WithFields(logrus.Fields{
		"error_type":  vErr.GCMErrorType,
		"method_name": methodName,
		"error":       errorStr,
	}).Error(vErr.GCMErrorType)
}

func UrlInfoError(tableName string, reason string, err error) {
	errorStr := "nil"
	if err != nil {
		errorStr = err.Error()
	}
	logrus.WithFields(logrus.Fields{
		"error_type": vErr.UrlInfoErrorType,
		"table_name": tableName,
		"reason":     reason,
		"error":      errorStr,
	}).Error(vErr.UrlInfoErrorType)
}

func EncryptionError(methodName string, err error) {
	errorStr := "nil"
	if err != nil {
		errorStr = err.Error()
	}
	logrus.WithFields(logrus.Fields{
		"error_type":  vErr.EncryptionErrType,
		"method_name": methodName,
		"error":       errorStr,
	}).Error(vErr.EncryptionErrType)
}

func UnAuthorizedError(methodName string, err error) {
	errorStr := "nil"
	if err != nil {
		errorStr = err.Error()
	}
	logrus.WithFields(logrus.Fields{
		"error_type":  vErr.InvalidTokenType,
		"method_name": methodName,
		"error":       errorStr,
	}).Error(vErr.InvalidTokenType)
}

func JwtError(methodName string, err error) {
	errorStr := "nil"
	if err != nil {
		errorStr = err.Error()
	}
	logrus.WithFields(logrus.Fields{
		"error_type":  vErr.JwtErrorType,
		"method_name": methodName,
		"error":       errorStr,
	}).Error(vErr.JwtErrorType)
}

func UnMarshalError(tableName string, methodName string, operationType string, err error) {
	errorStr := "nil"
	if err != nil {
		errorStr = err.Error()
	}
	logrus.WithFields(logrus.Fields{
		"error_type":     vErr.UnmarshalErrorType,
		"table_name":     tableName,
		"method_name":    methodName,
		"operation_type": operationType,
		"error":          errorStr,
	}).Error(vErr.UnmarshalErrorType)
}

func InvalidIdError(tableName string, methodName string, id interface{}) {

	logrus.WithFields(logrus.Fields{
		"error_type":  vErr.InvalidItemIdType,
		"table_name":  tableName,
		"method_name": methodName,
		"id_value":    id,
	}).Error(vErr.InvalidItemIdType)
}

func MarshalError(tableName string, methodName string, operationType string, err error) {
	errorStr := "nil"
	if err != nil {
		errorStr = err.Error()
	}
	logrus.WithFields(logrus.Fields{
		"error_type":     vErr.MarshalErrorType,
		"table_name":     tableName,
		"method_name":    methodName,
		"operation_type": operationType,
		"error":          errorStr,
	}).Error(vErr.MarshalErrorType)
}

func AuthInternalServerError(methodName string, err error) {
	errorStr := "nil"
	if err != nil {
		errorStr = err.Error()
	}
	logrus.WithFields(logrus.Fields{
		"error_type":  vErr.InternalServerErrorType,
		"method_name": methodName,
		"error":       errorStr,
	}).Error(vErr.InternalServerErrorType)
}