package expand

import (
	"bytes"
	"encoding/binary"
	"errors"
	"fmt"
	"github.com/JFJun/bifrost-go/uint128"
	"github.com/JFJun/bifrost-go/utils"
	"github.com/huandu/xstrings"
	"github.com/shopspring/decimal"
	"github.com/stafiprotocol/go-substrate-rpc-client/scale"
	"github.com/stafiprotocol/go-substrate-rpc-client/types"
	"io"
	"math/big"
	"reflect"
)

type Balance struct {
	Reader io.Reader
	Value  decimal.Decimal
}

func (b *Balance) Decode(decoder scale.Decoder) error {
	buf := &bytes.Buffer{}
	b.Reader = buf
	data := make([]byte, 16)
	err := decoder.Read(data)
	if err != nil {
		return fmt.Errorf("decode balance: read bytes error: %v", err)
	}
	buf.Write(data)
	c := make([]byte, 16)
	if utils.BytesToHex(c) == "ffffffffffffffffffffffffffffffff" {
		b.Value = decimal.Zero
		return nil
	}

	b.Value = decimal.NewFromBigInt(uint128.FromBytes(c).Big(), 0)
	return nil
}

type Vec struct {
	Value []interface{}
}

/*
sub type must be struct,not ptr
*/
func (v *Vec) ProcessVec(decoder scale.Decoder, subType interface{}) error {
	var u types.UCompact
	err := decoder.Decode(&u)
	if err != nil {
		return fmt.Errorf("decode Vec: get length error: %v", err)
	}
	length := int(utils.UCompactToBigInt(u).Int64())
	if length > 5000 {
		return fmt.Errorf("vec length %d exceeds %d", length, 1000)
	}
	for i := 0; i < length; i++ {
		st := reflect.TypeOf(subType)
		if st.Kind() != reflect.Struct {
			return errors.New("decode Vec: struct type is not struct")
		}
		tmp := reflect.New(st)
		subType := tmp.Interface()
		err = decoder.Decode(subType)
		if err != nil {
			return fmt.Errorf("decode Vec: decoder subtype error: %v", err)
		}
		v.Value = append(v.Value, subType)
	}
	return nil
}

/*
解码包的问题，所以这里只能根据需求写死
*/
type TransferCall struct {
	Value interface{}
}

func (t *TransferCall) Decode(decoder scale.Decoder) error {
	//1. 先获取callidx
	b := make([]byte, 2)
	err := decoder.Read(b)
	if err != nil {
		return fmt.Errorf("deode transfer call: read callIdx bytes error: %v", err)
	}
	callIdx := xstrings.RightJustify(utils.IntToHex(b[0]), 2, "0") + xstrings.RightJustify(utils.IntToHex(b[1]), 2, "0")
	result := map[string]interface{}{
		"call_index": callIdx,
	}
	var param []ExtrinsicParam
	// 0 ---> 	Address
	var address MultiAddress
	err = decoder.Decode(&address)
	if err != nil {
		return fmt.Errorf("decode call: decode Balances.transfer.Address error: %v", err)
	}
	param = append(param,
		ExtrinsicParam{
			Name:     "dest",
			Type:     "Address",
			Value:    utils.BytesToHex(address.AccountId[:]),
			ValueRaw: utils.BytesToHex(address.AccountId[:]),
		})
	// 1 ----> Compact<Balance>
	var bb types.UCompact

	err = decoder.Decode(&bb)
	if err != nil {
		return fmt.Errorf("decode call: decode Balances.transfer.Compact<Balance> error: %v", err)
	}
	v := utils.UCompactToBigInt(bb).Int64()
	param = append(param,
		ExtrinsicParam{
			Name:  "value",
			Type:  "Compact<Balance>",
			Value: v,
		})
	result["call_args"] = param
	t.Value = result
	return nil
}

type Address struct {
	AccountLength string `json:"account_length"`
	Value         string
}

func (a *Address) Decode(decoder scale.Decoder) error {
	al, err := decoder.ReadOneByte()
	if err != nil {
		return fmt.Errorf("decode address: get account length error: %v", err)
	}
	a.AccountLength = utils.BytesToHex([]byte{al})
	if a.AccountLength == "ff" && defaultSerDeOptions.SerDe.NoPalletIndices == false {
		data := make([]byte, 32)
		err = decoder.Read(data)
		if err != nil {
			return fmt.Errorf("decode address: get address 32 bytes error: %v", err)
		}
		a.Value = utils.BytesToHex(data)
		return nil
	}
	d := make([]byte, 31)
	err = decoder.Read(d)
	if err != nil {
		return fmt.Errorf("decode address: get address 31 bytes error: %v", err)
	}
	a.Value = utils.BytesToHex(append([]byte{al}, d...))
	return nil
}

//
type U32 struct {
	Value uint32
}

func (u *U32) Decode(decoder scale.Decoder) error {
	data := make([]byte, 4)
	err := decoder.Read(data)
	if err != nil {
		return fmt.Errorf("decode u32 : read 4 bytes error: %v", err)
	}
	u.Value = binary.LittleEndian.Uint32(data)
	return nil
}

// AccountInfo contains information of an account
type StafiAccountInfo struct {
	Nonce    types.U32
	RefCount types.U8
	Data     struct {
		Free       types.U128
		Reserved   types.U128
		MiscFrozen types.U128
		FreeFrozen types.U128
	}
}

type WeightToFeeCoefficient struct {
	CoeffInteger types.U128
	CoeffFrac    U32
	Negative     bool
	Degree       types.U8
}

