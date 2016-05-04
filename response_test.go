package aspixml

import (
	"encoding/xml"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestMarshalResponseFault(t *testing.T) {
	r := Response{
		Type:  FaultResp,
		Fault: &testFault1,
	}
	b, err := xml.MarshalIndent(&r, "", "  ")
	require.Nil(t, err)
	assert.Equal(t, []byte(testFault1XML), b)
}

func TestUnmarshalResponseFault(t *testing.T) {
	var r Response
	err := xml.Unmarshal([]byte(testFault1XML), &r)
	require.Nil(t, err)
	assert.Equal(t, FaultResp, r.Type)
	assert.Nil(t, r.MessageDelivery)
	assert.Equal(t, testFault1.Code, r.Fault.Code)
	assert.Equal(t, testFault1.Msg, r.Fault.Msg)
	assert.Empty(t, r.Fault.Detail)
}

func TestMarshalResponseMessageDelivery(t *testing.T) {
	r := Response{
		Type:            MessageDeliveryResp,
		MessageDelivery: &testMessageDelivery1,
	}
	b, err := xml.MarshalIndent(&r, "", "  ")
	require.Nil(t, err)
	assert.Equal(t, []byte(testMessageDelivery1XML), b)
}

func TestUnmarshalResponseMessageDelivery(t *testing.T) {
	var r Response
	err := xml.Unmarshal([]byte(testMessageDelivery1XML), &r)
	require.Nil(t, err)
	assert.Equal(t, MessageDeliveryResp, r.Type)
	assert.Nil(t, r.Fault)
	assert.Len(t, r.MessageDelivery.Messages, 2)
	// TODO
}
