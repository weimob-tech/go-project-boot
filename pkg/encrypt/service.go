package encrypt

import (
	"context"
	"fmt"
	"github.com/weimob-tech/go-project-base/pkg/codec"
	"github.com/weimob-tech/go-project-base/pkg/config"
	"github.com/weimob-tech/go-project-base/pkg/http"
	"github.com/weimob-tech/go-project-base/pkg/wlog"
	"github.com/weimob-tech/go-project-base/pkg/x"
)

const (
	securityBosPath   = "%s/apigw/bos/v2.0/security/%s?accesstoken=%s"
	batchSecurityPath = "%s/api/1_0/ec/order/%s?accesstoken=%s"
)

var (
	client  http.Client
	baseUrl string
)

type Request struct {
	Api   string
	Body  string
	Spec  string
	Token string
}

type RequestSource struct {
	Source string `json:"source,omitempty"`
}

type Data struct {
	Result string `json:"result,omitempty"`
}

type Response struct {
	Code         x.Code `json:"code,omitempty"`
	GlobalTicket string `json:"globalTicket,omitempty"`
	Data         Data   `json:"data,omitempty"`
}

func (response *Response) Unmarshall(res http.Response) error {
	return codec.Json.Unmarshal(res.Body(), response)
}

func Setup() {
	client = http.Global
	baseUrl = config.GetString("client.schema") + "://" + config.GetString("client.domain")
}

func buildRequestUrl(api, spec, token string) string {
	if spec == "bos" {
		return fmt.Sprintf(securityBosPath, baseUrl, api, token)
	} else {
		return fmt.Sprintf(batchSecurityPath, baseUrl, api, token)
	}
}

func DoRequest(request *Request) (response *Response, err error) {
	url := buildRequestUrl(request.Api, request.Spec, request.Token)
	res := client.NewResponse()

	req := client.NewRequest()
	req.GetHeader().SetMethod(http.MethodPost)
	req.GetHeader().SetContentTypeBytes(http.ContentTypeJsonByte)
	req.SetRequestURI(url)
	req.SetBody(codec.ToJsonByte(&RequestSource{request.Body}))

	err = client.Do(context.Background(), req, res)
	if err != nil {
		return nil, err
	}
	response = new(Response)
	err = response.Unmarshall(res)
	if err != nil {
		return nil, err
	}
	if response.Code.Errcode != "0" {
		wlog.Warnf("encrypt service request failed, code: %d, msg: %s", response.Code.Errcode, response.Code.Errmsg)
		return nil, x.ErrorOf(response.Code)
	}
	return
}
