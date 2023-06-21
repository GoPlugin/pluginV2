package web

import (
	"github.com/gin-gonic/gin"

	"github.com/pluginV2/core/chains/evm/types"
	"github.com/pluginV2/core/services/chainlink"
	"github.com/pluginV2/core/utils"
	"github.com/pluginV2/core/web/presenters"
)

func NewEVMNodesController(app chainlink.Application) NodesController {
	parse := func(s string) (id utils.Big, err error) {
		err = id.UnmarshalText([]byte(s))
		return
	}
	return newNodesController[utils.Big, types.Node, presenters.EVMNodeResource](
		app.GetChains().EVM, ErrEVMNotEnabled, parse, presenters.NewEVMNodeResource, func(c *gin.Context) (types.Node, error) {
			var request types.NewNode

			if err := c.ShouldBindJSON(&request); err != nil {
				return types.Node{}, err
			}
			return types.Node{
				Name:       request.Name,
				EVMChainID: request.EVMChainID,
				WSURL:      request.WSURL,
				HTTPURL:    request.HTTPURL,
				SendOnly:   request.SendOnly,
			}, nil
		},
		app.GetLogger(),
		app.GetAuditLogger(),
	)
}
