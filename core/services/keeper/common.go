package keeper

import (
	"time"

	"github.com/ethereum/go-ethereum/common"

	"github.com/GoPlugin/pluginV2/core/assets"
	"github.com/GoPlugin/pluginV2/core/chains/evm/types"
	"github.com/GoPlugin/pluginV2/core/gethwrappers/generated/keeper_registry_wrapper1_1"
	"github.com/GoPlugin/pluginV2/core/gethwrappers/generated/keeper_registry_wrapper1_2"
	"github.com/GoPlugin/pluginV2/core/gethwrappers/generated/keeper_registry_wrapper1_3"
	"github.com/GoPlugin/pluginV2/core/services/pg"
)

var Registry1_1ABI = types.MustGetABI(keeper_registry_wrapper1_1.KeeperRegistryABI)
var Registry1_2ABI = types.MustGetABI(keeper_registry_wrapper1_2.KeeperRegistryABI)
var Registry1_3ABI = types.MustGetABI(keeper_registry_wrapper1_3.KeeperRegistryABI)

type Config interface {
	EvmEIP1559DynamicFees() bool
	KeySpecificMaxGasPriceWei(addr common.Address) *assets.Wei
	KeeperDefaultTransactionQueueDepth() uint32
	KeeperGasPriceBufferPercent() uint16
	KeeperGasTipCapBufferPercent() uint16
	KeeperBaseFeeBufferPercent() uint16
	KeeperMaximumGracePeriod() int64
	KeeperRegistryCheckGasOverhead() uint32
	KeeperRegistryPerformGasOverhead() uint32
	KeeperRegistryMaxPerformDataSize() uint32
	KeeperRegistrySyncInterval() time.Duration
	KeeperRegistrySyncUpkeepQueueSize() uint32
	KeeperTurnLookBack() int64
	pg.QConfig
}

type RegistryGasChecker interface {
	KeeperRegistryCheckGasOverhead() uint32
	KeeperRegistryPerformGasOverhead() uint32
	KeeperRegistryMaxPerformDataSize() uint32
}
