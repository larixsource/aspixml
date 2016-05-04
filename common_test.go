package aspixml

import (
	"encoding/xml"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

var testFault1 = Fault{
	Code: 203,
	Msg:  "access denied - authentication failed",
}

var testFault1XML = `<Fault>
  <FaultCode>203</FaultCode>
  <FaultString>access denied - authentication failed</FaultString>
</Fault>`

func TestMarshalFault(t *testing.T) {
	b, err := xml.MarshalIndent(&testMessageDelivery1, "", "  ")
	require.Nil(t, err)
	expected := []byte(testMessageDelivery1XML)
	assert.Equal(t, expected, b)
}

func TestUnmarshalFault(t *testing.T) {
	var f Fault
	err := xml.Unmarshal([]byte(testFault1XML), &f)
	require.Nil(t, err)
	assert.Equal(t, testFault1.Code, f.Code)
	assert.Equal(t, testFault1.Msg, f.Msg)
	assert.Empty(t, f.Detail)
}

var testFault2XML = `<Fault>
  <FaultCode>203</FaultCode>
  <FaultString>access denied - authentication failed</FaultString>
  <Detail>chan!</Detail>
</Fault>`

func TestUnmarshalFaultWithDetail(t *testing.T) {
	var f Fault
	err := xml.Unmarshal([]byte(testFault2XML), &f)
	require.Nil(t, err)
	assert.Equal(t, 203, f.Code)
	assert.Equal(t, "access denied - authentication failed", f.Msg)
	assert.Equal(t, "chan!", f.Detail)
}
