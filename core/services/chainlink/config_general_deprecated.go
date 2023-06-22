package chainlink

import (
	"time"

	config "github.com/GoPlugin/pluginV2/core/config/v2"
)

func (g *generalConfig) AdvisoryLockCheckInterval() time.Duration { panic(config.ErrUnsupported) }

func (g *generalConfig) AdvisoryLockID() int64 { panic(config.ErrUnsupported) }

func (g *generalConfig) GetAdvisoryLockIDConfiguredOrDefault() int64 { panic(config.ErrUnsupported) }
