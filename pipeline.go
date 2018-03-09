package webber

type Pipeline func(result *Result)

var ConsolePipeline = func(result *Result) {
	logger.Info(result)
}
