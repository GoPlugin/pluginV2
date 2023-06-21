package web

import (
	"github.com/pluginV2-starknet/relayer/pkg/chainlink/db"

	"github.com/pluginV2/core/services/chainlink"
	"github.com/pluginV2/core/web/presenters"
)

func NewStarkNetChainsController(app chainlink.Application) ChainsController {
	return newChainsController[string, *db.ChainCfg]("starknet", app.GetChains().StarkNet, ErrStarkNetNotEnabled,
		func(s string) (string, error) { return s, nil }, presenters.NewStarkNetChainResource, app.GetLogger(), app.GetAuditLogger())
}
