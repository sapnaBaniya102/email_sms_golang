package constant

import "fmt"

type CMDStatus uint32

type CMDId uint32

const (
	// SMPP Protocol Version
	SMPP_VERSION = 0x34

	// Max PDU size to minimize some attack vectors
	MAX_PDU_SIZE = 4096 // 4KB

	// Sequence number start/end
	SEQUENCE_NUM_START = 0x00000001
	SEQUENCE_NUM_END   = 0x7FFFFFFF
)

const (
	// ESME Error Constants
	ESME_ROK              CMDStatus = 0x00000000 // OK!
	ESME_RINVMSGLEN       CMDStatus = 0x00000001 // Message Length is invalid
	ESME_RINVCMDLEN       CMDStatus = 0x00000002 // Command Length is invalid
	ESME_RINVCMDID        CMDStatus = 0x00000003 // Invalid Command ID
	ESME_RINVBNDSTS       CMDStatus = 0x00000004 // Incorrect BIND Status for given command
	ESME_RALYBND          CMDStatus = 0x00000005 // ESME Already in Bound State
	ESME_RINVPRTFLG       CMDStatus = 0x00000006 // Invalid Priority Flag
	ESME_RINVREGDLVFLG    CMDStatus = 0x00000007 // Invalid Registered Delivery Flag
	ESME_RSYSERR          CMDStatus = 0x00000008 // System Error
	ESME_RINVSRCADR       CMDStatus = 0x0000000A // Invalid Source Address
	ESME_RINVDSTADR       CMDStatus = 0x0000000B // Invalid Dest Addr
	ESME_RINVMSGID        CMDStatus = 0x0000000C // Message ID is invalid
	ESME_RBINDFAIL        CMDStatus = 0x0000000D // Bind Failed
	ESME_RINVPASWD        CMDStatus = 0x0000000E // Invalid Password
	ESME_RINVSYSID        CMDStatus = 0x0000000F // Invalid System ID
	ESME_RCANCELFAIL      CMDStatus = 0x00000011 // Cancel SM Failed
	ESME_RREPLACEFAIL     CMDStatus = 0x00000013 // Replace SM Failed
	ESME_RMSGQFUL         CMDStatus = 0x00000014 // Message Queue Full
	ESME_RINVSERTYP       CMDStatus = 0x00000015 // Invalid Service Type
	ESME_RINVNUMDESTS     CMDStatus = 0x00000033 // Invalid number of destinations
	ESME_RINVDLNAME       CMDStatus = 0x00000034 // Invalid Distribution List name
	ESME_RINVDESTFLAG     CMDStatus = 0x00000040 // Destination flag is invalid
	ESME_RINVSUBREP       CMDStatus = 0x00000042 // Invalid 'submit with replace' request
	ESME_RINVESMCLASS     CMDStatus = 0x00000043 // Invalid esm_class field data
	ESME_RCNTSUBDL        CMDStatus = 0x00000044 // Cannot Submit to Distribution List
	ESME_RSUBMITFAIL      CMDStatus = 0x00000045 // submit_sm or submit_multi failed
	ESME_RINVSRCTON       CMDStatus = 0x00000048 // Invalid Source address TON
	ESME_RINVSRCNPI       CMDStatus = 0x00000049 // Invalid Source address NPI
	ESME_RINVDSTTON       CMDStatus = 0x00000050 // Invalid Destination address TON
	ESME_RINVDSTNPI       CMDStatus = 0x00000051 // Invalid Destination address NPI
	ESME_RINVSYSTYP       CMDStatus = 0x00000053 // Invalid system_type field
	ESME_RINVREPFLAG      CMDStatus = 0x00000054 // Invalid replace_if_present flag
	ESME_RINVNUMMSGS      CMDStatus = 0x00000055 // Invalid number of messages
	ESME_RTHROTTLED       CMDStatus = 0x00000058 // Throttling error
	ESME_RINVSCHED        CMDStatus = 0x00000061 // Invalid Scheduled Delivery Time
	ESME_RINVEXPIRY       CMDStatus = 0x00000062 // Invalid message validity period (Expiry time)
	ESME_RINVDFTMSGID     CMDStatus = 0x00000063 // Predefined Message Invalid or Not Found
	ESME_RX_T_APPN        CMDStatus = 0x00000064 // ESME Receiver Temporary App Error Code
	ESME_RX_P_APPN        CMDStatus = 0x00000065 // ESME Receiver Permanent App Error Code
	ESME_RX_R_APPN        CMDStatus = 0x00000066 // ESME Receiver Reject Message Error Code
	ESME_RQUERYFAIL       CMDStatus = 0x00000067 // Query_sm request failed
	ESME_RINVOPTPARSTREAM CMDStatus = 0x000000C0 // Error in the optional part of the PDU Body
	ESME_ROPTPARNOTALLWD  CMDStatus = 0x000000C1 // Optional Parameter not allowed
	ESME_RINVPARLEN       CMDStatus = 0x000000C2 // Invalid Parameter Length
	ESME_RMISSINGOPTPARAM CMDStatus = 0x000000C3 // Expected Optional Parameter missing
	ESME_RINVOPTPARAMVAL  CMDStatus = 0x000000C4 // Invalid Optional Parameter Value
	ESME_RDELIVERYFAILURE CMDStatus = 0x000000FE // Delivery Failure (used for data_sm_resp)
	ESME_RUNKNOWNERR      CMDStatus = 0x000000FF // Unknown Error
)

