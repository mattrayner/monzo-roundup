package handler

import (
  "github.com/mattrayner/monzo-roundup/services/dynamodb"
  "github.com/mattrayner/monzo-roundup/services/monzoapi"
  _monzoapi "github.com/mattrayner/monzo-roundup/types/services/monzoapi"
  // _dynamodb "github.com/mattrayner/monzo-roundup/types/services/dynamodb"
  "github.com/mattrayner/monzo-roundup/types/services/handler"
  "log"
)

type Service interface {
  handler.Methods
  Start(dynamoDB dynamodb.Service, monzo monzoapi.Service, monzoClient *_monzoapi.Client)
}

type service struct {
  dynamoDB dynamodb.Service
  monzo monzoapi.Service
  monzoClient *_monzoapi.Client
}

func NewService() Service {
  return &service{}
}

func (s *service) Start(dynamoDB dynamodb.Service, monzo monzoapi.Service, monzoClient *_monzoapi.Client) {
  if dynamoDB == nil || monzo == nil || monzoClient == nil {
    panic("required arguments not given")
  }

  s.dynamoDB    = dynamoDB
  s.monzo       = monzo
  s.monzoClient = monzoClient

  log.Println("Handler service started")
}
