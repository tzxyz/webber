package webber

import (
	"strings"
	"net/http"
	"gopkg.in/xmlpath.v2"
)

type Response struct {
	code    int
	status  string
	headers map[string][]string
	cookies []*http.Cookie
	body    string
	request *Request
	node    *xmlpath.Node
	charset string
}

func NewResponse() *Response {
	return &Response{
		headers: make(map[string][]string),
		cookies: make([]*http.Cookie, 16),
	}
}

func (r *Response) GetUrl() string {
	return r.request.url
}

func (r *Response) Code(code int) *Response {
	r.code = code
	return r
}

func (r *Response) Status(status string) *Response {
	r.status = status
	return r
}

func (r *Response) Body(body string) *Response {
	r.body = body
	return r
}

func (r *Response) Headers(headers map[string][]string) *Response {
	r.headers = headers
	return r
}

func (r *Response) Cookies(cookies []*http.Cookie) *Response {
	r.cookies = cookies
	return r
}

func (r *Response) Request(request *Request) *Response {
	r.request = request
	return r
}

func (r *Response) Html() *Response {
	node, err := xmlpath.ParseHTML(strings.NewReader(r.body))
	// todo 不应该panic
	if err != nil {
		panic(err)
	}
	r.node = node
	return r
}

func (r *Response) Xpath(path string) []string {
	p := xmlpath.MustCompile(path)
	values := make([]string, 0)
	for n := p.Iter(r.node); n.Next(); {
		value := n.Node().String()
		values = append(values, value)
	}
	return values
}
