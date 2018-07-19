package handler

import (
  "github.com/mattrayner/monzo-roundup/types/services/handler"
  "github.com/mattrayner/monzo-roundup/types/services/monzoapi"
  "github.com/mattrayner/monzo-roundup/types/services/dynamodb"

  "errors"
)

func (s *service) Handle(input handler.HandleInput) (handler.HandleOutput, error) {
  user, err := s.dynamoDB.GetUser(&dynamodb.GetUserInput{ AccountID: input.AccountID }); if err != nil {
    return handler.HandleOutput{}, errors.New("unable to get user")
  }

  user, err = s.monzo.RefreshToken(user.RefreshToken, s.monzoClient); if err != nil {
    return handler.HandleOutput{}, errors.New("unable to refresh token")
  }

  user, err = s.dynamoDB.UpdateUser(user); if err != nil {
    return handler.HandleOutput{}, errors.New("unable to update user")
  }

  transaction, err := s.monzo.GetTransaction(monzoapi.GetTransactionInput{TransactionID: input.TransactionID}, s.monzoClient); if err != nil {
    return handler.HandleOutput{}, errors.New("unable to get transaction")
  }

  remainder := 100 + transaction.Amount

  if remainder <= 0 || remainder >= 100 {
    return handler.HandleOutput{}, errors.New("nothing to round")
  }

  pot, err := s.monzo.GetCoinJar(s.monzoClient); if err != nil {
    return handler.HandleOutput{}, errors.New("unable to get a Coin Jar pot")
  }

  updatedPot, err := s.monzo.Deposit(remainder, pot); if err != nil {
    return handler.HandleOutput{}, errors.New("unable to deposit into pot")
  }
  
  diff := updatedPot.Balance - pot.Balance
  if diff != remainder {
    return handler.HandleOutput{}, errors.New("pot balance did not increase by remainder")
  }
  
  return handler.HandleOutput{ Balance: updatedPot.Balance, Remainder: remainder }, nil
}
