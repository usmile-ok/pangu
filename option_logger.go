package pangu

import (
	`github.com/storezhang/simaqian`
)

var _ option = (*optionLogger)(nil)

type optionLogger struct {
	logger simaqian.Logger
}

// GLogger 配置日志
func GLogger(logger simaqian.Logger) *optionLogger {
	return &optionLogger{
		logger: logger,
	}
}

func (l *optionLogger) apply(options *options) {
	options.logger = l.logger
}
