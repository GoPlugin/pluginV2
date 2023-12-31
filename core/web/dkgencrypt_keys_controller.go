package web

import (
	"github.com/GoPlugin/pluginV2/core/services/chainlink"
	"github.com/GoPlugin/pluginV2/core/services/keystore/keys/dkgencryptkey"
	"github.com/GoPlugin/pluginV2/core/web/presenters"
)

func NewDKGEncryptKeysController(app chainlink.Application) KeysController {
	return NewKeysController[dkgencryptkey.Key, presenters.DKGEncryptKeyResource](
		app.GetKeyStore().DKGEncrypt(),
		app.GetLogger(),
		app.GetAuditLogger(),
		"dkgencryptKey",
		presenters.NewDKGEncryptKeyResource,
		presenters.NewDKGEncryptKeyResources)
}
