package telemetry

import (
	ocrtypes "github.com/GoPlugin/libocr/commontypes"

	"github.com/GoPlugin/pluginV2/core/services/synchronization"
)

type MonitoringEndpointGenerator interface {
	GenMonitoringEndpoint(contractID string, telemType synchronization.TelemetryType) ocrtypes.MonitoringEndpoint
}
