package manager

import (
	"bytes"
	"encoding/json"

	"github.com/KoyomiKun/grafana-cli/pkg/client"
	"github.com/KoyomiKun/grafana-cli/pkg/grafana"
	"github.com/KoyomiKun/grafana-cli/utils/log"
)

type FolderManager struct {
	client *client.Client
}

func NewFolderManager(client *client.Client) *FolderManager {
	return &FolderManager{
		client: client,
	}
}

func (fm *FolderManager) List() (grafana.Folders, error) {
	content, err := fm.client.GetApi("/api/folders", map[string]string{}, map[string]string{})
	if err != nil {
		log.Errorf("fail getting folders: %v", err)
		return nil, err
	}

	fds, err := grafana.NewFolders(content)
	if err != nil {
		log.Errorf("fail gettting folders: %v", err)
		return nil, err
	}

	return fds, nil
}

func (fm *FolderManager) CreateFolder(fd *grafana.Folder) error {

	content, err := json.Marshal(fd)
	if err != nil {
		log.Errorf("fail create folder %v : %v", fd, err)
		return err
	}

	resp, err := fm.client.PostApi("/api/folders", map[string]string{}, bytes.NewBuffer(content))
	if err != nil {
		log.Errorf("post create folder fails %s: %v", resp, err)
		return err
	}

	return nil
}
