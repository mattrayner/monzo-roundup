package monzoapi

import (
  "github.com/maraino/go-mock"
  "github.com/mattrayner/monzo-roundup/types/services/monzoapi"
)

type MockService struct {
  mock.Mock
}

func (s *MockService) Start(monzoClient *monzoapi.Client) {
  s.Called()
}

func (s *MockService) RefreshToken(input *monzoapi.RefreshTokenInput) (*monzoapi.RefreshTokenOutput, error) {
  ret := s.Called(input)
  return ret.Get(0).(*monzoapi.RefreshTokenOutput), ret.Error(1)
}

func (s *MockService) GetTransaction(input *monzoapi.GetTransactionInput) (*monzoapi.GetTransactionOutput, error) {
  ret := s.Called(input)
  return ret.Get(0).(*monzoapi.GetTransactionOutput), ret.Error(1)
}

func (s *MockService) GetCoinJar(input *monzoapi.GetCoinJarInput) (*monzoapi.GetCoinJarOutput, error) {
  ret := s.Called(input)
  return ret.Get(0).(*monzoapi.GetCoinJarOutput), ret.Error(1)
}

func (s *MockService) Deposit(input *monzoapi.DepositInput) (*monzoapi.DepositOutput, error) {
  ret := s.Called(input)
  return ret.Get(0).(*monzoapi.DepositOutput), ret.Error(1)
}
