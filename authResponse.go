package parasut

import "github.com/WEG-Technology/room"

type AuthResponse struct {
	AccessToken  string `json:"access_token"`
	TokenType    string `json:"token_type"`
	ExpiresIn    int    `json:"expires_in"`
	RefreshToken string `json:"refresh_token"`
	Scope        string `json:"scope"`
	CreatedAt    int    `json:"created_at"`
}

func (r AuthResponse) GetBearerToken(response room.IResponse) string {
	return response.Dto().(*AuthResponse).AccessToken
}
