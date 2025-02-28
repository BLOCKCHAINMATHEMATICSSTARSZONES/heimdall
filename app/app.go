package app

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"

	"github.com/maticnetwork/heimdall/x/bor"

	"github.com/cosmos/cosmos-sdk/baseapp"
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/grpc/tmservice"
	"github.com/cosmos/cosmos-sdk/client/rpc"
	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/cosmos/cosmos-sdk/codec/types"
	"github.com/cosmos/cosmos-sdk/server/api"
	"github.com/cosmos/cosmos-sdk/server/config"
	"github.com/cosmos/cosmos-sdk/testutil/testdata"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/module"
	"github.com/cosmos/cosmos-sdk/version"
	"github.com/cosmos/cosmos-sdk/x/auth"
	"github.com/cosmos/cosmos-sdk/x/auth/ante"
	authrest "github.com/cosmos/cosmos-sdk/x/auth/client/rest"
	authkeeper "github.com/cosmos/cosmos-sdk/x/auth/keeper"
	authsims "github.com/cosmos/cosmos-sdk/x/auth/simulation"
	authtx "github.com/cosmos/cosmos-sdk/x/auth/tx"
	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"
	"github.com/cosmos/cosmos-sdk/x/bank"
	bankkeeper "github.com/cosmos/cosmos-sdk/x/bank/keeper"
	banktypes "github.com/cosmos/cosmos-sdk/x/bank/types"
	"github.com/cosmos/cosmos-sdk/x/genutil"
	genutiltypes "github.com/cosmos/cosmos-sdk/x/genutil/types"
	"github.com/cosmos/cosmos-sdk/x/params"
	paramskeeper "github.com/cosmos/cosmos-sdk/x/params/keeper"
	paramstypes "github.com/cosmos/cosmos-sdk/x/params/types"
	"github.com/gorilla/mux"
	"github.com/rakyll/statik/fs"
	abci "github.com/tendermint/tendermint/abci/types"
	tmjson "github.com/tendermint/tendermint/libs/json"
	"github.com/tendermint/tendermint/libs/log"
	tmos "github.com/tendermint/tendermint/libs/os"
	dbm "github.com/tendermint/tm-db"

	// unnamed import of statik for swagger UI support
	_ "github.com/cosmos/cosmos-sdk/client/docs/statik"

	// "github.com/cosmos/cosmos-sdk/x/slashing"
	// slashingkeeper "github.com/cosmos/cosmos-sdk/x/slashing/keeper"
	// slashingtypes "github.com/cosmos/cosmos-sdk/x/slashing/types"

	hmparams "github.com/maticnetwork/heimdall/app/params"
	"github.com/maticnetwork/heimdall/helper"
	hmtypes "github.com/maticnetwork/heimdall/types"
	"github.com/maticnetwork/heimdall/types/common"
	hmCommonTypes "github.com/maticnetwork/heimdall/types/common"
	hmmodule "github.com/maticnetwork/heimdall/types/module"
	"github.com/maticnetwork/heimdall/x/chainmanager"
	chainKeeper "github.com/maticnetwork/heimdall/x/chainmanager/keeper"
	chainmanagerTypes "github.com/maticnetwork/heimdall/x/chainmanager/types"
	"github.com/maticnetwork/heimdall/x/checkpoint"
	checkpointkeeper "github.com/maticnetwork/heimdall/x/checkpoint/keeper"
	checkpointtypes "github.com/maticnetwork/heimdall/x/checkpoint/types"
	"github.com/maticnetwork/heimdall/x/clerk"
	clerkkeeper "github.com/maticnetwork/heimdall/x/clerk/keeper"
	clerktypes "github.com/maticnetwork/heimdall/x/clerk/types"
	"github.com/maticnetwork/heimdall/x/gov"
	govkeeper "github.com/maticnetwork/heimdall/x/gov/keeper"
	govtypes "github.com/maticnetwork/heimdall/x/gov/types"
	"github.com/maticnetwork/heimdall/x/sidechannel"
	sidechannelkeeper "github.com/maticnetwork/heimdall/x/sidechannel/keeper"
	sidechanneltypes "github.com/maticnetwork/heimdall/x/sidechannel/types"
	"github.com/maticnetwork/heimdall/x/staking"
	stakingkeeper "github.com/maticnetwork/heimdall/x/staking/keeper"
	stakingtypes "github.com/maticnetwork/heimdall/x/staking/types"
	"github.com/maticnetwork/heimdall/x/topup"
	topupkeeper "github.com/maticnetwork/heimdall/x/topup/keeper"
	topuptypes "github.com/maticnetwork/heimdall/x/topup/types"

	borkeeper "github.com/maticnetwork/heimdall/x/bor/keeper"
	bortypes "github.com/maticnetwork/heimdall/x/bor/types"

	// unnamed import of statik for swagger UI support
	_ "github.com/maticnetwork/heimdall/client/docs/statik"
)