func (d *WeightToFeeCoefficient) Decode(decoder scale.Decoder) error {
	err := decoder.Decode(&d.CoeffInteger)
	if err != nil {
		return fmt.Errorf("decode WeightToFeeCoefficient: decode CoeffInteger error: %v", err)
	}
	err = decoder.Decode(&d.CoeffFrac)
	if err != nil {
		return fmt.Errorf("decode WeightToFeeCoefficient: decode CoeffFrac error: %v", err)
	}
	err = decoder.Decode(&d.Negative)
	if err != nil {
		return fmt.Errorf("decode WeightToFeeCoefficient: decode Negative error: %v", err)
	}
	err = decoder.Decode(&d.Degree)
	if err != nil {
		return fmt.Errorf("decode WeightToFeeCoefficient: decode Degree error: %v", err)
	}
	return nil
}

/*
https://github.com/polkadot-js/api/blob/0e52e8fe23ed029b8101cf8bf82deccd5aa7a790/packages/types/src/generic/MultiAddress.ts
*/

type MultiAddress struct {
	GenericMultiAddress
}
type GenericMultiAddress struct {
	AccountId types.AccountID
	Index     types.UCompact
	Raw       types.Bytes
	Address32 types.H256
	Address20 types.H160
	types     int
}

func (d *GenericMultiAddress) Decode(decoder scale.Decoder) error {
	b, err := decoder.ReadOneByte()
	if err != nil {
		return fmt.Errorf("generic MultiAddress read on bytes error: %v", err)
	}
	switch int(b) {
	case 0:
		err = decoder.Decode(&d.AccountId)
	case 1:
		err = decoder.Decode(&d.Index)
	case 2:
		err = decoder.Decode(&d.Address32)
	case 3:
		err = decoder.Decode(&d.Address20)
	default:
		err = fmt.Errorf("generic MultiAddress unsupport type=%d ", b)
	}
	if err != nil {
		return err
	}
	d.types = int(b)
	return nil
}

func (d GenericMultiAddress) Encode(encoder scale.Encoder) error {
	t := types.NewU8(uint8(d.types))
	err := encoder.Encode(t)
	if err != nil {
		return err
	}
	switch d.types {
	case 0:
		if &d.AccountId == nil {
			err = fmt.Errorf("generic MultiAddress id is null:%v", d.AccountId)
		}
		err = encoder.Encode(d.AccountId)
	case 1:
		if &d.Index == nil {
			err = fmt.Errorf("generic MultiAddress index is null:%v", d.Index)
		}
		err = encoder.Encode(d.Index)
	case 2:
		if &d.Address32 == nil {
			err = fmt.Errorf("generic MultiAddress address32 is null:%v", d.Address32)
		}
		err = encoder.Encode(d.Address32)
	case 3:
		if &d.Address20 == nil {
			err = fmt.Errorf("generic MultiAddress address20 is null:%v", d.Address20)
		}
		err = encoder.Encode(d.Address20)
	default:
		err = fmt.Errorf("generic MultiAddress unsupport this types: %d", d.types)
	}
	if err != nil {
		return err
	}
	return nil
}

/*
没办法，底层解析就是这样做的，只能这样写了，虽然很不友好
*/
func (d *GenericMultiAddress) GetTypes() int {
	return d.types
}
func (d *GenericMultiAddress) SetTypes(types int) {
	d.types = types
}
func (d *GenericMultiAddress) GetAccountId() types.AccountID {
	return d.AccountId
}
func (d *GenericMultiAddress) GetIndex() types.UCompact {
	return d.Index
}
func (d *GenericMultiAddress) GetAddress32() types.H256 {
	return d.Address32
}
func (d *GenericMultiAddress) GetAddress20() types.H160 {
	return d.Address20
}

func (d *GenericMultiAddress) ToAddress() types.Address {
	if d.types != 0 {
		return types.Address{}
	}
	var ai []byte
	ai = append([]byte{0x00}, d.AccountId[:]...)
	return types.NewAddressFromAccountID(ai)
}

type ExtrinsicSignatureV4 struct {
	Signer    MultiAddress
	Signature types.MultiSignature
	Era       types.ExtrinsicEra // extra via system::CheckEra
	Nonce     types.UCompact     // extra via system::CheckNonce (Compact<Index> where Index is u32))
	Tip       types.UCompact     // extra via balances::TakeFees (Compact<Balance> where Balance is u128))
}

/*
https://github.com/polkadot-js/api/packages/types/src/interfaces/system/types.ts p13
*/
type AccountInfoWithProviders struct {
	Nonce       types.U32 `json:"nonce"`
	Consumers   types.U32 `json:"consumers"`
	Providers   types.U32 `json:"providers"`
	Sufficients types.U32 `json:"sufficients"`
	Data        struct {
		Free       types.U128 `json:"free"`
		Reserved   types.U128 `json:"reserved"`
		MiscFrozen types.U128 `json:"misc_frozen"`
		FreeFrozen types.U128 `json:"free_frozen"`
	} `json:"data"`
}

type FeeDetail struct {
	BaseFee           types.U128
	LenFee            types.U128
	AdjustedWeightFee types.U128
}

func (this *FeeDetail) FeeTotal() types.U128 {
	var result = big.NewInt(0)
	result = result.Add(this.BaseFee.Int, this.LenFee.Int)
	result = result.Add(result, this.AdjustedWeightFee.Int)

	return types.NewU128(*result)
}
