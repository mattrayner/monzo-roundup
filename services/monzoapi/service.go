package monzoapi

import "github.com/mattrayner/monzo-roundup/types/services/monzoapi"

type Service interface {
  monzoapi.Methods
  Start(client *monzoapi.Client)
}

type service struct {
  client *monzoapi.Client
}

func NewService() Service {
  return &service{}
}

func (s *service) Start(client *monzoapi.Client) {
  if client == nil {
    panic("Required argument not supplied")
  }

  s.client = client
}

