package base

import (
	"fmt"
	"github.com/JFJun/go-substrate-rpc-client/v3/scale"
	"github.com/JFJun/go-substrate-rpc-client/v3/types"
	"math/big"
)

type Outcome struct {
	IsComplete   bool
	AsComplete   types.Weight
	isIncomplete bool
	asIncomplete Incomplete
	IsError      bool
	AsError      XcmError
}

func (o *Outcome) Decode(decoder scale.Decoder) error {
	n, err := decoder.ReadOneByte()
	if err != nil {
		return err
	}
	switch n {
	case 0:
		o.IsComplete = true
		return decoder.Decode(&o.AsComplete)
	case 1:
		o.isIncomplete = true
		return decoder.Decode(&o.asIncomplete)
	case 2:
		o.IsError = true
		return decoder.Decode(&o.AsError)
	default:
		return fmt.Errorf("Outcome,unknow enum: %d", n)
	}
}

type Incomplete struct {
	Weight   types.Weight
	XcmError XcmError
}
type XcmError struct {
	Value string
}

func (x *XcmError) Decode(decoder scale.Decoder) error {
	b, err := decoder.ReadOneByte()
	if err != nil {
		return err
	}
	switch b {
	case 0:
		x.Value = "Undefined"
	case 1:
		x.Value = "Overflow"
	case 2:
		x.Value = "Unimplemented"
	case 3:
		x.Value = "UnhandledXcmVersion"
	case 4:
		x.Value = "UnhandledXcmMessage"
	case 5:
		x.Value = "UnhandledEffect"
	case 6:
		x.Value = "EscalationOfPrivilege"
	case 7:
		x.Value = "UntrustedReserveLocation"
	case 8:
		x.Value = "UntrustedTeleportLocation"
	case 9:
		x.Value = "DestinationBufferOverflow"
	case 10:
		x.Value = "SendFailed"
	case 11:
		x.Value = "CannotReachDestination"
		// (MultiLocation, Xcm)
		var u CannotReachDestination
		return decoder.Decode(&u)
	case 12:
		x.Value = "MultiLocationFull"
	case 13:
		x.Value = "FailedToDecode"
	case 14:
		x.Value = "BadOrigin"
	case 15:
		x.Value = "ExceedsMaxMessageSize"
	case 16:
		x.Value = "FailedToTransactAsset"
	case 17:
		x.Value = "WeightLimitReached"
		var weight types.Weight
		return decoder.Decode(&weight)
	case 18:
		x.Value = "TooMuchWeightRequired"
	case 19:
		x.Value = "NotHoldingFees"
	case 20:
		x.Value = "WeightNotComputable"
	case 21:
		x.Value = "NotWithdrawable"
	case 22:
		x.Value = "LocationCannotHold"
	case 23:
		x.Value = "Undefined"
	case 24:
		x.Value = "Undefined"
	default:
		return fmt.Errorf("xcnError,unkonw enum :%d", b)
	}
	return nil
}

type CannotReachDestination struct {
	MultiLocation MultiLocation
	Xcm           Xcm
}

type MultiLocation struct {
	Value string
}

