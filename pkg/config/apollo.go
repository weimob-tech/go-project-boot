package config

import (
	"github.com/apolloconfig/agollo/v4"
	"github.com/apolloconfig/agollo/v4/env/config"
	"github.com/apolloconfig/agollo/v4/storage"
	base "github.com/weimob-tech/go-project-base/pkg/config"
	"github.com/weimob-tech/go-project-boot/pkg/wlog"
)

var (
	client agollo.Client
)

func GetApolloClient() agollo.Client {
	return client
}

func setupApollo() (err error) {
	addr := base.GetString("apollo.meta")
	secret := base.GetString("apollo.secret")
	if addr == "" || secret == "" {
		return
	}
	c := &config.AppConfig{
		NamespaceName: "application",
		Cluster:       base.GetString("env"),
		AppID:         base.GetString("appId"),
		Secret:        secret,
		IP:            addr,
	}

	client, err = agollo.StartWithConfig(func() (*config.AppConfig, error) {
		return c, nil
	})
	wlog.I().Msg("初始化Apollo配置成功")

	// 添加默认监听
	client.AddChangeListener(&storeOverrider{})
	// 注入配置
	client.GetApolloConfigCache().Range(func(k, v interface{}) bool {
		base.Set(k.(string), v)
		return true
	})
	return
}

type storeOverrider struct{}

func (l *storeOverrider) OnChange(e *storage.ChangeEvent) {}

func (l *storeOverrider) OnNewestChange(e *storage.FullChangeEvent) {
	for k, v := range e.Changes {
		base.Set(k, v)
	}
}
