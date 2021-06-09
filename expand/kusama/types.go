package kusama

import (
	"fmt"
	"github.com/JFJun/bifrost-go/expand/base"
	"github.com/JFJun/go-substrate-rpc-client/v3/scale"
	"github.com/JFJun/go-substrate-rpc-client/v3/types"
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
	// 2021-06-09
	PhragmenElection_NewTerm                        []EventPhragmenElectionNewTerm
	PhragmenElection_EmptyTerm                      []EventPhragmenElectionEmptyTerm
	PhragmenElection_ElectionError                  []EventPhragmenElectionElectionError
	PhragmenElection_MemberKicked                   []EventPhragmenElectionMemberKicked
	PhragmenElection_Renounced                      []EventPhragmenElectionRenounced
	PhragmenElection_CandidateSlashed               []EventPhragmenElectionCandidateSlashed
	PhragmenElection_SeatHolderSlashed              []EventPhragmenElectionSeatHolderSlashed
	ElectionProviderMultiPhase_SolutionStored       []EventElectionProviderMultiPhaseSolutionStored
	ElectionProviderMultiPhase_ElectionFinalized    []EventElectionProviderMultiPhaseElectionFinalized
	ElectionProviderMultiPhase_Rewarded             []EventElectionProviderMultiPhaseRewarded
	ElectionProviderMultiPhase_Slashed              []EventElectionProviderMultiPhaseSlashed
	ElectionProviderMultiPhase_SignedPhaseStarted   []EventElectionProviderMultiPhaseSignedPhaseStarted
	ElectionProviderMultiPhase_UnsignedPhaseStarted []EventElectionProviderMultiPhaseUnsignedPhaseStarted
	Gilt_BidRetracted                               []EventGiltBidRetracted
	Gilt_BidPlaced                                  []EventGiltBidPlaced
	Gilt_GiltIssued                                 []EventGiltGiltIssued
	Gilt_GiltThawed                                 []EventGiltGiltThawed
	ParasInclusion_CandidateBacked                  []EventParasInclusionCandidateBacked
	ParasInclusion_CandidateIncluded                []EventParasInclusionCandidateIncluded
	ParasInclusion_CandidateTimedOut                []EventParasInclusionCandidateTimedOut
	Paras_CurrentCodeUpdated                        []EventParasCurrentCodeUpdated
	Paras_CurrentHeadUpdated                        []EventParasCurrentHeadUpdated
	Paras_CodeUpgradeScheduled                      []EventParasCodeUpgradeScheduled
	Paras_NewHeadNoted                              []EventParasNewHeadNoted
	Paras_ActionQueued                              []EventParasActionQueued
	ParasUmp_InvalidFormat                          []EventParasUmpInvalidFormat
	ParasUmp_UnsupportedVersion                     []EventParasUmpUnsupportedVersion
	ParasUmp_ExecutedUpward                         []EventParasUmpExecutedUpward
	ParasUmp_WeightExhausted                        []EventParasUmpWeightExhausted
	ParasUmp_UpwardMessagesReceived                 []EventParasUmpUpwardMessagesReceived
	ParasHrmp_OpenChannelRequested                  []EventParasHrmpOpenChannelRequested
	ParasHrmp_OpenChannelAccepted                   []EventParasHrmpOpenChannelAccepted
	ParasHrmp_ChannelClosed                         []EventParasHrmpChannelClosed
	Registrar_Registered                            []EventRegistrarRegistered
	Registrar_Deregistered                          []EventRegistrarDeregistered
	Registrar_Reserved                              []EventRegistrarReserved
	Slots_NewLeasePeriod                            []EventSlotsNewLeasePeriod
	Slots_Leased                                    []EventSlotsLeased
	Auctions_AuctionStarted                         []EventAuctionsAuctionStarted
	Auctions_AuctionClosed                          []EventAuctionsAuctionClosed
	Auctions_WonDeploy                              []EventAuctionsWonDeploy
	Auctions_WonRenewal                             []EventAuctionsWonRenewal
	Auctions_Reserved                               []EventAuctionsReserved
	Auctions_Unreserved                             []EventAuctionsUnreserved
	Auctions_ReserveConfiscated                     []EventAuctionsReserveConfiscated
	Auctions_BidAccepted                            []EventAuctionsBidAccepted
	Auctions_WinningOffset                          []EventAuctionsWinningOffset
	Crowdloan_Created                               []EventCrowdloanCreated
	Crowdloan_Contributed                           []EventCrowdloanContributed
	Crowdloan_Withdrew                              []EventCrowdloanWithdrew
	Crowdloan_PartiallyRefunded                     []EventCrowdloanPartiallyRefunded
	Crowdloan_AllRefunded                           []EventCrowdloanAllRefunded
	Crowdloan_Dissolved                             []EventCrowdloanDissolved
	Crowdloan_DeployDataFixed                       []EventCrowdloanDeployDataFixed
	Crowdloan_Onboarded                             []EventCrowdloanOnboarded
	Crowdloan_HandleBidResult                       []EventCrowdloanHandleBidResult
	Crowdloan_Edited                                []EventCrowdloanEdited
	Crowdloan_MemoUpdated                           []EventCrowdloanMemoUpdated
	Crowdloan_AddedToNewRaise                       []EventCrowdloanAddedToNewRaise
	XcmPallet_Attempted                             []EventXcmPalletAttempted
	XcmPallet_Sent                                  []EventXcmPalletSent
}
type EventXcmPalletSent struct {
	Phase          types.Phase
	MultiLocation1 base.MultiLocation
	MultiLocation2 base.MultiLocation
	Xcm            base.Xcm
	Topics         []types.Hash
}
type EventXcmPalletAttempted struct {
	Phase   types.Phase
	Outcome base.Outcome
	Topics  []types.Hash
}
type EventCrowdloanAddedToNewRaise struct {
	Phase  types.Phase
	ParaId types.U32
	Topics []types.Hash
}
type EventCrowdloanMemoUpdated struct {
	Phase     types.Phase
	AccountId types.AccountID
	ParaId    types.U32
	VecU8     types.Bytes
	Topics    []types.Hash
}
type EventCrowdloanEdited struct {
	Phase  types.Phase
	ParaId types.U32
	Topics []types.Hash
}
type EventCrowdloanHandleBidResult struct {
	Phase          types.Phase
	ParaId         types.U32
	DispatchResult types.DispatchResult
	Topics         []types.Hash
}
type EventCrowdloanOnboarded struct {
	Phase   types.Phase
	ParaId1 types.U32
	ParaId2 types.U32
	Topics  []types.Hash
}
type EventCrowdloanDeployDataFixed struct {
	Phase  types.Phase
	ParaId types.U32
	Topics []types.Hash
}
type EventCrowdloanDissolved struct {
	Phase  types.Phase
	ParaId types.U32
	Topics []types.Hash
}
type EventCrowdloanAllRefunded struct {
	Phase  types.Phase
	ParaId types.U32
	Topics []types.Hash
}
type EventCrowdloanPartiallyRefunded struct {
	Phase  types.Phase
	ParaId types.U32
	Topics []types.Hash
}
type EventCrowdloanWithdrew struct {
	Phase     types.Phase
	AccountId types.AccountID
	ParaId    types.U32
	Balance   types.U128
	Topics    []types.Hash
}