const appName = "Heimdall"

var (
	// DefaultNodeHome default home directories for the application daemon
	DefaultNodeHome string

	// ModuleBasics defines the module BasicManager is in charge of setting up basic,
	// non-dependant module elements, such as codec registration
	// and genesis verification.
	ModuleBasics = module.NewBasicManager(
		auth.AppModuleBasic{},
		genutil.AppModuleBasic{},
		bank.AppModuleBasic{},
		chainmanager.AppModuleBasic{},
		sidechannel.AppModuleBasic{},
		staking.AppModuleBasic{},
		params.AppModuleBasic{},
		gov.AppModuleBasic{},
		checkpoint.AppModuleBasic{},
		topup.AppModuleBasic{},
		clerk.AppModuleBasic{},
		bor.AppModuleBasic{},
	)

	// module account permissions
	maccPerms = map[string][]string{
		authtypes.FeeCollectorName: nil,
		govtypes.ModuleName:        {},
	}

	// module accounts that are allowed to receive tokens
	allowedReceivingModAcc = map[string]bool{
		// distrtypes.ModuleName: true,
	}
)

var _ App = (*HeimdallApp)(nil)

// HeimdallApp extends an ABCI application, but with most of its parameters exported.
// They are exported for convenience in creating helper functions, as object
// capabilities aren't needed for testing.
type HeimdallApp struct {
	*baseapp.BaseApp
	legacyAmino       *codec.LegacyAmino
	appCodec          codec.Marshaler
	interfaceRegistry types.InterfaceRegistry
	txDecoder         sdk.TxDecoder

	invCheckPeriod uint

	// keys to access the substores
	keys  map[string]*sdk.KVStoreKey
	tkeys map[string]*sdk.TransientStoreKey

	// keepers
	AccountKeeper     authkeeper.AccountKeeper
	BankKeeper        bankkeeper.Keeper
	ChainKeeper       chainKeeper.Keeper
	ClerkKeeper       clerkkeeper.Keeper
	SidechannelKeeper sidechannelkeeper.Keeper
	StakingKeeper     stakingkeeper.Keeper
	ParamsKeeper      paramskeeper.Keeper
	GovKeeper         govkeeper.Keeper
	CheckpointKeeper  checkpointkeeper.Keeper
	TopupKeeper       topupkeeper.Keeper
	BorKeeper         borkeeper.Keeper

	// side router
	sideRouter hmtypes.SideRouter

	// contract caller
	caller helper.ContractCaller

	// the module manager
	mm *module.Manager

	// simulation manager
	sm *module.SimulationManager
}

var logger = helper.Logger.With("module", "app")

func init() {
	userHomeDir, err := os.UserHomeDir()
	if err != nil {
		panic(err)
	}

	DefaultNodeHome = filepath.Join(userHomeDir, ".heimdalld")
}

