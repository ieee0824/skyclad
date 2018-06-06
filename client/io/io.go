package clientio

type GetContainerOutput struct {
	ContainerID   string
	ImageName     string
	ContainerName string
}

type GetContainerStatusOutput map[string]interface{}
