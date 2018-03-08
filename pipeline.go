package webber

import log "github.com/sirupsen/logrus"

type Pipeline func(result *Result)

var ConsolePipeline = func(result *Result) {
	log.Info(result)
}