func (m *MultiLocation) Decode(decoder scale.Decoder) error {
	b, err := decoder.ReadOneByte()
	if err != nil {
		return err
	}
	switch b {
	case 0:
		m.Value = "Null"
	case 1:
		m.Value = "X1"
		var j Junction
		return decoder.Decode(&j)
	case 2:
		m.Value = "X2"
		// (Junction,Junction)
		for i := 0; i < 2; i++ {
			var j Junction
			if err = decoder.Decode(&j); err != nil {
				return err
			}
		}
	case 3:
		m.Value = "X3"
		// (Junction,Junction,Junction)
		for i := 0; i < 3; i++ {
			var j Junction
			if err = decoder.Decode(&j); err != nil {
				return err
			}
		}
	case 4:
		m.Value = "X4"
		for i := 0; i < 4; i++ {
			var j Junction
			if err = decoder.Decode(&j); err != nil {
				return err
			}
		}
	case 5:
		m.Value = "X5"
		for i := 0; i < 5; i++ {
			var j Junction
			if err = decoder.Decode(&j); err != nil {
				return err
			}
		}
	case 6:
		m.Value = "X6"
		for i := 0; i < 6; i++ {
			var j Junction
			if err = decoder.Decode(&j); err != nil {
				return err
			}
		}
	case 7:
		m.Value = "X7"
		for i := 0; i < 7; i++ {
			var j Junction
			if err = decoder.Decode(&j); err != nil {
				return err
			}
		}
	case 8:
		m.Value = "X8"
		for i := 0; i < 8; i++ {
			var j Junction
			if err = decoder.Decode(&j); err != nil {
				return err
			}
		}
	default:
		return fmt.Errorf("MultiLocation,unknow enum: %d", b)
	}
	return nil
}

type Junction struct {
	Value string
}

func (j *Junction) Decode(decoder scale.Decoder) error {
	b, err := decoder.ReadOneByte()
	if err != nil {
		return err
	}
	switch b {
	case 0:
		j.Value = "Parent"
	case 1:
		j.Value = "Parachain"
		// Compact<u32>
		var cu types.UCompact
		return decoder.Decode(&cu)
	case 2:
		j.Value = "AccountId32"
		//      "AccountId32Junction"
		var aij AccountId32Junction
		return decoder.Decode(&aij)
	case 3:
		j.Value = "AccountIndex64"
		var aij AccountIndex64Junction
		return decoder.Decode(&aij)
	case 4:
		j.Value = "AccountKey20"
		var aij AccountKey20Junction
		return decoder.Decode(&aij)
	case 5:
		j.Value = "PalletInstance"
		var u types.U8
		return decoder.Decode(&u)
	case 6:
		j.Value = "GeneralIndex"
		// Compact<u128>
		var cu types.UCompact
		return decoder.Decode(&cu)
	case 7:
		j.Value = "GeneralKey"
		// Vec<u8>
		var cu types.Bytes
		return decoder.Decode(&cu)
	case 8:
		j.Value = "OnlyChild"
	case 9:
		j.Value = "Plurality"
		// PluralityJunction
		var aij PluralityJunction
		return decoder.Decode(&aij)
	default:
		return fmt.Errorf("Junction,unknow enum: %d", b)
	}
	return nil
}

type AccountId32Junction struct {
	Network NetworkId
	Id      types.AccountID
}

type NetworkId struct {
	Value string
}

func (n *NetworkId) Decode(decoder scale.Decoder) error {
	b, err := decoder.ReadOneByte()
	if err != nil {
		return err
	}
	switch b {
	case 0:
		n.Value = "Any"
	case 1:
		n.Value = "Named"
		var u types.Bytes
		return decoder.Decode(&u)
	case 2:
		n.Value = "Polkadot"
	case 3:
		n.Value = "Kusama"
	default:
		return fmt.Errorf("NetworkId,unknow enum: %d", b)
	}
	return nil
}

type AccountIndex64Junction struct {
	Network NetworkId
	Index   types.UCompact
}

type AccountKey20Junction struct {
	Network NetworkId
	Index   U160 //[u8,20]
}
type U160 struct {
	*big.Int
}

// NewU128 creates a new U160 type
func NewU160(i big.Int) U160 {
	return U160{&i}
}

// Decode implements decoding as per the Scale specification
func (i *U160) Decode(decoder scale.Decoder) error {
	bs := make([]byte, 20)
	err := decoder.Read(bs)
	if err != nil {
		return err
	}
	// reverse bytes, scale uses little-endian encoding, big.int's bytes are expected in big-endian
	scale.Reverse(bs)

	b, err := types.UintBytesToBigInt(bs)
	if err != nil {
		return err
	}

	// deal with zero differently to get a nil representation (this is how big.Int deals with 0)
	if b.Sign() == 0 {
		*i = U160{big.NewInt(0)}
		return nil
	}

	*i = U160{b}
	return nil
}

