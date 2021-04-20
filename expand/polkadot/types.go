package polkadot

import (
	"github.com/JFJun/bifrost-go/expand/base"
	"github.com/stafiprotocol/go-substrate-rpc-client/types"
)

type PolkadotEventRecords struct {
	base.BaseEventRecords
	Claims_Claimed                    []EventClaimsClaimed
	ElectionsPhragmen_VoterReported   []EventElectionsPhragmenVoterReported
	ElectionsPhragmen_MemberRenounced []EventElectionsPhragmenMemberRenounced
	ElectionsPhragmen_MemberKicked    []EventElectionsPhragmenMemberKicked
	ElectionsPhragmen_ElectionError   []EventElectionsPhragmenElectionError
	ElectionsPhragmen_EmptyTerm       []EventElectionsPhragmenEmptyTerm
	//ElectionsPhragmen_NewTerm		[]EventElectionsPhragmenNewTerm		暂不支持解析
	Democracy_Blacklisted []EventDemocracyBlacklisted
	Tips_TipClosing []EventTipsTipClosing
	Tips_NewTip []EventTreasuryNewTip
	ElectionProviderMultiPhase_UnsignedPhaseStarted []EventElectionProviderMultiPhaseUnsignedPhaseStarted
	ElectionProviderMultiPhase_SolutionStored []ElectionProviderMultiPhaseSolutionStored
	
}

type EventTreasuryNewTip struct {
	Phase  types.Phase
	Hash   types.Hash
	Topics []types.Hash
}

func (p PolkadotEventRecords) GetBalancesTransfer() []types.EventBalancesTransfer {
	return p.Balances_Transfer
}

func (p PolkadotEventRecords) GetSystemExtrinsicSuccess() []types.EventSystemExtrinsicSuccess {
	return p.System_ExtrinsicSuccess
}

func (p PolkadotEventRecords) GetSystemExtrinsicFailed() []types.EventSystemExtrinsicFailed {
	return p.System_ExtrinsicFailed
}

type EventElectionProviderMultiPhaseUnsignedPhaseStarted struct {
	Phase     types.Phase
	BountyIndex types.U32
	Topics    []types.Hash
}
type ElectionProviderMultiPhaseSolutionStored struct {
	Phase     types.Phase
	ElectionCompute types.ElectionCompute
	Topics    []types.Hash
}
// EventTreasuryTipClosing is emitted when a tip suggestion has reached threshold and is closing.
type EventTipsTipClosing struct {
	Phase  types.Phase
	Hash   types.Hash
	Topics []types.Hash
}

type EventDemocracyBlacklisted struct {
	Phase  types.Phase
	Hash   types.Hash
	Topics []types.Hash
}

//type EventElectionsPhragmenNewTerm struct {
//	Phase    types.Phase
//	Vec
//	Topics []types.Hash
//}
type EventElectionsPhragmenEmptyTerm struct {
	Phase types.Phase

	Topics []types.Hash
}
type EventElectionsPhragmenElectionError struct {
	Phase  types.Phase
	Topics []types.Hash
}
type EventElectionsPhragmenMemberKicked struct {
	Phase     types.Phase
	AccountId types.AccountID
	Topics    []types.Hash
}
type EventElectionsPhragmenMemberRenounced struct {
	Phase     types.Phase
	AccountId types.AccountID
	Topics    []types.Hash
}
type EventElectionsPhragmenVoterReported struct {
	Phase  types.Phase
	Who1   types.AccountID
	Who2   types.AccountID
	Bool   types.Bool
	Topics []types.Hash
}
type EventClaimsClaimed struct {
	Phase           types.Phase
	AccountId       types.AccountID
	EthereumAddress base.VecU8L20
	Balance         types.U128
	Topics          []types.Hash
}
