package webber

import (
	"sync"
	"time"
	"runtime"
	"fmt"
	"strings"
	"net/http"
)

type Webber struct {
	name       string
	startUrls  []string
	downloader Downloader
	scheduler  Scheduler
	processor  Processor
	pipelines  []Pipeline
	running    chan *Request
}

func New() *Webber {
	return &Webber{
		startUrls:  make([]string, 0),
		downloader: DefaultDownloader,
		scheduler:  DefaultScheduler,
		pipelines:  DefaultPipelines,
		running:    make(chan *Request, 16),
	}
}

func (w *Webber) Name(name string) *Webber {
	w.name = name
	return w
}

func (w *Webber) StartUrls(urls ...string) *Webber {
	w.startUrls = urls
	return w
}

func (w *Webber) Downloader(downloader Downloader) *Webber {
	w.downloader = downloader
	return w
}

func (w *Webber) Scheduler(scheduler Scheduler) *Webber {
	w.scheduler = scheduler
	return w
}

func (w *Webber) Processor(processor Processor) *Webber {
	w.processor = processor
	return w
}

func (w *Webber) Pipelines(pipelines ...Pipeline) *Webber {
	w.pipelines = pipelines
	return w
}

func (w *Webber) Start() {

	runtime.GOMAXPROCS(runtime.NumCPU())

	for _, url := range w.startUrls {
		r, err := http.NewRequest("GET", url, nil)
		if err != nil {

		}
		w.scheduler.Push(newRequest(r))
	}

	var wg sync.WaitGroup

	for {

		req := w.scheduler.Poll()

		if req == nil && len(w.running) == 0 {
			break
		} else if req == nil {
			time.Sleep(time.Second)
		} else {
			wg.Add(1)
			go func() {

				w.running <- req

				defer func() {
					<-w.running
					wg.Done()
				}()

				resp, errs := w.downloader(req)

				if errs != nil && len(errs) != 0 {
					var msg = make([]string, 0)
					for _, err := range errs {
						msg = append(msg, fmt.Sprint(err))
					}
					logger.Error(strings.Join(msg, ","))
				}

				result := w.processor(resp)

				for _, pipeline := range w.pipelines {
					pipeline(result)
				}

				if result.HasNextUrl() {

					for _, url := range result.NextUrls() {
						r, err := http.NewRequest("GET", url, nil)
						if err != nil {

						}
						w.scheduler.Push(newRequest(r))
					}
				}
			}()
			wg.Wait()
		}
	}
}
