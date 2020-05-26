package keeper

import (
	"encoding/json"
	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/enigmampc/EnigmaBlockchain/x/registration/internal/types"
)

// InitGenesis sets supply information for genesis.
//
func InitGenesis(ctx sdk.Context, keeper Keeper, data types.GenesisState) {

	if data.MasterCertificate != nil {
		// keeper.setMasterPublicKey(ctx, data.MasterPublic)
		keeper.setMasterCertificate(ctx, data.MasterCertificate)
		for _, storedRegInfo := range data.Registration {
			keeper.SetRegistrationInfo(ctx, storedRegInfo)
		}
	} else {
		panic("Cannot start without MasterCertificate set")
	}
}

// ExportGenesis returns a GenesisState for a given context and keeper.
func ExportGenesis(ctx sdk.Context, keeper Keeper) types.GenesisState {
	var genState types.GenesisState

	genState.MasterCertificate = *keeper.GetMasterCertificate(ctx)

	keeper.ListRegistrationInfo(ctx, func(pubkey []byte, regInfo types.RegistrationNodeInfo) bool {
		genState.Registration = append(genState.Registration, regInfo)
		return false
	})

	return genState
}

func GetGenesisStateFromAppState(cdc *codec.Codec, appState map[string]json.RawMessage) types.GenesisState {
	var genesisState types.GenesisState
	if appState[types.ModuleName] != nil {
		cdc.MustUnmarshalJSON(appState[types.ModuleName], &genesisState)
	}

	return genesisState
}