// NewHeimdallApp returns a reference to an initialized HeimdallApp.
func NewHeimdallApp(
	logger log.Logger,
	db dbm.DB,
	traceStore io.Writer,
	loadLatest bool,
	skipUpgradeHeights map[int64]bool,
	homePath string,
	invCheckPeriod uint,
	encodingConfig hmparams.EncodingConfig,
	baseAppOptions ...func(*baseapp.BaseApp),
) *HeimdallApp {
	// TODO: Remove legacyAmino in favor of appCodec once all modules are migrated.
	appCodec := encodingConfig.Marshaler
	legacyAmino := encodingConfig.Amino
	interfaceRegistry := encodingConfig.InterfaceRegistry
	txDecoder := encodingConfig.TxConfig.TxDecoder()

	bApp := baseapp.NewBaseApp(appName, logger, db, txDecoder, baseAppOptions...)
	bApp.SetCommitMultiStoreTracer(traceStore)
	bApp.SetAppVersion(version.Version)
	bApp.SetInterfaceRegistry(interfaceRegistry)

	//
	// Keys
	//

	keys := sdk.NewKVStoreKeys(
		authtypes.StoreKey,
		banktypes.StoreKey,
		chainmanagerTypes.StoreKey,
		clerktypes.StoreKey,
		sidechanneltypes.StoreKey,
		stakingtypes.StoreKey,
		checkpointtypes.StoreKey,
		// distrtypes.StoreKey,
		// slashingtypes.StoreKey,
		govtypes.StoreKey,
		paramstypes.StoreKey,
		topuptypes.StoreKey,
		bortypes.StoreKey,
	)
	tkeys := sdk.NewTransientStoreKeys(paramstypes.TStoreKey)

	app := &HeimdallApp{
		BaseApp:           bApp,
		legacyAmino:       legacyAmino,
		appCodec:          appCodec,
		interfaceRegistry: interfaceRegistry,
		invCheckPeriod:    invCheckPeriod,
		keys:              keys,
		tkeys:             tkeys,
		txDecoder:         txDecoder,
	}

	//
	// module communicator
	//

	moduleCommunicator := ModuleCommunicator{App: app}

	//
	// Keepers
	//

	// create params keeper
	app.ParamsKeeper = initParamsKeeper(appCodec, legacyAmino, keys[paramstypes.StoreKey], tkeys[paramstypes.TStoreKey])
	// set the BaseApp's parameter store
	bApp.SetParamStore(app.ParamsKeeper.Subspace(baseapp.Paramspace).WithKeyTable(paramskeeper.ConsensusParamsKeyTable()))

	//chainmanager keeper
	app.ChainKeeper = chainKeeper.NewKeeper(
		appCodec,
		keys[chainmanagerTypes.StoreKey], // target store
		app.GetSubspace(chainmanagerTypes.ModuleName),
		app.caller,
	)
	// account keeper
	app.AccountKeeper = authkeeper.NewAccountKeeper(
		appCodec,
		app.keys[authtypes.StoreKey],
		app.GetSubspace(authtypes.ModuleName),
		authtypes.ProtoBaseAccount,
		MacPerms(),
	)

	app.BankKeeper = bankkeeper.NewBaseKeeper(
		appCodec,
		keys[banktypes.StoreKey],
		app.AccountKeeper,
		app.GetSubspace(banktypes.ModuleName),
		app.BlockedAddrs(),
	)

	app.SidechannelKeeper = sidechannelkeeper.NewKeeper(
		appCodec,
		keys[sidechanneltypes.StoreKey],
	)

	app.StakingKeeper = stakingkeeper.NewKeeper(
		appCodec,
		keys[stakingtypes.StoreKey], // target store
		app.GetSubspace(stakingtypes.ModuleName),
		app.ChainKeeper,
		app.BankKeeper,
		moduleCommunicator,
	)
	app.CheckpointKeeper = checkpointkeeper.NewKeeper(
		appCodec,
		keys[checkpointtypes.StoreKey], // target store
		app.GetSubspace(checkpointtypes.ModuleName),
		app.StakingKeeper,
		app.ChainKeeper,
		moduleCommunicator,
	)

	govRouter := govtypes.NewRouter()
	govRouter.AddRoute(govtypes.RouterKey, govtypes.ProposalHandler)

	app.GovKeeper = govkeeper.NewKeeper(
		appCodec,
		keys[govtypes.StoreKey], // target store
		app.GetSubspace(govtypes.ModuleName),
		app.BankKeeper,
		govRouter,
		&app.StakingKeeper,
		app.AccountKeeper,
	)

	app.ClerkKeeper = clerkkeeper.NewKeeper(
		appCodec,
		keys[clerktypes.StoreKey], // target store
		app.ChainKeeper,
	)

	app.TopupKeeper = topupkeeper.NewKeeper(
		appCodec,
		keys[topuptypes.StoreKey],
		app.GetSubspace(topuptypes.ModuleName),
		app.ChainKeeper,
		app.BankKeeper,
		app.StakingKeeper,
	)

	app.BorKeeper = borkeeper.NewKeeper(
		appCodec, keys[bortypes.StoreKey],
		app.GetSubspace(bortypes.ModuleName),
		app.ChainKeeper,
		app.StakingKeeper,
		app.CheckpointKeeper,
		app.caller,
	)

	// Contract caller
	contractCallerObj, err := helper.NewContractCaller()
	if err != nil {
		tmos.Exit(err.Error())
	}
	app.caller = contractCallerObj

	// register the staking hooks
	// NOTE: stakingKeeper above is passed by reference, so that it will contain these hooks
	// app.StakingKeeper = *stakingKeeper.SetHooks(
	// 	stakingtypes.NewMultiStakingHooks(app.DistrKeeper.Hooks()),
	// )

	// NOTE: Any module instantiated in the module manager that is later modified
	// must be passed by reference here.
	app.mm = module.NewManager(
		genutil.NewAppModule(
			app.AccountKeeper,
			app.StakingKeeper,
			app.BaseApp.DeliverTx,
			encodingConfig.TxConfig,
		),
		auth.NewAppModule(appCodec, app.AccountKeeper, nil),
		bank.NewAppModule(appCodec, app.BankKeeper, app.AccountKeeper),
		sidechannel.NewAppModule(appCodec, app.SidechannelKeeper),
		chainmanager.NewAppModule(appCodec, app.ChainKeeper),
		staking.NewAppModule(appCodec, app.StakingKeeper, &app.caller),
		clerk.NewAppModule(appCodec, app.ClerkKeeper, &app.caller),
		gov.NewAppModule(appCodec, app.GovKeeper, app.AccountKeeper, app.BankKeeper),
		params.NewAppModule(app.ParamsKeeper),
		checkpoint.NewAppModule(appCodec, app.CheckpointKeeper, &app.caller),
		bor.NewAppModule(appCodec, app.BorKeeper, &app.caller),
		topup.NewAppModule(appCodec, app.TopupKeeper, &app.caller),
	)

	// During begin block slashing happens after distr.BeginBlocker so that
	// there is nothing left over in the validator fee pool, so as to keep the
	// CanWithdrawInvariant invariant.
	// NOTE: staking module is required if HistoricalEntries param > 0
	app.mm.SetOrderBeginBlockers(
		sidechanneltypes.ModuleName,
		stakingtypes.ModuleName,
	)
	app.mm.SetOrderEndBlockers(
		sidechanneltypes.ModuleName,
		stakingtypes.ModuleName,
		govtypes.ModuleName,
	)
	app.mm.SetOrderEndBlockers(
		stakingtypes.ModuleName,
		govtypes.ModuleName,
	)

	// NOTE: The genutils module must occur after staking so that pools are
	// properly initialized with tokens from genesis accounts.
	// NOTE: Capability module must occur first so that it can initialize any capabilities
	// so that other modules that want to create or claim capabilities afterwards in InitChain
	// can do so safely.
	app.mm.SetOrderInitGenesis(
		authtypes.ModuleName,
		banktypes.ModuleName,
		sidechanneltypes.ModuleName,
		chainmanagerTypes.ModuleName,
		stakingtypes.ModuleName,
		checkpointtypes.ModuleName,
		clerktypes.ModuleName,
		genutiltypes.ModuleName,
		govtypes.ModuleName,
		bortypes.ModuleName,
		topuptypes.ModuleName,
	)

	// app.mm.RegisterInvariants(&app.CrisisKeeper)
	app.mm.RegisterRoutes(app.Router(), app.QueryRouter(), encodingConfig.Amino)
	app.mm.RegisterServices(module.NewConfigurator(app.MsgServiceRouter(), app.GRPCQueryRouter()))

	// side router
	app.sideRouter = hmtypes.NewSideRouter()
	for _, m := range app.mm.Modules {
		if m.Route().Path() != "" { //nolint
			if sm, ok := m.(hmmodule.SideModule); ok {
				app.sideRouter.AddRoute(m.Route().Path(), &hmtypes.SideHandlers{ //nolint
					SideTxHandler: sm.NewSideTxHandler(),
					PostTxHandler: sm.NewPostTxHandler(),
				})
			}
		}
	}
	app.sideRouter.Seal()

	// add test gRPC service for testing gRPC queries in isolation
	testdata.RegisterQueryServer(app.GRPCQueryRouter(), testdata.QueryImpl{})

	// create the simulation manager and define the order of the modules for deterministic simulations
	//
	// NOTE: this is not required apps that don't use the simulator for fuzz testing
	// transactions

	// TODO : replace nil with staking.NewAppModule
	app.sm = module.NewSimulationManager(
		auth.NewAppModule(appCodec, app.AccountKeeper, authsims.RandomGenesisAccounts),
		bank.NewAppModule(appCodec, app.BankKeeper, app.AccountKeeper),
		sidechannel.NewAppModule(appCodec, app.SidechannelKeeper),
		// staking.NewAppModule(appCodec),
		nil,
		params.NewAppModule(app.ParamsKeeper),
	)

	// TODO : uncomment later
	// app.sm.RegisterStoreDecoders()

	// initialize stores
	app.MountKVStores(keys)
	app.MountTransientStores(tkeys)

	// initialize BaseApp
	app.SetInitChainer(app.InitChainer)
	app.SetBeginBlocker(app.BeginBlocker)
	app.SetAnteHandler(
		ante.NewAnteHandler(
			app.AccountKeeper,
			app.BankKeeper,
			ante.DefaultSigVerificationGasConsumer,
			encodingConfig.TxConfig.SignModeHandler(),
		),
	)
	app.SetEndBlocker(app.EndBlocker)

	// side-tx processor
	app.SetPostDeliverTxHandler(app.PostDeliverTxHandler)
	app.SetBeginSideBlocker(app.BeginSideBlocker)
	app.SetDeliverSideTxHandler(app.DeliverSideTxHandler)

	if loadLatest {
		if err := app.LoadLatestVersion(); err != nil {
			tmos.Exit(err.Error())
		}

		// Initialize and seal the capability keeper so all persistent capabilities
		// are loaded in-memory and prevent any further modules from creating scoped
		// sub-keepers.
		// This must be done during creation of baseapp rather than in InitChain so
		// that in-memory capabilities get regenerated on app restart.
		// Note that since this reads from the store, we can only perform it when
		// `loadLatest` is set to true.
		// ctx := app.BaseApp.NewUncachedContext(true, tmproto.Header{})
		// app.CapabilityKeeper.InitializeAndSeal(ctx)
	}

	return app
}

