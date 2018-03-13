package webber

import (
	"net/http"
	"gopkg.in/xmlpath.v2"
)

type Response struct {
	req     *Request
	resp    *http.Response
	url     string
	node    *xmlpath.Node
	charset string
}

func newResponse(req *Request, resp *http.Response) *Response {
	return &Response{
		req:  req,
		resp: resp,
		url: req.url,
	}
}

func (r *Response) Url() string {
	return r.url
}

func (r *Response) Html() *Response {
	node, err := xmlpath.ParseHTML(r.resp.Body)
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
