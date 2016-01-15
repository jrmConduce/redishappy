package api

import (
	"net/http"

	"github.com/jrmConduce/redishappy/sentinel"
	"github.com/jrmConduce/redishappy/util"
)

type SentinelApi struct {
	Manager *sentinel.SentinelManager
}

func (s *SentinelApi) Get(w http.ResponseWriter, r *http.Request) {

	responseChannel := make(chan sentinel.SentinelTopology)
	s.Manager.GetState(sentinel.TopologyRequest{ReplyChannel: responseChannel})
	sentinelState := <-responseChannel
	util.WriteResponseAsJSON(w, sentinelState)
}
