package bifrost

import (
	"fmt"
	"github.com/JFJun/bifrost-go/expand/base"
	"github.com/stafiprotocol/go-substrate-rpc-client/scale"
	"github.com/stafiprotocol/go-substrate-rpc-client/types"
)

type BifrostEventRecords struct {
	base.BaseEventRecords
	Democracy_Blacklisted		[]EventDemocracyBlacklisted
	Elections_ElectionError		[]EventElectionsElectionError
	Elections_CandidateSlashed		[]EventElectionsCandidateSlashed
	Elections_SeatHolderSlashed		[]EventElectionsSeatHolderSlashed
	Assets_Created		[]EventAssetsCreated
	Assets_AccountAssetCreated		[]EventAssetsAccountAssetCreated
	Assets_AccountAssetDestroy		[]EventAssetsAccountAssetDestroy
	Assets_UnlockedAsset		[]EventAssetsUnlockedAsset
	Convert_UpdateConvertSuccess		[]EventConvertUpdateConvertSuccess
	Convert_UpdateRatePerBlockSuccess		[]EventConvertUpdateRatePerBlockSuccess
	Convert_ConvertTokenToVTokenSuccess		[]EventConvertConvertTokenToVTokenSuccess
	Convert_ConvertVTokenToTokenSuccess		[]EventConvertConvertVTokenToTokenSuccess
	Convert_RedeemedPointsSuccess		[]EventConvertRedeemedPointsSuccess
	Convert_UpdateConvertPoolSuccess		[]EventConvertUpdateConvertPoolSuccess
	BridgeEos_InitSchedule		[]EventBridgeEosInitSchedule
	BridgeEos_ChangeSchedule		[]EventBridgeEosChangeSchedule
	BridgeEos_ProveAction		[]EventBridgeEosProveAction
	BridgeEos_RelayBlock		[]EventBridgeEosRelayBlock
	BridgeEos_Deposit		[]EventBridgeEosDeposit
	BridgeEos_DepositFail		[]EventBridgeEosDepositFail
	BridgeEos_Withdraw		[]EventBridgeEosWithdraw
	BridgeEos_WithdrawFail		[]EventBridgeEosWithdrawFail
	BridgeEos_SentCrossChainTransaction		[]EventBridgeEosSentCrossChainTransaction
	BridgeEos_FailToSendCrossChainTransaction		[]EventBridgeEosFailToSendCrossChainTransaction
	BridgeEos_GrantedCrossChainPrivilege		[]EventBridgeEosGrantedCrossChainPrivilege
	BridgeEos_RemovedCrossChainPrivilege		[]EventBridgeEosRemovedCrossChainPrivilege
	BridgeEos_UnsignedTrx		[]EventBridgeEosUnsignedTrx
	BridgeIost_InitSchedule		[]EventBridgeIostInitSchedule
	BridgeIost_ChangeSchedule		[]EventBridgeIostChangeSchedule
	BridgeIost_ProveAction		[]EventBridgeIostProveAction
	BridgeIost_RelayBlock		[]EventBridgeIostRelayBlock
	BridgeIost_Deposit		[]EventBridgeIostDeposit
	BridgeIost_DepositFail		[]EventBridgeIostDepositFail
	BridgeIost_Withdraw		[]EventBridgeIostWithdraw
	BridgeIost_WithdrawFail		[]EventBridgeIostWithdrawFail
	BridgeIost_SendTransactionSuccess		[]EventBridgeIostSendTransactionSuccess
	BridgeIost_SendTransactionFailure		[]EventBridgeIostSendTransactionFailure
	BridgeIost_GrantedCrossChainPrivilege		[]EventBridgeIostGrantedCrossChainPrivilege
	BridgeIost_RemovedCrossChainPrivilege		[]EventBridgeIostRemovedCrossChainPrivilege
	Swap_AddLiquiditySuccess		[]EventSwapAddLiquiditySuccess
	Swap_RemoveLiquiditySuccess		[]EventSwapRemoveLiquiditySuccess
	Swap_AddSingleLiquiditySuccess		[]EventSwapAddSingleLiquiditySuccess
	Swap_RemoveSingleLiquiditySuccess		[]EventSwapRemoveSingleLiquiditySuccess
	Swap_SwapTokenSuccess		[]EventSwapSwapTokenSuccess
	Swap_CreatePoolSuccess		[]EventSwapCreatePoolSuccess
	Voucher_IssuedVoucher		[]EventVoucherIssuedVoucher
	Voucher_DestroyedVoucher		[]EventVoucherDestroyedVoucher
}
type EventVoucherDestroyedVoucher struct {
	Phase    types.Phase
	Who 	types.AccountID
	Balance 	types.U128
	Topics []types.Hash
}
type EventVoucherIssuedVoucher struct {
	Phase    types.Phase
	Who 	types.AccountID
	Balance 	types.U128
	Topics []types.Hash
}
type EventSwapCreatePoolSuccess struct {
	Phase    types.Phase

	Topics []types.Hash
}
type EventSwapSwapTokenSuccess struct {
	Phase    types.Phase
	Balance1	types.U128
	Balance2	types.U128
	Topics []types.Hash
}
type EventSwapRemoveSingleLiquiditySuccess struct {
	Phase    types.Phase

	Topics []types.Hash
}
type EventSwapAddSingleLiquiditySuccess struct {
	Phase    types.Phase

	Topics []types.Hash
}
type EventSwapRemoveLiquiditySuccess struct {
	Phase    types.Phase

	Topics []types.Hash
}
type EventSwapAddLiquiditySuccess struct {
	Phase    types.Phase

	Topics []types.Hash
}
type EventBridgeIostRemovedCrossChainPrivilege struct {
	Phase    types.Phase
	Who 	types.AccountID
	Topics []types.Hash
}
type EventBridgeIostGrantedCrossChainPrivilege struct {
	Phase    types.Phase
	Who 	types.AccountID
	Topics []types.Hash
}
type EventBridgeIostSendTransactionFailure struct {
	Phase    types.Phase

	Topics []types.Hash
}
type EventBridgeIostSendTransactionSuccess struct {
	Phase    types.Phase

	Topics []types.Hash
}
type EventBridgeIostWithdrawFail struct {
	Phase    types.Phase

	Topics []types.Hash
}
type EventBridgeIostWithdraw struct {
	Phase    types.Phase
	Who 	types.AccountID
	Data 	types.Bytes
	Topics []types.Hash
}
type EventBridgeIostDepositFail struct {
	Phase    types.Phase

	Topics []types.Hash
}
type EventBridgeIostDeposit struct {
	Phase    types.Phase
	Data 	types.Bytes
	Who 	types.AccountID
	Topics []types.Hash
}

