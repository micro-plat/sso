package mp

type AccessToken struct {
	token string
}

func NewAccessToken(token string) *AccessToken {
	return &AccessToken{token: token}
}
func (t *AccessToken) Token() (token string, err error) {
	return t.token, nil
}
func (t *AccessToken) RefreshToken(currentToken string) (token string, err error) {
	t.token = currentToken
	return t.token, nil
}
