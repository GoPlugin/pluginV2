package evm

import (
	"github.com/smartcontractkit/sqlx"

	"github.com/GoPlugin/pluginV2/core/chains"
	"github.com/GoPlugin/pluginV2/core/chains/evm/types"
	"github.com/GoPlugin/pluginV2/core/logger"
	"github.com/GoPlugin/pluginV2/core/services/pg"
	"github.com/GoPlugin/pluginV2/core/utils"
)

// NewORM returns a new EVM ORM
func NewORM(db *sqlx.DB, lggr logger.Logger, cfg pg.QConfig) types.ORM {
	q := pg.NewQ(db, lggr.Named("EVMORM"), cfg)
	return chains.NewORM[utils.Big, *types.ChainCfg, types.Node](q, "evm", "ws_url", "http_url", "send_only")
}