type EventBridgeIostRelayBlock struct {
	Phase    types.Phase

	Topics []types.Hash
}
type EventBridgeIostProveAction struct {
	Phase    types.Phase

	Topics []types.Hash
}
type EventBridgeIostChangeSchedule struct {
	Phase    types.Phase
	VersionId1	VersionId
	VersionId2	VersionId
	Topics []types.Hash
}
type EventBridgeIostInitSchedule struct {
	Phase    types.Phase
	VersionId	VersionId
	Topics []types.Hash
}
type EventBridgeEosUnsignedTrx struct {
	Phase    types.Phase

	Topics []types.Hash
}
type EventBridgeEosRemovedCrossChainPrivilege struct {
	Phase    types.Phase
	Who 	types.AccountID
	Topics []types.Hash
}
type EventBridgeEosGrantedCrossChainPrivilege struct {
	Phase    types.Phase
	Who 	types.AccountID
	Topics []types.Hash
}
type EventBridgeEosFailToSendCrossChainTransaction struct {
	Phase    types.Phase

	Topics []types.Hash
}
type EventBridgeEosSentCrossChainTransaction struct {
	Phase    types.Phase

	Topics []types.Hash
}
type EventBridgeEosWithdrawFail struct {
	Phase    types.Phase

	Topics []types.Hash
}
type EventBridgeEosWithdraw struct {
	Phase    types.Phase
	Who 	types.AccountID
	Data 	types.Bytes
	Topics []types.Hash
}
type EventBridgeEosDepositFail struct {
	Phase    types.Phase

	Topics []types.Hash
}
type EventBridgeEosDeposit struct {
	Phase    types.Phase
	Data 	types.Bytes
	Who 	types.AccountID
	Topics []types.Hash
}
type EventBridgeEosRelayBlock struct {
	Phase    types.Phase

	Topics []types.Hash
}
type EventBridgeEosProveAction struct {
	Phase    types.Phase

	Topics []types.Hash
}
type EventBridgeEosChangeSchedule struct {
	Phase    types.Phase
	VersionId1	VersionId
	VersionId2	VersionId
	Topics []types.Hash
}
type EventBridgeEosInitSchedule struct {
	Phase    types.Phase
	VersionId	VersionId
	Topics []types.Hash
}
type EventConvertUpdateConvertPoolSuccess struct {
	Phase    types.Phase

	Topics []types.Hash
}
type EventConvertRedeemedPointsSuccess struct {
	Phase    types.Phase

	Topics []types.Hash
}
type EventConvertConvertVTokenToTokenSuccess struct {
	Phase    types.Phase

	Topics []types.Hash
}

