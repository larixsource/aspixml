package aspixml

import (
	"encoding/xml"
)

type MessageStatus struct {
	XMLName xml.Name `xml:"MessageStatus"`
	Code    int      `xml:"code,attr"`
	Time    ASPITime `xml:"time,attr"`
	Value   string   `xml:",chardata"`
}

type ForwardMessage struct {
	XMLName       xml.Name      `xml:"ForwardMessage"`
	NID           uint64        `xml:"nid,attr"`
	FID           uint64        `xml:"fid,attr"`
	Limit         int           `xml:"limit,attr,omitempty"`
	MessageStatus MessageStatus `xml:"MessageStatus"`
}

type AdC struct {
	XMLName xml.Name `xml:"AdC"`
	Ocean   string   `xml:"ocean,attr"`
	Value   string   `xml:",chardata"`
}

type MessageData struct {
	XMLName xml.Name `xml:"MessageData"`
	Value   string   `xml:",chardata"`
}

type Flags struct {
	XMLName xml.Name `xml:"Flags"`
	LES     bool     `xml:"les,attr,omitempty"`
	App     bool     `xml:"app,attr,omitempty"`
}

type QoS struct {
	XMLName xml.Name `xml:"QoS"`
	Channel uint     `xml:"channel,attr"`
	Frame   uint     `xml:"frame,attr"`
	RSN     uint     `xml:"rsn,attr"`
	Level   float64  `xml:"level,attr"`
	Offset  int      `xml:"offset,attr"`
	Doppler int      `xml:"doppler,attr"`
	Errors  uint     `xml:"errors,attr"`
	SNR     float64  `xml:"snr,attr"`
}

type ReturnMessage struct {
	XMLName       xml.Name      `xml:"ReturnMessage"`
	RID           uint64        `xml:"rid,attr"`
	AdC           *AdC          `xml:"AdC,omitempty"`
	MessageData   *MessageData  `xml:"MessageData,omitempty"`
	MessageStatus MessageStatus `xml:"MessageStatus"`
	Flags         *Flags        `xml:"Flags,omitempty"`
	QoS           *QoS          `xml:"QoS,omitempty"`
}

type MessageDelivery struct {
	XMLName         xml.Name         `xml:"MessageDelivery"`
	ForwardMessages []ForwardMessage `xml:"ForwardMessage"`
	ReturnMessages  []ReturnMessage  `xml:"ReturnMessage"`
}