// MakeCodecs constructs the *std.Codec and *codec.LegacyAmino instances used by
// HeimdallApp. It is useful for tests and clients who do not want to construct the
// full HeimdallApp
func MakeCodecs() (codec.Marshaler, *codec.LegacyAmino) {
	encodingConfig := MakeEncodingConfig()
	return encodingConfig.Marshaler, encodingConfig.Amino
}

// Name returns the name of the App
func (app *HeimdallApp) Name() string { return app.BaseApp.Name() }

// BeginBlocker application updates every begin block
func (app *HeimdallApp) BeginBlocker(ctx sdk.Context, req abci.RequestBeginBlock) abci.ResponseBeginBlock {
	if len(req.Header.GetProposerAddress()) != 0 {
		proposerAddress, _ := sdk.AccAddressFromHex(string(req.Header.GetProposerAddress()))
		app.ChainKeeper.SetBlockProposer(
			ctx,
			proposerAddress,
		)
	}

	return app.mm.BeginBlock(ctx, req)
}

// EndBlocker application updates every end block
func (app *HeimdallApp) EndBlocker(ctx sdk.Context, req abci.RequestEndBlock) abci.ResponseEndBlock {

	if proposer, ok := app.ChainKeeper.GetBlockProposer(ctx); ok {
		moduleAccount := app.AccountKeeper.GetModuleAccount(ctx, authtypes.FeeCollectorName)
		amount := app.BankKeeper.GetBalance(ctx, moduleAccount.GetAddress(), stakingtypes.FeeToken)
		if !amount.IsZero() {
			coins := sdk.Coins{sdk.Coin{Denom: stakingtypes.FeeToken, Amount: sdk.NewInt(1000)}} //check amount
			if err := app.BankKeeper.SendCoinsFromModuleToAccount(ctx, authtypes.FeeCollectorName, proposer, coins); err != nil {
				logger.Error("EndBlocker | SendCoinsFromModuleToAccount", "Error", err)
			}
		}

		// remove block proposer
		app.ChainKeeper.RemoveBlockProposer(ctx)
	}
	var tmValUpdates []abci.ValidatorUpdate

	// --- Start update to new validators
	currentValidatorSet := app.StakingKeeper.GetValidatorSet(ctx)
	allValidators := app.StakingKeeper.GetAllValidators(ctx)
	ackCount := app.CheckpointKeeper.GetACKCount(ctx)

	// get validator updates
	setUpdates := helper.GetUpdatedValidators(
		currentValidatorSet, // pointer to current validator set -- UpdateValidators will modify it
		allValidators,       // All validators
		ackCount,            // ack count
	)

	if len(setUpdates) > 0 {
		// create new validator set
		if err := currentValidatorSet.UpdateWithChangeSet(setUpdates); err != nil {
			// return with nothing
			logger.Error("Unable to update current validator set", "Error", err)
			return abci.ResponseEndBlock{}
		}

		// increment proposer priority
		currentValidatorSet.IncrementProposerPriority(1)

		// validator set change
		logger.Debug("[ENDBLOCK] Updated current validator set", "proposer", currentValidatorSet.GetProposer())

		// save set in store
		if err := app.StakingKeeper.UpdateValidatorSetInStore(ctx, currentValidatorSet); err != nil {
			// return with nothing
			logger.Error("Unable to update current validator set in state", "Error", err)
			return abci.ResponseEndBlock{}
		}

		// convert updates from map to array
		for _, v := range setUpdates {
			tmValUpdates = append(tmValUpdates, abci.ValidatorUpdate{
				Power:  int64(v.VotingPower),
				PubKey: hmCommonTypes.NewPubKeyFromHex(v.PubKey).TMProtoCryptoPubKey(),
			})
		}
	}

	// end block
	app.mm.EndBlock(ctx, req)

	// send validator updates to peppermint
	return abci.ResponseEndBlock{
		ValidatorUpdates: tmValUpdates,
	}
}

