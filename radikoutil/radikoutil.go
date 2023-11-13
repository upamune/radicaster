package radikoutil

import (
	"context"

	"github.com/cockroachdb/errors"
	"github.com/yyoshiki41/go-radiko"
)

// https://github.com/uru2/rec_radiko_ts/blob/3b2f334cf05b6616ba6e87a1bee8d82e7d2df86c/rec_radiko_ts.sh#L10-L11
const defaultAuthToken = "bcd151073c03b352e1ef2fd66c32209da9ca0afa"

type Option func(o *options)

type options struct {
	areaID          string
	isPremium       bool
	email, password string
}

func evaluateOptions(opts []Option) *options {
	o := &options{}
	for _, fn := range opts {
		fn(o)
	}
	return o
}

func WithPremium(email, password string) Option {
	return func(o *options) {
		o.isPremium = true
		o.email = email
		o.password = password
	}
}

func WithAreaID(areaID string) Option {
	return func(o *options) {
		o.areaID = areaID
	}
}

func NewClient(ctx context.Context, opts ...Option) (*radiko.Client, error) {
	opt := evaluateOptions(opts)
	c, err := radiko.New(defaultAuthToken)
	if err != nil {
		return nil, errors.Wrap(err, "failed to construct a radiko Client")
	}

	currentAreaID, err := radiko.AreaID()
	if err != nil {
		return nil, errors.Wrap(err, "failed to get current areaID via API")
	}

	if opt.areaID != "" {
		c.SetAreaID(opt.areaID)
	} else {
		// NOTE: エリアIDが指定されてない時は今のエリアのIDにする
		c.SetAreaID(currentAreaID)
	}

	// NOTE: プレミアム会員かつ、現在のエリア外の時はログインする
	if opt.isPremium && isCurrentAreaID(currentAreaID, opt.areaID) {
		status, err := c.Login(ctx, opt.email, opt.password)
		if err != nil {
			return nil, errors.Wrap(err, "failed to login to radiko")
		}
		if status.StatusCode() != 200 {
			return nil, errors.Errorf("failed to login to radiko: %d", status.StatusCode())
		}
	}

	if _, err := c.AuthorizeToken(ctx); err != nil {
		return nil, errors.Wrap(err, "failed to authorize token")
	}

	return c, nil
}

func isCurrentAreaID(currentAreaID, areaID string) bool {
	// NOTE: 未指定の時はエリア内
	if areaID == "" {
		return true
	}
	return currentAreaID == areaID
}
