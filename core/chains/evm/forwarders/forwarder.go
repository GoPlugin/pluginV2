package forwarders

import (
	"time"

	"github.com/ethereum/go-ethereum/common"

	"github.com/GoPlugin/pluginV2/core/utils"
)

// Forwarder is the struct for Forwarder Addresses
type Forwarder struct {
	ID         int64
	Address    common.Address
	EVMChainID utils.Big
	CreatedAt  time.Time
	UpdatedAt  time.Time
}
