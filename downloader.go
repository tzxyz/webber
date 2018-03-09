package webber

import "net/http"
import requests "github.com/parnurzeal/gorequest"

type Downloader func(request *Request) *Response

var HttpDownloader = func(request *Request) *Response {
	var resp *http.Response
	var errs []error

	logger.Debug("Starting download url: " + request.url)

	resp, body, errs := requests.New().Get(request.url).EndBytes()

	if len(errs) != 0 {
		panic(errs[0])
	}

	return NewResponse().
		Code(resp.StatusCode).
		Status(resp.Status).
		Body(string(body)).
		Headers(resp.Header).
		Cookies(resp.Cookies()).
		Request(request)
}
