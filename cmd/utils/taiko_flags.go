package utils

import (
	"os"

	"github.com/ethereum/go-ethereum/eth"
	"github.com/ethereum/go-ethereum/eth/ethconfig"
	"github.com/ethereum/go-ethereum/node"
	"github.com/ethereum/go-ethereum/rpc"
	"github.com/urfave/cli/v2"
        "github.com/ethereum/go-ethereum/log"
)

var (
	TaikoFlag = cli.BoolFlag{
		Name:  "taiko",
		Usage: "Taiko network",
	}
)

// RegisterTaikoAPIs initializes and registers the Taiko RPC APIs.
func RegisterTaikoAPIs(stack *node.Node, cfg *ethconfig.Config, backend *eth.Ethereum) {
	log.Warn("Register Taiko API")
	if os.Getenv("TAIKO_TEST") != "" {
		return
	}
	log.Warn("Really Register Taiko API")
	// Add methods under "taiko_" RPC namespace to the available APIs list
	stack.RegisterAPIs([]rpc.API{
		{
			Namespace: "taiko",
			Service:   eth.NewTaikoAPIBackend(backend),
			Public:    true,
		},
		{
			Namespace:     "taikoAuth",
			Service:       eth.NewTaikoAuthAPIBackend(backend),
			Authenticated: true, 
			Public:        true,
		},
	})
}
