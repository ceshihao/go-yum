package yum

import "time"

type UpdateNotice struct {
	Author      string               `xml:"author,attr"`
	From        string               `xml:"from,attr"`
	Status      string               `xml:"status,attr"`
	Type        string               `xml:"type,attr"`
	Version     string               `xml:"version,attr"`
	ID          string               `xml:"id"`
	Title       string               `xml:"title"`
	Issued      *UpdateInfoTime      `xml:"issued"`
	Updated     *UpdateInfoTime      `xml:"updated"`
	Rights      string               `xml:"rights"`
	Release     string               `xml:"release"`
	Severity    string               `xml:"severity"`
	Summary     string               `xml:"summary"`
	Description string               `xml:"description"`
	References  *ReferenceCollection `xml:"references"`
	PkgList     *PkgList             `xml:"pkglist"`
}

type UpdateInfoTime struct {
	Date string `xml:"date,attr"`
	Time int64  `xml:"time,attr"`
}

type ReferenceCollection struct {
	References []*Reference `xml:"reference"`
}

type Reference struct {
	Href  string `xml:"href,attr"`
	ID    string `xml:"id,attr"`
	Type  string `xml:"type,attr"`
	Title string `xml:"title,attr"`
}

type PkgList struct {
	Collection *Collection `xml:"collection"`
	Name       string      `xml:"name"`
	Packages   []*Package  `xml:"package"`
}

type Collection struct {
	Short string `xml:"short,attr"`
}

type Package struct {
	Name     string `xml:"name,attr"`
	Version  string `xml:"version,attr"`
	Release  string `xml:"release,attr"`
	Epoch    string `xml:"epoch,attr"`
	Arch     string `xml:"arch,attr"`
	Src      string `xml:"src,attr"`
	FileName string `xml:"filename"`
}

func (u *UpdateNotice) UpdatedTime() time.Time {
	return time.Unix(u.Updated.Time, 0)
}
