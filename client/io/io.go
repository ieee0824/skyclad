package clientio

import "time"

type GetContainerOutput struct {
	ContainerID   string     `json:"container_id"`
	ImageName     string     `json:"image_name"`
	ContainerName string     `json:"container_name"`
	Uptime        *time.Time `json:"uptime"`
}

type GetContainerStatusOutput map[string]interface{}
