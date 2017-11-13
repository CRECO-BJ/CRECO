package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"reflect"
	"unicode"

	"github.com/creco/go-creco/mycreco/utils"
	"github.com/creco/go-creco/node"
	"github.com/go-ethereum/params"
	"github.com/naoina/toml"
	cli "gopkg.in/urfave/cli.v1"
)

const (
	clientIdentifier = "geth" // Client identifier to advertise over the network
)

// These settings ensure that TOML keys use the same names as Go struct fields.
var tomlSettings = toml.Config{
	NormFieldName: func(rt reflect.Type, key string) string {
		return key
	},
	FieldToKey: func(rt reflect.Type, field string) string {
		return field
	},
	MissingField: func(rt reflect.Type, field string) error {
		link := ""
		if unicode.IsUpper(rune(rt.Name()[0])) && rt.PkgPath() != "main" {
			link = fmt.Sprintf(", see https://godoc.org/%s#%s for available fields", rt.PkgPath(), rt.Name())
		}
		return fmt.Errorf("field '%s' is not defined in %s%s", field, rt.String(), link)
	},
}

type gethConfig struct {
	Node node.Config
}

func makeFullNode(ctx *cli.Context) *node.Node {
	stack, cfg := makeConfigNode(ctx)

	//utils.RegisterEthService(stack, &cfg.Eth)

	// Whisper must be explicitly enabled by specifying at least 1 whisper flag or in dev mode
	// shhEnabled := enableWhisper(ctx)
	// shhAutoEnabled := !ctx.GlobalIsSet(utils.WhisperEnabledFlag.Name) && ctx.GlobalIsSet(utils.DeveloperFlag.Name)
	// if shhEnabled || shhAutoEnabled {
	// 	if ctx.GlobalIsSet(utils.WhisperMaxMessageSizeFlag.Name) {
	// 		cfg.Shh.MaxMessageSize = uint32(ctx.Int(utils.WhisperMaxMessageSizeFlag.Name))
	// 	}
	// 	if ctx.GlobalIsSet(utils.WhisperMinPOWFlag.Name) {
	// 		cfg.Shh.MinimumAcceptedPOW = ctx.Float64(utils.WhisperMinPOWFlag.Name)
	// 	}
	// 	utils.RegisterShhService(stack, &cfg.Shh)
	// }

	// Add the Ethereum Stats daemon if requested.
	// if cfg.Ethstats.URL != "" {
	// 	utils.(stack, cfg.Ethstats.URL)
	// }

	// Add the release oracle service so it boots along with node.
	// if err := stack.Register(func(ctx *node.ServiceContext) (node.Service, error) {
	// 	config := release.Config{
	// 		Oracle: relOracle,
	// 		Major:  uint32(params.VersionMajor),
	// 		Minor:  uint32(params.VersionMinor),
	// 		Patch:  uint32(params.VersionPatch),
	// 	}
	// 	commit, _ := hex.DecodeString(gitCommit)
	// 	copy(config.Commit[:], commit)
	// 	return release.NewReleaseService(ctx, config)
	// }); err != nil {
	// 	utils.Fatalf("Failed to register the Geth release oracle service: %v", err)
	// }
	fmt.Println("datadir:", stack.DataDir())
	fmt.Println("cfg:", cfg.Node)
	return stack
}

func makeConfigNode(ctx *cli.Context) (*node.Node, gethConfig) {
	// Load defaults.
	cfg := gethConfig{
		//Eth:  eth.DefaultConfig,
		//Shh:  whisper.DefaultConfig,
		Node: defaultNodeConfig(),
	}

	// Load config file.
	if file := ctx.GlobalString(configFileFlag.Name); file != "" {
		if err := loadConfig(file, &cfg); err != nil {
			utils.Fatalf("%v", err)
		}
	}

	// Apply flags.
	utils.SetNodeConfig(ctx, &cfg.Node)
	stack, err := node.New(&cfg.Node)
	if err != nil {
		utils.Fatalf("Failed to create the protocol stack: %v", err)
	}
	// utils.SetEthConfig(ctx, stack, &cfg.Eth)
	// if ctx.GlobalIsSet(utils.EthStatsURLFlag.Name) {
	// 	cfg.Ethstats.URL = ctx.GlobalString(utils.EthStatsURLFlag.Name)
	// }

	//utils.SetShhConfig(ctx, stack, &cfg.Shh)

	return stack, cfg
}

func defaultNodeConfig() node.Config {
	cfg := node.DefaultConfig
	cfg.Name = clientIdentifier
	cfg.Version = params.VersionWithCommit(gitCommit)
	//cfg.HTTPModules = append(cfg.HTTPModules, "eth", "shh")
	cfg.HTTPModules = append(cfg.HTTPModules)
	//cfg.WSModules = append(cfg.WSModules, "eth", "shh")
	cfg.WSModules = append(cfg.WSModules)
	cfg.IPCPath = "geth.ipc"
	return cfg
}

func loadConfig(file string, cfg *gethConfig) error {
	f, err := os.Open(file)
	if err != nil {
		return err
	}
	defer f.Close()

	err = tomlSettings.NewDecoder(bufio.NewReader(f)).Decode(cfg)
	// Add file name to errors that have a line number.
	if _, ok := err.(*toml.LineError); ok {
		err = errors.New(file + ", " + err.Error())
	}
	return err
}
