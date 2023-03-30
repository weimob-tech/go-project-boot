package config

import (
	. "github.com/smartystreets/goconvey/convey"
	"github.com/spf13/viper"
	"github.com/weimob-tech/go-project-base/pkg/config"
	"testing"
)

func TestCachedConfig(t *testing.T) {
	Convey("test cached config", t, func(c C) {
		store := &viperStore{viper.New()}
		store.SetDefault("test.foo", "a")
		store.SetDefault("test.bar", "b")

		var expect = map[string]interface{}{"foo": "a", "bar": "b"}

		c.Convey("should get string map", func(c C) {
			So(store.GetStringMap("test"), ShouldResemble, expect)
		})

		c.Convey("should get string map cached", func(c C) {
			cached := config.Cached(store)
			So(cached.GetStringMap("test"), ShouldResemble, expect)
		})
	})
}
