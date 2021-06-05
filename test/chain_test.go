package test

import (
	"fmt"
	"github.com/JFJun/bifrost-go/client"
	"github.com/JFJun/bifrost-go/expand"
	"github.com/JFJun/bifrost-go/expand/base"
	"github.com/JFJun/bifrost-go/expand/polkadot"
	"github.com/stafiprotocol/go-substrate-rpc-client/types"
	"reflect"
	"testing"
)

func Test_Chain(t *testing.T) {
	c, err := client.New("wss://rpc.polkadot.io")
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

func Test_Bew(t *testing.T) {
	b := polkadot.PolkadotEventRecords{}
	//tp:=reflect.TypeOf(b)
	//fmt.Println(tp.NumField())
	//fmt.Println(tp.Field(1).Name)
	getEventTypesFieldName(b)
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
	c, err := client.New("wss://rpc.polkadot.io")
	if err != nil {
		t.Fatal(err)
	}

	for _, mod := range c.Meta.AsMetadataV12.Modules {
		if mod.HasEvents {
			for _, event := range mod.Events {

				if len(event.Args) == 0 {
					//fmt.Printf("type Event%s%s struct { \n	Phase    types.Phase\n	\n	Topics []types.Hash\n}\n", mod.Name, event.Name)
				} else {
					//as:=""
					for _, arg := range event.Args {
						fmt.Println(arg)
					}

				}

				//fmt.Println(event.Args)

			}
		}
	}
}
