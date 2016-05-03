package aspixml

import (
	"bytes"
	"encoding/xml"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestMarshalRequestDelivery(t *testing.T) {
	msg := RequestDelivery{
		Authentication: Authentication{
			ID:       "666",
			Password: "muahahaha",
		},
		ForwardMessage: ForwardMessageReq{
			NID:   6,
			Limit: 100,
		},
		ReturnMessage: ReturnMessageReq{
			RID: 21,
		},
	}
	var buf bytes.Buffer
	err := xml.NewEncoder(&buf).Encode(&msg)
	require.Nil(t, err)
	expectedXML := `<RequestDelivery><Authentication id="666">muahahaha</Authentication><ForwardMessage nid="6" limit="100"></ForwardMessage><ReturnMessage rid="21"></ReturnMessage></RequestDelivery>`
	assert.Equal(t, expectedXML, buf.String())
}
