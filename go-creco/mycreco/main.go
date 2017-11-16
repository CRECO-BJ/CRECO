package main

import (
	"fmt"
	"os"

	"github.com/creco/go-creco/common"
	"github.com/creco/go-creco/mycreco/utils"
	"github.com/creco/go-creco/node"
	"github.com/creco/go-creco/p2p/simulations"
	"github.com/creco/go-creco/p2p/simulations/adapters"
	cli "gopkg.in/urfave/cli.v1"
)

var (
	configFileFlag = cli.StringFlag{
		Name:  "config",
		Usage: "TOML configuration file",
	}
)

var client *simulations.Client

var (
	// Git SHA1 commit hash of the release (set via linker flags)
	gitCommit = ""

	// Ethereum address of the Geth release oracle.
	relOracle = common.HexToAddress("0xfa7b9770ca4cb04296cac84f37736d4041251cdf")
	// The app that holds all commands and flags.
	//app = utils.NewApp(gitCommit, "the go-ethereum command line interface")
	// flags that configure the node
	nodeFlags = []cli.Flag{
		configFileFlag,
	}
)

func main() {
	app := cli.NewApp()
	app.Usage = "devp2p simulation command-line client"
	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:   "api",
			Value:  "http://localhost:8888",
			Usage:  "simulation API URL",
			EnvVar: "P2PSIM_API_URL",
		},
	}
	app.Before = func(ctx *cli.Context) error {
		client = simulations.NewClient(ctx.GlobalString("api"))
		return nil
	}
	app.Run(os.Args)

	client = simulations.NewClient("http://localhost:8888")

	network, err := client.GetNetwork()

	fmt.Println("network:", network.ID)
	// cfg := gethConfig{
	// 	//Eth:  eth.DefaultConfig,
	// 	//Shh:  whisper.DefaultConfig,
	// 	Node: defaultNodeConfig(),
	// }
	// stack, err := node.New(&cfg.Node)
	// if err != nil {
	// 	utils.Fatalf("Failed to create the protocol stack: %v", err)
	// }
	// utils.StartNode(stack)

	config := &adapters.NodeConfig{
		Name: "node01",
	}

	node, err := client.CreateNode(config)
	if err != nil {
		fmt.Println("something errror:", err)
		return
	}
	fmt.Println("Created", node.Name)

}

// geth is the main entry point into the system if no special subcommand is ran.
// It creates a default node based on the command line arguments and runs it in
// blocking mode, waiting for it to be shut down.
// func geth(ctx *cli.Context) error {
// 	node := makeFullNode(ctx)
// 	startNode(ctx, node)
// 	node.Wait()
// 	return nil
// }

// startNode boots up the system node and all registered protocols, after which
// it unlocks any requested accounts, and starts the RPC/IPC interfaces and the
// miner.
func startNode(ctx *cli.Context, stack *node.Node) {
	// Start up the node itself
	utils.StartNode(stack)

	// Unlock any account specifically requested
	//ks := stack.AccountManager().Backends(keystore.KeyStoreType)[0].(*keystore.KeyStore)

	//passwords := utils.MakePasswordList(ctx)
	//unlocks := strings.Split(ctx.GlobalString(utils.UnlockedAccountFlag.Name), ",")
	// for i, account := range unlocks {
	// 	if trimmed := strings.TrimSpace(account); trimmed != "" {
	// 		unlockAccount(ctx, ks, trimmed, i, passwords)
	// 	}
	// }
	// Register wallet event handlers to open and auto-derive wallets
	// events := make(chan accounts.WalletEvent, 16)
	// stack.AccountManager().Subscribe(events)

	go func() {
		fmt.Println("empty go...")
		// Create an chain state reader for self-derivation
		// rpcClient, err := stack.Attach()
		// if err != nil {
		// 	utils.Fatalf("Failed to attach to self: %v", err)
		// }
		//stateReader := ethclient.NewClient(rpcClient)

		// Open any wallets already attached
		// for _, wallet := range stack.AccountManager().Wallets() {
		// 	if err := wallet.Open(""); err != nil {
		// 		log.Warn("Failed to open wallet", "url", wallet.URL(), "err", err)
		// 	}
		// }
		// Listen for wallet event till termination
		// for event := range events {
		// 	switch event.Kind {
		// 	case accounts.WalletArrived:
		// 		if err := event.Wallet.Open(""); err != nil {
		// 			log.Warn("New wallet appeared, failed to open", "url", event.Wallet.URL(), "err", err)
		// 		}
		// 	case accounts.WalletOpened:
		// 		status, _ := event.Wallet.Status()
		// 		log.Info("New wallet appeared", "url", event.Wallet.URL(), "status", status)
		//
		// 		if event.Wallet.URL().Scheme == "ledger" {
		// 			event.Wallet.SelfDerive(accounts.DefaultLedgerBaseDerivationPath, stateReader)
		// 		} else {
		// 			event.Wallet.SelfDerive(accounts.DefaultBaseDerivationPath, stateReader)
		// 		}
		//
		// 	case accounts.WalletDropped:
		// 		log.Info("Old wallet dropped", "url", event.Wallet.URL())
		// 		event.Wallet.Close()
		// 	}
		// }
	}()
	// Start auxiliary services if enabled
	// if ctx.GlobalBool(utils.MiningEnabledFlag.Name) || ctx.GlobalBool(utils.DeveloperFlag.Name) {
	// 	// Mining only makes sense if a full Ethereum node is running
	// 	var ethereum *eth.Ethereum
	// 	if err := stack.Service(&ethereum); err != nil {
	// 		utils.Fatalf("ethereum service not running: %v", err)
	// 	}
	// 	// Use a reduced number of threads if requested
	// 	if threads := ctx.GlobalInt(utils.MinerThreadsFlag.Name); threads > 0 {
	// 		type threaded interface {
	// 			SetThreads(threads int)
	// 		}
	// 		if th, ok := ethereum.Engine().(threaded); ok {
	// 			th.SetThreads(threads)
	// 		}
	// 	}
	// 	// Set the gas price to the limits from the CLI and start mining
	// 	ethereum.TxPool().SetGasPrice(utils.GlobalBig(ctx, utils.GasPriceFlag.Name))
	// 	if err := ethereum.StartMining(true); err != nil {
	// 		utils.Fatalf("Failed to start mining: %v", err)
	// 	}
	// }
}
