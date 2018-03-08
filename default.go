package webber

var (
	DefaultDownloader = HttpDownloader
	DefaultScheduler  = InMemoryScheduler
	DefaultPipelines  = []Pipeline{ConsolePipeline}
)

