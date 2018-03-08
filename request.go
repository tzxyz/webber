package webber

type Request struct {
	url     string
	method  string
	body    string
	headers map[string]string
	cookies map[string]string
}

func NewRequest() *Request {
	return &Request{
		headers: make(map[string]string),
		cookies: make(map[string]string),
	}
}

func (r *Request) Url(url string) *Request {
	r.url = url
	return r
}

func (r *Request) Method(method string) *Request {
	r.method = method
	return r
}

func (r *Request) Body(body string) *Request {
	r.body = body
	return r
}

func (r *Request) Header(name string, value string) *Request {
	r.headers[name] = value
	return r
}

func (r *Request) Headers(headers map[string]string) *Request {
	r.headers = headers
	return r
}

func (r *Request) Cookie(name string, value string) *Request {
	r.cookies[name] = value
	return r
}

func (r *Request) Cookies(cookies map[string]string) *Request {
	r.cookies = cookies
	return r
}
