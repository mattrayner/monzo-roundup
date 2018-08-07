package monzoapi

import "github.com/mattrayner/monzo-roundup/types/services/monzoapi"

func (s *service) RefreshToken (input *monzoapi.RefreshTokenInput) (*monzoapi.RefreshTokenOutput, error) {
  return &monzoapi.RefreshTokenOutput{}, nil
}

func (s *service) GetTransaction (input *monzoapi.GetTransactionInput) (*monzoapi.GetTransactionOutput, error) {
  return &monzoapi.GetTransactionOutput{}, nil
}

func (s *service) GetCoinJar (input *monzoapi.GetCoinJarInput) (*monzoapi.GetCoinJarOutput, error) {
  return &monzoapi.GetCoinJarOutput{}, nil
}

func (s *service) Deposit (input *monzoapi.DepositInput) (*monzoapi.DepositOutput, error) {
  return &monzoapi.DepositOutput{}, nil
}