type PluralityJunction struct {
	Id   BodyId
	Part BodyPart
}

type BodyId struct {
	Value string
}

func (b *BodyId) Decode(decoder scale.Decoder) error {
	n, err := decoder.ReadOneByte()
	if err != nil {
		return err
	}
	switch n {
	case 0:
		b.Value = "Unit"
	case 1:
		b.Value = "Named"
		var u types.Bytes
		return decoder.Decode(&u)
	case 2:
		b.Value = "Index"
		var u types.UCompact
		return decoder.Decode(&u)
	case 3:
		b.Value = "Executive"
	case 4:
		b.Value = "Technical"
	case 5:
		b.Value = "Legislative"
	case 6:
		b.Value = "Judicial"
	default:
		return fmt.Errorf("BodyId,unknow enum: %d", n)
	}
	return nil
}

type BodyPart struct {
	Value string
}

func (b *BodyPart) Decode(decoder scale.Decoder) error {
	n, err := decoder.ReadOneByte()
	if err != nil {
		return err
	}
	switch n {
	case 0:
		b.Value = "Voice"
	case 1:
		b.Value = "Members"
		var u types.UCompact
		return decoder.Decode(&u)
	case 2:
		b.Value = "Fraction"
		var u BodyPartFraction
		return decoder.Decode(&u)
	case 3:
		b.Value = "AtLeastProportion"
		var u BodyPartAtLeastProportion
		return decoder.Decode(&u)
	case 4:
		b.Value = "MoreThanProportion"
		var u BodyPartMoreThanProportion
		return decoder.Decode(&u)
	default:
		return fmt.Errorf("BodyPart,unknow enum: %d", n)

	}
	return nil
}

type BodyPartFraction struct {
	Nom   types.UCompact
	Denom types.UCompact
}

type BodyPartAtLeastProportion struct {
	Nom   types.UCompact
	Denom types.UCompact
}

type BodyPartMoreThanProportion struct {
	Nom   types.UCompact
	Denom types.UCompact
}

type Xcm struct {
	Value string
}

func (x *Xcm) Decode(decoder scale.Decoder) error {
	n, err := decoder.ReadOneByte()
	if err != nil {
		return err
	}
	switch n {
	case 0:
		x.Value = "WithdrawAsset"
		var u XcmWithdrawAsset
		return decoder.Decode(&u)
	case 1:
		x.Value = "ReserveAssetDeposit"
		var u XcmReserveAssetDeposit
		return decoder.Decode(&u)
	case 2:
		x.Value = "TeleportAsset"
		var u XcmTeleportAsset
		return decoder.Decode(&u)
	case 3:
		x.Value = "QueryResponse"
		var u XcmQueryResponse
		return decoder.Decode(&u)
	case 4:
		x.Value = "TransferAsset"
		var u XcmTransferAsset
		return decoder.Decode(&u)
	case 5:
		x.Value = "TransferReserveAsset"
		var u XcmTransferReserveAsset
		return decoder.Decode(&u)
	case 6:
		x.Value = "Transact"
		var u XcmTransact
		return decoder.Decode(&u)
	case 7:
		x.Value = "HrmpNewChannelOpenRequest"
		var u XcmHrmpNewChannelOpenRequest
		return decoder.Decode(&u)
	case 8:
		x.Value = "HrmpChannelAccepted"
		var u XcmHrmpChannelAccepted
		return decoder.Decode(&u)
	case 9:
		x.Value = "HrmpChannelClosing"
		var u XcmHrmpChannelClosing
		return decoder.Decode(&u)
	case 10:
		x.Value = "RelayedFrom"
		var u XcmRelayedFrom
		return decoder.Decode(&u)
	default:
		return fmt.Errorf("Xcm,unkonw enum: %d", n)
	}
}

