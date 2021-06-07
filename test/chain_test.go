package test

import (
	"fmt"
	"github.com/JFJun/bifrost-go/client"
	"github.com/JFJun/bifrost-go/expand"
	"github.com/JFJun/bifrost-go/expand/base"
	"github.com/JFJun/bifrost-go/expand/polkadot"
	"github.com/JFJun/go-substrate-rpc-client/v3/types"
	"reflect"
	"strings"
	"testing"
)

/*
公共节点： https://polkadot.js.org/apps/#/explorer, wss://api.crust.network/

*/
func Test_Chain(t *testing.T) {
	c, err := client.New("wss://crab-rpc.darwinia.network/")
	if err != nil {
		t.Fatal(err)
	}

	b := polkadot.PolkadotEventRecords{}
	existMap := getEventTypesFieldName(b)

	fmt.Println(c.ChainName)
	for _, mod := range c.Meta.AsMetadataV12.Modules {
		if mod.HasEvents {
			for _, event := range mod.Events {
				typeName := fmt.Sprintf("%s_%s", mod.Name, event.Name)
				if IsExist(typeName, existMap) {
					continue
				}
				fmt.Printf("%s		[]Event%s%s\n", typeName, mod.Name, event.Name)
				if len(event.Args) == 0 {
					fmt.Printf("type Event%s%s struct { \n	Phase    types.Phase\n	\n	Topics []types.Hash\n}\n", mod.Name, event.Name)
				} else {
					as := ""
					for _, arg := range event.Args {
						s := fmt.Sprintf("	%s    types.\n", arg)
						as = as + s
					}
					fmt.Printf("type Event%s%s struct { \n	Phase    types.Phase\n%v	Topics []types.Hash\n}\n", mod.Name, event.Name, as)
				}

				//fmt.Println(event.Args)
				fmt.Println("------------------------------------------------")
			}
		}
	}
}

func Test_New(t *testing.T) {
	//b := polkadot.PolkadotEventRecords{}
	////tp:=reflect.TypeOf(b)
	////fmt.Println(tp.NumField())
	////fmt.Println(tp.Field(1).Name)
	//getEventTypesFieldName(b)

	s := "sp_std::marker::PhantomData<(AccountId, Event)>"
	namespaceSlice := strings.Split(s, "::")
	fmt.Println(len(namespaceSlice))
	fmt.Println(namespaceSlice)
}

func getEventTypesFieldName(ier expand.IEventRecords) []string {
	var existMap []string
	//first types.EventRecords
	te := types.EventRecords{}
	tep := reflect.TypeOf(te)
	for i := 0; i < tep.NumField(); i++ {
		existMap = append(existMap, tep.Field(i).Name)
	}

	// second
	be := base.BaseEventRecords{}
	bep := reflect.TypeOf(be)
	for i := 0; i < bep.NumField(); i++ {
		if bep.Field(i).Name == "EventRecords" {
			continue
		}
		existMap = append(existMap, tep.Field(i).Name)
	}
	// third parse IEventRecords
	ierp := reflect.TypeOf(ier)
	for i := 0; i < ierp.NumField(); i++ {
		if ierp.Field(i).Name == "BaseEventRecords" {
			continue
		}

		existMap = append(existMap, ierp.Field(i).Name)
	}
	return existMap
}

func IsExist(typeName string, existTypes []string) bool {
	for _, v := range existTypes {
		if typeName == v {
			return true
		}
	}
	return false
}

func Test_GetAllType(t *testing.T) {
	c, err := client.New("wss://crab-rpc.darwinia.network/")
	if err != nil {
		t.Fatal(err)
	}
	var tmpLists []string
	for _, mod := range c.Meta.AsMetadataV12.Modules {

		if mod.HasEvents {
			for _, event := range mod.Events {

				if len(event.Args) == 0 {
					//fmt.Printf("type Event%s%s struct { \n	Phase    types.Phase\n	\n	Topics []types.Hash\n}\n", mod.Name, event.Name)
				} else {
					//as:=""
					for _, arg := range event.Args {
						norPrint := false
						for _, as := range tmpLists {
							if as == string(arg) {
								norPrint = true
							}
						}
						if norPrint {
							continue
						}
						tmpLists = append(tmpLists, string(arg))
						fmt.Println(arg)
						//fmt.Println(ConvertType(string(arg)))
						//fmt.Println("========================================")

					}

				}
				//fmt.Println(event.Args)
			}
		}
	}

}

type ChainTypes struct {
	ChainName string                   `json:"chain_name"`
	Modules   []map[string]interface{} `json:"modules"`
}

/*

{
	"ModName":{
		"type":"string,struct,enum",
		"value":map[string]string,
	}
}


*/
func ConvertType(name string) string {

	name = strings.TrimSpace(name)
	name = strings.ReplaceAll(name, "T::", "")
	name = strings.ReplaceAll(name, "VecDeque<", "Vec<")
	name = strings.ReplaceAll(name, "<T>", "")
	name = strings.ReplaceAll(name, "<T as Trait>::", "")
	name = strings.ReplaceAll(name, "<T, I>", "")
	name = strings.ReplaceAll(name, "\n", " ")
	name = strings.ReplaceAll(name, `&'static[u8]`, "Bytes")

	switch name {
	case "()", "<InherentOfflineReport as InherentOfflineReport>::Inherent":
		name = "Null"
	case "Vec<u8>":
		name = "Bytes"
	case "<Lookup as StaticLookup>::Source":
		name = "Address"
	case "<Balance as HasCompact>::Type":
		name = "Compact<Balance>"
	case "<BlockNumber as HasCompact>::Type":
		name = "Compact<BlockNumber>"
	case "<Moment as HasCompact>::Type":
		name = "Compact<Moment>"
	case "<T as Trait<I>>::Proposal":
		name = "Proposal"
	case "wasm::PrefabWasmModule":
		name = "PrefabWasmModule"
	}
	if strings.Contains(name, "::") {
		subNames := strings.Split(name, "::")
		name = subNames[len(subNames)-1]
	}
	return name
}
