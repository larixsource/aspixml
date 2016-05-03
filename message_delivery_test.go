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
	require.Len(t, msgDelivery.ForwardMessages, 2)
	require.Len(t, msgDelivery.ReturnMessages, 3)

	// fwd 0
	fwd := msgDelivery.ForwardMessages[0]
	assert.EqualValues(t, 78212893, fwd.NID)
	assert.EqualValues(t, 81502503, fwd.FID)
	assert.Equal(t, 105, fwd.MessageStatus.Code)
	assert.Equal(t, ASPITime{time.Date(2015, 1, 23, 18, 25, 0, 0, time.UTC)}, fwd.MessageStatus.Time)
	assert.Equal(t, "transmission to terminal complete", fwd.MessageStatus.Value)

	// fwd 1
	fwd = msgDelivery.ForwardMessages[1]
	assert.EqualValues(t, 78213327, fwd.NID)
	assert.EqualValues(t, 81502933, fwd.FID)
	assert.Equal(t, 105, fwd.MessageStatus.Code)
	assert.Equal(t, ASPITime{time.Date(2015, 1, 23, 18, 32, 30, 0, time.UTC)}, fwd.MessageStatus.Time)
	assert.Equal(t, "transmission to terminal complete", fwd.MessageStatus.Value)

	// return 0
	ret := msgDelivery.ReturnMessages[0]
	assert.EqualValues(t, 1915861563, ret.RID)
	assert.EqualValues(t, "AORWGL", ret.AdC.Ocean)
	assert.EqualValues(t, "DST009F2BD69", ret.AdC.Value)
	assert.EqualValues(t, "902190000000000000000", ret.MessageData)
	assert.Equal(t, 100, ret.MessageStatus.Code)
	assert.Equal(t, ASPITime{time.Date(2015, 5, 4, 13, 18, 36, 0, time.UTC)}, ret.MessageStatus.Time)
	assert.Equal(t, "status ok", ret.MessageStatus.Value)
	assert.False(t, ret.Flags.LES)
	assert.False(t, ret.Flags.App)

	// return 1
	ret = msgDelivery.ReturnMessages[1]
	assert.EqualValues(t, 1915861943, ret.RID)
	assert.EqualValues(t, "AORWGL", ret.AdC.Ocean)
	assert.EqualValues(t, "DST009F2BD69", ret.AdC.Value)
	assert.EqualValues(t, "901199d52f9f0e6003833", ret.MessageData)
	assert.Equal(t, 100, ret.MessageStatus.Code)
	assert.Equal(t, ASPITime{time.Date(2015, 5, 4, 13, 19, 15, 0, time.UTC)}, ret.MessageStatus.Time)
	assert.Equal(t, "status ok", ret.MessageStatus.Value)
	assert.False(t, ret.Flags.LES)
	assert.False(t, ret.Flags.App)

	// return 2
	ret = msgDelivery.ReturnMessages[2]
	assert.EqualValues(t, 1915918463, ret.RID)
	assert.EqualValues(t, "AORWGL", ret.AdC.Ocean)
	assert.EqualValues(t, "DST009F2BD69", ret.AdC.Value)
	assert.EqualValues(t, "902190000000000000000", ret.MessageData)
	assert.Equal(t, 100, ret.MessageStatus.Code)
	assert.Equal(t, ASPITime{time.Date(2015, 5, 4, 15, 8, 16, 0, time.UTC)}, ret.MessageStatus.Time)
	assert.Equal(t, "status ok", ret.MessageStatus.Value)
	assert.False(t, ret.Flags.LES)
	assert.False(t, ret.Flags.App)
}
