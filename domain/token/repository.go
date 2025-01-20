package token

// Repository トークンの永続化を担当するリポジトリのインターフェース
type Repository interface {
	// Save トークンを保存する
	Save(token *Token) error

	// FindByAccessToken アクセストークンからトークンを取得する
	FindByAccessToken(accessToken string) (*Token, error)

	// FindByRefreshToken リフレッシュトークンからトークンを取得する
	FindByRefreshToken(refreshToken string) (*Token, error)

	// Delete トークンを削除する
	Delete(accessToken string) error
}
