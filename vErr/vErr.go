package vErr

import (
	"gopkg.in/go-playground/validator.v9"
)

const (
	InvalidParameterType      = "invalid_parameter"
	FailedStatusType          = "failed"
	ItemNotFoundType          = "item_not_found"
	ItemAlreadyExistsType     = "item_already_exists"
	InvalidInputDataType      = "invalid_input_data"
	DatabaseRequestFailedType = "database_request_failed"
	DatabaseErrorType         = "database_error"
	InternalServerErrorType   = "internal_server_error"
	BadRequestType            = "bad_request"
	MissingRequiredFieldsType = "missing_required_fields"
	ForbidenAccessType        = "forbidden_access"
	MarshalErrorType          = "marshal_error"
	UnmarshalErrorType        = "un_marshal_error"
	UserIdEmptyType           = "user_id_empty"
	InvalidTimestampType      = "invalid_timestamp"
	InvalidItemIdType         = "invalid_item_id"
	OAuthErrorType            = "oauth"
	InvalidTokenType          = "invalid_token"
	JwtErrorType              = "jwt_error"
	EncryptionErrType         = "encryption_error"
	UnauthorizedErrorType     = "unauthorized"
	NoUserIdInContextType     = "no_user_id_in_context_error"
	NoUserDateInContextType   = "no_user_data_in_context_error"
	UpdateFailedType          = "update_failed"
	GCMErrorType              = "google_cloud_messaging_error"
	GoogleOAuthType           = "google_oauth_error"
	LWAOAuthType              = "lwa_oauth_error"
	UrlInfoErrorType          = "url_info_error"
	VcryptErrorType           = "vcrypt_error"
	TokenInvalidatedErrorType = "invalidated_token_received"
	HttpRequestCreationFailed = "failed_to_create_http_request"
	HttpRequestFailed         = "http_request_failed"
	HttpResponseNotOk         = "http_response_not_200"

	//Token refresh error types
	RefreshTokenNotAvailable = "refresh_token_not_avaiable"
	TokenRefreshFailed       = "token_refresh_failed"
)

var (
	InternalServerErr        = New(InternalServerErrorType, "")
	DatabaseRequestFailedErr = New(DatabaseRequestFailedType, "")
	BadRequestErr            = New(BadRequestType, "")
	MissingRequiredFieldsErr = New(MissingRequiredFieldsType, "")
	ItemNotFoundErr          = New(ItemNotFoundType, "")
	UpdateFailedErr          = New(UpdateFailedType, "")
	DatabaseErr              = New(DatabaseErrorType, "")
	UnMarshalErr             = New(UnmarshalErrorType, "")
	MarshalErr               = New(MarshalErrorType, "")
	ItemAlreadyExistsErr     = New(ItemAlreadyExistsType, "")
	UnauthorizedErr          = New(UnauthorizedErrorType, "")

	UserIdEmptyVErr      = New(UserIdEmptyType, "")
	InvalidTimestampVErr = New(InvalidTimestampType, "")
	InvalidItemIdVErr    = New(InvalidItemIdType, "")

	ErrInvalidToken = New(InvalidTokenType, "Invalid Token")
)

type (
	Error interface {
		// Satisfy the generic error interface.
		error

		// Returns the short phrase depicting the classification of the error.
		Type() string

		// Returns the error details message.
		Message() string

		Json() jsonResponseErr
	}

	baseError struct {
		errorType string
		message   string
	}

	jsonResponseErr struct {
		ErrorType string `json:"error"`
		Message   string `json:"message,omitempty"`
	}
)

func (err baseError) Type() string {
	return err.errorType
}

func (err baseError) Message() string {
	return err.message
}

func (err baseError) Json() jsonResponseErr {
	return jsonResponseErr{err.Type(), err.Message()}
}

func (err baseError) Error() string {
	//FTODO: better format
	return err.errorType + " : " + err.message
}

func New(typeStr string, message string) Error {
	return &baseError{typeStr, message}
}

func InvalidParameterError(message string) Error {
	return New(InvalidParameterType, message)
}

func RequiredParamMissingError(paramName string) Error {
	message := paramName + " should not be empty"
	return New(InvalidInputDataType, message)
}

func InvalidInputDataError(message string) Error {
	return New(InvalidInputDataType, message)
}

func InternalServerError(message string) Error {
	return New(InternalServerErrorType, message)
}

func SendError(err error) Error {
	return New(err.Error(), "")
}

func ValidatorErr(err error) Error {
	return New(InvalidInputDataType, GetValidatorMsg(err))
}

func GetValidatorMsg(err error) string {
	message := ""
	for key, err := range err.(validator.ValidationErrors) {
		if key == 0 {
			message = "Invalid input for field '" + err.Field() + "', tag details = " + err.Tag() + ":" + err.Param()
			break
		}
	}
	return message
}
