package dynamodb

import (
  "github.com/maraino/go-mock"
  "github.com/mattrayner/monzo-roundup/types/services/dynamodb"
)

type MockService struct {
  mock.Mock
}

func (s *MockService) Start() (error) {
  ret := s.Called()
  return ret.Get(0).(error)
}

func (s *MockService) Stop() {
  s.Called()
}

func (s *MockService) IsStarted() bool {
  return s.Called().Bool(0)
}

func (s *MockService) GetUser(input *dynamodb.GetUserInput) (*dynamodb.GetUserOutput, error) {
  ret := s.Called(input)
  return ret.Get(0).(*dynamodb.GetUserOutput), ret.Error(1)
}

func (s *MockService) UpdateUser(input *dynamodb.UpdateUserInput) (*dynamodb.UpdateUserOutput, error) {
  ret := s.Called(input)
  return ret.Get(0).(*dynamodb.UpdateUserOutput), ret.Error(1)
}

