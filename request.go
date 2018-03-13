package webber

import "io"
import "net/http"

type Request struct {
	req *http.Request
	url string
	method string
	header http.Header
	body io.ReadCloser
}

func newRequest(req *http.Request) *Request {
	return &Request{
		req: req,
		url: req.URL.String(),
		method: req.Method,
		header: req.Header,
		body: req.Body,
	}
}