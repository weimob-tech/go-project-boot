package codec

import "github.com/bytedance/sonic"

type sonicCodec struct{}

func (s sonicCodec) Marshal(v any) ([]byte, error) {
	return sonic.Marshal(v)
}

func (s sonicCodec) MarshalString(v any) (string, error) {
	return sonic.MarshalString(v)
}

func (s sonicCodec) Unmarshal(data []byte, v any) error {
	return sonic.Unmarshal(data, v)
}

func (s sonicCodec) UnmarshalString(data string, v any) error {
	return sonic.UnmarshalString(data, v)
}

func (s sonicCodec) GetString(data []byte, keys ...any) (string, error) {
	node, err := sonic.Get(data, keys...)
	if err != nil {
		return "", err
	}
	return node.String()
}
