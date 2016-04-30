package aspixml

import (
	"encoding/xml"
	"time"
)

type Authentication struct {
	XMLName  xml.Name `xml:"Authentication"`
	ID       int      `xml:"id,attr"`
	Password string   `xml:",chardata"`
}

type ASPITime struct {
	time.Time
}

func (at *ASPITime) UnmarshalXMLAttr(attr xml.Attr) (err error) {
	at.Time, err = time.Parse("2006-01-02 15:04:05", attr.Value)
	return
}
