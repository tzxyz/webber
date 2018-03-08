package webber

type Processor func(response *Response) *Result
