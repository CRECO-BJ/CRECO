package node

import (
	"io/ioutil"
	"log"
	"encoding/json"
	"strings"
	"path/filepath"

	"github.com/ethereum/go-ethereum/eth"
	"github.com/ethereum/go-ethereum/eth/downloader"
	"github.com/ethereum/go-ethereum/core"
	ethnode "github.com/ethereum/go-ethereum/node"
	"github.com/ethereum/go-ethereum/p2p/discv5"
	"github.com/ethereum/go-ethereum/p2p/nat"
	"github.com/ethereum/go-ethereum/params"
)

const (
	ropstenID = 3
	kovanID = 42
	rinkebyID = 4
	sokolID = 77
)

var ropstenBootNodes = []string{
	"enode://94c15d1b9e2fe7ce56e458b9a3b672ef11894ddedd0c6f247e0f1d3487f52b66208fb4aeb8179fce6e3a749ea93ed147c37976d67af557508d199d9594c35f09@192.81.208.223:30303",
	"enode://6332792c4a00e3e4ee0926ed89e0d27ef985424d97b6a45bf0f23e51f0dcb5e66b875777506458aea7af6f9e4ffb69f43f3778ee73c81ed9d34c51c4b16b0b0f@52.232.243.152:30303",
}

// Node ...
type Node struct {
}

// NewNode ...
func NewNode() (*Node, error) {
	// Convert the bootnodes to internal enode representations
	var enodes []*discv5.Node
	for _, boot := range ropstenBootNodes {
		if url, err := discv5.ParseNode(boot); err == nil {
			enodes = append(enodes, url)
		} else {
			log.Error("Failed to parse bootnode URL", "url", boot, "err", err)
		}
	}

	stack, err := ethnode.New(&ethnode.Config{
		Name:    "geth",
		Version: params.VersionWithMeta,
		DataDir: filepath.Join(os.Getenv("HOME"), ".testWallet"),
		P2P: p2p.Config{
			NAT:              nat.Any(),
			NoDiscovery:      true,
			DiscoveryV5:      true,
			ListenAddr:       fmt.Sprintf(":%d", 30303),
			MaxPeers:         25,
			BootstrapNodesV5: enodes,
		},
	})
	if err != nil {
		return nil, err
	}

	// Assemble the Ethereum light client protocol
	if err := stack.Register(func(ctx *ethnode.ServiceContext) (ethnode.Service, error) {
		cfg := eth.DefaultConfig
		cfg.SyncMode = downloader.LightSync
		cfg.NetworkId = network
		return les.New(ctx, &cfg)
	}); err != nil {
		return nil, err
	}

	return stack, nil
}

// Run ...
func (n *Node) Run(stact *ethnode.Node, enodes []*discv5.Node) error {
	// Assemble and start the faucet light service
	// Boot up the client and ensure it connects to bootnodes
	if err := stack.Start(); err != nil {
		return nil, err
	}
	for _, boot := range ropstenBootNodes {
		old, _ := discover.ParseNode(boot)
		stack.Server().AddPeer(old)
	}
	// Attach to the client and retrieve and interesting metadatas
	api, err := stack.Attach()
	if err != nil {
		stack.Stop()
		return nil, err
	}
	client := ethclient.NewClient(api)

	return nil
}
