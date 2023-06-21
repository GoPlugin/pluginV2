package log

import (
	"github.com/ethereum/go-ethereum/core/types"

	evmclient "github.com/pluginV2/core/chains/evm/client"
	evmtypes "github.com/pluginV2/core/chains/evm/types"
	"github.com/pluginV2/core/logger"
	"github.com/pluginV2/core/utils"
)

// NewTestBroadcaster creates a broadcaster with Pause/Resume enabled.
func NewTestBroadcaster(orm ORM, ethClient evmclient.Client, config Config, lggr logger.Logger, highestSavedHead *evmtypes.Head, mailMon *utils.MailboxMonitor) *broadcaster {
	b := NewBroadcaster(orm, ethClient, config, lggr, highestSavedHead, mailMon)
	b.testPause, b.testResume = make(chan struct{}), make(chan struct{})
	return b
}

func (b *broadcaster) ExportedAppendLogChannel(ch1, ch2 <-chan types.Log) chan types.Log {
	return b.appendLogChannel(ch1, ch2)
}
