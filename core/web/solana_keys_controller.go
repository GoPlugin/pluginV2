package web

import (
	"github.com/GoPlugin/pluginV2/core/services/chainlink"
	"github.com/GoPlugin/pluginV2/core/services/keystore/keys/solkey"
	"github.com/GoPlugin/pluginV2/core/web/presenters"
)

func NewSolanaKeysController(app chainlink.Application) KeysController {
	return NewKeysController[solkey.Key, presenters.SolanaKeyResource](app.GetKeyStore().Solana(), app.GetLogger(), app.GetAuditLogger(),
		"solanaKey", presenters.NewSolanaKeyResource, presenters.NewSolanaKeyResources)
}
