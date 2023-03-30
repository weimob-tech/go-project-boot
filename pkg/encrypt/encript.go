package encrypt

func NewBosEncryptRequest(token, body string) *Request {
	return &Request{
		Api:   "encrypt",
		Body:  body,
		Token: token,
		Spec:  "bos",
	}
}

func NewBosIsEncryptRequest(token, body string) *Request {
	return &Request{
		Api:   "isEncrypt",
		Body:  body,
		Token: token,
		Spec:  "bos",
	}
}

func NewEncryptRequest(token, body string) *Request {
	return &Request{
		Api:   "encrypt",
		Body:  body,
		Token: token,
	}
}

func NewIsEncryptRequest(token, body string) *Request {
	return &Request{
		Api:   "isEncrypt",
		Body:  body,
		Token: token,
	}
}
