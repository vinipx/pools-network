package cmd

import (
	"encoding/json"
	"fmt"
	"net"
	"os"
	"path/filepath"
	"time"

	types4 "github.com/cosmos/cosmos-sdk/x/auth/types"

	banktypes "github.com/cosmos/cosmos-sdk/x/bank/types"

	types3 "github.com/tendermint/tendermint/types"

	"github.com/cosmos/cosmos-sdk/x/genutil/types"

	config2 "github.com/cosmos/cosmos-sdk/server/config"

	"github.com/bloxapp/pools-network/app/params"

	"github.com/cosmos/iavl/common"

	"github.com/cosmos/cosmos-sdk/crypto/hd"

	keyringTypes "github.com/cosmos/cosmos-sdk/crypto/keyring"

	tx2 "github.com/cosmos/cosmos-sdk/client/tx"

	"github.com/cosmos/cosmos-sdk/client"

	types2 "github.com/cosmos/cosmos-sdk/x/staking/types"

	os2 "github.com/tendermint/tendermint/libs/os"

	"github.com/cosmos/cosmos-sdk/x/genutil"

	"github.com/tendermint/tendermint/crypto"

	"github.com/tendermint/tendermint/config"

	"github.com/spf13/viper"

	"github.com/cosmos/cosmos-sdk/client/flags"

	"github.com/cosmos/cosmos-sdk/server"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/module"
	"github.com/spf13/cobra"
)

var (
	flagNodeDirPrefix     = "node-dir-prefix"
	flagNumValidators     = "v"
	flagOutputDir         = "output-dir"
	flagNodeDaemonHome    = "node-daemon-home"
	flagNodeCLIHome       = "node-cli-home"
	flagStartingIPAddress = "starting-ip-address"
)

type testnetConfig struct {
	ctxConfig         *config.Config
	Encoding          params.EncodingConfig
	BasicAppManager   module.BasicManager
	BalancesIterator  banktypes.GenesisBalancesIterator
	OutputDir         string
	ChainID           string
	MinGasPrices      string
	NodeDirPrefix     string
	NodeDaemonHome    string
	NodeCLIHome       string
	StartingIPAddress string
	NumValidators     int
}

const nodeDirPerm = 0755

// get cmd to initialize all files for tendermint testnet and application
func testnetCmd(ctx *server.Context, encoding params.EncodingConfig,
	mbm module.BasicManager,
	genBalIterator banktypes.GenesisBalancesIterator,
) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "testnet",
		Short: "Initialize files for a pool testnet",
		Long: `testnet will create "v" number of directories and populate each with
necessary files (private validator, genesis, config, etc.).

Note, strict routability for addresses is turned off in the config file.

Example:
	pools testnet --v 4 --output-dir ./output --starting-ip-address 192.168.10.2
	`,
		RunE: func(cmd *cobra.Command, args []string) error {
			return runLocalTestnetCmd(cmd, &testnetConfig{
				ctxConfig:         ctx.Config,
				Encoding:          encoding,
				BasicAppManager:   mbm,
				BalancesIterator:  genBalIterator,
				OutputDir:         viper.GetString(flagOutputDir),
				ChainID:           viper.GetString(flags.FlagChainID),
				MinGasPrices:      viper.GetString(server.FlagMinGasPrices),
				NodeDirPrefix:     viper.GetString(flagNodeDirPrefix),
				NodeDaemonHome:    viper.GetString(flagNodeDaemonHome),
				NodeCLIHome:       viper.GetString(flagNodeCLIHome),
				StartingIPAddress: viper.GetString(flagStartingIPAddress),
				NumValidators:     viper.GetInt(flagNumValidators),
			})
		},
	}

	cmd.Flags().Int(flagNumValidators, 4,
		"Number of validators to initialize the testnet with")
	cmd.Flags().StringP(flagOutputDir, "o", "./mytestnet",
		"Directory to store initialization data for the testnet")
	cmd.Flags().String(flagNodeDirPrefix, "node",
		"Prefix the directory name for each node with (node results in node0, node1, ...)")
	cmd.Flags().String(flagNodeDaemonHome, "poolsd",
		"Home directory of the node's daemon configuration")
	cmd.Flags().String(flagNodeCLIHome, "poolscli",
		"Home directory of the node's cli configuration")
	cmd.Flags().String(flagStartingIPAddress, "192.168.0.1",
		"Starting IP address (192.168.0.1 results in persistent peers list ID0@192.168.0.1:46656, ID1@192.168.0.2:46656, ...)")
	cmd.Flags().String(flags.FlagChainID, "", "genesis file chain-id, if left blank will be randomly created")
	cmd.Flags().String(
		server.FlagMinGasPrices, fmt.Sprintf("0.000006%s", sdk.DefaultBondDenom),
		"Minimum gas prices to accept for transactions; All fees in a tx must meet this minimum (e.g. 0.01photino,0.001stake)")
	return cmd
}

