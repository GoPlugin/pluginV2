package web

import (
	"github.com/smartcontractkit/chainlink-solana/pkg/solana/db"

	"github.com/GoPlugin/pluginV2/core/services/chainlink"
	"github.com/GoPlugin/pluginV2/core/web/presenters"
)

func NewSolanaChainsController(app chainlink.Application) ChainsController {
	return newChainsController[string, *db.ChainCfg]("solana", app.GetChains().Solana, ErrSolanaNotEnabled,
		func(s string) (string, error) { return s, nil }, presenters.NewSolanaChainResource, app.GetLogger(), app.GetAuditLogger())
}
