package parasut

import (
	"github.com/WEG-Technology/room"
	"os"
)

type AuthRequestPayload struct {
	Username    string `json:"username"`
	Password    string `json:"password"`
	RedirectUri string `json:"redirect_uri"`
	GrantType   string `json:"grant_type"`
}

func NewAuthRequest() room.IRequest {
	r, e := room.NewPostRequest(
		room.WithEndPoint(authEndpoint),
		room.WithBody(AuthRequestPayload{
			Username:    os.Getenv(UsernameEnv),
			Password:    os.Getenv(PasswordEnv),
			RedirectUri: "urn:ietf:wg:oauth:2.0:oob",
			GrantType:   "password",
		}),
		room.WithDto(&AuthResponse{}),
	)

	if e != nil {
		panic(e)
	}

	return r
}
