package clientiface

import (
	"time"

	"github.com/ieee0824/skyclad/client/io"
)

// ClientAPI is docker client interface
type ClientAPI interface {
	GetContainers() ([]clientio.GetContainerOutput, error)
	GetStatus(cID string) (clientio.GetContainerStatusOutput, error)
	GetStatuses(cIDs []string) ([]clientio.GetContainerStatusOutput, error)
	GetUptime(cID string) (time.Time, error)
}
