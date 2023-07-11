package code_gen

import "testing"

func TestGUnStruct(t *testing.T) {
	GenModel("coin", "Coin")
}

func TestGenRpc(t *testing.T) {
	rpcCommon := RpcCommon{
		PackageName: "mclient",
		ModuleName:  "Market",
		ServiceName: "Market",
		GrpcPackage: "market",
	}
	// rpc FindSymbolThumb(MarketReq) returns(SymbolThumbRes);
	//  rpc FindSymbolThumbTrend(MarketReq) returns(SymbolThumbRes);
	//  rpc FindSymbolInfo(MarketReq) returns(ExchangeCoin);
	rpc1 := Rpc{
		FunName: "FindSymbolThumb",
		Resp:    "SymbolThumbRes",
		Req:     "MarketReq",
	}
	rpc2 := Rpc{
		FunName: "FindSymbolThumbTrend",
		Resp:    "SymbolThumbRes",
		Req:     "MarketReq",
	}
	rpc3 := Rpc{
		FunName: "FindSymbolInfo",
		Resp:    "ExchangeCoin",
		Req:     "MarketReq",
	}
	rpcList := []Rpc{}
	rpcList = append(rpcList, rpc1, rpc2, rpc3)
	result := RpcResult{
		RpcCommon: rpcCommon,
		Rpc:       rpcList,
	}
	GenZeroRpc(result)
}