func runLocalTestnetCmd(cmd *cobra.Command, config *testnetConfig) error {
	// context and builders
	clientCtx := client.GetClientContextFromCmd(cmd)
	clientCtx, err := client.ReadTxCommandFlags(clientCtx, cmd.Flags())
	if err != nil {
		return err
	}

	chainId := config.ChainID
	if chainId == "" {
		chainId = "chain-" + common.RandStr(6)
	}

	monikers := make([]string, 0)
	nodeIDs := make([]string, 0)
	nodeDirs := make([]string, 0)
	valPubKeys := make([]crypto.PubKey, 0)
	//genTxs := make([]types4.GenesisAccount, 0)

	var (
		accounts        []types4.GenesisAccount
		accountBalances []banktypes.Balance
		genFiles        []string
	)

	// generate private keys, node IDs, and initial transactions
	for i := 0; i < config.NumValidators; i++ {
		nodeDirName := fmt.Sprintf("%s%d", config.NodeDirPrefix, i)
		nodeDir := filepath.Join(config.OutputDir, nodeDirName, config.NodeDaemonHome)

		// generate validator info
		keyring, memo, nodeId, account, accountBalance, msg, err := generateValidator(i,
			config.StartingIPAddress,
			nodeDir,
			config.OutputDir,
			nodeDirName,
			config.NodeDaemonHome,
			config.ctxConfig,
		)
		if err != nil {
			return err
		}

		// get key
		key, err := keyring.Key(nodeDirName)
		if err != nil {
			return err
		}

		// add relevant data to arrays
		monikers = append(monikers, nodeDirName)
		nodeIDs = append(nodeIDs, nodeId)
		valPubKeys = append(valPubKeys, key.GetPubKey())
		accounts = append(accounts, account)
		accountBalances = append(accountBalances, *accountBalance)
		genFiles = append(genFiles, config.ctxConfig.GenesisFile())
		nodeDirs = append(nodeDirs, nodeDir)

		// build and sign tx
		txFactory := tx2.NewFactoryCLI(clientCtx, cmd.Flags()).WithChainID(chainId).WithKeybase(keyring)
		txBuilder := clientCtx.TxConfig.NewTxBuilder()
		txBuilder.SetMemo(memo)
		if err := txBuilder.SetMsgs(msg); err != nil {
			return err
		}

		if err := tx2.Sign(txFactory, nodeDirName, txBuilder); err != nil {
			return err
		}

		tx := txBuilder.GetTx()
		txBytes, err := clientCtx.TxConfig.TxJSONEncoder()(tx)
		if err != nil {
			return err
		}

		// gather gentxs folder
		gentxsDir := filepath.Join(config.OutputDir, "gentxs")
		if err := writeFile(fmt.Sprintf("%v.json", nodeDirName), gentxsDir, txBytes); err != nil {
			_ = os.RemoveAll(config.OutputDir)
			return err
		}

		// write app config
		gaiaConfigFilePath := filepath.Join(nodeDir, "config/app.toml")
		config2.WriteConfigFile(gaiaConfigFilePath, config2.DefaultConfig())
	}

	if err := initGenFiles(config, clientCtx, accounts, accountBalances, chainId, genFiles); err != nil {
		return err
	}

	if err := collectGenFiles(config, clientCtx.TxConfig, chainId, monikers, nodeIDs, nodeDirs, valPubKeys); err != nil {
		return err
	}

	return nil
}

// initGenFiles initializeds a tendermint state with the created accounts
func initGenFiles(
	config *testnetConfig,
	clientCtx client.Context,
	genAccounts []types4.GenesisAccount,
	genBalances []banktypes.Balance,
	chainId string,
	genFiles []string,
) error {
	// Generate default state and set the accounts in the genesis state
	appGenState := config.BasicAppManager.DefaultGenesis(config.Encoding.Marshaler)
	var authGenState types4.GenesisState
	clientCtx.JSONMarshaler.MustUnmarshalJSON(appGenState[types4.ModuleName], &authGenState)

	accounts, err := types4.PackAccounts(genAccounts)
	if err != nil {
		return err
	}
	authGenState.Accounts = accounts
	appGenState[types4.ModuleName] = clientCtx.JSONMarshaler.MustMarshalJSON(&authGenState)

	// set the balances in the genesis state
	var bankGenState banktypes.GenesisState
	clientCtx.JSONMarshaler.MustUnmarshalJSON(appGenState[banktypes.ModuleName], &bankGenState)
	bankGenState.Balances = genBalances
	appGenState[banktypes.ModuleName] = clientCtx.JSONMarshaler.MustMarshalJSON(&bankGenState)

	// save updated state for each node
	appGenStateJSON, err := json.MarshalIndent(appGenState, "", "  ")
	if err != nil {
		return err
	}

	// GenesisDoc defines the initial conditions for a tendermint blockchain, in particular its validator set.
	doc := types3.GenesisDoc{
		ChainID:  chainId,
		AppState: appGenStateJSON,
	}

	// generate empty genesis files for each validator and save
	for i := 0; i < config.NumValidators; i++ {
		if err := doc.SaveAs(genFiles[i]); err != nil {
			return err
		}
	}
	return nil
}

