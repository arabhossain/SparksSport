package providers

import (
	"SparksSport/cmd/app/core/config"
	"github.com/gorilla/mux"
	"gorm.io/gorm"
	"net/http"
)

type ServiceProvider interface {
	Register(*ServiceContainer)
}

type Response struct {
	Success func(http.ResponseWriter, *http.Request, interface{}, string, int)
	Error   func(http.ResponseWriter, *http.Request, string, int)
}

type ServiceContainer struct {
	Config   core_config.Config
	Router   *mux.Router
	DB       *gorm.DB
	Response Response
	Services map[string]interface{}
}

func (c *ServiceContainer) RegisterService(name string, service interface{}) {
	c.Services[name] = service
}

func (c *ServiceContainer) ResolveService(name string) interface{} {
	return c.Services[name]
}
