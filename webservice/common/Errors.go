package common

import "errors"

var (
	ERR_CONNECTION_LOSS = errors.New("ERR_CONNECTION_LOSS")

	ERR_SEND_MESSAGE_FULL = errors.New("ERR_SEND_MESSAGE_FULL")

	ERR_JOIN_ROOM_TWICE = errors.New("ERR_JOIN_ROOM_TWICE")

	ERR_NOT_IN_ROOM = errors.New("ERR_NOT_IN_ROOM")

	ERR_ROOM_ID_INVALID = errors.New("ERR_ROOM_ID_INVALID")

	ERR_DISPATCH_CHANNEL_FULL = errors.New("ERR_DISPATCH_CHANNEL_FULL")

	ERR_MERGE_CHANNEL_FULL = errors.New("ERR_MERGE_CHANNEL_FULL")

	ERR_CERT_INVALID = errors.New("ERR_CERT_INVALID")

	ERR_LOGIC_DISPATCH_CHANNEL_FULL = errors.New("ERR_LOGIC_DISPATCH_CHANNEL_FULL")
)

