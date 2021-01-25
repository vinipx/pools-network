package cmd

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"time"

	github_com_bloxapp_pools_network_shared_types "github.com/bloxapp/pools-network/shared/types"

	types5 "github.com/bloxapp/pools-network/x/poolsnetwork/types"

	types4 "github.com/cosmos/cosmos-sdk/x/auth/types"

	banktypes "github.com/cosmos/cosmos-sdk/x/bank/types"

	"github.com/tendermint/tendermint/config"
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

	tendermintConfig "github.com/tendermint/tendermint/config"

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
	Encoding         params.EncodingConfig
	BasicAppManager  module.BasicManager
	BalancesIterator banktypes.GenesisBalancesIterator
	OutputDir        string
	ChainID          string
	MinGasPrices     string
	NodeDirPrefix    string
	NodeDaemonHome   string
	NodeCLIHome      string
	NumValidators    int
}

const (
	nodeDirPerm      = 0755
	localHost        = "127.0.0.1"
	P2PPortTemplate  = "500%d"
	RPCPortTemplate  = "600%d"
	GRPCPortTemplate = "700%d"
)

// get cmd to initialize all files for tendermint testnet and application
func testnetCmd(encoding params.EncodingConfig,
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
				Encoding:         encoding,
				BasicAppManager:  mbm,
				BalancesIterator: genBalIterator,
				OutputDir:        viper.GetString(flagOutputDir),
				ChainID:          viper.GetString(flags.FlagChainID),
				MinGasPrices:     viper.GetString(server.FlagMinGasPrices),
				NodeDirPrefix:    viper.GetString(flagNodeDirPrefix),
				NodeDaemonHome:   viper.GetString(flagNodeDaemonHome),
				NodeCLIHome:      viper.GetString(flagNodeCLIHome),
				NumValidators:    viper.GetInt(flagNumValidators),
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
	nodeConfigs := make([]*tendermintConfig.Config, 0)
	appConfigs := make([]*config2.Config, 0)

	var (
		accounts        []types4.GenesisAccount
		genOperators    []types5.Operator
		accountBalances []banktypes.Balance
		genFiles        []string
	)

	// generate private keys, node IDs, and initial transactions
	for i := 0; i < config.NumValidators; i++ {
		nodeDirName := fmt.Sprintf("%s%d", config.NodeDirPrefix, i)
		nodeDir := filepath.Join(config.OutputDir, nodeDirName, config.NodeDaemonHome)

		ctxConfig := tendermintConfig.DefaultConfig()

		// generate validator info
		keyring, memo, nodeId, account, accountBalance, msg, operator, err := generateOperator(i,
			nodeDir,
			config.OutputDir,
			nodeDirName,
			ctxConfig,
		)
		if err != nil {
			return err
		}

		// get key
		key, err := keyring.Key(nodeDirName)
		if err != nil {
			return err
		}

		// write app config
		configFilePath := filepath.Join(nodeDir, "config/app.toml")
		appConfig := config2.DefaultConfig()
		config2.WriteConfigFile(configFilePath, appConfig)

		// add relevant data to arrays
		monikers = append(monikers, nodeDirName)
		nodeIDs = append(nodeIDs, nodeId)
		valPubKeys = append(valPubKeys, key.GetPubKey())
		accounts = append(accounts, account)
		accountBalances = append(accountBalances, *accountBalance)
		genFiles = append(genFiles, ctxConfig.GenesisFile())
		nodeDirs = append(nodeDirs, nodeDir)
		nodeConfigs = append(nodeConfigs, ctxConfig)
		appConfigs = append(appConfigs, appConfig)
		genOperators = append(genOperators, *operator)

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
	}

	if err := initGenFiles(config, clientCtx, accounts, accountBalances, chainId, genFiles); err != nil {
		return err
	}

	if err := collectGenFiles(config, nodeConfigs, clientCtx, chainId, monikers, nodeIDs, nodeDirs, valPubKeys, genOperators); err != nil {
		return err
	}

	initNetworkConfig(config, nodeConfigs, appConfigs, nodeDirs)

	return nil
}

