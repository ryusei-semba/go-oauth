package usecase

type TokenResponse struct {
	Token string `json:"token"`
}

func RequestToken() (TokenResponse, error) {
	return TokenResponse{
		"tokenSample",
	}, nil
}
