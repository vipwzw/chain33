package api

import (
	"testing"
	"time"

	"github.com/33cn/chain33/client/mocks"
	qmocks "github.com/33cn/chain33/queue/mocks"
	"github.com/33cn/chain33/rpc"
	"github.com/33cn/chain33/types"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"

	_ "github.com/33cn/chain33/system"
)

func TestAPI(t *testing.T) {
	api := new(mocks.QueueProtocolAPI)
	eapi := New(api, "")
	param := &types.ReqHashes{
		Hashes: [][]byte{[]byte("hello")},
	}
	api.On("GetBlockByHashes", mock.Anything).Return(&types.BlockDetails{}, nil)
	detail, err := eapi.GetBlockByHashes(param)
	assert.Nil(t, err)
	assert.Equal(t, detail, &types.BlockDetails{})

	param2 := &types.ReqRandHash{
		ExecName: "ticket",
		BlockNum: 5,
		Hash:     []byte("hello"),
	}
	api.On("Query", "ticket", "RandNumHash", mock.Anything).Return(&types.ReplyHash{Hash: []byte("hello")}, nil)
	randhash, err := eapi.GetRandNum(param2)
	assert.Nil(t, err)
	assert.Equal(t, randhash, []byte("hello"))

	types.SetTitleOnlyForTest("user.p.wzw.")
	//testnode setup
	rpcCfg := new(types.RPC)
	rpcCfg.GrpcBindAddr = "127.0.0.1:8002"
	rpcCfg.JrpcBindAddr = "127.0.0.1:8001"
	rpcCfg.MainnetJrpcAddr = rpcCfg.JrpcBindAddr
	rpcCfg.Whitelist = []string{"127.0.0.1", "0.0.0.0"}
	rpcCfg.JrpcFuncWhitelist = []string{"*"}
	rpcCfg.GrpcFuncWhitelist = []string{"*"}
	rpc.InitCfg(rpcCfg)
	server := rpc.NewGRpcServer(&qmocks.Client{}, api)
	assert.NotNil(t, server)
	go server.Listen()
	time.Sleep(time.Second)

	eapi = New(api, "")
	detail, err = eapi.GetBlockByHashes(param)
	assert.Nil(t, err)
	assert.Equal(t, detail, &types.BlockDetails{})
	randhash, err = eapi.GetRandNum(param2)
	assert.Nil(t, err)
	assert.Equal(t, randhash, []byte("hello"))
}
