package webber

import requests "github.com/parnurzeal/gorequest"

type Downloader func(request *Request) (*Response, []error)

var HttpDownloader = func(request *Request) (*Response, []error) {

	logger.Debug("Starting download url: " + request.req.URL.String())

	resp, _, errs := requests.New().Get(request.url).EndBytes()

	if errs != nil && len(errs) != 0 {
		return nil, errs
	}

	return newResponse(request, resp), nil
}
