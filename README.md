# aspixml
ASPI-XML/3.2 types

This package provides types to marshal/unmarshal messages of the Honeywell ASPI-XML/3.2 protocol.

Not all messages are supported yet, check the table below:

Msg | Kind | Supported
--- | --- | ---
SubmitMessage | Request | No
StatusReport | Response | No
RequestStatus | Request | No
CancelMessage | Request | No
RequestDelivery | Request | Yes
MessageDelivery | Response | Yes
GetMessageCopy | Request | No
MessageCopy | Response | No
GetLastIdentifier | Request | No
LastIdentifier | Response | No
Fault | Response | No
