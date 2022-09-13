package gmapstruct

import (
	"gtags"
	"reflect"
	"testing"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/rpc"
	"gotest.tools/assert"
)

////////
type BasedRequest struct {
	// Order   map[string]string
	OrderBy string `json:"order_by" d:"desc" v:"in:asc,desc"`
	// maxlimit 1000
	Limit  int    `json:"limit" d:"200" p:"limit"`
	Curosr string `json:"cursor"`
	//legacy
	NextPage string `json:"next"`
	PrivPage string `json:"previous"`

	// Order map[string]string
	// idxDef []idx.IdxDef
}

type InTxRange struct {
	BasedRequest
	Address   common.Address        `json:"address" d:"0x1000000000000000000000000000000000000000"`
	BlockHash rpc.BlockNumberOrHash `json:"block_hash" d:"earliest"`
	FromBlock rpc.BlockNumber       `json:"from_block" d:"earliest"`
	ToBlock   rpc.BlockNumber       `json:"to_block" d:"latest"`
	Kind      []string              `json:"kind" d:"erc20,erc721"`
}

func Test_EthBlockNumberOrHash(t *testing.T) {
	bh := rpc.BlockNumberOrHash{}
	typ := reflect.TypeOf(bh)
	ok := gtags.TypMethod(typ, "UnmarshalJSON")
	assert.Equal(t, ok, true)
}
func Test_EthRequest_DefVal(t *testing.T) {
	val := &InTxRange{}
	dmap := map[string]any{}
	err := Map2Struct(val, dmap)

	assert.Equal(t, err, nil)
	assert.Equal(t, val.OrderBy, "desc")
	assert.Equal(t, val.FromBlock, rpc.EarliestBlockNumber)
	assert.Equal(t, val.ToBlock, rpc.LatestBlockNumber)
	assert.Equal(t, val.Kind[0], "erc20")
	assert.Equal(t, val.Kind[1], "erc721")
	assert.Equal(t, *val.BlockHash.BlockNumber, rpc.EarliestBlockNumber)
}
