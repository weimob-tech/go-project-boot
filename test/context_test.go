package test

import (
	"fmt"
	"testing"

	"github.com/bytedance/sonic"
	jsoniter "github.com/json-iterator/go"
	. "github.com/smartystreets/goconvey/convey"
	"github.com/weimob-tech/go-project-base/pkg/codec"
	"github.com/weimob-tech/go-project-base/pkg/x"
)

func TestCtx(t *testing.T) {
	Convey("test json", t, func() {
		intJson := `{"code":{"errcode": 123, "errmsg": "success"}}`
		strJson := `{"code":{"errcode": "123", "errmsg": "success"}}`

		Convey("using jsoniter", func() {
			var err error
			var intResult x.EmptyResult
			var strResult x.EmptyResult
			err = jsoniter.Unmarshal([]byte(intJson), &intResult)
			So(err, ShouldBeNil)
			err = jsoniter.Unmarshal([]byte(strJson), &strResult)
			So(err, ShouldBeNil)

			fmt.Println(codec.ToJson(&intResult))
			fmt.Println(codec.ToJson(&strResult))
		})

		Convey("using sonic", func() {
			var err error
			var intResult x.EmptyResult
			var strResult x.EmptyResult
			err = sonic.Unmarshal([]byte(intJson), &intResult)
			So(err, ShouldBeNil)
			err = sonic.Unmarshal([]byte(strJson), &strResult)
			So(err, ShouldBeNil)

			fmt.Println(codec.ToJson(&intResult))
			fmt.Println(codec.ToJson(&strResult))
		})
	})
}
