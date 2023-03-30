package http

import (
	"context"
	"fmt"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
	"github.com/weimob-tech/go-project-base/pkg/http"
	"github.com/weimob-tech/go-project-base/pkg/x"
)

func (s *hertzServer) AddExtendCallback(config *http.ExtendCallbackConfig) {
	s.Handle(config.Method, config.Path, func(c context.Context, ctx *app.RequestContext) {
		cc := &http.ExtendCallbackContext{
			Path:    string(ctx.Request.Path()),
			Method:  string(ctx.Method()),
			Context: c,
		}
		// extract headers
		if len(config.Headers) > 0 {
			var headers = make(map[string]string, len(config.Headers))
			for _, key := range config.Headers {
				headers[key] = string(ctx.GetHeader(key))
			}
			cc.Headers = headers
		}
		// extract path param
		if len(config.Params) > 0 && len(ctx.Params) > 0 {
			params := make(map[string]string, len(config.Params))
			for _, key := range config.Params {
				params[key] = ctx.Param(key)
			}
			cc.Params = params
		}
		// extract url queries
		if len(config.Queries) > 0 && ctx.QueryArgs().Len() > 0 {
			queries := make(map[string]string, len(config.Queries))
			for _, key := range config.Queries {
				queries[key] = ctx.Query(key)
			}
			cc.Queries = queries
		}
		// extract body
		payload, err := ctx.Body()
		if err != nil {
			ctx.JSON(consts.StatusBadRequest, x.Fail("90400", fmt.Sprintf("bad request %v", err)))
			return
		}
		cc.Body = payload
		// 执行回调方法
		res, err := config.Callback(cc)
		if err != nil {
			ctx.JSON(consts.StatusOK, err)
			return
		}
		if config.NullToEmpty && res == nil {
			ctx.Data(consts.StatusOK, "application/json", nil)
			return
		}
		ctx.JSON(consts.StatusOK, res)
	})
}