// InitChainer application update at chain initialization
func (app *HeimdallApp) InitChainer(ctx sdk.Context, req abci.RequestInitChain) abci.ResponseInitChain {
	var genesisState GenesisState
	if err := tmjson.Unmarshal(req.AppStateBytes, &genesisState); err != nil {
		panic(err)
	}
	// get validator updates
	//if err := ModuleBasics.ValidateGenesis(app.appCodec,ctx, genesisState); err != nil {
	//	panic(err)
	//}

	// check fee collector module account
	if moduleAcc := app.AccountKeeper.GetModuleAccount(ctx, authtypes.FeeCollectorName); moduleAcc == nil {
		panic(fmt.Sprintf("%s module account has not been set", authtypes.FeeCollectorName))
	}

	// Init genesis
	app.mm.InitGenesis(ctx, app.AppCodec(), genesisState)

	// get staking state
	stakingState := stakingtypes.GetGenesisStateFromAppState(app.AppCodec(), genesisState)
	checkpointState := checkpointtypes.GetGenesisStateFromAppState(app.AppCodec(), genesisState)

	// check if validator is current validator
	// add to val updates else skip
	var valUpdates []abci.ValidatorUpdate
	for _, validator := range stakingState.Validators {
		// TODO use checkpoint state to get current validator set once checkpoint module is ready

		if validator.IsCurrentValidator(checkpointState.AckCount) {
			// convert to Validator Update
			updateVal := abci.ValidatorUpdate{
				Power:  int64(validator.VotingPower),
				PubKey: common.NewPubKeyFromHex(validator.PubKey).TMProtoCryptoPubKey(),
			}
			// Add validator to validator updated to be processed below
			valUpdates = append(valUpdates, updateVal)
		}
	}

	// TODO make sure old validtors dont go in validator updates ie deactivated validators have to be removed
	// udpate validators
	return abci.ResponseInitChain{
		// validator updates
		Validators: valUpdates,
	}
}

