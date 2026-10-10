package gssion

import gonfig "github.com/vrianta/gonfig/v1"

var config = gonfig.New[struct {
	Https           bool `env:"SESSION_HTTPS" arg:"session-https"`
	MaxSessionCount int  `env:"SESSION_MAX_SESSION_COUNT" arg:"session-max-session-count" default:"10000"`
}](false)
