package authUtil

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/sanathp/api-to-ai-util/logger"
	"github.com/sanathp/api-to-ai-util/vErr"
)

func GetDetailsFromGooglePlusToken(token string) (GooglePlusOAuthResponse, error) {
	var googleOAuthResponse GooglePlusOAuthResponse
	response, err := http.Get("https://www.googleapis.com/oauth2/v1/userinfo?alt=json&access_token=" + token)
	if err != nil {
		logger.ErrorOfType(vErr.GoogleOAuthType, "GetDetailsFromGooglePlusToken", err)
		return googleOAuthResponse, err
	}

	defer response.Body.Close()
	contents, err := ioutil.ReadAll(response.Body)
	if err != nil {
		logger.ErrorOfType(vErr.GoogleOAuthType, "GetDetailsFromGooglePlusToken", err)
		return googleOAuthResponse, err
	}

	err2 := json.Unmarshal(contents, &googleOAuthResponse)
	logger.Debug("Contents is ", string(contents))

	if err2 != nil || len(googleOAuthResponse.Email) == 0 {
		logger.ErrorOfType(vErr.GoogleOAuthType, "GetDetailsFromGooglePlusToken", err2)
		return googleOAuthResponse, err
	}

	logger.Debug(googleOAuthResponse.DisplayName, "---", googleOAuthResponse.ID, "---email ", googleOAuthResponse.Email)

	return googleOAuthResponse, nil
}
