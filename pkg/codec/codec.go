package codec

import "github.com/weimob-tech/go-project-base/pkg/codec"

func SetupDefault() {
	codec.Json = &sonicCodec{}
}
