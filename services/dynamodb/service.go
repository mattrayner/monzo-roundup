package dynamodb

import (
  "github.com/mattrayner/monzo-roundup/types/services/dynamodb"
  _dynamodb "github.com/aws/aws-sdk-go/service/dynamodb"
  "log"
)

type Service interface {
  dynamodb.Methods
  Start()
}

type service struct {
  srv *_dynamodb.DynamoDB
}

func (s *service) Start() {
  log.Println("DynamoDB service started")
}

func NewService() Service {
  return &service{}
}
