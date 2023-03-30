package encrypt

func NewBosDecryptRequest(token, body string) *Request {
	return &Request{
		Api:   "decrypt",
		Body:  body,
		Token: token,
		Spec:  "bos",
	}
}

func NewDecryptRequest(token, body string) *Request {
	return &Request{
		Api:   "decrypt",
		Body:  body,
		Token: token,
	}
}
