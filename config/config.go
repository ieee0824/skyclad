package config

import (
	"regexp"
	"time"
)

type Config struct {
	AlertLimit time.Duration
	Interval   time.Duration
	Ignore     *regexp.Regexp
	Notifer    string
}

func New(
	al string,
	interval string,
	ignoreRule string,
	notifer string,
) *Config {
	ald, err := time.ParseDuration(al)
	if err != nil {
		panic(err)
	}
	id, err := time.ParseDuration(interval)
	if err != nil {
		panic(err)
	}
	var ignoreRuleRegexp *regexp.Regexp
	if ignoreRule != "" {
		ignoreRuleRegexp = regexp.MustCompilePOSIX(ignoreRule)
	}

	return &Config{
		AlertLimit: ald,
		Interval:   id,
		Ignore:     ignoreRuleRegexp,
		Notifer:    notifer,
	}
}
