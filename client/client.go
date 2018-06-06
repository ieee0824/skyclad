package client

import (
	"encoding/json"
	"errors"
	"log"
	"os/exec"
	"strings"
	"time"

	"github.com/ieee0824/skyclad/client/io"
)

var DefaultClient = &Client{}

func GetContainers() ([]clientio.GetContainerOutput, error) {
	return DefaultClient.GetContainers()
}

func GetStatus(cID string) (clientio.GetContainerStatusOutput, error) {
	return DefaultClient.GetStatus(cID)
}

func GetStatuses(cIDs []string) ([]clientio.GetContainerStatusOutput, error) {
	return DefaultClient.GetStatuses(cIDs)
}

func GetUptime(cID string) (time.Time, error) {
	return DefaultClient.GetUptime(cID)
}

// Client is docker client
type Client struct {
}

func (*Client) GetContainers() ([]clientio.GetContainerOutput, error) {
	cmd := exec.Command("docker", "ps", "--format", `table {{.ID}} {{.Image}} {{.Names}}`)
	out, err := cmd.Output()
	if err != nil {
		log.Println("error: docker ps")
		return nil, err
	}
	containers := strings.Split(strings.TrimSuffix(string(out), "\n"), "\n")[1:]
	ret := make([]clientio.GetContainerOutput, len(containers))

	for i, container := range containers {
		infos := strings.Split(container, " ")
		ret[i].ContainerID = infos[0]
		ret[i].ImageName = infos[1]
		ret[i].ContainerName = infos[2]
	}

	return ret, nil
}

func (*Client) GetStatus(cID string) (clientio.GetContainerStatusOutput, error) {
	cmd := exec.Command("docker", "inspect", cID)
	ret := []map[string]interface{}{}

	out, err := cmd.Output()
	if err != nil {
		log.Println("error: docker inspect")
		return nil, err
	}

	if err := json.Unmarshal(out, &ret); err != nil {
		return nil, err
	}
	if len(ret) == 0 {
		return map[string]interface{}{}, nil
	}
	return ret[0], nil
}

func (c *Client) GetStatuses(cIDs []string) ([]clientio.GetContainerStatusOutput, error) {
	if len(cIDs) == 0 {
		return nil, errors.New("status is empty")
	}
	ret := make([]clientio.GetContainerStatusOutput, 0, len(cIDs))
	for _, id := range cIDs {
		status, err := c.GetStatus(id)
		if err != nil {
			continue
		}
		ret = append(ret, status)
	}
	if len(ret) == 0 {
		return nil, errors.New("status is empty")
	}

	return ret, nil
}

func (c *Client) GetUptime(cID string) (time.Time, error) {
	status, err := c.GetStatus(cID)
	if err != nil {
		return time.Time{}, err
	}
	state, ok := status["State"].(map[string]interface{})
	if !ok {
		return time.Time{}, errors.New("not found State")
	}

	t, err := time.Parse(time.RFC3339, state["StartedAt"].(string))
	if err != nil {
		return time.Time{}, err
	}

	return t.In(time.FixedZone("Asia/Tokyo", 9*60*60)), nil
}
