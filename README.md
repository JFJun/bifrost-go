# bifrost go sdk
## fork
    https://github.com/JFJun/stafi-substrate-go.git
## transfer
    from := ""
    to := ""
    amount := uint64(50000000000000)
    c, err := client.New("wss://testnet.liebi.com")
    if err != nil {
    	t.Fatal(err)
    }
    v, err := c.C.RPC.State.GetRuntimeVersionLatest()
    if err != nil {
    	t.Fatal(err)
    }
    
    acc, err := c.GetAccountInfo(from)
    if err != nil {
    	t.Fatal(err)
    }
    c.SetPrefix(ss58.BifrostPrefix)
    nonce := uint64(acc.Nonce)
    types.SetSerDeOptions(types.SerDeOptions{NoPalletIndices: true})
    transaction:=tx.NewSubstrateTransaction(from,nonce)
    transaction.SetGenesisHashAndBlockHash(c.GetGenesisHash(),
    	c.GetGenesisHash())
    ed, err := expand.NewMetadataExpand(c.Meta)
    if err != nil {
    	t.Fatal(err)
    }
    call,err :=ed.BalanceTransferKeepAliveCall(to,amount)
    if err != nil {
    	t.Fatal(err)
    }
    transaction.SetCall(call)
    transaction.SetSpecVersionAndCallId(uint32(v.SpecVersion), uint32(v.TransactionVersion))
    tt, err := transaction.SignTranaction("", crypto.Sr25519Type)
    if err != nil {
    	t.Fatal(err)
    }
  
    var result interface{}
    err = c.C.Client.Call(&result, "author_submitExtrinsic", tt)
    if err != nil {
    	t.Fatal(err)
    }
    fmt.Println(result)