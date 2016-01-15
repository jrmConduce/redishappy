package api

import (
	"net/http"

	"github.com/jrmConduce/redishappy/configuration"
	"github.com/jrmConduce/redishappy/util"
)

type ConfigurationApi struct {
	ConfigurationManager *configuration.ConfigurationManager
}

func (s *ConfigurationApi) Get(w http.ResponseWriter, r *http.Request) {
	config := s.ConfigurationManager.GetCurrentConfiguration()
	util.WriteResponseAsJSON(w, config)
}
