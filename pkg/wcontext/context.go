package wcontext

import (
	"context"
	"fmt"
	"github.com/redis/go-redis/v9"
	"github.com/weimob-tech/go-project-base/pkg/config"
	"github.com/weimob-tech/go-project-boot/pkg/auth"
	"github.com/weimob-tech/go-project-boot/pkg/db"
	"gorm.io/gorm"
)

type ContextKey string

const (
	ContextKeyId ContextKey = "wm_context"
)

var ctx *Context

type Context struct {
	context.Context
	auth  *auth.OAuth
	mysql *gorm.DB
	redis *redis.Client
}

func Global() *Context {
	return ctx
}

func Setup() {
	ctx = &Context{
		Context: context.Background(),
	}
	// setup repositories
	if config.GetBool("mysql.enable") {
		db, err := db.NewMysql(ctx)
		if err != nil {
			panic(fmt.Errorf("fatal error when create mysql client: %w", err))
		}
		ctx.mysql = db
	}
	if config.GetBool("redis.enable") {
		db, err := db.NewRedis(ctx)
		if err != nil {
			panic(fmt.Errorf("fatal error when create redis client: %w", err))
		}
		ctx.redis = db
	}
	// setup oauth client
	// todo: later
	ctx.auth = &auth.OAuth{}
}

func (c *Context) With(ctx context.Context) *Context {
	return &Context{
		Context: ctx,
		auth:    c.auth,
		mysql:   c.mysql,
		redis:   c.redis,
	}
}

func (c *Context) Mysql() *gorm.DB {
	return c.mysql
}

func (c *Context) Redis() *redis.Client {
	return c.redis
}