// initNetworkConfig sets each node's RPC and P2P addresses correctly.
func initNetworkConfig(
	testnetConfig *testnetConfig,
	nodeConfigs []*tendermintConfig.Config,
	appConfigs []*config2.Config,
	nodeDirs []string,
) {
	for i := 0; i < testnetConfig.NumValidators; i++ {
		configNode := nodeConfigs[i]
		configNode.SetRoot(nodeDirs[i])
		configNode.RPC.ListenAddress = tcpLocalHostWithPort(fmt.Sprintf(RPCPortTemplate, i))
		configNode.P2P.ListenAddress = tcpLocalHostWithPort(fmt.Sprintf(P2PPortTemplate, i))
		configNode.P2P.AddrBookStrict = false
		configNode.P2P.AllowDuplicateIP = true
		config.WriteConfigFile(fmt.Sprintf("%s/config/config.toml", nodeDirs[i]), configNode)

		appConfigs[i].GRPC.Address = localHostWithPort(fmt.Sprintf(GRPCPortTemplate, i))
		configFilePath := filepath.Join(nodeDirs[i], "config/app.toml")
		config2.WriteConfigFile(configFilePath, appConfigs[i])
	}
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

// collectGenFiles takes the gentx's created and puts
// them into the genutil module which will run on node start and generate the relevant operators and validators
func collectGenFiles(
	config *testnetConfig,
	nodeConfigs []*tendermintConfig.Config,
	clientCtx client.Context,
	chainId string,
	monikers, nodeIds, nodeDirs []string,
	valPks []crypto.PubKey,
	genOperators []types5.Operator,
) error {
	genesisTime := time.Now()
	for i := 0; i < config.NumValidators; i++ {
		gentxsDir := filepath.Join(config.OutputDir, "gentxs")
		configNode := nodeConfigs[i]
		configNode.SetRoot(nodeDirs[i])
		configNode.Moniker = monikers[i]

		// config for inserting validator data into the genesis state
		initConfig := types.NewInitConfig(chainId, gentxsDir, nodeIds[i], valPks[i])

		genDoc, err := types3.GenesisDocFromFile(configNode.GenesisFile())
		if err != nil {
			return err
		}

		// create pools operators for generated validators
		appGenesisState, err := types.GenesisStateFromGenDoc(*genDoc)
		if err != nil {
			return err
		}
		var poolsGenState types5.GenesisState
		clientCtx.JSONMarshaler.MustUnmarshalJSON(appGenesisState[types5.ModuleName], &poolsGenState)
		poolsGenState.Operators = genOperators
		appGenesisState[types5.ModuleName] = clientCtx.JSONMarshaler.MustMarshalJSON(&poolsGenState)
		// put back into gendoc
		genDoc.AppState, err = json.MarshalIndent(appGenesisState, "", "  ")
		if err != nil {
			return err
		}

		// places gen txs into the genesis state
		nodeAppState, err := genutil.GenAppStateFromConfig(
			config.Encoding.Marshaler,
			clientCtx.TxConfig,
			configNode,
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

func generateOperator(
	indx int,
	nodeDir string,
	outputDir, nodeDirName string,
	ctxConfig *config.Config,
) (
	keyring keyringTypes.Keyring,
	memo string,
	nodeId string,
	account types4.GenesisAccount,
	accountBalance *banktypes.Balance,
	msgValidator *types2.MsgCreateValidator,
	operator *types5.Operator,
	err error,
) {
	ctxConfig.SetRoot(nodeDir)

	// make config dir
	err = os.MkdirAll(filepath.Join(nodeDir, "config"), nodeDirPerm)
	if err != nil {
		_ = os.RemoveAll(outputDir)
		return
	}

	// moniker
	ctxConfig.Moniker = nodeDirName

	// generate node id and validator pk
	nodeID, valPubKey, err := genutil.InitializeNodeValidatorFiles(ctxConfig)
	if err != nil {
		_ = os.RemoveAll(outputDir)
		return
	}

	// When collecting gen txs cosmos SDK parses the memo to get the node's persistent address
	memo = fmt.Sprintf("%s@%s", nodeID, localHostWithPort(fmt.Sprintf(P2PPortTemplate, indx)))

	// generate keyring
	keyring, err = keyringTypes.New("", keyringTypes.BackendMemory, outputDir, nil)
	if err != nil {
		return
	}
	path := hd.CreateHDPath(118, 0, 0).String()
	info, _, err := keyring.NewMnemonic(nodeDirName, keyringTypes.English, path, hd.Secp256k1)
	if err != nil {
		return
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
	msgValidator = types2.NewMsgCreateValidator(
		sdk.ValAddress(info.GetAddress()),
		valPubKey,
		sdk.NewCoin(sdk.DefaultBondDenom,
			sdk.TokensFromConsensusPower(100)),
		types2.NewDescription(nodeDirName, "", "", "", ""),
		types2.NewCommissionRates(sdk.ZeroDec(), sdk.ZeroDec(), sdk.ZeroDec()),
		sdk.OneInt(),
	)

	// create operator which refers the created validator
	encodedPk, err := sdk.Bech32ifyPubKey(sdk.Bech32PubKeyTypeConsPub, valPubKey)
	if err != nil {
		return
	}
	operator = &types5.Operator{
		EthereumAddress:  github_com_bloxapp_pools_network_shared_types.EthereumAddress{},
		ConsensusAddress: github_com_bloxapp_pools_network_shared_types.ConsensusAddress(valPubKey.Address()),
		ConsensusPk:      encodedPk,
		EthStake:         msgValidator.Value.Amount.Uint64(),
	}

	return
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

func localHostWithPort(port string) string {
	return fmt.Sprintf("%s:%s", localHost, port)
}

func tcpLocalHostWithPort(port string) string {
	return fmt.Sprintf("tcp://%s", localHostWithPort(port))
}
