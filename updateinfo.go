package yum

import (
	"encoding/xml"
	"io"
	"os"
)

type UpdateInfoDatabase struct {
	updateInfoPath string
	reader         io.Reader
}

type UpdateInfo struct {
	XMLName    xml.Name       `xml:"updates"`
	UpdateInfo []*UpdateEntry `xml:"update"`
}

func OpenUpdateInfoDB(path string) (*UpdateInfoDatabase, error) {
	updateInfoFile, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	return &UpdateInfoDatabase{
		updateInfoPath: path,
		reader:         updateInfoFile,
	}, nil
}

func (u *UpdateInfoDatabase) Updates() ([]*UpdateEntry, error) {
	decoder := xml.NewDecoder(u.reader)
	updateInfo := &UpdateInfo{}
	if err := decoder.Decode(updateInfo); err != nil {
		return nil, err
	}
	return updateInfo.UpdateInfo, nil
}