// LoadHeight loads a particular height
func (app *HeimdallApp) LoadHeight(height int64) error {
	return app.LoadVersion(height)
}

// ModuleAccountAddrs returns all the app's module account addresses.
func (app *HeimdallApp) ModuleAccountAddrs() map[string]bool {
	modAccAddrs := make(map[string]bool)
	for acc := range maccPerms {
		modAccAddrs[authtypes.NewModuleAddress(acc).String()] = true
	}

	return modAccAddrs
}

// BlockedAddrs returns all the app's module account addresses that are not
// allowed to receive external tokens.
func (app *HeimdallApp) BlockedAddrs() map[string]bool {
	blockedAddrs := make(map[string]bool)

	for acc := range maccPerms {
		blockedAddrs[authtypes.NewModuleAddress(acc).String()] = !allowedReceivingModAcc[acc]
	}

	return blockedAddrs
}

// LegacyAmino returns HeimdallApp's amino codec.
//
// NOTE: This is solely to be used for testing purposes as it may be desirable
// for modules to register their own custom testing types.
func (app *HeimdallApp) LegacyAmino() *codec.LegacyAmino {
	return app.legacyAmino
}

// AppCodec returns HeimdallApp's app codec.
//
// NOTE: This is solely to be used for testing purposes as it may be desirable
// for modules to register their own custom testing types.
func (app *HeimdallApp) AppCodec() codec.Marshaler {
	return app.appCodec
}

