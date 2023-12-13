package boot

import (
	hconf "github.com/cloudwego/hertz/pkg/common/config"
	httpBoot "github.com/weimob-tech/go-project-boot/pkg/http"
)

// WithHertzServer is a HertzOption that sets the hertz server.
// 允许用户自定义 hertz server 行为
func WithHertzHttpServer(opts ...hconf.Option) Option {
	return func(c *Configure) {
		c.NewServer = func() {
			c.Server = httpBoot.NewHertzServer(opts...)
		}
	}
}
