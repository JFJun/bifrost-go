package expand

import (
	"github.com/JFJun/bifrost-go/expand/bifrost"
	"github.com/stafiprotocol/go-substrate-rpc-client/types"
)

type IEventRecords interface {
	GetBalancesTransfer() []types.EventBalancesTransfer
	GetSystemExtrinsicSuccess() []types.EventSystemExtrinsicSuccess
	GetSystemExtrinsicFailed() []types.EventSystemExtrinsicFailed
}
/*
扩展： 解析event
*/
func DecodeEventRecords(meta *types.Metadata, rawData string, chainName string) (IEventRecords, error) {
	e := types.EventRecordsRaw(types.MustHexDecodeString(rawData))
	var ier IEventRecords
	var events bifrost.BifrostEventRecords
	err := e.DecodeEventRecords(meta, &events)
	if err != nil {
		return nil, err
	}
	ier = &events
	return ier, nil
}