// InterfaceRegistry returns HeimdallApp's InterfaceRegistry
func (app *HeimdallApp) InterfaceRegistry() types.InterfaceRegistry {
	return app.interfaceRegistry
}

// GetKey returns the KVStoreKey for the provided store key.
//
// NOTE: This is solely to be used for testing purposes.
func (app *HeimdallApp) GetKey(storeKey string) *sdk.KVStoreKey {
	return app.keys[storeKey]
}

// GetTKey returns the TransientStoreKey for the provided store key.
//
// NOTE: This is solely to be used for testing purposes.
func (app *HeimdallApp) GetTKey(storeKey string) *sdk.TransientStoreKey {
	return app.tkeys[storeKey]
}

// GetSideRouter returns side-tx router
func (app *HeimdallApp) GetSideRouter() hmtypes.SideRouter {
	return app.sideRouter
}

// SetSideRouter sets side-tx router
//
// NOTE: This is solely to be used for testing purposes.
func (app *HeimdallApp) SetSideRouter(router hmtypes.SideRouter) {
	app.sideRouter = router
}

// SetTxDecoder sets tx decoder
//
// NOTE: This is solely to be used for testing purposes.
func (app *HeimdallApp) SetTxDecoder(decoder sdk.TxDecoder) {
	app.txDecoder = decoder
}

