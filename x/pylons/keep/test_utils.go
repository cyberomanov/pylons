package keep

import (
	abci "github.com/tendermint/tendermint/abci/types"
	dbm "github.com/tendermint/tendermint/libs/db"
	"github.com/tendermint/tendermint/libs/log"

	codec "github.com/cosmos/cosmos-sdk/codec"
	"github.com/cosmos/cosmos-sdk/store"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/x/auth"
	"github.com/cosmos/cosmos-sdk/x/bank"
	"github.com/cosmos/cosmos-sdk/x/params"

	"github.com/MikeSofaer/pylons/x/pylons/types"
)

type TestCoinInput struct {
	Cdc  *codec.Codec
	Ctx  sdk.Context
	Ak   auth.AccountKeeper
	Pk   params.Keeper
	Bk   bank.Keeper
	FcK  auth.FeeCollectionKeeper
	PlnK Keeper
}

func GenItem(cbID string, sender sdk.AccAddress, name string) *types.Item {
	return types.NewItem(
		cbID,
		(types.DoubleInputParamList{types.DoubleInputParam{Key: "endurance", DoubleWeightTable: types.DoubleWeightTable{WeightRanges: []types.DoubleWeightRange{
			types.DoubleWeightRange{
				Lower:  100.00,
				Upper:  500.00,
				Weight: 6,
			},
			types.DoubleWeightRange{
				Lower:  501.00,
				Upper:  800.00,
				Weight: 2,
			},
		}}}}).Actualize(),
		(types.LongInputParamList{types.LongInputParam{Key: "HP", IntWeightTable: types.IntWeightTable{WeightRanges: []types.IntWeightRange{
			types.IntWeightRange{
				Lower:  100,
				Upper:  500,
				Weight: 6,
			},
			types.IntWeightRange{
				Lower:  501,
				Upper:  800,
				Weight: 2,
			},
		}}}}).Actualize(),
		(types.StringInputParamList{types.StringInputParam{Key: "Name", Value: name}}).Actualize(),
		sender,
	)
}

func SetupTestCoinInput() TestCoinInput {
	db := dbm.NewMemDB()

	cdc := codec.New()
	auth.RegisterBaseAccount(cdc)

	authCapKey := sdk.NewKVStoreKey("authCapKey")
	fckCapKey := sdk.NewKVStoreKey("fckCapKey")
	keyParams := sdk.NewKVStoreKey("params")
	tkeyParams := sdk.NewTransientStoreKey("transient_params")

	fcKey := sdk.NewKVStoreKey("fee_collection")
	cbKey := sdk.NewKVStoreKey("pylons")
	rcKey := sdk.NewKVStoreKey("pylons_recipe")
	itKey := sdk.NewKVStoreKey("pylons_item")
	execKey := sdk.NewKVStoreKey("pylons_execution")

	ms := store.NewCommitMultiStore(db)
	ms.MountStoreWithDB(authCapKey, sdk.StoreTypeIAVL, db)
	ms.MountStoreWithDB(fckCapKey, sdk.StoreTypeIAVL, db)
	ms.MountStoreWithDB(keyParams, sdk.StoreTypeIAVL, db)
	ms.MountStoreWithDB(fcKey, sdk.StoreTypeIAVL, db)
	ms.MountStoreWithDB(cbKey, sdk.StoreTypeIAVL, db)
	ms.MountStoreWithDB(rcKey, sdk.StoreTypeIAVL, db)
	ms.MountStoreWithDB(itKey, sdk.StoreTypeIAVL, db)
	ms.MountStoreWithDB(execKey, sdk.StoreTypeIAVL, db)
	ms.MountStoreWithDB(tkeyParams, sdk.StoreTypeTransient, db)
	ms.LoadLatestVersion()

	ms.GetKVStore(cbKey)

	pk := params.NewKeeper(cdc, keyParams, tkeyParams)
	ak := auth.NewAccountKeeper(
		cdc, authCapKey, pk.Subspace(auth.DefaultParamspace), auth.ProtoBaseAccount,
	)
	ctx := sdk.NewContext(ms, abci.Header{ChainID: "test-chain-id"}, false, log.NewNopLogger())

	ak.SetParams(ctx, auth.DefaultParams())

	bk := bank.NewBaseKeeper(
		ak,
		pk.Subspace(bank.DefaultParamspace),
		bank.DefaultCodespace,
	)

	fcK := auth.NewFeeCollectionKeeper(cdc, fcKey)

	plnK := NewKeeper(
		bk,
		cbKey,   // cookbook
		rcKey,   // recipe
		itKey,   // item
		execKey, // exec
		cdc,
	)

	return TestCoinInput{Cdc: cdc, Ctx: ctx, Ak: ak, Pk: pk, Bk: bk, FcK: fcK, PlnK: plnK}
}
