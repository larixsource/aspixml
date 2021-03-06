package aspixml

import (
	"encoding/xml"
	"time"
)

type Authentication struct {
	XMLName  xml.Name `xml:"Authentication"`
	ID       string   `xml:"id,attr"`
	Password string   `xml:",chardata"`
}

type ASPITime struct {
	time.Time
}

func (at *ASPITime) UnmarshalXMLAttr(attr xml.Attr) (err error) {
	at.Time, err = time.Parse("2006-01-02 15:04:05", attr.Value)
	return
}

func (at *ASPITime) MarshalXMLAttr(name xml.Name) (xml.Attr, error) {
	return xml.Attr{
		Name:  name,
		Value: at.Time.Format("2006-01-02 15:04:05"),
	}, nil
}

type Fault struct {
	Code   int    `xml:"FaultCode"`
	Msg    string `xml:"FaultString"`
	Detail string `xml:"Detail,omitempty"`
}
