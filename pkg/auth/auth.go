package auth

import (
	"github.com/weimob-tech/go-project-base/pkg/auth"
	"github.com/weimob-tech/go-project-base/pkg/hook"
)

type OAuth struct {
	AccessToken     string
	PublicAccountId string
	BusinessId      string
}

func init() {
	hook.AddPostStartHook(func() {
		auth.DefaultStore = Cached(auth.NewHttpStore())
	})
}
