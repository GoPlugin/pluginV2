package types

import (
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/lib/pq"

	"gopkg.in/guregu/null.v2"

	"github.com/GoPlugin/pluginV2/core/store/models"
	"github.com/GoPlugin/pluginV2/core/utils"
)

type MercuryConfig struct {
	FeedID          common.Hash   `json:"feedID"`
	URL             *models.URL   `json:"url"`
	ServerPubKey    hexutil.Bytes `json:"serverPubKey"`
	ClientPrivKeyID string        `json:"clientPrivKeyID"`
}

type RelayConfig struct {
	MercuryConfig               *MercuryConfig
	ChainID                     *utils.Big     `json:"chainID"`
	FromBlock                   uint64         `json:"fromBlock"`
	EffectiveTransmitterAddress null.String    `json:"effectiveTransmitterAddress"`
	SendingKeys                 pq.StringArray `json:"sendingKeys"`
}
