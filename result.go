package webber

type Items map[string]interface{}

type Result struct {
	items    Items
	nextUrls []string
}

func NewResult() *Result {
	return &Result{
		items:    make(Items),
		nextUrls: make([]string, 0, 16),
	}
}

func (r *Result) PushItem(key string, value interface{}) *Result {
	r.items[key] = value
	return r
}

func (r *Result) PushUrls(urls ...string) *Result {
	r.nextUrls = urls
	return r
}

func (r *Result) Items() Items {
	return r.items
}

func (r *Result) NextUrls() []string {
	return r.nextUrls
}

func (r *Result) HasNextUrl() bool {
	return len(r.nextUrls) > 0
}