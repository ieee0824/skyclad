package main

import (
	"flag"

	_ "github.com/ieee0824/skyclad/plugins/slack"

	"github.com/ieee0824/skyclad/config"
	"github.com/ieee0824/skyclad/observer"
)

func main() {
	alertLimit := flag.String(
		"lt",
		"24h",
		"alert limit duration.",
	)
	observeInterval := flag.String(
		"it",
		"10s",
		"monitoring interval",
	)
	ignoreRule := flag.String(
		"ignore",
		"",
		"ignore container rule",
	)
	notifer := flag.String(
		"n",
		"",
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
