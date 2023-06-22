package v2

import (
	"testing"

	"github.com/GoPlugin/pluginV2/core/services/chainlink/cfgtest"
)

func TestCoreDefaults_notNil(t *testing.T) {
	cfgtest.AssertFieldsNotNil(t, &defaults)
}