const (
	// PDU Types
	GENERIC_NACK          CMDId = 0x80000000
	BIND_RECEIVER         CMDId = 0x00000001
	BIND_RECEIVER_RESP    CMDId = 0x80000001
	BIND_TRANSMITTER      CMDId = 0x00000002
	BIND_TRANSMITTER_RESP CMDId = 0x80000002
	QUERY_SM              CMDId = 0x00000003
	QUERY_SM_RESP         CMDId = 0x80000003
	SUBMIT_SM             CMDId = 0x00000004
	SUBMIT_SM_RESP        CMDId = 0x80000004
	DELIVER_SM            CMDId = 0x00000005
	DELIVER_SM_RESP       CMDId = 0x80000005
	UNBIND                CMDId = 0x00000006
	UNBIND_RESP           CMDId = 0x80000006
	REPLACE_SM            CMDId = 0x00000007
	REPLACE_SM_RESP       CMDId = 0x80000007
	CANCEL_SM             CMDId = 0x00000008
	CANCEL_SM_RESP        CMDId = 0x80000008
	BIND_TRANSCEIVER      CMDId = 0x00000009
	BIND_TRANSCEIVER_RESP CMDId = 0x80000009
	OUTBIND               CMDId = 0x0000000B
	ENQUIRE_LINK          CMDId = 0x00000015
	ENQUIRE_LINK_RESP     CMDId = 0x80000015
	SUBMIT_MULTI          CMDId = 0x00000021
	SUBMIT_MULTI_RESP     CMDId = 0x80000021
	ALERT_NOTIFICATION    CMDId = 0x00000102
	DATA_SM               CMDId = 0x00000103
	DATA_SM_RESP          CMDId = 0x80000103
)

const (
	// FIELDS
	SYSTEM_ID               = "system_id"
	PASSWORD                = "password"
	SYSTEM_TYPE             = "system_type"
	INTERFACE_VERSION       = "interface_version"
	ADDR_TON                = "addr_ton"
	ADDR_NPI                = "addr_npi"
	ADDRESS_RANGE           = "address_range"
	SERVICE_TYPE            = "service_type"
	SOURCE_ADDR_TON         = "source_addr_ton"
	SOURCE_ADDR_NPI         = "source_addr_npi"
	SOURCE_ADDR             = "source_addr"
	DEST_ADDR_TON           = "dest_addr_ton"
	DEST_ADDR_NPI           = "dest_addr_npi"
	DESTINATION_ADDR        = "destination_addr"
	ESM_CLASS               = "esm_class"
	PROTOCOL_ID             = "protocol_id"
	PRIORITY_FLAG           = "priority_flag"
	SCHEDULE_DELIVERY_TIME  = "schedule_delivery_time"
	VALIDITY_PERIOD         = "validity_period"
	REGISTERED_DELIVERY     = "registered_delivery"
	REPLACE_IF_PRESENT_FLAG = "replace_if_present_flag"
	DATA_CODING             = "data_coding"
	SM_DEFAULT_MSG_ID       = "sm_default_msg_id"
	SM_LENGTH               = "sm_length"
	SHORT_MESSAGE           = "short_message"
	MESSAGE_ID              = "message_id"
	FINAL_DATE              = "final_date"
	MESSAGE_STATE           = "message_state"
	ERROR_CODE              = "error_code"
)