type XcmRelayedFrom struct {
	Who     MultiLocation
	Message Xcm
}
type XcmHrmpChannelClosing struct {
	Initiator types.UCompact
	Sender    types.UCompact
	Recipient types.UCompact
}
type XcmHrmpChannelAccepted struct {
	Recipient types.UCompact
}
type XcmHrmpNewChannelOpenRequest struct {
	Sender         types.UCompact
	MaxMessageSize types.UCompact
	MaxCapacity    types.UCompact
}
type XcmTransact struct {
	OriginType          XcmOriginKind
	requireWeightAtMost types.U64
	Call                DoubleEncodedCall
}

type XcmOriginKind struct {
	Value string
}
type DoubleEncodedCall struct {
	Encoded types.Bytes
}

func (x *XcmOriginKind) Decode(decoder scale.Decoder) error {
	n, err := decoder.ReadOneByte()
	if err != nil {
		return err
	}
	switch n {
	case 0:
		x.Value = "Native"
	case 1:
		x.Value = "SovereignAccount"
	case 2:
		x.Value = "Superuser"
	case 3:
		x.Value = "Xcm"
	default:
		return fmt.Errorf("XcmOriginKind,unknown enum: %d", n)
	}
	return nil
}

type XcmTransferReserveAsset struct {
	Assets  []MultiAsset
	Dest    MultiLocation
	Effects []XcmOrder
}
type XcmTransferAsset struct {
	Assets []MultiAsset
	Dest   MultiLocation
}
type XcmQueryResponse struct {
	QueryId  types.UCompact
	Response XcmResponse
}

type XcmResponse struct {
	Value string
}

func (x *XcmResponse) Decode(decoder scale.Decoder) error {
	n, err := decoder.ReadOneByte()
	if err != nil {
		return err
	}
	switch n {
	case 0:
		x.Value = "Assets"
		var u []MultiAsset
		return decoder.Decode(&u)
	default:
		return fmt.Errorf("XcmResponse,unknown enum: %d", n)
	}
}

type XcmWithdrawAsset XcmAssetEffects
type XcmReserveAssetDeposit XcmAssetEffects
type XcmTeleportAsset XcmAssetEffects

type XcmAssetEffects struct {
	Assents []MultiAsset
	Effects []XcmOrder
}

type XcmOrder struct {
	Value string
}

func (x *XcmOrder) Decode(decoder scale.Decoder) error {
	n, err := decoder.ReadOneByte()
	if err != nil {
		return err
	}
	switch n {
	case 0:
		x.Value = "Null"
	case 1:
		x.Value = "DepositAsset"
		var u XcmOrderDepositAsset
		return decoder.Decode(&u)
	case 2:
		x.Value = "DepositReserveAsset"
		var u XcmOrderDepositReserveAsset
		return decoder.Decode(&u)
	case 3:
		x.Value = "ExchangeAsset"
		var u XcmOrderExchangeAsset
		return decoder.Decode(&u)
	case 4:
		x.Value = "InitiateReserveWithdraw"
		var u XcmOrderInitiateReserveWithdraw
		return decoder.Decode(&u)
	case 5:
		x.Value = "InitiateTeleport"
		var u XcmOrderInitiateTeleport
		return decoder.Decode(&u)
	case 6:
		x.Value = "QueryHolding"
		var u XcmOrderQueryHolding
		return decoder.Decode(&u)
	case 7:
		x.Value = "BuyExecution"
		var u XcmOrderBuyExecution
		return decoder.Decode(&u)
	default:
		return fmt.Errorf("XcmOrder,unkown enum:%d", n)
	}
	return nil
}

