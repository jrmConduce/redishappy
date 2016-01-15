package sentinel

import (
	"testing"

	"github.com/jrmConduce/redishappy/configuration"
	"github.com/jrmConduce/redishappy/types"
)

func TestBasicEventChannel(t *testing.T) {

	switchmasterchannel := make(chan types.MasterSwitchedEvent)

	manager := NewManager(switchmasterchannel, configuration.NewConfigurationManager(configuration.Configuration{}))

	manager.Notify(&SentinelAdded{Sentinel: types.Sentinel{Host: "10.1.1.1", Port: 12345}})

	responseChannel := make(chan SentinelTopology)

	manager.GetState(TopologyRequest{ReplyChannel: responseChannel})
	topologyState := <-responseChannel

	if len(topologyState.Sentinels) != 1 {
		t.Error("Topology count should be 1")
	}
}
