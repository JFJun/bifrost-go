package base

import "github.com/JFJun/go-substrate-rpc-client/v3/types"

type HrmpChannelId struct {
	Sender   types.U32
	Receiver types.U32
}
