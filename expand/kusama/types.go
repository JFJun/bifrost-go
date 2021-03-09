package kusama

import (
	"github.com/JFJun/bifrost-go/expand/base"
	"github.com/stafiprotocol/go-substrate-rpc-client/types"
)

func (p KusamaEventRecords) GetBalancesTransfer() []types.EventBalancesTransfer {
	return p.Balances_Transfer
}

func (p KusamaEventRecords) GetSystemExtrinsicSuccess() []types.EventSystemExtrinsicSuccess {
	return p.System_ExtrinsicSuccess
}

func (p KusamaEventRecords) GetSystemExtrinsicFailed() []types.EventSystemExtrinsicFailed {
	return p.System_ExtrinsicFailed
}

type KusamaEventRecords struct {
	base.BaseEventRecords
	Staking_Kicked                      []EventStakingKicked
	Democracy_Blacklisted               []EventDemocracyBlacklisted
	ElectionsPhragmen_NewTerm           []EventElectionsPhragmenNewTerm
	ElectionsPhragmen_EmptyTerm         []EventElectionsPhragmenEmptyTerm
	ElectionsPhragmen_ElectionError     []EventElectionsPhragmenElectionError
	ElectionsPhragmen_MemberKicked      []EventElectionsPhragmenMemberKicked
	ElectionsPhragmen_Renounced         []EventElectionsPhragmenRenounced
	ElectionsPhragmen_CandidateSlashed  []EventElectionsPhragmenCandidateSlashed
	ElectionsPhragmen_SeatHolderSlashed []EventElectionsPhragmenSeatHolderSlashed
	Bounties_BountyProposed             []EventBountiesBountyProposed
	Bounties_BountyRejected             []EventBountiesBountyRejected
	Bounties_BountyBecameActive         []EventBountiesBountyBecameActive
	Bounties_BountyAwarded              []EventBountiesBountyAwarded
	Bounties_BountyClaimed              []EventBountiesBountyClaimed
	Bounties_BountyCanceled             []EventBountiesBountyCanceled
	Bounties_BountyExtended             []EventBountiesBountyExtended
	Tips_NewTip                         []EventTipsNewTip
	Tips_TipClosing                     []EventTipsTipClosing
	Tips_TipClosed                      []EventTipsTipClosed
	Tips_TipRetracted                   []EventTipsTipRetracted
	Tips_TipSlashed                     []EventTipsTipSlashed
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
type EventElectionsPhragmenMemberKicked struct {
	Phase     types.Phase
	AccountId types.AccountID
	Topics    []types.Hash
}
type EventElectionsPhragmenElectionError struct {
	Phase types.Phase

	Topics []types.Hash
}
type EventStakingKicked struct {
	Phase      types.Phase
	AccountId1 types.AccountID
	AccountId  types.AccountID
	Topics     []types.Hash
}
type EventDemocracyBlacklisted struct {
	Phase  types.Phase
	Hash   types.Hash
	Topics []types.Hash
}
type EventElectionsPhragmenNewTerm struct {
	Phase  types.Phase
	Abs    []AccountBalance
	Topics []types.Hash
}

type AccountBalance struct {
	Who     types.AccountID
	Balance types.U128
}

type EventElectionsPhragmenEmptyTerm struct {
	Phase types.Phase

	Topics []types.Hash
}
