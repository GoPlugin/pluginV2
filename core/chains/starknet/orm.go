package starknet

import (
	"github.com/smartcontractkit/sqlx"

	starknetdb "github.com/smartcontractkit/chainlink-starknet/relayer/pkg/chainlink/db"

	"github.com/GoPlugin/pluginV2/core/chains"
	"github.com/GoPlugin/pluginV2/core/chains/starknet/types"
	"github.com/GoPlugin/pluginV2/core/logger"
	"github.com/GoPlugin/pluginV2/core/services/pg"
)

// https://app.shortcut.com/chainlinklabs/story/33622/remove-legacy-config
func NewORM(db *sqlx.DB, lggr logger.Logger, cfg pg.QConfig) types.ORM {
	q := pg.NewQ(db, lggr.Named("ORM"), cfg)
	return chains.NewORM[string, *starknetdb.ChainCfg, starknetdb.Node](q, "starknet", "url")
}

func NewORMImmut(cfgs chains.ChainConfig[string, *starknetdb.ChainCfg, starknetdb.Node]) types.ORM {
	return chains.NewORMImmut(cfgs)
}
