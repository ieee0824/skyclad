package main

import (
	"flag"

	_ "github.com/ieee0824/skyclad/plugins/slack"

	"github.com/ieee0824/getenv"
	"github.com/ieee0824/skyclad/config"
	"github.com/ieee0824/skyclad/observer"
)

func main() {
	alertLimit := flag.String(
		"lt",
		getenv.String("LIMIT_DURATION", "24h"),
		"alert limit duration.",
	)
	observeInterval := flag.String(
		"it",
		getenv.String("MONITORING_INTERVAL", "10m"),
		"monitoring interval",
	)
	ignoreRule := flag.String(
		"ignore",
		getenv.String("IGNORE_RULE"),
		"ignore container rule. Can use posix regexp.",
	)
	notifer := flag.String(
		"n",
		getenv.String("NOTIFER_TYPE"),
		"Notification destination. If it is empty, stdout.",
	)
	flag.Parse()

	cfg := config.New(
		*alertLimit,
		*observeInterval,
		*ignoreRule,
		*notifer,
	)

	o := observer.New(cfg)

	if err := o.Observe(); err != nil {
		panic(err)
	}
}
