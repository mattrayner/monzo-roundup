package monzoapi

import "github.com/mattrayner/monzo-roundup/types/services/monzoapi"

type Service interface {
  monzoapi.Methods
  Start(client *monzoapi.MonzoClient)
}

type service struct {
  client *monzoapi.MonzoClient
}

func NewService() Service {
  return &service{}
}

func (s *service) Start(client *monzoapi.MonzoClient) {
  if client == nil {
    panic("Required argument not supplied")
  }

  s.client = client
}

