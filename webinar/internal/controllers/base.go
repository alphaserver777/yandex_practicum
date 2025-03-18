package controllers

import (
	"net/http"

	"github.com/go-chi/chi/v5"
)

type Logger interface {
	Info(args ...interface{})
}

type BaseController struct {
	logger Logger
}

func NewBaseController(logger Logger) *BaseController {
	return &BaseController{
		logger: logger,
	}
}

func (c *BaseController) Route() *chi.Mux {
	r := chi.NewRouter()
	r.Get("/", c.handleMain)
	r.Get("/{name}", c.handleName)
	return r
}

func (c *BaseController) handleMain(writer http.ResponseWriter, request *http.Request) {
	c.logger.Info("main")
	writer.Write([]byte("Hello - main"))
}

func (c *BaseController) handleName(writer http.ResponseWriter, request *http.Request) {
	c.logger.Info("name")
	name := chi.URLParam(request, "name")
	writer.Write([]byte("Hello - " + name))
}