type EventCrowdloanContributed struct {
	Phase     types.Phase
	AccountId types.AccountID
	ParaId    types.U32
	Balance   types.U128
	Topics    []types.Hash
}
type EventCrowdloanCreated struct {
	Phase  types.Phase
	ParaId types.U32
	Topics []types.Hash
}
type EventAuctionsWinningOffset struct {
	Phase        types.Phase
	AuctionIndex types.U32
	BlockNumber  types.BlockNumber
	Topics       []types.Hash
}
type EventAuctionsBidAccepted struct {
	Phase        types.Phase
	AccountId    types.AccountID
	ParaId       types.U32
	Balance      types.U128
	LeasePeriod1 types.BlockNumber
	LeasePeriod2 types.BlockNumber
	Topics       []types.Hash
}
type EventAuctionsReserveConfiscated struct {
	Phase     types.Phase
	ParaId    types.U32
	AccountId types.AccountID
	Balance   types.U128
	Topics    []types.Hash
}
type EventAuctionsUnreserved struct {
	Phase     types.Phase
	AccountId types.AccountID
	Balance   types.U128
	Topics    []types.Hash
}
type EventAuctionsReserved struct {
	Phase     types.Phase
	AccountId types.AccountID
	Balance1  types.U128
	Balance2  types.U128
	Topics    []types.Hash
}
type EventAuctionsWonRenewal struct {
	Phase        types.Phase
	ParaId       types.U32
	LeasePeriod1 types.BlockNumber
	LeasePeriod2 types.BlockNumber
	Balance      types.U128
	Topics       []types.Hash
}
type EventAuctionsWonDeploy struct {
	Phase     types.Phase
	AccountId types.AccountID
	SlotRange SlotRange
	ParaId    types.U32
	Balance   types.U128
	Topics    []types.Hash
}
type EventAuctionsAuctionClosed struct {
	Phase        types.Phase
	AuctionIndex types.U32
	Topics       []types.Hash
}
type EventAuctionsAuctionStarted struct {
	Phase        types.Phase
	AuctionIndex types.U32
	LeasePeriod  types.BlockNumber
	BlockNumber  types.BlockNumber
	Topics       []types.Hash
}
type EventSlotsLeased struct {
	Phase        types.Phase
	ParaId       types.U32
	AccountId    types.AccountID
	LeasePeriod1 types.BlockNumber
	LeasePeriod2 types.BlockNumber
	Balance1     types.U128
	Balance2     types.U128
	Topics       []types.Hash
}
type EventSlotsNewLeasePeriod struct {
	Phase       types.Phase
	LeasePeriod types.BlockNumber
	Topics      []types.Hash
}
type EventRegistrarReserved struct {
	Phase     types.Phase
	ParaId    types.U32
	AccountId types.AccountID
	Topics    []types.Hash
}
type EventRegistrarDeregistered struct {
	Phase  types.Phase
	ParaId types.U32
	Topics []types.Hash
}
type EventRegistrarRegistered struct {
	Phase     types.Phase
	ParaId    types.U32
	AccountId types.AccountID
	Topics    []types.Hash
}
type EventParasHrmpChannelClosed struct {
	Phase         types.Phase
	ParaId        types.U32
	HrmpChannelId base.HrmpChannelId
	Topics        []types.Hash
}
type EventParasHrmpOpenChannelAccepted struct {
	Phase   types.Phase
	ParaId1 types.U32
	ParaId2 types.U32
	Topics  []types.Hash
}
type EventParasHrmpOpenChannelRequested struct {
	Phase   types.Phase
	ParaId1 types.U32
	ParaId2 types.U32
	U32_1   types.U32
	U32_2   types.U32
	Topics  []types.Hash
}
type EventParasUmpUpwardMessagesReceived struct {
	Phase  types.Phase
	ParaId types.U32
	U32    types.U32
	U      types.U32
	Topics []types.Hash
}
type EventParasUmpWeightExhausted struct {
	Phase     types.Phase
	MessageId types.U256
	Weight1   types.Weight
	Weight2   types.Weight
	Topics    []types.Hash
}
type EventParasUmpExecutedUpward struct {
	Phase     types.Phase
	MessageId types.U256
	Outcome   base.Outcome
	Topics    []types.Hash
}
type EventParasUmpUnsupportedVersion struct {
	Phase     types.Phase
	MessageId types.U256
	Topics    []types.Hash
}
type EventParasUmpInvalidFormat struct {
	Phase     types.Phase
	MessageId types.U256 //[u8,32]
	Topics    []types.Hash
}
type EventParasActionQueued struct {
	Phase        types.Phase
	ParaId       types.U32
	SessionIndex types.U32
	Topics       []types.Hash
}
type EventParasNewHeadNoted struct {
	Phase  types.Phase
	ParaId types.U32
	Topics []types.Hash
}
type EventParasCodeUpgradeScheduled struct {
	Phase  types.Phase
	ParaId types.U32
	Topics []types.Hash
}
type EventParasCurrentHeadUpdated struct {
	Phase  types.Phase
	ParaId types.U32
	Topics []types.Hash
}
type EventParasCurrentCodeUpdated struct {
	Phase  types.Phase
	ParaId types.U32
	Topics []types.Hash
}
type EventParasInclusionCandidateTimedOut struct {
	Phase     types.Phase
	CRH       CandidateReceipt
	HeadData  types.Bytes
	CoreIndex types.U32
	Topics    []types.Hash
}
type EventParasInclusionCandidateIncluded struct {
	Phase      types.Phase
	CRH        CandidateReceipt
	HeadData   types.Bytes
	CoreIndex  types.U32
	GroupIndex types.U32
	Topics     []types.Hash
}
type EventParasInclusionCandidateBacked struct {
	Phase      types.Phase
	CRH        CandidateReceipt
	HeadData   types.Bytes
	CoreIndex  types.U32
	GroupIndex types.U32
	Topics     []types.Hash
}
type EventGiltGiltThawed struct {
	Phase       types.Phase
	ActiveIndex types.U32
	AccountId   types.AccountID
	BalanceOf1  types.U128
	BalanceOf2  types.U128
	Topics      []types.Hash
}
type EventGiltGiltIssued struct {
	Phase       types.Phase
	ActiveIndex types.U32
	BlockNumber types.BlockNumber
	AccountId   types.AccountID
	BalanceOf   types.U128
	Topics      []types.Hash
}
type EventGiltBidRetracted struct {
	Phase     types.Phase
	AccountId types.AccountID
	BalanceOf types.U128 // Balance
	U32       types.U32
	Topics    []types.Hash
}
type EventGiltBidPlaced struct {
	Phase     types.Phase
	AccountId types.AccountID
	BalanceOf types.U128
	U32       types.U32
	Topics    []types.Hash
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
type EventPhragmenElectionSeatHolderSlashed struct {
	Phase     types.Phase
	AccountId types.AccountID
	Balance   types.U128
	Topics    []types.Hash
}
type EventPhragmenElectionCandidateSlashed struct {
	Phase     types.Phase
	AccountId types.AccountID
	Balance   types.U128
	Topics    []types.Hash
}
type EventPhragmenElectionRenounced struct {
	Phase     types.Phase
	AccountId types.AccountID
	Topics    []types.Hash
}
type EventPhragmenElectionMemberKicked struct {
	Phase     types.Phase
	AccountId types.AccountID
	Topics    []types.Hash
}
type EventPhragmenElectionElectionError struct {
	Phase types.Phase

	Topics []types.Hash
}
type EventPhragmenElectionEmptyTerm struct {
	Phase types.Phase

	Topics []types.Hash
}
type EventPhragmenElectionNewTerm struct {
	Phase  types.Phase
	Vab    []TupleAccountBalance
	Topics []types.Hash
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
	Abs    []base.AccountBalance
	Topics []types.Hash
}

type EventElectionsPhragmenEmptyTerm struct {
	Phase types.Phase

	Topics []types.Hash
}

type TupleAccountBalance struct {
	AccountId types.AccountID
	Balance   types.U128
}

type SlotRange struct {
	Value string
}

func (s *SlotRange) Decode(decoder scale.Decoder) error {
	n, err := decoder.ReadOneByte()
	if err != nil {
		return err
	}
	switch n {
	case 0:
		s.Value = "ZeroZero"
	case 1:
		s.Value = "ZeroOne"
	case 2:
		s.Value = "ZeroTwo"
	case 3:
		s.Value = "ZeroThree"
	case 4:
		s.Value = "OneOne"
	case 5:
		s.Value = "OneTwo"
	case 6:
		s.Value = "OneThree"
	case 7:
		s.Value = "TwoTwo"
	case 8:
		s.Value = "TwoThree"
	case 9:
		s.Value = "ThreeThree"
	default:
		return fmt.Errorf("SlotRange,unknoww enum: %d", n)

	}
	return nil
}

type CandidateReceipt struct {
	Descriptor      CandidateDescriptor
	CommitmentsHash types.Hash
}

type CandidateDescriptor struct {
	ParaId                      types.U32
	RelayChainHash              types.Hash
	CollatorId                  types.H256
	PersistedValidationDataHash types.Hash
	PovHash                     types.Hash
	ErasureRoot                 types.Hash
	Signature                   types.Signature
	ParaHead                    types.Hash
	ValidationCodeHash          types.Hash
}
