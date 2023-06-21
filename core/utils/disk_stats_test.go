package utils_test

import (
	"testing"

	"github.com/pluginV2/core/utils"
	"github.com/stretchr/testify/assert"
)

func TestDiskStatsProvider_AvailableSpace(t *testing.T) {
	t.Parallel()

	provider := utils.NewDiskStatsProvider()

	size, err := provider.AvailableSpace(".")
	assert.NoError(t, err)
	assert.NotZero(t, size)

	_, err = provider.AvailableSpace("")
	assert.Error(t, err)
}
