package http

import (
	"context"
	"crypto/tls"
	"fmt"
	"github.com/cloudwego/hertz/pkg/protocol"
	"github.com/weimob-tech/go-project-base/pkg/config"
	"github.com/weimob-tech/go-project-base/pkg/http"
	"github.com/weimob-tech/go-project-base/pkg/wlog"
	"io"
	"time"

	"github.com/cloudwego/hertz/pkg/app/client"
	"github.com/cloudwego/hertz/pkg/network/standard"
)

type hertzRequest struct {
	*protocol.Request
}

func (h *hertzRequest) GetHeader() http.Header {
	return &h.Header
}

func (h *hertzRequest) GetRequest() any {
	return h.Request
}

func (h *hertzRequest) SetFile(formName, fileName string, file io.Reader) {
	h.Request.SetFileReader(formName, fileName, file)
	h.Request.SetMultipartFormData(map[string]string{"name": fileName})
}

type hertzResponse struct {
	*protocol.Response
}

func (h *hertzResponse) GetResponse() any {
	return h.Response
}

type hertzClient struct {
	*client.Client
	logLvl http.LogLvl
}

func (hc *hertzClient) NewRequest() http.Request {
	return &hertzRequest{&protocol.Request{}}
}

func (hc *hertzClient) NewResponse() http.Response {
	return &hertzResponse{&protocol.Response{}}
}

func (hc *hertzClient) Do(ctx context.Context, request http.Request, response http.Response) (err error) {
	return elapsed(hc.logLvl, func() error {
		input := request.GetRequest().(*protocol.Request)
		output := response.GetResponse().(*protocol.Response)

		// log before
		if hc.logLvl >= http.LogLvlBase {
			wlog.Infof("> HTTP Request: %s", input.Method())
			wlog.Infof("> HTTP Request: %s", input.URI().String())
		}
		if hc.logLvl >= http.LogLvlHeader {
			wlog.Infof("> HTTP Request header:")
			input.Header.VisitAllCustomHeader(func(k, v []byte) {
				wlog.Infof("> > %s=%s", string(k), string(v))
			})
		}
		if hc.logLvl >= http.LogLvlBody {
			wlog.Infof("> HTTP Request body: %s", string(input.BodyBytes()))
		}

		// do request
		err := hc.Client.Do(ctx,
			request.GetRequest().(*protocol.Request),
			response.GetResponse().(*protocol.Response))

		// log after
		if hc.logLvl >= http.LogLvlBase {
			wlog.Infof(">")
			wlog.Infof("> HTTP Response: %d", output.StatusCode())
		}
		if hc.logLvl >= http.LogLvlHeader {
			wlog.Infof("> HTTP Response header:")
			output.Header.VisitAll(func(k, v []byte) {
				wlog.Infof("> > %s=%s", string(k), string(v))
			})
		}
		if hc.logLvl >= http.LogLvlBody {
			wlog.Infof("> HTTP Response body: %s", string(output.BodyBytes()))
		}
		return err
	})
}

func elapsed(lvl http.LogLvl, fn func() error) error {
	if lvl == http.LogLvlWarn {
		return fn()
	}
	var start = time.Now()
	err := fn()
	wlog.Infof("> HTTP: elapsed %v", time.Since(start))
	wlog.Info(">")
	return err
}

func (hc *hertzClient) GetClient() any {
	return hc.Client
}

func NewHertzClient() (cli *hertzClient) {
	// todo add timeout, proxy, async
	var (
		hc  *client.Client
		err error
	)
	// todo add middleware trace, signer, logger
	if config.GetString("client.schema") == "https" {
		clientCfg := &tls.Config{
			InsecureSkipVerify: true,
		}
		hc, err = client.NewClient(
			client.WithTLSConfig(clientCfg),
			client.WithDialer(standard.NewDialer()),
		)
	} else {
		hc, err = client.NewClient(client.WithDialer(standard.NewDialer()))
	}
	if err != nil {
		panic(fmt.Errorf("create http client failed %v", err))
	}
	return &hertzClient{hc, http.GetLevel(config.GetString("client.log.lvl"))}
}
