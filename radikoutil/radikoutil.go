package radikoutil

import (
	"context"

	"github.com/cockroachdb/errors"
	"github.com/samber/mo"
	"github.com/yyoshiki41/go-radiko"
)

// https://github.com/uru2/rec_radiko_ts/blob/3b2f334cf05b6616ba6e87a1bee8d82e7d2df86c/rec_radiko_ts.sh#L10-L11
const defaultAuthToken = "bcd151073c03b352e1ef2fd66c32209da9ca0afa"

func newClient(email, password mo.Option[string]) (*radiko.Client, error) {
	token := defaultAuthToken

	if email.IsPresent() && password.IsPresent() {
		token = ""
	}

	c, err := radiko.New(token)
	if err != nil {
		return nil, errors.Wrap(err, "failed to construct a radiko Client")
	}
	return c, nil
}

type ClientManager struct {
	email, password mo.Option[string]
	currentAreaID   string
}

func NewClientManger(email, password mo.Option[string]) (*ClientManager, error) {
	areaID, err := radiko.AreaID()
	if err != nil {
		return nil, errors.Wrap(err, "failed to get areaID")
	}
	return &ClientManager{
		email:         email,
		password:      password,
		currentAreaID: areaID,
	}, nil
}

func (cm *ClientManager) isCurrentArea(areaID mo.Option[string]) bool {
	id, ok := areaID.Get()
	if !ok {
		// NOTE: 指定なしの時は現在のエリアIDと同じとみなす
		return true
	}
	return cm.currentAreaID == id
}

func (cm *ClientManager) Get(ctx context.Context, areaID mo.Option[string]) (*radiko.Client, error) {
	email, password := cm.email, cm.password
	c, err := newClient(email, password)
	if err != nil {
		return nil, errors.Wrap(err, "failed to construct a radiko Client")
	}

	// NOTE: 現在のエリア以外かつ、メールアドレスとパスワードが指定されている場合はログインする
	if !cm.isCurrentArea(areaID) && email.IsPresent() && password.IsPresent() {
		status, err := c.Login(ctx, email.OrElse(""), password.OrElse(""))
		if err != nil {
			return nil, errors.Wrap(err, "failed to login to radiko")
		}
		if status.StatusCode() != 200 {
			return nil, errors.Errorf("failed to login to radiko: %d", status.StatusCode())
		}
	}

	c.SetAreaID(areaID.OrElse(cm.currentAreaID))

	if _, err := c.AuthorizeToken(ctx); err != nil {
		return nil, errors.Wrap(err, "failed to authorize token")
	}

	return c, nil
}
