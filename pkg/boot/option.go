package boot

import (
	codecBase "github.com/weimob-tech/go-project-base/pkg/codec"
	"github.com/weimob-tech/go-project-base/pkg/hook"
	httpBase "github.com/weimob-tech/go-project-base/pkg/http"
	logBase "github.com/weimob-tech/go-project-base/pkg/wlog"
	"github.com/weimob-tech/go-project-boot/pkg/config"
	httpBoot "github.com/weimob-tech/go-project-boot/pkg/http"
	"github.com/weimob-tech/go-project-boot/pkg/wcontext"
	logBoot "github.com/weimob-tech/go-project-boot/pkg/wlog"
)

var (
	defaultClientCreator = func() {
		httpBase.NewHttpClient = func() httpBase.Client {
			httpBase.Global = httpBoot.NewHertzClient()
			return httpBase.Global
		}
	}

	defaultPreStartHook  = hook.ApplyPreStartHook
	defaultPostStartHook = hook.ApplyPostStartHook
)

type Option func(c *Configure)

type ServerOption func(s httpBase.Server)

type Configure struct {
	Server           httpBase.Server
	BeforeSetup      func()
	SetupConfig      func()
	SetupLog         func()
	SetupContainer   func()
	AfterSetup       func()
	NewServer        func()
	NewClient        func()
	NewCodec         func()
	SetupServer      func()
	PreStarterHook   func()
	BlockingStarters []BlockingStarter
	PostStarterHook  func()
}

func configureFrom(opts ...Option) *Configure {
	conf := defaultConfigure()
	if len(opts) == 0 {
		return conf
	}

	for _, opt := range opts {
		opt(conf)
	}
	return conf
}

func defaultConfigure() *Configure {
	return &Configure{
		BeforeSetup:     func() {},
		SetupConfig:     config.Setup,
		SetupLog:        logBoot.Setup,
		SetupContainer:  wcontext.Setup,
		AfterSetup:      func() {},
		NewClient:       defaultClientCreator,
		PreStarterHook:  defaultPreStartHook,
		PostStarterHook: defaultPostStartHook,
	}
}

func RunBeforeSetup(blk func()) Option {
	return func(c *Configure) {
		c.BeforeSetup = blk
	}
}

func RunAfterSetup(blk func()) Option {
	return func(c *Configure) {
		c.AfterSetup = blk
	}
}

func WithHttpServer() Option {
	return func(c *Configure) {
		c.NewServer = func() {
			c.Server = httpBoot.NewHertzServer()
		}
	}
}

func WithHttpServerFrom(server httpBase.Server) Option {
	return func(c *Configure) {
		c.Server = server
	}
}

func WithHttpClient() Option {
	return func(c *Configure) {
		c.NewClient = defaultClientCreator
	}
}

func WithHttpClientFrom(creator httpBase.HttpClientFactory) Option {
	return func(c *Configure) {
		c.NewClient = func() {
			httpBase.NewHttpClient = func() httpBase.Client {
				return creator()
			}
		}
	}
}

func ConfigureHttpServer(opts ...ServerOption) Option {
	return func(c *Configure) {
		c.SetupServer = func() {
			if c.Server == nil && len(opts) == 0 {
				return
			}

			for _, opt := range opts {
				opt(c.Server)
			}
		}
	}
}

func WithStarter(starter ...BlockingStarter) Option {
	return func(c *Configure) {
		c.BlockingStarters = append(c.BlockingStarters, starter...)
	}
}

func WithJsonCodec(json codecBase.JsonCodec) Option {
	return func(c *Configure) {
		c.NewCodec = func() {
			codecBase.Json = json
		}
	}
}

func WithLoggerFrom(l logBase.FullLogger) Option {
	return func(c *Configure) {
		c.SetupLog = func() {
			logBase.SetLogger(l)
		}
	}
}

func WithPreStartHook(fn func()) Option {
	return func(c *Configure) {
		c.PreStarterHook = fn
	}
}

func WithPostStartHook(fn func()) Option {
	return func(c *Configure) {
		c.PostStarterHook = fn
	}
}