const (
	// Optional Field Tags
	DEST_ADDR_SUBUNIT           = 0x0005
	DEST_NETWORK_TYPE           = 0x0006
	DEST_BEARER_TYPE            = 0x0007
	DEST_TELEMATICS_ID          = 0x0008
	SOURCE_ADDR_SUBUNIT         = 0x000D
	SOURCE_NETWORK_TYPE         = 0x000E
	SOURCE_BEARER_TYPE          = 0x000F
	SOURCE_TELEMATICS_ID        = 0x0010
	QOS_TIME_TO_LIVE            = 0x0017
	PAYLOAD_TYPE                = 0x0019
	ADDITIONAL_STATUS_INFO_TEXT = 0x001D
	RECEIPTED_MESSAGE_ID        = 0x001E
	MS_MSG_WAIT_FACILITIES      = 0x0030
	PRIVACY_INDICATOR           = 0x0201
	SOURCE_SUBADDRESS           = 0x0202
	DEST_SUBADDRESS             = 0x0203
	USER_MESSAGE_REFERENCE      = 0x0204
	USER_RESPONSE_CODE          = 0x0205
	SOURCE_PORT                 = 0x020A
	DESTINATION_PORT            = 0x020B
	SAR_MSG_REF_NUM             = 0x020C
	LANGUAGE_INDICATOR          = 0x020D
	SAR_TOTAL_SEGMENTS          = 0x020E
	SAR_SEGMENT_SEQNUM          = 0x020F
	SC_INTERFACE_VERSION        = 0x0210
	CALLBACK_NUM_PRES_IND       = 0x0302
	CALLBACK_NUM_ATAG           = 0x0303
	NUMBER_OF_MESSAGES          = 0x0304
	CALLBACK_NUM                = 0x0381
	DPF_RESULT                  = 0x0420
	SET_DPF                     = 0x0421
	MS_AVAILABILITY_STATUS      = 0x0422
	NETWORK_ERROR_CODE          = 0x0423
	MESSAGE_PAYLOAD             = 0x0424
	DELIVERY_FAILURE_REASON     = 0x0425
	MORE_MESSAGES_TO_SEND       = 0x0426
	DR_MESSAGE_STATE            = 0x0427
	USSD_SERVICE_OP             = 0x0501
	DISPLAY_TIME                = 0x1201
	SMS_SIGNAL                  = 0x1203
	MS_VALIDITY                 = 0x1204
	ALERT_ON_MESSAGE_DELIVERY   = 0x130C
	ITS_REPLY_TYPE              = 0x1380
	ITS_SESSION_INFO            = 0x1383
)

const (
	// Encoding Types
	ENCODING_DEFAULT   = 0x00 // SMSC Default
	ENCODING_IA5       = 0x01 // IA5 (CCITT T.50)/ASCII (ANSI X3.4)
	ENCODING_BINARY    = 0x02 // Octet unspecified (8-bit binary)
	ENCODING_ISO88591  = 0x03 // Latin 1 (ISO-8859-1)
	ENCODING_BINARY2   = 0x04 // Octet unspecified (8-bit binary)
	ENCODING_JIS       = 0x05 // JIS (X 0208-1990)
	ENCODING_ISO88595  = 0x06 // Cyrillic (ISO-8859-5)
	ENCODING_ISO88598  = 0x07 // Latin/Hebrew (ISO-8859-8)
	ENCODING_ISO10646  = 0x08 // UCS2 (ISO/IEC-10646)
	ENCODING_PICTOGRAM = 0x09 // Pictogram Encoding
	ENCODING_ISO2022JP = 0x0A // ISO-2022-JP (Music Codes)
	ENCODING_EXTJIS    = 0x0D // Extended Kanji JIS (X 0212-1990)
	ENCODING_KSC5601   = 0x0E // KS C 5601
)

