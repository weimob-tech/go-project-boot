package wlog

import (
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"github.com/weimob-tech/go-project-base/pkg/config"
	"github.com/weimob-tech/go-project-base/pkg/wlog"
)

var Logger *zeroLogger

func init() {
	Logger = &zeroLogger{Logger: log.Logger}
}

func Setup() {
	if config.Debug("log") {
		Logger = NewPrettyLog().(*zeroLogger)
	} else {
		Logger = NewJsonLog().(*zeroLogger)
	}
	wlog.SetLogger(Logger)
}

func L() *zerolog.Logger {
	return &Logger.Logger
}

func D() *zerolog.Event {
	return L().Debug().Caller(1)
}

func I() *zerolog.Event {
	return L().Info().Caller(1)
}

func W(err ...error) *zerolog.Event {
	l := L().Warn().Caller(1)
	for _, e := range err {
		l = l.Err(e)
	}
	return l
}

func E(err ...error) *zerolog.Event {
	l := L().Error().Caller(1)
	for _, e := range err {
		l = l.Err(e)
	}
	return l
}

func F(err ...error) *zerolog.Event {
	l := L().Fatal().Caller(1)
	for _, e := range err {
		l = l.Err(e)
	}
	return l
}

func P(err ...error) *zerolog.Event {
	l := L().Panic().Caller(1)
	for _, e := range err {
		l = l.Err(e)
	}
	return l
}
