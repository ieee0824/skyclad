package observer

import (
	"strings"
	"time"

	"github.com/ieee0824/skyclad/client"
	"github.com/ieee0824/skyclad/client/iface"
	"github.com/ieee0824/skyclad/client/io"
	"github.com/ieee0824/skyclad/config"
	"github.com/ieee0824/skyclad/notifer"
)

type Observer struct {
	_      struct{}
	client clientiface.ClientAPI
	config *config.Config
}

func New(cfg *config.Config) *Observer {
	return &Observer{
		client: &client.Client{},
		config: cfg,
	}
}

func (o *Observer) observe(info clientio.GetContainerOutput) (bool, error) {
	now := time.Now()
	uptime, err := o.client.GetUptime(info.ContainerID)
	if err != nil {
		return false, err
	}

	if o.config.AlertLimit < now.Sub(uptime) {
		return true, nil
	}
	return false, nil
}

func (o *Observer) Observe() error {
	ticker := time.NewTicker(o.config.Interval)

	for {
		select {
		case <-ticker.C:
			containers, err := o.client.GetContainers()
			if err != nil {
				return err
			}
			oldContainers := []clientio.GetContainerOutput{}
			for _, container := range containers {
				if strings.Contains(container.ImageName, "skyclad") || strings.Contains(container.ContainerName, "skyclad") {
					continue
				}
				if o.config.Ignore != nil && o.config.Ignore.MatchString(container.ImageName) {
					continue
				}
				t, err := o.observe(container)
				if err != nil {
					continue
				}
				if t {
					oldContainers = append(oldContainers, container)
				}
			}
			if err := notifer.GetNotifer("").Notice(oldContainers); err != nil {
				panic(err)
			}
		}
	}
}
