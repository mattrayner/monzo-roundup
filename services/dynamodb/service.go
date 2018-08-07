package dynamodb

import (
  "github.com/mattrayner/monzo-roundup/types/services/dynamodb"
  _dynamodb "github.com/aws/aws-sdk-go/service/dynamodb"
  "log"
  "os"
  "github.com/aws/aws-sdk-go/aws"
  "github.com/aws/aws-sdk-go/aws/session"
  "errors"
)

type Service interface {
  dynamodb.Methods
  Start() (error)
}

type service struct {
  srv *_dynamodb.DynamoDB
}

func (s *service) Start() (error) {
  AWS_REGION := os.Getenv("AWS_REGION")
  config := &aws.Config{Region: aws.String(AWS_REGION)}

  sess, err := session.NewSession(config)
  if err != nil {
    return errors.New("error creating new session object")
  }

  s.srv = _dynamodb.New(sess)

  log.Println("DynamoDB service started")

  return nil
}

func NewService() Service {
  return &service{}
}
