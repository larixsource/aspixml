package aspixml

import (
	"encoding/xml"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestMarshalFault(t *testing.T) {
	f := Fault{
		Code: 203,
		Msg:  "access denied - authentication failed",
	}
	b, err := xml.Marshal(&f)
	require.Nil(t, err)
	expected := []byte(`<Fault><FaultCode>203</FaultCode><FaultString>access denied - authentication failed</FaultString></Fault>`)
	assert.Equal(t, expected, b)
}

func TestUnmarshalFault(t *testing.T) {
	msg := []byte(`<Fault><FaultCode>203</FaultCode><FaultString>access denied - authentication failed</FaultString></Fault>`)
	var f Fault
	err := xml.Unmarshal(msg, &f)
	require.Nil(t, err)
	assert.Equal(t, 203, f.Code)
	assert.Equal(t, "access denied - authentication failed", f.Msg)
	assert.Empty(t, f.Detail)
}

func TestUnmarshalFaultWithDetail(t *testing.T) {
	msg := []byte(`<Fault><FaultCode>203</FaultCode><FaultString>access denied - authentication failed</FaultString><Detail>chan!</Detail></Fault>`)
	var f Fault
	err := xml.Unmarshal(msg, &f)
	require.Nil(t, err)
	assert.Equal(t, 203, f.Code)
	assert.Equal(t, "access denied - authentication failed", f.Msg)
	assert.Equal(t, "chan!", f.Detail)
}
