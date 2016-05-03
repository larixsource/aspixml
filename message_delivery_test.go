package aspixml

import (
	"encoding/xml"
	"os"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestUnmarshalMessageDelivery(t *testing.T) {
	f, err := os.Open("message_delivery.xml")
	require.Nil(t, err)

	var msgDelivery MessageDelivery
	err = xml.NewDecoder(f).Decode(&msgDelivery)
	require.Nil(t, err)

	// check
	require.Len(t, msgDelivery.Messages, 5)

	// fwd 0
	fwd := msgDelivery.Messages[0].Forward
	assert.EqualValues(t, 78212893, fwd.NID)
	assert.EqualValues(t, 81502503, fwd.FID)
	assert.Equal(t, 105, fwd.MessageStatus.Code)
	assert.Equal(t, ASPITime{time.Date(2015, 1, 23, 18, 25, 0, 0, time.UTC)}, fwd.MessageStatus.Time)
	assert.Equal(t, "transmission to terminal complete", fwd.MessageStatus.Value)

	// return 0
	ret := msgDelivery.Messages[1].Return
	assert.EqualValues(t, 1915861563, ret.RID)
	assert.EqualValues(t, "AORWGL", ret.AdC.Ocean)
	assert.EqualValues(t, "DST009F2BD69", ret.AdC.Value)
	assert.EqualValues(t, "902190000000000000000", ret.MessageData)
	assert.Equal(t, 100, ret.MessageStatus.Code)
	assert.Equal(t, ASPITime{time.Date(2015, 5, 4, 13, 18, 36, 0, time.UTC)}, ret.MessageStatus.Time)
	assert.Equal(t, "status ok", ret.MessageStatus.Value)
	assert.Zero(t, ret.Flags.LES)
	assert.Zero(t, ret.Flags.App)

	// return 1
	ret = msgDelivery.Messages[2].Return
	assert.EqualValues(t, 1915861943, ret.RID)
	assert.EqualValues(t, "AORWGL", ret.AdC.Ocean)
	assert.EqualValues(t, "DST009F2BD69", ret.AdC.Value)
	assert.EqualValues(t, "901199d52f9f0e6003833", ret.MessageData)
	assert.Equal(t, 100, ret.MessageStatus.Code)
	assert.Equal(t, ASPITime{time.Date(2015, 5, 4, 13, 19, 15, 0, time.UTC)}, ret.MessageStatus.Time)
	assert.Equal(t, "status ok", ret.MessageStatus.Value)
	assert.Zero(t, ret.Flags.LES)
	assert.Zero(t, ret.Flags.App)

	// fwd 1
	fwd = msgDelivery.Messages[3].Forward
	assert.EqualValues(t, 78213327, fwd.NID)
	assert.EqualValues(t, 81502933, fwd.FID)
	assert.Equal(t, 105, fwd.MessageStatus.Code)
	assert.Equal(t, ASPITime{time.Date(2015, 1, 23, 18, 32, 30, 0, time.UTC)}, fwd.MessageStatus.Time)
	assert.Equal(t, "transmission to terminal complete", fwd.MessageStatus.Value)

	// return 2
	ret = msgDelivery.Messages[4].Return
	assert.EqualValues(t, 1915918463, ret.RID)
	assert.EqualValues(t, "AORWGL", ret.AdC.Ocean)
	assert.EqualValues(t, "DST009F2BD69", ret.AdC.Value)
	assert.EqualValues(t, "902190000000000000000", ret.MessageData)
	assert.Equal(t, 100, ret.MessageStatus.Code)
	assert.Equal(t, ASPITime{time.Date(2015, 5, 4, 15, 8, 16, 0, time.UTC)}, ret.MessageStatus.Time)
	assert.Equal(t, "status ok", ret.MessageStatus.Value)
	assert.Zero(t, ret.Flags.LES)
	assert.Zero(t, ret.Flags.App)
}

func TestMarshalMessageDelivery(t *testing.T) {
	md := MessageDelivery{
		Messages: []FwdOrReturnMsg{
			FwdOrReturnMsg{
				Forward: &ForwardMessage{
					NID: 78213327,
					FID: 81502933,
					MessageStatus: MessageStatus{
						Code:  105,
						Time:  ASPITime{time.Date(2015, 1, 23, 18, 25, 0, 0, time.UTC)},
						Value: "transmission to terminal complete",
					},
				},
			},
			FwdOrReturnMsg{
				Return: &ReturnMessage{
					RID: 1915861563,
					AdC: &AdC{
						Ocean: "AORWGL",
						Value: "DST009F2BD69",
					},
					MessageData: "902190000000000000000",
					MessageStatus: MessageStatus{
						Code:  100,
						Time:  ASPITime{time.Date(2015, 5, 4, 13, 18, 36, 0, time.UTC)},
						Value: "status ok",
					},
					Flags: &Flags{
						LES: 0,
						App: 0,
					},
				},
			},
		},
	}
	expected := `<MessageDelivery>
  <ForwardMessage nid="78213327" fid="81502933">
    <MessageStatus code="105" time="2015-01-23 18:25:00">transmission to terminal complete</MessageStatus>
  </ForwardMessage>
  <ReturnMessage rid="1915861563">
    <AdC ocean="AORWGL">DST009F2BD69</AdC>
    <MessageData>902190000000000000000</MessageData>
    <MessageStatus code="100" time="2015-05-04 13:18:36">status ok</MessageStatus>
    <Flags les="0" app="0"></Flags>
  </ReturnMessage>
</MessageDelivery>`
	msg, err := xml.MarshalIndent(&md, "", "  ")
	require.Nil(t, err)
	assert.Equal(t, []byte(expected), msg)
}
