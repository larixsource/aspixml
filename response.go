package aspixml

import (
	"encoding/xml"
	"fmt"
)

type ResponseType int

const (
	UnknownResp ResponseType = iota
	FaultResp
	MessageDeliveryResp
)

// Response provides an easy way to receive a response from MHS, recognizing its type, and if there was an error. The
// field Type is set in the unmarshalling process according to the message type: i.e. if it's a Fault, the Fault field
// will be non-nil and the remaining fields will be nil, if it's a MessageDelivery, the MessageDelivery field will be
// non-nil and the rest will be nil, and so on.
type Response struct {
	Type            ResponseType
	Fault           *Fault
	MessageDelivery *MessageDelivery
}

func (r *Response) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	switch r.Type {
	case FaultResp:
		return e.Encode(r.Fault)
	case MessageDeliveryResp:
		return e.Encode(r.MessageDelivery)
	default:
		return fmt.Errorf("Invalid Response: %+v", r)
	}
}

func (r *Response) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	switch start.Name.Local {
	case "Fault":
		r.Type = FaultResp
		r.Fault = &Fault{}
		return d.DecodeElement(r.Fault, &start)
	case "MessageDelivery":
		r.Type = MessageDeliveryResp
		r.MessageDelivery = &MessageDelivery{}
		return d.DecodeElement(r.MessageDelivery, &start)
	default:
		r.Type = UnknownResp
		return fmt.Errorf("Unknown element %s", start.Name)
	}
}
