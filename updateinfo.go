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
	XMLName       xml.Name        `xml:"updates"`
	UpdateNotices []*UpdateNotice `xml:"update"`
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

func (u *UpdateInfoDatabase) UpdateNotices() ([]*UpdateNotice, error) {
	decoder := xml.NewDecoder(u.reader)
	updateInfo := &UpdateInfo{}
	if err := decoder.Decode(updateInfo); err != nil {
		return nil, err
	}
	return updateInfo.UpdateNotices, nil
}
