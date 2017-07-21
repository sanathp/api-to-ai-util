package authUtil

type (
	UserId struct {
		Id      string `json:"userId"`
		Created int    `json:"created"`
		Expires int    `json:"expires"`
	}

	GooglePlusOAuthResponse struct {
		Email         string `json:"email"`
		DisplayName   string `json:"name"`
		Gender        string `json:"gender"`
		GivenName     string `json:"given_name"`
		FamilyName    string `json:"family_name"`
		Locale        string `json:"locale"`
		Picture       string `json:"picture"`
		ID            string `json:"id"`
		Url           string `json:"url"`
		VerifiedEmail bool   `json:"verified_email"`
	}
)

var AuthEndPoint string

const (
	UserIdPath           = "/userid"
	Authorization        = "Authorization"
	BearerAuthScheme     = "Bearer"
	BearerTokenMinLength = 20
)