// GetSubspace returns a param subspace for a given module name.
//
// NOTE: This is solely to be used for testing purposes.
func (app *HeimdallApp) GetSubspace(moduleName string) paramstypes.Subspace {
	subspace, _ := app.ParamsKeeper.GetSubspace(moduleName)
	return subspace
}

// SimulationManager implements the SimulationApp interface
func (app *HeimdallApp) SimulationManager() *module.SimulationManager {
	return app.sm
}

// RegisterAPIRoutes registers all application module routes with the provided
// API server.
func (app *HeimdallApp) RegisterAPIRoutes(apiSvr *api.Server, apiConfig config.APIConfig) {
	clientCtx := apiSvr.ClientCtx
	rpc.RegisterRoutes(clientCtx, apiSvr.Router)
	// Register legacy tx routes.
	authrest.RegisterTxRoutes(clientCtx, apiSvr.Router)

	// Register new tx routes from grpc-gateway.
	authtx.RegisterGRPCGatewayRoutes(clientCtx, apiSvr.GRPCRouter)
	// Register new tendermint queries routes from grpc-gateway.
	tmservice.RegisterGRPCGatewayRoutes(clientCtx, apiSvr.GRPCRouter)

	ModuleBasics.RegisterRESTRoutes(clientCtx, apiSvr.Router)
	ModuleBasics.RegisterGRPCGatewayRoutes(clientCtx, apiSvr.GRPCRouter)

	// register swagger API from root so that other applications can override easily
	if apiConfig.Swagger {
		RegisterSwaggerAPI(clientCtx, apiSvr.Router)
	}
}

// RegisterTxService implements the Application.RegisterTxService method.
func (app *HeimdallApp) RegisterTxService(clientCtx client.Context) {
	authtx.RegisterTxService(app.BaseApp.GRPCQueryRouter(), clientCtx, app.BaseApp.Simulate, app.interfaceRegistry)
}

// RegisterTendermintService implements the Application.RegisterTendermintService method.
func (app *HeimdallApp) RegisterTendermintService(clientCtx client.Context) {
	tmservice.RegisterTendermintService(app.BaseApp.GRPCQueryRouter(), clientCtx, app.interfaceRegistry)
}

// RegisterSwaggerAPI registers swagger route with API Server
func RegisterSwaggerAPI(ctx client.Context, rtr *mux.Router) {
	statikFS, err := fs.New()
	if err != nil {
		panic(err)
	}

	staticServer := http.FileServer(statikFS)
	rtr.PathPrefix("/swagger/").Handler(http.StripPrefix("/swagger/", staticServer))
}

// GetMaccPerms returns a copy of the module account permissions
func GetMaccPerms() map[string][]string {
	dupMaccPerms := make(map[string][]string)
	for k, v := range maccPerms {
		dupMaccPerms[k] = v
	}
	return dupMaccPerms
}

// initParamsKeeper init params keeper and its subspaces
func initParamsKeeper(appCodec codec.BinaryMarshaler, legacyAmino *codec.LegacyAmino, key, tkey sdk.StoreKey) paramskeeper.Keeper {
	paramsKeeper := paramskeeper.NewKeeper(appCodec, legacyAmino, key, tkey)

	paramsKeeper.Subspace(authtypes.ModuleName)
	paramsKeeper.Subspace(banktypes.ModuleName)
	paramsKeeper.Subspace(stakingtypes.ModuleName)
	paramsKeeper.Subspace(chainmanagerTypes.ModuleName)
	paramsKeeper.Subspace(sidechanneltypes.ModuleName)
	paramsKeeper.Subspace(checkpointtypes.ModuleName)
	paramsKeeper.Subspace(govtypes.ModuleName).WithKeyTable(govtypes.ParamKeyTable())
	paramsKeeper.Subspace(bortypes.ModuleName)
	paramsKeeper.Subspace(clerktypes.ModuleName)
	paramsKeeper.Subspace(topuptypes.ModuleName)

	return paramsKeeper
}
