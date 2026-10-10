package gssion

import gonfig "github.com/vrianta/gonfig/v1"

var config = gonfig.New[struct {
	Https           bool `env:"SESSION_HTTPS" arg:"session-https"`                 // https need enabled if the application is served through https keep it false mostly in local environments
	MaxSessionCount int  `env:"SESSION_HTTPS" arg:"session-https" default:"10000"` // cap to 10000
}](false)
