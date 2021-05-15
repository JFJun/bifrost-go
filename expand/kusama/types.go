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
	Staking_Kicked                                  []EventStakingKicked
	Democracy_Blacklisted                           []EventDemocracyBlacklisted
	ElectionsPhragmen_NewTerm                       []EventElectionsPhragmenNewTerm
	ElectionsPhragmen_EmptyTerm                     []EventElectionsPhragmenEmptyTerm
	ElectionsPhragmen_ElectionError                 []EventElectionsPhragmenElectionError
	ElectionsPhragmen_MemberKicked                  []EventElectionsPhragmenMemberKicked
	ElectionsPhragmen_Renounced                     []EventElectionsPhragmenRenounced
	ElectionsPhragmen_CandidateSlashed              []EventElectionsPhragmenCandidateSlashed
	ElectionsPhragmen_SeatHolderSlashed             []EventElectionsPhragmenSeatHolderSlashed
	ElectionProviderMultiPhase_SolutionStored       []ElectionProviderMultiPhaseSolutionStored
	ElectionProviderMultiPhase_UnsignedPhaseStarted []EventElectionProviderMultiPhaseUnsignedPhaseStarted
	ElectionProviderMultiPhase_Rewarded             []EventElectionProviderMultiPhaseRewarded
	ElectionProviderMultiPhase_Slashed              []EventElectionProviderMultiPhaseSlashed
	ElectionProviderMultiPhase_SignedPhaseStarted   []EventElectionProviderMultiPhaseSignedPhaseStarted
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
	PhragmenElection_NewTerm                        []EventPhragmenElectionNewTerm
	PhragmenElection_EmptyTerm                      []EventPhragmenElectionEmptyTerm
	PhragmenElection_ElectionError                  []EventPhragmenElectionElectionError
	PhragmenElection_MemberKicked                   []EventPhragmenElectionMemberKicked
	PhragmenElection_Renounced                      []EventPhragmenElectionRenounced
	PhragmenElection_CandidateSlashed               []EventPhragmenElectionCandidateSlashed
	PhragmenElection_SeatHolderSlashed              []EventPhragmenElectionSeatHolderSlashed
	ElectionProviderMultiPhase_ElectionFinalized    []EventElectionProviderMultiPhaseElectionFinalized
	Gilt_BidPlaced                                  []EventGiltBidPlaced
	Gilt_BidRetracted                               []EventGiltBidRetracted
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
type EventElectionProviderMultiPhaseUnsignedPhaseStarted struct {
	Phase       types.Phase
	BountyIndex types.U32
	Topics      []types.Hash
}
type ElectionProviderMultiPhaseSolutionStored struct {
	Phase           types.Phase
	ElectionCompute types.ElectionCompute
	Topics          []types.Hash
}

