package web

import (
	"github.com/GoPlugin/pluginV2/core/services/chainlink"
	"github.com/GoPlugin/pluginV2/core/services/keystore/keys/dkgsignkey"
	"github.com/GoPlugin/pluginV2/core/web/presenters"
)

func NewDKGSignKeysController(app chainlink.Application) KeysController {
	return NewKeysController[dkgsignkey.Key, presenters.DKGSignKeyResource](
		app.GetKeyStore().DKGSign(),
		app.GetLogger(),
		app.GetAuditLogger(),
		"dkgsignKey",
		presenters.NewDKGSignKeyResource,
		presenters.NewDKGSignKeyResources)
}
