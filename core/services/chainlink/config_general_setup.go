package chainlink

import (
	v2 "github.com/GoPlugin/pluginV2/core/config/v2"
)

func (g *generalConfig) EthereumNodes() string { panic(v2.ErrUnsupported) }
func (g *generalConfig) SolanaNodes() string   { panic(v2.ErrUnsupported) }
func (g *generalConfig) StarkNetNodes() string { panic(v2.ErrUnsupported) }
