package web

import (
	"github.com/GoPlugin/pluginV2/core/chains/evm/types"
	"github.com/GoPlugin/pluginV2/core/services/chainlink"
	"github.com/GoPlugin/pluginV2/core/utils"
	"github.com/GoPlugin/pluginV2/core/web/presenters"
)

var ErrEVMNotEnabled = errChainDisabled{name: "EVM", envVar: "EVM_ENABLED"}

func NewEVMChainsController(app chainlink.Application) ChainsController {
	parse := func(s string) (id utils.Big, err error) {
		err = id.UnmarshalText([]byte(s))
		return
	}
	return newChainsController[utils.Big, *types.ChainCfg, presenters.EVMChainResource](
		"evm", app.GetChains().EVM, ErrEVMNotEnabled, parse, presenters.NewEVMChainResource, app.GetLogger(), app.GetAuditLogger())
}
