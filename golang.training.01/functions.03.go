type UserOffer struct {
	GlobalUserID int64  `dynamodb:"global_user_id"`
	CreatedAt    int64  `dynamodb:"created_at"`
	OfferId      string `dynamodb:"offer_id"`
	Consumed     bool   `dynamodb:"consumed"`
	Whitelisted  bool   `dynamodb:"whitelisted"`
	ExpiredAt    int64  `dynamodb:"expired_at"`
}

type UserOffers []*UserOffer

func (ua *UserOffer) Persist(client *Client) error {
	// DynamoDB persists stuff
}

func (uo *UserOffers) ByGlobalUserID(client *Client, globalUserId int64) error {
	// DynamoDB query stuff
}
