package polkadot

import (
	"github.com/JFJun/bifrost-go/expand/base"
	"github.com/JFJun/go-substrate-rpc-client/v3/types"
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
	//2021-06-09
	Staking_Kicked                                  []EventStakingKicked
	ElectionsPhragmen_NewTerm                       []EventElectionsPhragmenNewTerm
	ElectionsPhragmen_Renounced                     []EventElectionsPhragmenRenounced
	ElectionsPhragmen_CandidateSlashed              []EventElectionsPhragmenCandidateSlashed
	ElectionsPhragmen_SeatHolderSlashed             []EventElectionsPhragmenSeatHolderSlashed
	Bounties_BountyProposed                         []EventBountiesBountyProposed
	Bounties_BountyRejected                         []EventBountiesBountyRejected
	Bounties_BountyBecameActive                     []EventBountiesBountyBecameActive
	Bounties_BountyAwarded                          []EventBountiesBountyAwarded
	Bounties_BountyClaimed                          []EventBountiesBountyClaimed
	Bounties_BountyCanceled                         []EventBountiesBountyCanceled
	Bounties_BountyExtended                         []EventBountiesBountyExtended
	Tips_NewTip                                     []EventTipsNewTip
	Tips_TipClosing                                 []EventTipsTipClosing
	Tips_TipClosed                                  []EventTipsTipClosed
	Tips_TipRetracted                               []EventTipsTipRetracted
	Tips_TipSlashed                                 []EventTipsTipSlashed
	ElectionProviderMultiPhase_SolutionStored       []EventElectionProviderMultiPhaseSolutionStored
	ElectionProviderMultiPhase_ElectionFinalized    []EventElectionProviderMultiPhaseElectionFinalized
	ElectionProviderMultiPhase_Rewarded             []EventElectionProviderMultiPhaseRewarded
	ElectionProviderMultiPhase_Slashed              []EventElectionProviderMultiPhaseSlashed
	ElectionProviderMultiPhase_SignedPhaseStarted   []EventElectionProviderMultiPhaseSignedPhaseStarted
	ElectionProviderMultiPhase_UnsignedPhaseStarted []EventElectionProviderMultiPhaseUnsignedPhaseStarted
}

type EventElectionProviderMultiPhaseUnsignedPhaseStarted struct {
	Phase  types.Phase
	U32    types.U32
	Topics []types.Hash
}
type EventElectionProviderMultiPhaseSignedPhaseStarted struct {
	Phase  types.Phase
	U32    types.U32
	Topics []types.Hash
}
type EventElectionProviderMultiPhaseSlashed struct {
	Phase     types.Phase
	AccountId types.AccountID
	Topics    []types.Hash
}
type EventElectionProviderMultiPhaseRewarded struct {
	Phase     types.Phase
	AccountId types.AccountID
	Topics    []types.Hash
}
type EventElectionProviderMultiPhaseElectionFinalized struct {
	Phase                  types.Phase
	Option_ElectionCompute base.OptionElectionCompute
	Topics                 []types.Hash
}
type EventElectionProviderMultiPhaseSolutionStored struct {
	Phase           types.Phase
	ElectionCompute types.ElectionCompute
	Topics          []types.Hash
}
type EventTipsTipSlashed struct {
	Phase     types.Phase
	Hash      types.Hash
	AccountId types.AccountID
	Balance   types.U128
	Topics    []types.Hash
}
type EventTipsTipRetracted struct {
	Phase  types.Phase
	Hash   types.Hash
	Topics []types.Hash
}
type EventTipsTipClosed struct {
	Phase     types.Phase
	Hash      types.Hash
	AccountId types.AccountID
	Balance   types.U128
	Topics    []types.Hash
}
type EventTipsTipClosing struct {
	Phase  types.Phase
	Hash   types.Hash
	Topics []types.Hash
}
type EventTipsNewTip struct {
	Phase  types.Phase
	Hash   types.Hash
	Topics []types.Hash
}
type EventBountiesBountyExtended struct {
	Phase       types.Phase
	BountyIndex types.U32
	Topics      []types.Hash
}
type EventBountiesBountyCanceled struct {
	Phase       types.Phase
	BountyIndex types.U32
	Topics      []types.Hash
}
type EventBountiesBountyClaimed struct {
	Phase       types.Phase
	BountyIndex types.U32
	Balance     types.U128
	AccountId   types.AccountID
	Topics      []types.Hash
}
type EventBountiesBountyAwarded struct {
	Phase       types.Phase
	BountyIndex types.U32
	AccountId   types.AccountID
	Topics      []types.Hash
}
type EventBountiesBountyBecameActive struct {
	Phase       types.Phase
	BountyIndex types.U32
	Topics      []types.Hash
}
type EventBountiesBountyRejected struct {
	Phase       types.Phase
	BountyIndex types.U32
	Balance     types.U128
	Topics      []types.Hash
}
type EventBountiesBountyProposed struct {
	Phase       types.Phase
	BountyIndex types.U32
	Topics      []types.Hash
}
type EventElectionsPhragmenSeatHolderSlashed struct {
	Phase     types.Phase
	AccountId types.AccountID
	Balance   types.U128
	Topics    []types.Hash
}
type EventElectionsPhragmenCandidateSlashed struct {
	Phase     types.Phase
	AccountId types.AccountID
	Balance   types.U128
	Topics    []types.Hash
}
type EventElectionsPhragmenRenounced struct {
	Phase     types.Phase
	AccountId types.AccountID
	Topics    []types.Hash
}
type EventElectionsPhragmenNewTerm struct {
	Phase  types.Phase
	Abs    []base.AccountBalance
	Topics []types.Hash
}
type EventStakingKicked struct {
	Phase      types.Phase
	AccountId1 types.AccountID
	AccountId2 types.AccountID
	Topics     []types.Hash
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
