package spec

import (
  . "github.com/onsi/ginkgo"
  . "github.com/onsi/gomega"
  uut "github.com/mattrayner/monzo-roundup/services/handler"
  "github.com/mattrayner/monzo-roundup/types/services/handler"
  _monzoapi "github.com/mattrayner/monzo-roundup/types/services/monzoapi"
  _dynamodb "github.com/mattrayner/monzo-roundup/types/services/dynamodb"
  "github.com/mattrayner/monzo-roundup/services/dynamodb"
  "github.com/mattrayner/monzo-roundup/services/monzoapi"
  "errors"
)

var _ = Describe("Service", func() {
  var (
    service uut.Service
    ddb *dynamodb.MockService
    monzo *monzoapi.MockService
    client *_monzoapi.Client
  )

  BeforeEach(func() {
    service = uut.NewService()

    ddb    = &dynamodb.MockService{}
    monzo  = &monzoapi.MockService{}
    client = &_monzoapi.Client{ ClientID: "id_1234", ClientSecret: "secret_1234" }

    service.Start(ddb, monzo, client)
  })

  Describe("Handler", func() {
    Context("with full data", func() {
      It("should return the expected balance and remainder", func() {
        ddb.When("GetUser",  &_dynamodb.GetUserInput{ AccountID: "acc_1234" }).Return(&_dynamodb.GetUserOutput{ User: &_dynamodb.User{ AccountID: "acc_1234", AuthToken: "auth_1234", RefreshToken: "refresh_1234" } }, nil).Times(1)
        monzo.When("RefreshToken",  &_monzoapi.RefreshTokenInput{ RefreshToken: "refresh_1234", MonzoClient: client }).Return(&_monzoapi.RefreshTokenOutput{ AuthToken: "auth_5678", RefreshToken: "refresh_5678" }, nil).Times(1)
        ddb.When("UpdateUser",  &_dynamodb.UpdateUserInput{ User: &_dynamodb.User{ AccountID: "acc_1234", AuthToken: "auth_5678", RefreshToken: "refresh_5678" } }).Return(&_dynamodb.UpdateUserOutput{ User: &_dynamodb.User{ AccountID: "acc_1234", AuthToken: "auth_1234", RefreshToken: "refresh_1234" } }, nil).Times(1)
        monzo.When("GetTransaction", &_monzoapi.GetTransactionInput{ TransactionID: "tx_1234", MonzoClient: client, AuthKey: "auth_5678" }).Return(&_monzoapi.GetTransactionOutput{ 	TransactionID: "tx_1234", Created: "2018-07-31T00:00:00", Description: "Test transaction", Amount: -79, LocalAmount: -79, Originator: false }, nil).Times(1)
        monzo.When("GetCoinJar", &_monzoapi.GetCoinJarInput{ MonzoClient: client, AuthKey: "auth_5678" }).Return(&_monzoapi.GetCoinJarOutput{ 	Pot: &_monzoapi.Pot{ PotID: "pot_1234", Name: "Coin Jar", Type: "pig", Balance: 79, Currency: "GBP", Created: "2018-06-30T00:00:00", Updated: "2018-07-04T00:00:00", Deleted: false } }, nil).Times(1)
        monzo.When("Deposit", &_monzoapi.DepositInput{ Amount: 21, Pot: &_monzoapi.Pot{ PotID: "pot_1234", Name: "Coin Jar", Type: "pig", Balance: 79, Currency: "GBP", Created: "2018-06-30T00:00:00", Updated: "2018-07-04T00:00:00", Deleted: false }, MonzoClient: client, AuthKey: "auth_5678" }).Return(&_monzoapi.DepositOutput{ 	Pot: &_monzoapi.Pot{ PotID: "pot_1234", Name: "Coin Jar", Type: "pig", Balance: 100, Currency: "GBP", Created: "2018-06-30T00:00:00", Updated: "2018-07-04T00:00:00", Deleted: false } }, nil).Times(1)

        out, err := service.Handle(&handler.HandleInput{ AccountID: "acc_1234", TransactionID: "tx_1234" })
        Expect(err).ToNot(HaveOccurred())

        Expect(out.Balance).To(Equal(int32(100)))
        Expect(out.Remainder).To(Equal(int32(21)))
      })
    })

    Context("with a failed deposit", func() {
      It("should return the expected error", func() {
        ddb.When("GetUser",  &_dynamodb.GetUserInput{ AccountID: "acc_1234" }).Return(&_dynamodb.GetUserOutput{ User: &_dynamodb.User{ AccountID: "acc_1234", AuthToken: "auth_1234", RefreshToken: "refresh_1234" } }, nil).Times(1)
        monzo.When("RefreshToken",  &_monzoapi.RefreshTokenInput{ RefreshToken: "refresh_1234", MonzoClient: client }).Return(&_monzoapi.RefreshTokenOutput{ AuthToken: "auth_5678", RefreshToken: "refresh_5678" }, nil).Times(1)
        ddb.When("UpdateUser",  &_dynamodb.UpdateUserInput{ User: &_dynamodb.User{ AccountID: "acc_1234", AuthToken: "auth_5678", RefreshToken: "refresh_5678" } }).Return(&_dynamodb.UpdateUserOutput{ User: &_dynamodb.User{ AccountID: "acc_1234", AuthToken: "auth_1234", RefreshToken: "refresh_1234" } }, nil).Times(1)
        monzo.When("GetTransaction", &_monzoapi.GetTransactionInput{ TransactionID: "tx_1234", MonzoClient: client, AuthKey: "auth_5678" }).Return(&_monzoapi.GetTransactionOutput{ 	TransactionID: "tx_1234", Created: "2018-07-31T00:00:00", Description: "Test transaction", Amount: -79, LocalAmount: -79, Originator: false }, nil).Times(1)
        monzo.When("GetCoinJar", &_monzoapi.GetCoinJarInput{ MonzoClient: client, AuthKey: "auth_5678" }).Return(&_monzoapi.GetCoinJarOutput{ 	Pot: &_monzoapi.Pot{ PotID: "pot_1234", Name: "Coin Jar", Type: "pig", Balance: 79, Currency: "GBP", Created: "2018-06-30T00:00:00", Updated: "2018-07-04T00:00:00", Deleted: false } }, nil).Times(1)
        monzo.When("Deposit", &_monzoapi.DepositInput{ Amount: 21, Pot: &_monzoapi.Pot{ PotID: "pot_1234", Name: "Coin Jar", Type: "pig", Balance: 79, Currency: "GBP", Created: "2018-06-30T00:00:00", Updated: "2018-07-04T00:00:00", Deleted: false }, MonzoClient: client, AuthKey: "auth_5678" }).Return(&_monzoapi.DepositOutput{}, errors.New("deposit error")).Times(1)

        _, err := service.Handle(&handler.HandleInput{ AccountID: "acc_1234", TransactionID: "tx_1234" })
        Expect(err).To(HaveOccurred())
        Expect(err.Error()).To(Equal("unable to deposit into pot"))
      })
    })

    Context("error getting coin jar", func() {
      It("should return the expected error", func() {
        ddb.When("GetUser",  &_dynamodb.GetUserInput{ AccountID: "acc_1234" }).Return(&_dynamodb.GetUserOutput{ User: &_dynamodb.User{ AccountID: "acc_1234", AuthToken: "auth_1234", RefreshToken: "refresh_1234" } }, nil).Times(1)
        monzo.When("RefreshToken",  &_monzoapi.RefreshTokenInput{ RefreshToken: "refresh_1234", MonzoClient: client }).Return(&_monzoapi.RefreshTokenOutput{ AuthToken: "auth_5678", RefreshToken: "refresh_5678" }, nil).Times(1)
        ddb.When("UpdateUser",  &_dynamodb.UpdateUserInput{ User: &_dynamodb.User{ AccountID: "acc_1234", AuthToken: "auth_5678", RefreshToken: "refresh_5678" } }).Return(&_dynamodb.UpdateUserOutput{ User: &_dynamodb.User{ AccountID: "acc_1234", AuthToken: "auth_1234", RefreshToken: "refresh_1234" } }, nil).Times(1)
        monzo.When("GetTransaction", &_monzoapi.GetTransactionInput{ TransactionID: "tx_1234", MonzoClient: client, AuthKey: "auth_5678" }).Return(&_monzoapi.GetTransactionOutput{ 	TransactionID: "tx_1234", Created: "2018-07-31T00:00:00", Description: "Test transaction", Amount: -79, LocalAmount: -79, Originator: false }, nil).Times(1)
        monzo.When("GetCoinJar", &_monzoapi.GetCoinJarInput{ MonzoClient: client, AuthKey: "auth_5678" }).Return(&_monzoapi.GetCoinJarOutput{}, errors.New("error finding coin jar")).Times(1)

        _, err := service.Handle(&handler.HandleInput{ AccountID: "acc_1234", TransactionID: "tx_1234" })
        Expect(err).To(HaveOccurred())
        Expect(err.Error()).To(Equal("unable to get a Coin Jar pot"))
      })
    })

    Context("error getting transaction", func() {
      It("should return the expected error", func() {
        ddb.When("GetUser",  &_dynamodb.GetUserInput{ AccountID: "acc_1234" }).Return(&_dynamodb.GetUserOutput{ User: &_dynamodb.User{ AccountID: "acc_1234", AuthToken: "auth_1234", RefreshToken: "refresh_1234" } }, nil).Times(1)
        monzo.When("RefreshToken",  &_monzoapi.RefreshTokenInput{ RefreshToken: "refresh_1234", MonzoClient: client }).Return(&_monzoapi.RefreshTokenOutput{ AuthToken: "auth_5678", RefreshToken: "refresh_5678" }, nil).Times(1)
        ddb.When("UpdateUser",  &_dynamodb.UpdateUserInput{ User: &_dynamodb.User{ AccountID: "acc_1234", AuthToken: "auth_5678", RefreshToken: "refresh_5678" } }).Return(&_dynamodb.UpdateUserOutput{ User: &_dynamodb.User{ AccountID: "acc_1234", AuthToken: "auth_1234", RefreshToken: "refresh_1234" } }, nil).Times(1)
        monzo.When("GetTransaction", &_monzoapi.GetTransactionInput{ TransactionID: "tx_1234", MonzoClient: client, AuthKey: "auth_5678" }).Return(&_monzoapi.GetTransactionOutput{}, errors.New("error getting transaction")).Times(1)

        _, err := service.Handle(&handler.HandleInput{ AccountID: "acc_1234", TransactionID: "tx_1234" })
        Expect(err).To(HaveOccurred())
        Expect(err.Error()).To(Equal("unable to get transaction"))
      })
    })

    Context("error updating user", func() {
      It("should return the expected error", func() {
        ddb.When("GetUser",  &_dynamodb.GetUserInput{ AccountID: "acc_1234" }).Return(&_dynamodb.GetUserOutput{ User: &_dynamodb.User{ AccountID: "acc_1234", AuthToken: "auth_1234", RefreshToken: "refresh_1234" } }, nil).Times(1)
        monzo.When("RefreshToken",  &_monzoapi.RefreshTokenInput{ RefreshToken: "refresh_1234", MonzoClient: client }).Return(&_monzoapi.RefreshTokenOutput{ AuthToken: "auth_5678", RefreshToken: "refresh_5678" }, nil).Times(1)
        ddb.When("UpdateUser",  &_dynamodb.UpdateUserInput{ User: &_dynamodb.User{ AccountID: "acc_1234", AuthToken: "auth_5678", RefreshToken: "refresh_5678" } }).Return(&_dynamodb.UpdateUserOutput{}, errors.New("error updating user")).Times(1)

        _, err := service.Handle(&handler.HandleInput{ AccountID: "acc_1234", TransactionID: "tx_1234" })
        Expect(err).To(HaveOccurred())
        Expect(err.Error()).To(Equal("unable to update user"))
      })
    })

    Context("error refreshing token", func() {
      It("should return the expected error", func() {
        ddb.When("GetUser",  &_dynamodb.GetUserInput{ AccountID: "acc_1234" }).Return(&_dynamodb.GetUserOutput{ User: &_dynamodb.User{ AccountID: "acc_1234", AuthToken: "auth_1234", RefreshToken: "refresh_1234" } }, nil).Times(1)
        monzo.When("RefreshToken",  &_monzoapi.RefreshTokenInput{ RefreshToken: "refresh_1234", MonzoClient: client }).Return(&_monzoapi.RefreshTokenOutput{}, errors.New("error refreshing token")).Times(1)

        _, err := service.Handle(&handler.HandleInput{ AccountID: "acc_1234", TransactionID: "tx_1234" })
        Expect(err).To(HaveOccurred())
        Expect(err.Error()).To(Equal("unable to refresh token"))
      })
    })

    Context("error getting user", func() {
      It("should return the expected error", func() {
        ddb.When("GetUser",  &_dynamodb.GetUserInput{ AccountID: "acc_1234" }).Return(&_dynamodb.GetUserOutput{}, errors.New("error getting user")).Times(1)

        _, err := service.Handle(&handler.HandleInput{ AccountID: "acc_1234", TransactionID: "tx_1234" })
        Expect(err).To(HaveOccurred())
        Expect(err.Error()).To(Equal("unable to get user"))
      })
    })
  })
})