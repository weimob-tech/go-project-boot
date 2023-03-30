package db

import (
	"context"
	. "github.com/smartystreets/goconvey/convey"
	"testing"
	"time"
)

func TestNewRedisClient(t *testing.T) {
	Convey("Create redis client should success", t, func() {
		rdb, err := NewRedis(context.Background())
		So(err, ShouldBeNil)

		Convey("Redis get string should success", func() {
			key := "go-ability-core"
			_, err := rdb.Set(context.Background(), key, "value", time.Second).
				Result()
			So(err, ShouldBeNil)

			val, err := rdb.Get(context.Background(), key).Result()
			So(err, ShouldBeNil)
			So(val, ShouldEqual, "value")
		})
	})
}
