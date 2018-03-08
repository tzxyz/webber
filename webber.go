package webber

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

func (engine *Webber) Name(name string) *Webber {
	engine.name = name
	return engine
}

func (engine *Webber) StartUrls(urls ...string) *Webber {
	engine.startUrls = urls
	return engine
}

func (engine *Webber) Downloader(downloader Downloader) *Webber {
	engine.downloader = downloader
	return engine
}

func (engine *Webber) Scheduler(scheduler Scheduler) *Webber {
	engine.scheduler = scheduler
	return engine
}

func (engine *Webber) Processor(processor Processor) *Webber {
	engine.processor = processor
	return engine
}

func (engine *Webber) Pipelines(pipelines ...Pipeline) *Webber {
	engine.pipelines = pipelines
	return engine
}
