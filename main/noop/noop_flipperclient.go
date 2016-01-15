package main

import (
	"github.com/jrmConduce/redishappy/services/logger"
	"github.com/jrmConduce/redishappy/types"
	"github.com/jrmConduce/redishappy/util"
)

type NoOpFlipperClient struct {
}

func NewNoOpFlipper() *NoOpFlipperClient {
	return &NoOpFlipperClient{}
}

func (*NoOpFlipperClient) InitialiseRunningState(state *types.MasterDetailsCollection) {
	logger.NoteWorthy.Printf("InitialiseRunningState called : %s", util.String(state.Items()))
}

func (*NoOpFlipperClient) Orchestrate(switchEvent types.MasterSwitchedEvent) {
	logger.NoteWorthy.Printf("Orchestrate called : %s", util.String(switchEvent))
}