type XcmOrderBuyExecution struct {
	Fees        MultiAsset
	Weight      types.U64
	Debt        types.U64
	HaltOnError types.Bool
	Xcm         []Xcm
}
type XcmOrderQueryHolding struct {
	QueryId types.UCompact
	Dest    MultiLocation
	Assets  []MultiAsset
}
type XcmOrderInitiateTeleport struct {
	Assets  []MultiAsset
	Dest    MultiLocation
	Effects []XcmOrder
}
type XcmOrderInitiateReserveWithdraw struct {
	Assets  []MultiAsset
	Reserve MultiLocation
	Effects []XcmOrder
}
type XcmOrderExchangeAsset struct {
	Give    []MultiAsset
	Receive []MultiAsset
}
type XcmOrderDepositAsset struct {
	Assets []MultiAsset
	Dest   MultiLocation
}
type XcmOrderDepositReserveAsset struct {
	Assets  []MultiAsset
	Dest    MultiLocation
	Effects []XcmOrder
}

type MultiAsset struct {
	Value string
}

func (m *MultiAsset) Decode(decoder scale.Decoder) error {
	n, err := decoder.ReadOneByte()
	if err != nil {
		return err
	}
	switch n {
	case 0:
		m.Value = "Null"
	case 1:
		m.Value = "All"
	case 2:
		m.Value = "AllFungible"
	case 3:
		m.Value = "AllNonFungible"
	case 4:
		m.Value = "AllAbstractFungible"
		var u types.Bytes
		return decoder.Decode(&u)
	case 5:
		m.Value = "AllAbstractNonFungible"
		var u types.Bytes
		return decoder.Decode(&u)
	case 6:
		m.Value = "AllConcreteFungible"
		var u MultiLocation
		return decoder.Decode(&u)
	case 7:
		m.Value = "AllConcreteNonFungible"
		var u MultiLocation
		return decoder.Decode(&u)
	case 8:
		m.Value = "AbstractFungible"
		var u MultiAssetAbstractFungible
		return decoder.Decode(&u)
	case 9:
		m.Value = "AbstractNonFungible"
		var u MultiAssetAbstractNonFungible
		return decoder.Decode(&u)
	case 10:
		m.Value = "ConcreteFungible"
		var u MultiAssetConcreteFungible
		return decoder.Decode(&u)
	case 11:
		m.Value = "ConcreteNonFungible"
		var u MultiAssetConcreteNonFungible
		return decoder.Decode(&u)

	default:
		return fmt.Errorf("MultiAsset,unkown enum: %d", err)
	}
	return nil
}

type MultiAssetAbstractFungible struct {
	Id       types.Bytes
	Instance types.UCompact
}
type MultiAssetAbstractNonFungible struct {
	Class    types.Bytes
	Instance AssetInstance
}

type MultiAssetConcreteFungible struct {
	Id     MultiLocation
	Amount types.UCompact
}
type MultiAssetConcreteNonFungible struct {
	Class    types.Bytes
	Instance AssetInstance
}
type AssetInstance struct {
	Value string
}

func (a *AssetInstance) Decode(decoder scale.Decoder) error {
	n, err := decoder.ReadOneByte()
	if err != nil {
		return err
	}
	switch n {
	case 0:
		a.Value = "Undefined"
	case 1:
		a.Value = "Index8"
		var u types.U8
		return decoder.Decode(&u)
	case 2:
		a.Value = "Index16"
		fallthrough
	case 3:
		a.Value = "Index32"
		fallthrough
	case 4:
		a.Value = "Index64"
		fallthrough
	case 5:
		if a.Value == "" {
			a.Value = "Index128"
		}
		var u types.UCompact
		return decoder.Decode(&u)
	case 6:
		a.Value = "Array4"
		// [u8,4]
		var u types.U32
		return decoder.Decode(&u)
	case 7:
		a.Value = "Array8"
		// [u8,4]
		var u types.U64
		return decoder.Decode(&u)
	case 8:
		a.Value = "Array16"
		// [u8,4]
		var u types.U128
		return decoder.Decode(&u)
	case 9:
		a.Value = "Array132"
		// [u8,4]
		var u types.U256
		return decoder.Decode(&u)
	case 10:
		a.Value = "Blob"
		// [u8,4]
		var u types.Bytes
		return decoder.Decode(&u)
	default:
		return fmt.Errorf("AssetInstance,unkown enum: %d", err)
	}
	return nil
}
