package config

import (
	"fmt"
	"github.com/spf13/viper"
	"github.com/weimob-tech/go-project-base/pkg/config"
)

var store *viperStore

type viperStore struct {
	*viper.Viper
}

func (config *viperStore) GetConfig() any {
	return config.Viper
}

func Setup() {
	v := viper.New()
	setupDefaults(v)

	name := "application"
	env := v.GetString("env")
	if len(env) > 0 {
		name = fmt.Sprintf("%s-%s", name, env)
	}

	v.SetConfigName(name)
	v.SetConfigType("properties")
	v.AddConfigPath(".")
	v.AddConfigPath("./configs")
	err := v.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("fatal error configs file: %w", err))
	}
	v.AutomaticEnv()

	store = &viperStore{v}
	// 防呆，viper 在高并发下有性能问题
	config.SetStore(config.Cached(store))
	// 接入配置服务
	err = setupApollo()
	if err != nil {
		// 配置服务不可用，直接panic
		panic(err)
	}
}

func setupDefaults(v *viper.Viper) {
	_ = v.BindEnv("env", "env")
	_ = v.BindEnv("appId", "appId")
	_ = v.BindEnv("apollo.meta", "apollo.meta")

	v.SetDefault("client.schema", "https")
	v.SetDefault("client.domain", "dopen.weimob.com")
}
