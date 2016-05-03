package aspixml

import (
	"encoding/xml"
	"fmt"
)

type MessageStatus struct {
	XMLName xml.Name `xml:"MessageStatus"`
	Code    int      `xml:"code,attr"`
	Time    ASPITime `xml:"time,attr"`
	Value   string   `xml:",chardata"`
}

type ForwardMessage struct {
	XMLName       xml.Name `xml:"ForwardMessage"`
	NID           uint64   `xml:"nid,attr"`
	FID           uint64   `xml:"fid,attr"`
	Limit         int      `xml:"limit,attr,omitempty"`
	MessageStatus MessageStatus
}

type AdC struct {
	Ocean string `xml:"ocean,attr"`
	Value string `xml:",chardata"`
}

type Flags struct {
	LES int `xml:"les,attr"`
	App int `xml:"app,attr"`
}

type QoS struct {
	Channel uint    `xml:"channel,attr"`
	Frame   uint    `xml:"frame,attr"`
	RSN     uint    `xml:"rsn,attr"`
	Level   float64 `xml:"level,attr"`
	Offset  int     `xml:"offset,attr"`
	Doppler int     `xml:"doppler,attr"`
	Errors  uint    `xml:"errors,attr"`
	SNR     float64 `xml:"snr,attr"`
}

type ReturnMessage struct {
	XMLName       xml.Name `xml:"ReturnMessage"`
	RID           uint64   `xml:"rid,attr"`
	AdC           *AdC     `xml:"AdC,omitempty"`
	MessageData   string   `xml:"MessageData,omitempty"`
	MessageStatus MessageStatus
	Flags         *Flags `xml:"Flags,omitempty"`
	QoS           *QoS   `xml:"QoS,omitempty"`
}

type FwdOrReturnMsg struct {
	Return  *ReturnMessage
	Forward *ForwardMessage
}

func (fr *FwdOrReturnMsg) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	switch {
	case fr.Forward != nil:
		return e.Encode(fr.Forward)
	case fr.Return != nil:
		return e.Encode(fr.Return)
	default:
		return fmt.Errorf("Invalid FwdOrReturnMsg: %+v", fr)
	}
}

func (fr *FwdOrReturnMsg) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	switch start.Name.Local {
	case "ForwardMessage":
		fr.Forward = &ForwardMessage{}
		return d.DecodeElement(fr.Forward, &start)
	case "ReturnMessage":
		fr.Return = &ReturnMessage{}
		return d.DecodeElement(fr.Return, &start)
	default:
		return fmt.Errorf("Invalid element %s", start.Name)
	}
}

type MessageDelivery struct {
	XMLName  xml.Name         `xml:"MessageDelivery"`
	Messages []FwdOrReturnMsg `xml:",any"`
}