const (
	// ESM_CLASS Types
	ESM_CLASS_MSGMODE_DEFAULT      = 0x00 // Default SMSC mode (e.g. Store and Forward)
	ESM_CLASS_MSGMODE_DATAGRAM     = 0x01 // Datagram mode
	ESM_CLASS_MSGMODE_FORWARD      = 0x02 // Forward (i.e. Transaction) mode
	ESM_CLASS_MSGMODE_STOREFORWARD = 0x03 // Store and Forward mode (use this to select Store and Forward mode if Default mode is not Store and Forward)

	ESM_CLASS_MSGTYPE_DEFAULT     = 0x00 // Default message type (i.e. normal message)
	ESM_CLASS_MSGTYPE_DELIVERYACK = 0x08 // Message containts ESME Delivery Acknowledgement
	ESM_CLASS_MSGTYPE_USERACK     = 0x10 // Message containts ESME Manual/User Acknowledgement

	ESM_CLASS_GSMFEAT_NONE          = 0x00 // No specific features selected
	ESM_CLASS_GSMFEAT_UDHI          = 0x40 // UDHI Indicator (only relevant for MT msgs)
	ESM_CLASS_GSMFEAT_REPLYPATH     = 0x80 // Set Reply Path (only relevant for GSM net)
	ESM_CLASS_GSMFEAT_UDHIREPLYPATH = 0xC0 // Set UDHI and Reply Path (for GSM net)
)

func (s CMDId) Error() string {
	switch s {
	case GENERIC_NACK:
		return "GENERIC_NACK"
	case BIND_RECEIVER:
		return "BIND_RECEIVER"
	case BIND_RECEIVER_RESP:
		return "BIND_RECEIVER_RESP"
	case BIND_TRANSMITTER:
		return "BIND_TRANSMITTER"
	case BIND_TRANSMITTER_RESP:
		return "BIND_TRANSMITTER_RESP"
	case QUERY_SM:
		return "QUERY_SM"
	case QUERY_SM_RESP:
		return "QUERY_SM_RESP"
	case SUBMIT_SM:
		return "SUBMIT_SM"
	case SUBMIT_SM_RESP:
		return "SUBMIT_SM_RESP"
	case DELIVER_SM:
		return "DELIVER_SM"
	case DELIVER_SM_RESP:
		return "DELIVER_SM_RESP"
	case UNBIND:
		return "UNBIND"
	case UNBIND_RESP:
		return "UNBIND_RESP"
	case REPLACE_SM:
		return "REPLACE_SM"
	case REPLACE_SM_RESP:
		return "REPLACE_SM_RESP"
	case CANCEL_SM:
		return "CANCEL_SM"
	case CANCEL_SM_RESP:
		return "CANCEL_SM_RESP"
	case BIND_TRANSCEIVER:
		return "BIND_TRANSCEIVER"
	case BIND_TRANSCEIVER_RESP:
		return "BIND_TRANSCEIVER_RESP"
	case OUTBIND:
		return "OUTBIND"
	case ENQUIRE_LINK:
		return "ENQUIRE_LINK"
	case ENQUIRE_LINK_RESP:
		return "ENQUIRE_LINK_RESP"
	case SUBMIT_MULTI:
		return "SUBMIT_MULTI"
	case SUBMIT_MULTI_RESP:
		return "SUBMIT_MULTI_RESP"
	case ALERT_NOTIFICATION:
		return "ALERT_NOTIFICATION"
	case DATA_SM:
		return "DATA_SM"
	case DATA_SM_RESP:
		return "DATA_SM_RESP"
	default:
		return fmt.Sprint("Unknown PDU Type. ID:", uint32(s))
	}
}

