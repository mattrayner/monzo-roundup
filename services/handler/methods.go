package handler

import (
  "github.com/mattrayner/monzo-roundup/types/services/handler"
  "github.com/mattrayner/monzo-roundup/types/services/monzoapi"
  "github.com/mattrayner/monzo-roundup/types/services/dynamodb"

  "errors"
)

func (s *service) Handle(input *handler.HandleInput) (handler.HandleOutput, error) {
  getUserOutput, err := s.dynamoDB.GetUser(&dynamodb.GetUserInput{ AccountID: input.AccountID }); if err != nil {
    return handler.HandleOutput{}, errors.New("unable to get user")
  }

  refreshTokenOutput, err := s.monzo.RefreshToken(&monzoapi.RefreshTokenInput{ RefreshToken: getUserOutput.User.RefreshToken, MonzoClient: s.monzoClient }); if err != nil {
    return handler.HandleOutput{}, errors.New("unable to refresh token")
  }

  user := &dynamodb.User{ AccountID: input.AccountID, RefreshToken: refreshTokenOutput.RefreshToken, AuthToken: refreshTokenOutput.AuthToken }
  _, err = s.dynamoDB.UpdateUser(&dynamodb.UpdateUserInput{ User: user }); if err != nil {
    return handler.HandleOutput{}, errors.New("unable to update user")
  }

  transaction, err := s.monzo.GetTransaction(&monzoapi.GetTransactionInput{ TransactionID: input.TransactionID, MonzoClient: s.monzoClient, AuthKey: user.AuthToken }); if err != nil {
    return handler.HandleOutput{}, errors.New("unable to get transaction")
  }

  remainder := 100 + transaction.Amount

  if remainder <= 0 || remainder >= 100 {
    return handler.HandleOutput{}, errors.New("nothing to round")
  }

  getCoinJarOutput, err := s.monzo.GetCoinJar(&monzoapi.GetCoinJarInput{ MonzoClient: s.monzoClient, AuthKey: user.AuthToken }); if err != nil {
    return handler.HandleOutput{}, errors.New("unable to get a Coin Jar pot")
  }

  depositOutput, err := s.monzo.Deposit(&monzoapi.DepositInput{ Amount: remainder, Pot: getCoinJarOutput.Pot, MonzoClient: s.monzoClient, AuthKey: user.AuthToken }); if err != nil {
    return handler.HandleOutput{}, errors.New("unable to deposit into pot")
  }
  
  diff := depositOutput.Pot.Balance - getCoinJarOutput.Pot.Balance
  if diff != remainder {
    return handler.HandleOutput{}, errors.New("pot balance did not increase by remainder")
  }
  
  return handler.HandleOutput{ Balance: depositOutput.Pot.Balance, Remainder: remainder }, nil
}
