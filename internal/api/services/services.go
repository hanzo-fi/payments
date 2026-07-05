package services

import (
	"github.com/hanzo-fi/payments/internal/connectors/engine"
	"github.com/hanzo-fi/payments/internal/storage"
)

type Service struct {
	storage storage.Storage
	engine  engine.Engine
	debug   bool
}

func New(storage storage.Storage, engine engine.Engine, debug bool) *Service {
	return &Service{
		storage: storage,
		engine:  engine,
		debug:   debug,
	}
}
