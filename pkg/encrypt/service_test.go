package encrypt

import (
	"context"
	"github.com/weimob-tech/go-project-base/pkg/auth"
	"github.com/weimob-tech/go-project-base/pkg/codec"
	"github.com/weimob-tech/go-project-base/pkg/config"
	"github.com/weimob-tech/go-project-base/pkg/hook"
	"testing"
)

func init() {
	config.SetDefault("client.log.lvl", "body")
	config.SetDefault("client.schema", "http")
	config.SetDefault("client.domain", "dopen.weimob.com")
	config.SetDefault("client.oauth.domain", "dopen.weimob.com")
	config.SetDefault("weimob.cloud.foo.clientId", "F580B0441D733E0A8B34367D168772DD")
	config.SetDefault("weimob.cloud.foo.clientSecret", "E51E9913EAC1C0461354B7B74184189A")

	hook.ApplyPostStartHook()
	Setup()
}

func TestEncryptService(t *testing.T) {
	token := auth.GetCCToken(context.Background(), "foo", "", "")
	req := NewBosEncryptRequest(token, "18612341234")
	res, err := DoRequest(req)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(codec.ToJson(res))

	req = NewBosDecryptRequest(token, res.Data.Result)
	res, err = DoRequest(req)
	if err != nil {
		t.Fatal(err)
	}
	if res.Data.Result != "18612341234" {
		t.Errorf("decrypt failed, expect: 18612341234, got: %s", res.Data.Result)
	}
	t.Log(codec.ToJson(res))
}
