package auth

import (
	"context"
	"fmt"
	"github.com/weimob-tech/go-project-base/pkg/auth"
	"sync"
	"time"
)

func Cached(store auth.Store) auth.Store {
	return &cachedStore{
		layer: store,
	}
}

type cachedStore struct {
	sync.Map
	layer auth.Store
}

func (c *cachedStore) GetCCToken(ctx context.Context, product, shopId, shopType string) (response *auth.OAuthResponse, err error) {
	key := fmt.Sprintf("cc-%s/%s/%s", product, shopId, shopType)
	return c.cached(key, func() (*auth.OAuthResponse, error) {
		return c.layer.GetCCToken(ctx, product, shopId, shopType)
	})
}

func (c *cachedStore) GetProductCCToken(ctx context.Context, cid, cse, shopId, shopType string) (response *auth.OAuthResponse, err error) {
	key := fmt.Sprintf("pcc-%s/%s/%s", cid, shopId, shopType)
	return c.cached(key, func() (*auth.OAuthResponse, error) {
		return c.layer.GetProductCCToken(ctx, cid, cse, shopId, shopType)
	})
}

func (c *cachedStore) cached(key string, fallback func() (*auth.OAuthResponse, error)) (response *auth.OAuthResponse, err error) {
	if v, ok := c.Load(key); ok {
		response = v.(*auth.OAuthResponse)
		if response.ExpiresAt.Before(time.Now()) {
			// refresh the token
			c.Delete(key)
		} else {
			return
		}
	}
	response, err = fallback()
	if err != nil {
		return
	}
	c.Store(key, response)
	return
}
