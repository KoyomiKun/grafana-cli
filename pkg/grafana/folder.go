package grafana

import (
	"encoding/json"

	"github.com/KoyomiKun/grafana-cli/utils/log"
)

type Folders []*Folder

type Folder struct {
	Id    int    `json:"id,omitempty"`
	Uid   string `json:"uid,omitempty"`
	Title string `json:"title,omitempty"`
}

func NewFolders(content []byte) (Folders, error) {

	fds := Folders{}
	if err := json.Unmarshal(content, &fds); err != nil {
		log.Errorf("fail creating folders: %v", err)
		return nil, err
	}
	return fds, nil

}
