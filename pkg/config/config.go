package config

import (
	"fmt"
	"github.com/spf13/viper"
	"github.com/weimob-tech/go-project-base/pkg/config"
	"github.com/weimob-tech/go-project-boot/pkg/wlog"
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
	v.AutomaticEnv()

	// read application
	v.SetConfigName(name)
	v.AddConfigPath(".")
	v.AddConfigPath("./configs")
	v.AddConfigPath("./cmd/configs")
	err := v.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("fatal error configs file: %w", err))
	}
	// read env application
	if len(env) > 0 {
		name = fmt.Sprintf("%s-%s", name, env)
		v.SetConfigName(name)
		err = v.MergeInConfig()
		if err != nil {
			switch err.(type) {
			case viper.ConfigFileNotFoundError:
				wlog.W().Msgf("env config file not found: %s", name)
			default:
				panic(fmt.Errorf("fatal error configs file: %w", err))
			}
		}
	}

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
