package webber

import (
	"sync"
	"time"
	log "github.com/sirupsen/logrus"
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

	for _, url := range w.startUrls {
		w.scheduler.Push(NewRequest().Url(url))
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

				resp := w.downloader(req)

				log.Info("Downloader download Url: ", req.url)

				result := w.processor(resp)

				log.Info("Processor process Url: ", req.url)

				for _, pipeline := range w.pipelines {
					pipeline(result)
				}

				if result.HasNextUrl() {

					for _, url := range result.NextUrls() {
						w.scheduler.Push(NewRequest().Url(url))
					}
				}
			}()
			wg.Wait()
		}
	}
}