type EventElectionProviderMultiPhaseElectionFinalized struct {
	Phase           types.Phase
	ElectionCompute types.ElectionCompute
	Topics          []types.Hash
}
type EventElectionProviderMultiPhaseRewarded struct {
	Phase     types.Phase
	AccountId types.AccountID
	Topics    []types.Hash
}
type EventElectionProviderMultiPhaseSlashed struct {
	Phase     types.Phase
	AccountId types.AccountID
	Topics    []types.Hash
}
type EventElectionProviderMultiPhaseSignedPhaseStarted struct {
	Phase    types.Phase
	RoundVal types.U32
	Topics   []types.Hash
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

type EventPhragmenElectionNewTerm struct {
	Phase    types.Phase
	TermList []PhragmenElectionNewTermItem
	Topics   []types.Hash
}

type PhragmenElectionNewTermItem struct {
	AccountId types.AccountID
	Balance   types.U128
}
type EventPhragmenElectionEmptyTerm struct {
	Phase  types.Phase
	Topics []types.Hash
}
type EventPhragmenElectionElectionError struct {
	Phase  types.Phase
	Topics []types.Hash
}
type EventPhragmenElectionMemberKicked struct {
	Phase     types.Phase
	AccountId types.AccountID
	Topics    []types.Hash
}
type EventPhragmenElectionRenounced struct {
	Phase     types.Phase
	AccountId types.AccountID
	Topics    []types.Hash
}
type EventPhragmenElectionCandidateSlashed struct {
	Phase     types.Phase
	AccountId types.AccountID
	Balance   types.U128
	Topics    []types.Hash
}
type EventPhragmenElectionSeatHolderSlashed struct {
	Phase     types.Phase
	AccountId types.AccountID
	Balance   types.U128
	Topics    []types.Hash
}

type EventGiltBidPlaced struct {
	Phase     types.Phase
	AccountId types.AccountID
	Balance   types.U128
	Duration  types.U32
	Topics    []types.Hash
}
type EventGiltBidRetracted struct {
	Phase     types.Phase
	AccountId types.AccountID
	Balance   types.U128
	Duration  types.U32
	Topics    []types.Hash
}
type EventGiltGiltIssued struct {
	Phase       types.Phase
	ActiveIndex types.U32
	BlockNumber types.BlockNumber
	AccountId   types.AccountID
	Balance     types.U128
	Topics      []types.Hash
}
type EventGiltGiltThawed struct {
	Phase            types.Phase
	ActiveIndex      types.U32
	AccountId        types.AccountID
	OriginalAmount   types.U128
	AdditionalAmount types.U128
	Duration         types.U32
	Topics           []types.Hash
}
type CandidateReceipt struct {
	Descriptor      CandidateDescriptor
	CommitmentsHash types.Hash
}
type CandidateDescriptor struct {
	/// The ID of the para this is a candidate for.
	ParaId types.U32
	/// The hash of the relay-chain block this is executed in the context of.
	RelayParent types.Hash
	/// The collator's sr25519 public key.
	Collator types.Hash
	/// The blake2-256 hash of the persisted validation data. This is extra data derived from
	/// relay-chain state which may vary based on bitfields included before the candidate.
	/// Thus it cannot be derived entirely from the relay-parent.
	PersistedValidationDataHash types.Hash
	/// The blake2-256 hash of the pov.
	PovHash types.Hash
	/// The root of a block's erasure encoding Merkle tree.
	ErasureRoot types.Hash
	/// Signature on blake2-256 of components of this receipt:
	/// The parachain index, the relay parent, the validation data hash, and the pov_hash.
	Signature types.Hash
	/// Hash of the para header that is being generated by this candidate.
	ParaHead types.Hash
	/// The blake2-256 hash of the validation code bytes.
	ValidationCodeHash types.Hash
}
type EventParasInclusionCandidateBacked struct {
	Phase            types.Phase
	CandidateReceipt CandidateReceipt
	HeadData         types.Bytes
	CoreIndex        types.U32
	GroupIndex       types.U32
	Topics           []types.Hash
}
type EventParasInclusionCandidateIncluded struct {
	Phase            types.Phase
	CandidateReceipt CandidateReceipt
	HeadData         types.Bytes
	CoreIndex        types.U32
	GroupIndex       types.U32
	Topics           []types.Hash
}
type EventParasInclusionCandidateTimedOut struct {
	Phase            types.Phase
	CandidateReceipt CandidateReceipt
	HeadData         types.Bytes
	CoreIndex        types.U32
	Topics           []types.Hash
}

type EventParasCurrentCodeUpdated struct {
	Phase  types.Phase
	ParaId ParaId
	Topics []types.Hash
}
type EventParasCurrentHeadUpdated struct {
	Phase  types.Phase
	ParaId ParaId
	Topics []types.Hash
}
type EventParasCodeUpgradeScheduled struct {
	Phase  types.Phase
	ParaId ParaId
	Topics []types.Hash
}
type EventParasNewHeadNoted struct {
	Phase  types.Phase
	ParaId ParaId
	Topics []types.Hash
}
type EventParasActionQueued struct {
	Phase  types.Phase
	ParaId ParaId
	Topics []types.Hash
}
type ParaId types.U32
type EventParasHrmpOpenChannelRequested struct {
	Phase                  types.Phase
	Sender                 ParaId
	Recipient              ParaId
	ProposedMaxCapacity    types.U32
	ProposedMaxMessageSize types.U32
	Topics                 []types.Hash
}
type EventParasHrmpOpenChannelAccepted struct {
	Phase     types.Phase
	Sender    ParaId
	Recipient ParaId
	Topics    []types.Hash
}
type EventParasHrmpChannelClosed struct {
	Phase       types.Phase
	ByParachain ParaId
	ChannelId   HrmpChannelId
	Topics      []types.Hash
}
type HrmpChannelId struct {
	Sender    ParaId
	Recipient ParaId
}
type EventRegistrarRegistered struct {
	Phase     types.Phase
	ParaId    ParaId
	AccountId types.AccountID
	Topics    []types.Hash
}
type EventRegistrarDeregistered struct {
	Phase  types.Phase
	ParaId ParaId
	Topics []types.Hash
}
type EventRegistrarReserved struct {
	Phase     types.Phase
	ParaId    ParaId
	AccountId types.AccountID
	Topics    []types.Hash
}

type LeasePeriod types.BlockNumber
type Balance types.U128
type EventSlotsNewLeasePeriod struct {
	Phase       types.Phase
	LeasePeriod LeasePeriod
	Topics      []types.Hash
}
type EventSlotsLeased struct {
	Phase        types.Phase
	ParachainId  ParaId
	Leaser       types.AccountID
	PeriodBegin  LeasePeriod
	PeriodCount  LeasePeriod
	ExtraReseved Balance
	TotalAmount  Balance
	Topics       []types.Hash
}
type AuctionIndex types.U32
type EventAuctionsAuctionStarted struct {
	Phase        types.Phase
	AuctionIndex AuctionIndex
	LeasePeriod  LeasePeriod
	Ending       types.BlockNumber
	Topics       []types.Hash
}
type EventAuctionsAuctionClosed struct {
	Phase        types.Phase
	AuctionIndex AuctionIndex
	Topics       []types.Hash
}
type SlotRange types.U8
type EventAuctionsWonDeploy struct {
	Phase       types.Phase
	Bidder      types.AccountID
	Range       SlotRange
	ParachainId ParaId
	Amount      Balance
	Topics      []types.Hash
}
type EventAuctionsWonRenewal struct {
	Phase       types.Phase
	ParachainId ParaId
	Begin       LeasePeriod
	Count       LeasePeriod
	TotalAmount Balance
	Topics      []types.Hash
}
type EventAuctionsReserved struct {
	Phase         types.Phase
	Bidder        types.AccountID
	ExtraReserved Balance
	TotalAmount   Balance
	Topics        []types.Hash
}
type EventAuctionsUnreserved struct {
	Phase  types.Phase
	Bidder types.AccountID
	Amount Balance
	Topics []types.Hash
}
type EventAuctionsReserveConfiscated struct {
	Phase       types.Phase
	ParachainId ParaId
	Leaser      types.AccountID
	Amount      Balance
	Topics      []types.Hash
}
type EventAuctionsBidAccepted struct {
	Phase     types.Phase
	Who       types.AccountID
	ParaId    ParaId
	Amount    Balance
	FirstSlot LeasePeriod
	LastSlot  LeasePeriod
	Topics    []types.Hash
}
type EventAuctionsWinningOffset struct {
	Phase        types.Phase
	AuctionIndex AuctionIndex
	BlockNumber  types.BlockNumber
	Topics       []types.Hash
}
type EventCrowdloanCreated struct {
	Phase     types.Phase
	FundIndex ParaId
	Topics    []types.Hash
}
type EventCrowdloanContributed struct {
	Phase     types.Phase
	Who       types.AccountID
	FundIndex ParaId
	Amount    Balance
	Topics    []types.Hash
}
type EventCrowdloanWithdrew struct {
	Phase     types.Phase
	Who       types.AccountID
	FundIndex ParaId
	Amount    Balance
	Topics    []types.Hash
}
type EventCrowdloanPartiallyRefunded struct {
	Phase     types.Phase
	FundIndex ParaId
	Topics    []types.Hash
}
type EventCrowdloanAllRefunded struct {
	Phase     types.Phase
	FundIndex ParaId
	Topics    []types.Hash
}
type EventCrowdloanDissolved struct {
	Phase     types.Phase
	FundIndex ParaId
	Topics    []types.Hash
}
type EventCrowdloanDeployDataFixed struct {
	Phase     types.Phase
	FundIndex ParaId
	Topics    []types.Hash
}
type EventCrowdloanOnboarded struct {
	Phase       types.Phase
	FundIndex   ParaId
	ParachainId ParaId
	Topics      []types.Hash
}
type DispatchResult types.U8
type EventCrowdloanHandleBidResult struct {
	Phase          types.Phase
	FundIndex      ParaId
	DispatchResult DispatchResult
	Topics         []types.Hash
}
type EventCrowdloanEdited struct {
	Phase     types.Phase
	FundIndex ParaId
	Topics    []types.Hash
}
type EventCrowdloanMemoUpdated struct {
	Phase     types.Phase
	Who       types.AccountID
	FundIndex ParaId
	Memo      types.Bytes
	Topics    []types.Hash
}
type EventCrowdloanAddedToNewRaise struct {
	Phase     types.Phase
	Parachain ParaId
	Topics    []types.Hash
}
type EventXcmPalletAttempted struct {
	Phase   types.Phase
	Outcome types.U8
	Topics  []types.Hash
}

type MultiLocation types.U8
type EventXcmPalletSent struct {
	Phase          types.Phase
	MultiLocation1 MultiLocation
	MultiLocation2 MultiLocation
	Xcm            types.U8
	Topics         []types.Hash
}
