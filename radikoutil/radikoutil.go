package radikoutil

import (
	"context"

	"github.com/cockroachdb/errors"
	"github.com/yyoshiki41/go-radiko"
)

// https://github.com/uru2/rec_radiko_ts/blob/3b2f334cf05b6616ba6e87a1bee8d82e7d2df86c/rec_radiko_ts.sh#L10-L11
const defaultAuthToken = "bcd151073c03b352e1ef2fd66c32209da9ca0afa"

func NewClient(ctx context.Context) (*radiko.Client, error) {
	c, err := radiko.New(defaultAuthToken)
	if err != nil {
		return nil, errors.Wrap(err, "failed to construct a radiko Client")
	}

	areaID, err := radiko.AreaID()
	if err != nil {
		return nil, errors.Wrap(err, "failed to get areaID")
	}
	c.SetAreaID(areaID)

	if _, err := c.AuthorizeToken(ctx); err != nil {
		return nil, errors.Wrap(err, "failed to authorize token")
	}

	return c, nil
}
