package test

import (
	"bytes"
	"encoding/hex"
	"fmt"
	"github.com/JFJun/go-substrate-rpc-client/v3/scale"
	"github.com/JFJun/go-substrate-rpc-client/v3/types"
	"regexp"
	"testing"
)

func Test_DecodeU32(t *testing.T) {
	src := "20df0e0b00000000"
	data, _ := hex.DecodeString(src)
	var u types.U32
	decoder := scale.NewDecoder(bytes.NewReader(data))
	decoder.Decode(&u)
	fmt.Println(u)
}

func Test_DecodeU128(t *testing.T) {
	s := "Option<U8>"
	reg := regexp.MustCompile("^([^<]*)<(.+)>$")
	typeParts := reg.FindStringSubmatch(s)
	fmt.Println(typeParts)
}

func Test_DecodeBytes(t *testing.T) {
	src := "3dce1a3afabc5b656fd95278ac36b2b989657a6f7fce27463dfff849169e6824fcf77029d631090765ea964cfd2589b9b4bdb483c48dc2461ced321d6709bf46292b1f93260025d4ec59d9686aa3bf98587299a2d5ee0d312a568e87b6f0d05e4752b42821a3b024b9e6a4d55386e161e1248c24d2f357df4fb00e0ae2f42dac652df412ec2f60e8d16038e11ab04e27ae8d52740fcea88521efa8063b8f93205d0f366dc052783efef25f54f76c2a209488d6b55327aee9fa12ee147137d9b8d9610e85ce36c641dd56f03abd4ce651e6252c486dc057f4ea0fcd9473d9012e5e53319a20162c710c4c02ab787717f755ca3cf3b23d433a9104f711f91ca1960ec8d6bd570c093cb4f785021c5a398f5aa7f23e35eea1093339a3af70912e0937ecd1bae902d871cfc37e45767866e3ede6cbc805a8245efa05c88087a46f2f92ca913500b76a1106001fd8a6779d040c899e00af53ad6df694dd5d60c951797663e9c345216f2d8a2c151264892879e2c51169a7c9d914d98990705f33252565a9dc5e0068e17f3c2a08066175726120e40410080000000005617572610101142fd1c4edfa0d87435853b55a269b2c5eb052598d2f10e61ce31b84458aa74bf74555c4d8784244971635a0797fe7cb8b23d46ba81cdbd4aacd4e61d0b39681"
	data, _ := hex.DecodeString(src)
	fmt.Println(len(data))
	// 3*16+13 = 48+13 = 61
	decoder := scale.NewDecoder(bytes.NewReader(data))
	var b types.Bytes
	decoder.Decode(&b)
	fmt.Println(len(b))
	fmt.Println(hex.EncodeToString(b))
}
