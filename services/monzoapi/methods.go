package monzoapi

import "github.com/mattrayner/monzo-roundup/types/services/monzoapi"

func (s *service) RefreshToken (input *monzoapi.RefreshTokenInput) (*monzoapi.RefreshTokenOutput, error) {}

func (s *service) GetTransaction (input *monzoapi.GetTransactionInput) (*monzoapi.GetTransactionOutput, error) {}

func (s *service) GetCoinJar (input *monzoapi.GetCoinJarInput) (*monzoapi.GetCoinJarOutput, error) {}

func (s *service) Deposit (input *monzoapi.DepositInput) (*monzoapi.DepositOutput, error) {
  
}