func collectGenFiles(
	config *testnetConfig,
	clientCtx client.TxConfig,
	chainId string,
	monikers, nodeIds, nodeDirs []string,
	valPks []crypto.PubKey,
) error {
	genesisTime := time.Now()
	for i := 0; i < config.NumValidators; i++ {
		gentxsDir := filepath.Join(config.OutputDir, "gentxs")
		configNode := config.ctxConfig
		configNode.SetRoot(nodeDirs[i])
		configNode.Moniker = monikers[i]

		initConfig := types.NewInitConfig(chainId, gentxsDir, nodeIds[i], valPks[i])

		genDoc, err := types3.GenesisDocFromFile(configNode.GenesisFile())
		if err != nil {
			return err
		}

		nodeAppState, err := genutil.GenAppStateFromConfig(
			config.Encoding.Marshaler,
			clientCtx,
			config.ctxConfig,
			initConfig,
			*genDoc,
			config.BalancesIterator,
		)
		if err != nil {
			return err
		}

		genFile := configNode.GenesisFile()
		// overwrite each validator's genesis file to have a canonical genesis time
		if err := genutil.ExportGenesisFileWithTime(genFile, chainId, nil, nodeAppState, genesisTime); err != nil {
			return err
		}
	}
	return nil
}

func generateValidator(
	indx int,
	startingIPAddress string,
	nodeDir string,
	outputDir, nodeDirName, nodeDaemonHome string,
	ctxConfig *config.Config,
) (keyring keyringTypes.Keyring, memo string, nodeId string, account types4.GenesisAccount, accountBalance *banktypes.Balance, msg *types2.MsgCreateValidator, err error) {
	ctxConfig.SetRoot(nodeDir)
	ctxConfig.RPC.ListenAddress = "tcp://0.0.0.0:26657"

	// make config dir
	if err := os.MkdirAll(filepath.Join(nodeDir, "config"), nodeDirPerm); err != nil {
		_ = os.RemoveAll(outputDir)
		return nil, "", "", nil, nil, nil, err
	}

	// moniker
	ctxConfig.Moniker = nodeDirName

	// generate node id and validator pk
	nodeID, _, err := genutil.InitializeNodeValidatorFiles(ctxConfig)
	if err != nil {
		_ = os.RemoveAll(outputDir)
		return nil, "", "", nil, nil, nil, err
	}

	// ip and memo
	ip, err := getIP(indx, startingIPAddress)
	if err != nil {
		_ = os.RemoveAll(outputDir)
		return nil, "", "", nil, nil, nil, err
	}
	memo = fmt.Sprintf("%s@%s:26656", nodeID, ip)

	// generate keyring
	keyring, err = keyringTypes.New("", keyringTypes.BackendMemory, outputDir, nil)
	if err != nil {
		return nil, "", "", nil, nil, nil, err
	}
	path := hd.CreateHDPath(118, 0, 0).String()
	info, _, err := keyring.NewMnemonic(nodeDirName, keyringTypes.English, path, hd.Secp256k1)
	if err != nil {
		return nil, "", "", nil, nil, nil, err
	}

	// generate account
	accountBalance = &banktypes.Balance{
		Address: info.GetAddress().String(),
		Coins: sdk.Coins{
			sdk.NewCoin(fmt.Sprintf("%stoken", nodeDirName),
				sdk.TokensFromConsensusPower(1000)),
			sdk.NewCoin(sdk.DefaultBondDenom,
				sdk.TokensFromConsensusPower(500)),
		},
	}
	account = types4.NewBaseAccount(info.GetAddress(), nil, 0, 0)

	// validator account and create validator tx
	msg = types2.NewMsgCreateValidator(
		sdk.ValAddress(info.GetAddress()),
		info.GetPubKey(),
		sdk.NewCoin(sdk.DefaultBondDenom,
			sdk.TokensFromConsensusPower(100)),
		types2.NewDescription(nodeDirName, "", "", "", ""),
		types2.NewCommissionRates(sdk.ZeroDec(), sdk.ZeroDec(), sdk.ZeroDec()),
		sdk.OneInt(),
	)

	return keyring, memo, nodeID, account, accountBalance, msg, nil
}

func getIP(i int, startingIPAddr string) (ip string, err error) {
	if len(startingIPAddr) == 0 {
		ip, err = server.ExternalIP()
		if err != nil {
			return "", err
		}
		return ip, nil
	}
	return calculateIP(startingIPAddr, i)
}

func calculateIP(ip string, i int) (string, error) {
	ipv4 := net.ParseIP(ip).To4()
	if ipv4 == nil {
		return "", fmt.Errorf("%v: non ipv4 address", ip)
	}

	for j := 0; j < i; j++ {
		ipv4[3]++
	}

	return ipv4.String(), nil
}

func writeFile(name string, dir string, contents []byte) error {
	writePath := filepath.Join(dir)
	file := filepath.Join(writePath, name)

	err := os2.EnsureDir(writePath, 0700)
	if err != nil {
		return err
	}

	err = os2.WriteFile(file, contents, 0600)
	if err != nil {
		return err
	}

	return nil
}
