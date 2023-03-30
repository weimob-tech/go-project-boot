package wlog

import (
	"fmt"
	"github.com/rs/zerolog"
	"github.com/weimob-tech/go-project-base/pkg/wlog"
	"io"
	"os"
	"time"
)

type zeroLogger struct {
	zerolog.Logger
	level wlog.Level
}

func NewPrettyLog() wlog.FullLogger {
	zerolog.TimeFieldFormat = time.RFC3339Nano
	var l = zerolog.New(zerolog.ConsoleWriter{Out: os.Stdout, TimeFormat: time.RFC3339Nano}).
		With().
		Timestamp().
		Caller().
		Logger()
	return &zeroLogger{Logger: l}
}

func NewJsonLog() wlog.FullLogger {
	zerolog.TimeFieldFormat = time.RFC3339Nano
	var l = zerolog.New(os.Stdout).
		With().
		Timestamp().
		Caller().
		Logger()
	return &zeroLogger{Logger: l}
}

func (ll *zeroLogger) SetOutput(w io.Writer) {
	ll.Logger = ll.Logger.Output(w)
}

func (ll *zeroLogger) SetLevel(lv wlog.Level) {
	ll.level = lv
}

func (ll *zeroLogger) GetLogger() any {
	return ll.Logger
}

func (ll *zeroLogger) l() *zerolog.Logger {
	return &ll.Logger
}

func (ll *zeroLogger) Fatal(v ...interface{}) {
	ll.l().Fatal().CallerSkipFrame(2).Msg(fmt.Sprint(v...))
}

func (ll *zeroLogger) Error(v ...interface{}) {
	ll.l().Error().CallerSkipFrame(2).Msg(fmt.Sprint(v...))
}

func (ll *zeroLogger) Warn(v ...interface{}) {
	ll.l().Warn().CallerSkipFrame(2).Msg(fmt.Sprint(v...))
}

func (ll *zeroLogger) Info(v ...interface{}) {
	ll.l().Info().CallerSkipFrame(2).Msg(fmt.Sprint(v...))
}

func (ll *zeroLogger) Debug(v ...interface{}) {
	ll.l().Debug().CallerSkipFrame(2).Msg(fmt.Sprint(v...))
}

func (ll *zeroLogger) Trace(v ...interface{}) {
	ll.l().Trace().CallerSkipFrame(2).Msg(fmt.Sprint(v...))
}

func (ll *zeroLogger) Fatalf(format string, v ...interface{}) {
	ll.l().Fatal().CallerSkipFrame(2).Msgf(format, v...)
}

func (ll *zeroLogger) Errorf(format string, v ...interface{}) {
	ll.l().Error().CallerSkipFrame(2).Msgf(format, v...)
}

func (ll *zeroLogger) Warnf(format string, v ...interface{}) {
	ll.l().Warn().CallerSkipFrame(2).Msgf(format, v...)
}

func (ll *zeroLogger) Infof(format string, v ...interface{}) {
	ll.l().Info().CallerSkipFrame(2).Msgf(format, v...)
}

func (ll *zeroLogger) Debugf(format string, v ...interface{}) {
	ll.l().Debug().CallerSkipFrame(2).Msgf(format, v...)
}

func (ll *zeroLogger) Tracef(format string, v ...interface{}) {
	ll.l().Trace().CallerSkipFrame(2).Msgf(format, v...)
}