type EventConvertConvertTokenToVTokenSuccess struct {
	Phase    types.Phase

	Topics []types.Hash
}
type EventConvertUpdateRatePerBlockSuccess struct {
	Phase    types.Phase

	Topics []types.Hash
}
type EventConvertUpdateConvertSuccess struct {
	Phase    types.Phase
	Topics []types.Hash
}
type EventAssetsUnlockedAsset struct {
	Phase    types.Phase
	Who 	types.AccountID
	TokenSymbol TokenSymbol
	Balance	types.U128
	Topics []types.Hash
}
type EventAssetsAccountAssetDestroy struct {
	Phase    types.Phase
	Who 	types.AccountID
	AssetId	AssetId
	Topics []types.Hash
}
type EventAssetsAccountAssetCreated struct {
	Phase    types.Phase
	Who 	types.AccountID
	AssetId	AssetId
	Topics []types.Hash
}
type EventAssetsCreated struct {
	Phase    types.Phase
	AssetId	 AssetId
	TokenBalance types.U128
	Topics []types.Hash
}
type EventElectionsSeatHolderSlashed struct {
	Phase    types.Phase
	Who 	types.AccountID
	Balance 	types.U128
	Topics []types.Hash
}
type EventElectionsCandidateSlashed struct {
	Phase    types.Phase
	Who 	types.AccountID
	Balance 	types.U128
	Topics []types.Hash
}

type EventDemocracyBlacklisted struct {
	Phase    types.Phase
	Hash	types.Hash
	Topics []types.Hash
}
type EventElectionsElectionError struct {
	Phase    types.Phase

	Topics []types.Hash
}

/*
Bifrost types: https://github.com/bifrost-finance/bifrost/blob/multiaddress-parachain-doc/docs/developer_setting.json
*/

type VersionId types.U32

type TokenSymbol struct {
	enum []string
	Value string
}
func (d *TokenSymbol)Decode(decoder scale.Decoder)error{
	d.enum = []string{"aUSD", "DOT", "vDOT", "KSM", "vKSM", "EOS", "vEOS", "IOST", "vIOST"}
	b,err:=decoder.ReadOneByte()
	if err != nil {
		return err
	}
	if int(b)>len(d.enum) || int(b)<0 {
		return fmt.Errorf("types=[TokenSymbol]  don not have this enum: %d",b)
	}
	d.Value = d.enum[int(b)]
	return nil
}

type AssetId types.U32