func (s CMDStatus) Error() string {
	switch s {
	default:
		return fmt.Sprint("Unknown Status:", uint32(s))
	case ESME_ROK:
		return fmt.Sprint("No Error")
	case ESME_RINVMSGLEN:
		return fmt.Sprint("Message Length is invalid")
	case ESME_RINVCMDLEN:
		return fmt.Sprint("Command Length is invalid")
	case ESME_RINVCMDID:
		return fmt.Sprint("Invalid Command ID")
	case ESME_RINVBNDSTS:
		return fmt.Sprint("Incorrect BIND Status for given command")
	case ESME_RALYBND:
		return fmt.Sprint("ESME Already in Bound State")
	case ESME_RINVPRTFLG:
		return fmt.Sprint("Invalid Priority Flag")
	case ESME_RINVREGDLVFLG:
		return fmt.Sprint("Invalid Registered Delivery Flag")
	case ESME_RSYSERR:
		return fmt.Sprint("System Error")
	case ESME_RINVSRCADR:
		return fmt.Sprint("Invalid Source Address")
	case ESME_RINVDSTADR:
		return fmt.Sprint("Invalid Dest Addr")
	case ESME_RINVMSGID:
		return fmt.Sprint("Message ID is invalid")
	case ESME_RBINDFAIL:
		return fmt.Sprint("Bind Failed")
	case ESME_RINVPASWD:
		return fmt.Sprint("Invalid Password")
	case ESME_RINVSYSID:
		return fmt.Sprint("Invalid System ID")
	case ESME_RCANCELFAIL:
		return fmt.Sprint("Cancel SM Failed")
	case ESME_RREPLACEFAIL:
		return fmt.Sprint("Replace SM Failed")
	case ESME_RMSGQFUL:
		return fmt.Sprint("Message Queue Full")
	case ESME_RINVSERTYP:
		return fmt.Sprint("Invalid Service Type")
	case ESME_RINVNUMDESTS:
		return fmt.Sprint("Invalid number of destinations")
	case ESME_RINVDLNAME:
		return fmt.Sprint("Invalid Distribution List name")
	case ESME_RINVDESTFLAG:
		return fmt.Sprint("Destination flag is invalid")
	case ESME_RINVSUBREP:
		return fmt.Sprint("Invalid 'submit with replace' request")
	case ESME_RINVESMCLASS:
		return fmt.Sprint("Invalid esm_class field data")
	case ESME_RCNTSUBDL:
		return fmt.Sprint("Cannot Submit to Distribution List")
	case ESME_RSUBMITFAIL:
		return fmt.Sprint("submit_sm or submit_multi failed")
	case ESME_RINVSRCTON:
		return fmt.Sprint("Invalid Source address TON")
	case ESME_RINVSRCNPI:
		return fmt.Sprint("Invalid Source address NPI")
	case ESME_RINVDSTTON:
		return fmt.Sprint("Invalid Destination address TON")
	case ESME_RINVDSTNPI:
		return fmt.Sprint("Invalid Destination address NPI")
	case ESME_RINVSYSTYP:
		return fmt.Sprint("Invalid system_type field")
	case ESME_RINVREPFLAG:
		return fmt.Sprint("Invalid replace_if_present flag")
	case ESME_RINVNUMMSGS:
		return fmt.Sprint("Invalid number of messages")
	case ESME_RTHROTTLED:
		return fmt.Sprint("Throttling error (ESME has exceeded allowed message limit")
	case ESME_RINVSCHED:
		return fmt.Sprint("Invalid Scheduled Delivery Time")
	case ESME_RINVEXPIRY:
		return fmt.Sprint("Invalid message validity period (Expiry time)")
	case ESME_RINVDFTMSGID:
		return fmt.Sprint("Predefined Message Invalid or Not Found")
	case ESME_RX_T_APPN:
		return fmt.Sprint("ESME Receiver Temporary App Error Code")
	case ESME_RX_P_APPN:
		return fmt.Sprint("ESME Receiver Permanent App Error Code")
	case ESME_RX_R_APPN:
		return fmt.Sprint("ESME Receiver Reject Message Error Code")
	case ESME_RQUERYFAIL:
		return fmt.Sprint("Query_sm request failed")
	case ESME_RINVOPTPARSTREAM:
		return fmt.Sprint("Error in the optional part of the PDU Body.")
	case ESME_ROPTPARNOTALLWD:
		return fmt.Sprint("Optional Parameter not allowed")
	case ESME_RINVPARLEN:
		return fmt.Sprint("Invalid Parameter Length.")
	case ESME_RMISSINGOPTPARAM:
		return fmt.Sprint("Expected Optional Parameter missing")
	case ESME_RINVOPTPARAMVAL:
		return fmt.Sprint("Invalid Optional Parameter Value")
	case ESME_RDELIVERYFAILURE:
		return fmt.Sprint("Delivery Failure (used for data_sm_resp)")
	case ESME_RUNKNOWNERR:
		return fmt.Sprint("Unknown Error")
	}
}
