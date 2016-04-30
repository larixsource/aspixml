package aspixml

import "encoding/xml"

type ForwardMessageReq struct {
	XMLName xml.Name `xml:"ForwardMessage"`
	NID     int      `xml:"nid,attr"`
	Limit   int      `xml:"limit,attr,omitempty"`
}

type ReturnMessageReq struct {
	XMLName xml.Name `xml:"ReturnMessage"`
	RID     int      `xml:"rid,attr"`
	Limit   int      `xml:"limit,attr,omitempty"`
}

type RequestDelivery struct {
	XMLName        xml.Name `xml:"RequestDelivery"`
	QOS            bool     `xml:"qos,attr,omitempty"`
	Authentication Authentication
	ForwardMessage ForwardMessageReq
	ReturnMessage  ReturnMessageReq
}
