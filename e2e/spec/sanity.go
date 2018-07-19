package spec

import (
  . "github.com/onsi/ginkgo"
  . "github.com/onsi/gomega"
  "gopkg.in/jarcoal/httpmock.v1"
  "github.com/mattrayner/monzo-roundup"
  "net/http"
  "errors"
)

var _ = Describe("Monzo Roundup", func() {
  Context("with a known user", func() {
    var _ = BeforeEach(func() {
      httpmock.RegisterResponder("POST", "https://dynamodb.eu-west-1.amazonaws.com/",
        httpmock.NewStringResponder(200, `{"ConsumedCapacity": {"CapacityUnits": 25,"Table": {"CapacityUnits": 25},"TableName": "monzo-roundup"},"Item": {"id" : {"S": "acc_1234"},"auth_token" : {"S": "1234"},"refresh_token" : {"S": "5678"}}}`))
    })

    Context("with a successful token refresh", func() {
      var _ = BeforeEach(func() {
        httpmock.RegisterResponder("POST", "https://api.monzo.com/oauth2/token",
          httpmock.NewStringResponder(200, `{"access_token": "access_1234","client_id": "client_1234","expires_in": 21600,"refresh_token": "refresh_1234","token_type": "Bearer","user_id": "usr_1234"}`))
      })

      Context("with a successful user update", func() {
        var _ = BeforeEach(func() {
          count := 0

          httpmock.RegisterResponder("POST", "https://dynamodb.eu-west-1.amazonaws.com/",
            func(req *http.Request) (*http.Response, error) {
              count++

              switch count {
              case 1:
                return httpmock.NewStringResponse(200, `{"ConsumedCapacity": {"CapacityUnits": 25,"Table": {"CapacityUnits": 25},"TableName": "monzo-roundup"},"Item": {"id" : {"S": "acc_1234"},"auth_token" : {"S": "1234"},"refresh_token" : {"S": "5678"}}}`), nil
              case 2:
                return httpmock.NewStringResponse(200, `{"Attributes": { "refresh_token": "refresh_1234", "auth_token": "auth_1234" }, "ConsumedCapacity": {"CapacityUnits": 25,"Table": {"CapacityUnits": 25},"TableName": "monzo-roundup"}}`), nil
              }

              return nil, errors.New("fail")
            },
          )
        })


        Context("with a known transaction", func() {
          var _ = BeforeEach(func() {
            httpmock.RegisterResponder("GET", "https://api.monzo.com/transactions/tx_1234",
              httpmock.NewStringResponder(200, `{"transaction": {"account_balance": 13013,"amount": -79,"created": "2015-08-22T12:20:18Z","currency": "GBP","local_amount": -79,"description": "TEST PAYMENT","id": "tx_1234","merchant": "merch_1234","metadata": {},"notes": "A note","is_load": false,"settled": "2015-08-23T12:20:18Z"}}`))
          })

          Context("with a Coin Jar", func() {
            var _ = BeforeEach(func() {
              httpmock.RegisterResponder("GET", "https://api.monzo.com/pots",
                httpmock.NewStringResponder(200, `{"pots": [{"id": "pot_1234","name": "Coin Jar","style": "beach_ball","balance": 120,"currency": "GBP","created": "2017-11-09T12:30:53.695Z","updated": "2017-11-09T12:30:53.695Z","deleted": false}]}`))
            })

            Context("with a successful deposit", func() {
              var _ = BeforeEach(func() {
                httpmock.RegisterResponder("PUT", "https://api.monzo.com/pots/pot_1234/deposit",
                  httpmock.NewStringResponder(200, `{"id": "pot_1234","name": "Coin Jar","style": "beach_ball","balance": 141,"currency": "GBP","created": "2017-11-09T12:30:53.695Z","updated": "2017-11-09T12:30:53.695Z","deleted": false}`))
              })

              It("responds with the new pot balance message", func() {
                // expect(handler.main(handler.HandlerInput{ User: dynamodb.GetUserInput{ AccountID: "acc_12345" } }))
                resp, err := monzo_roundup.Handler(monzo_roundup.Request{AccountID: "acc_1234", TransactionID: "tx_1234"})

                Expect(err).ToNot(HaveOccurred())
                Expect(resp).To(Equal("Updated pot (pot_1234) old balance: 120, new balance is: 141"))
              })
            })

            Context("with an erroring deposit", func() {
              var _ = BeforeEach(func() {
                httpmock.RegisterResponder("PUT", "https://api.monzo.com/pots/pot_1234/deposit",
                  httpmock.NewStringResponder(401, `{ "error": true }`))
              })

              It("responds with the new expected error message", func() {
                // expect(handler.main(handler.HandlerInput{ User: dynamodb.GetUserInput{ AccountID: "acc_12345" } }))
                _, err := monzo_roundup.Handler(monzo_roundup.Request{AccountID: "acc_1234", TransactionID: "tx_1234"})

                Expect(err).To(HaveOccurred())
                Expect(err.Error()).To(Equal("error depositing into Pot: Non-200 Status code. Status code: (401), body: { \"error\": true }"))
              })
            })

            Context("with a non-json response", func() {
              var _ = BeforeEach(func() {
                httpmock.RegisterResponder("PUT", "https://api.monzo.com/pots/pot_1234/deposit",
                  httpmock.NewStringResponder(200, `Random non-json error`))
              })

              It("responds with the new expected error message", func() {
                // expect(handler.main(handler.HandlerInput{ User: dynamodb.GetUserInput{ AccountID: "acc_12345" } }))
                _, err := monzo_roundup.Handler(monzo_roundup.Request{AccountID: "acc_1234", TransactionID: "tx_1234"})

                Expect(err).To(HaveOccurred())
                Expect(err.Error()).To(Equal("error depositing into Pot: Error parsing JSON: invalid character 'R' looking for beginning of value"))
              })
            })
          })

          Context("with a deleted Coin Jar", func() {
            var _ = BeforeEach(func() {
              httpmock.RegisterResponder("GET", "https://api.monzo.com/pots",
                httpmock.NewStringResponder(200, `{"pots": [{"id": "pot_1234","name": "Coin Jar","style": "beach_ball","balance": 120,"currency": "GBP","created": "2017-11-09T12:30:53.695Z","updated": "2017-11-09T12:30:53.695Z","deleted": true}, {"id": "pot_1234","name": "Savings","style": "beach_ball","balance": 120,"currency": "GBP","created": "2017-11-09T12:30:53.695Z","updated": "2017-11-09T12:30:53.695Z","deleted": false}]}`))
            })

            It("responds with the new expected error message", func() {
              // expect(handler.main(handler.HandlerInput{ User: dynamodb.GetUserInput{ AccountID: "acc_12345" } }))
              _, err := monzo_roundup.Handler(monzo_roundup.Request{AccountID: "acc_1234", TransactionID: "tx_1234"})

              Expect(err).To(HaveOccurred())
              Expect(err.Error()).To(Equal("error getting Coin Jar: unable to find an active pot named 'Coin Jar'"))
            })
          })

          Context("with no Coin Jar", func() {
            var _ = BeforeEach(func() {
              httpmock.RegisterResponder("GET", "https://api.monzo.com/pots",
                httpmock.NewStringResponder(200, `{"pots": []}`))
            })

            It("responds with the new expected error message", func() {
              // expect(handler.main(handler.HandlerInput{ User: dynamodb.GetUserInput{ AccountID: "acc_12345" } }))
              _, err := monzo_roundup.Handler(monzo_roundup.Request{AccountID: "acc_1234", TransactionID: "tx_1234"})

              Expect(err).To(HaveOccurred())
              Expect(err.Error()).To(Equal("error getting Coin Jar: unable to find an active pot named 'Coin Jar'"))
            })
          })

          Context("with an error getting pots", func() {
            var _ = BeforeEach(func() {
              httpmock.RegisterResponder("GET", "https://api.monzo.com/pots",
                httpmock.NewStringResponder(500, `{"error": true}`))
            })

            It("responds with the new expected error message", func() {
              // expect(handler.main(handler.HandlerInput{ User: dynamodb.GetUserInput{ AccountID: "acc_12345" } }))
              _, err := monzo_roundup.Handler(monzo_roundup.Request{AccountID: "acc_1234", TransactionID: "tx_1234"})

              Expect(err).To(HaveOccurred())
              Expect(err.Error()).To(Equal("error getting Coin Jar: error getting pots from Monzo: Non-200 Status code. Status code: (500), body: {\"error\": true}"))
            })
          })

          Context("with a non-json response", func() {
            var _ = BeforeEach(func() {
              httpmock.RegisterResponder("GET", "https://api.monzo.com/pots",
                httpmock.NewStringResponder(200, `Random non-json response`))
            })

            It("responds with the new expected error message", func() {
              // expect(handler.main(handler.HandlerInput{ User: dynamodb.GetUserInput{ AccountID: "acc_12345" } }))
              _, err := monzo_roundup.Handler(monzo_roundup.Request{AccountID: "acc_1234", TransactionID: "tx_1234"})

              Expect(err).To(HaveOccurred())
              Expect(err.Error()).To(Equal("error getting Coin Jar: error parsing JSON: invalid character 'R' looking for beginning of value"))
            })
          })
        })

        Context("with an error getting transaction", func() {
          var _ = BeforeEach(func() {
            httpmock.RegisterResponder("GET", "https://api.monzo.com/transactions/tx_1234",
              httpmock.NewStringResponder(500, `{"error": true}`))
          })

          It("responds with the expected error", func() {
            // expect(handler.main(handler.HandlerInput{ User: dynamodb.GetUserInput{ AccountID: "acc_12345" } }))
            _, err := monzo_roundup.Handler(monzo_roundup.Request{AccountID: "acc_1234", TransactionID: "tx_1234"})

            Expect(err).To(HaveOccurred())
            Expect(err.Error()).To(MatchRegexp(""))
          })
        })

        Context("with an unknown transaction", func() {
          var _ = BeforeEach(func() {
            httpmock.RegisterResponder("GET", "https://api.monzo.com/transactions/tx_1234",
              httpmock.NewStringResponder(404, `{"code": "not_found.transaction_not_found","message": "Transaction not found"}`))
          })

          It("responds with the expected error", func() {
            // expect(handler.main(handler.HandlerInput{ User: dynamodb.GetUserInput{ AccountID: "acc_12345" } }))
            _, err := monzo_roundup.Handler(monzo_roundup.Request{AccountID: "acc_1234", TransactionID: "tx_1234"})

            Expect(err).To(HaveOccurred())
            Expect(err.Error()).To(MatchRegexp(""))
          })
        })

        Context("with a non-json response", func() {
          var _ = BeforeEach(func() {
            httpmock.RegisterResponder("GET", "https://api.monzo.com/transactions/tx_1234",
              httpmock.NewStringResponder(200, `Random non-json response`))
          })

          It("responds with the expected error", func() {
            // expect(handler.main(handler.HandlerInput{ User: dynamodb.GetUserInput{ AccountID: "acc_12345" } }))
            _, err := monzo_roundup.Handler(monzo_roundup.Request{AccountID: "acc_1234", TransactionID: "tx_1234"})

            Expect(err).To(HaveOccurred())
            Expect(err.Error()).To(MatchRegexp(""))
          })
        })
      })

      Context("with no change in DynamoDB", func() {
        var _ = BeforeEach(func() {
          count := 0

          httpmock.RegisterResponder("POST", "https://dynamodb.eu-west-1.amazonaws.com/",
            func(req *http.Request) (*http.Response, error) {
              count++

              switch count {
              case 1:
                return httpmock.NewStringResponse(200, `{"ConsumedCapacity": {"CapacityUnits": 25,"Table": {"CapacityUnits": 25},"TableName": "monzo-roundup"},"Item": {"id" : {"S": "acc_1234"},"auth_token" : {"S": "1234"},"refresh_token" : {"S": "5678"}}}`), nil
              case 2:
                return httpmock.NewStringResponse(200, `{"ConsumedCapacity": {"CapacityUnits": 25,"Table": {"CapacityUnits": 25},"TableName": "monzo-roundup"}}`), nil
              }

              return nil, errors.New("fail")
            },
          )
        })

        It("responds with the expected error", func() {
          // expect(handler.main(handler.HandlerInput{ User: dynamodb.GetUserInput{ AccountID: "acc_12345" } }))
          _, err := monzo_roundup.Handler(monzo_roundup.Request{AccountID: "acc_1234", TransactionID: "tx_1234"})

          Expect(err).To(HaveOccurred())
          Expect(err.Error()).To(Equal("Update to DynamoDB did nothing (id: acc_1234)"))
        })
      })

      Context("with an error updating user", func() {
        var _ = BeforeEach(func() {
          count := 0

          httpmock.RegisterResponder("POST", "https://dynamodb.eu-west-1.amazonaws.com/",
            func(req *http.Request) (*http.Response, error) {
              count++

              switch count {
              case 1:
                return httpmock.NewStringResponse(200, `{"ConsumedCapacity": {"CapacityUnits": 25,"Table": {"CapacityUnits": 25},"TableName": "monzo-roundup"},"Item": {"id" : {"S": "acc_1234"},"auth_token" : {"S": "1234"},"refresh_token" : {"S": "5678"}}}`), nil
              case 2:
                return httpmock.NewStringResponse(400, `{"ConsumedCapacity": {"CapacityUnits": 25,"Table": {"CapacityUnits": 25},"TableName": "monzo-roundup"}}`), nil
              }

              return nil, errors.New("fail")
            },
          )
        })

        It("responds with the expected error", func() {
          // expect(handler.main(handler.HandlerInput{ User: dynamodb.GetUserInput{ AccountID: "acc_12345" } }))
          _, err := monzo_roundup.Handler(monzo_roundup.Request{AccountID: "acc_1234", TransactionID: "tx_1234"})

          Expect(err).To(HaveOccurred())
          Expect(err.Error()).To(Equal("Error updating item in DynamoDB: : \n\tstatus code: 400, request id: "))
        })
      })

      Context("with a non-json response", func() {
        var _ = BeforeEach(func() {
          count := 0

          httpmock.RegisterResponder("POST", "https://dynamodb.eu-west-1.amazonaws.com/",
            func(req *http.Request) (*http.Response, error) {
              count++

              switch count {
              case 1:
                return httpmock.NewStringResponse(200, `{"ConsumedCapacity": {"CapacityUnits": 25,"Table": {"CapacityUnits": 25},"TableName": "monzo-roundup"},"Item": {"id" : {"S": "acc_1234"},"auth_token" : {"S": "1234"},"refresh_token" : {"S": "5678"}}}`), nil
              case 2:
                return httpmock.NewStringResponse(200, `Random non-json response`), nil
              }

              return nil, errors.New("fail")
            },
          )
        })

        It("responds with the expected error", func() {
          // expect(handler.main(handler.HandlerInput{ User: dynamodb.GetUserInput{ AccountID: "acc_12345" } }))
          _, err := monzo_roundup.Handler(monzo_roundup.Request{AccountID: "acc_1234", TransactionID: "tx_1234"})

          Expect(err).To(HaveOccurred())
          Expect(err.Error()).To(Equal("Error updating item in DynamoDB: SerializationError: failed decoding JSON RPC response\ncaused by: invalid character 'R' looking for beginning of value"))
        })
      })
    })

    Context("with an error refreshing token", func() {
      var _ = BeforeEach(func() {
        httpmock.RegisterResponder("POST", "https://api.monzo.com/oauth2/token",
          httpmock.NewStringResponder(500, `{"error": true}`))
      })

      It("responds with the expected error", func() {
        // expect(handler.main(handler.HandlerInput{ User: dynamodb.GetUserInput{ AccountID: "acc_12345" } }))
        _, err := monzo_roundup.Handler(monzo_roundup.Request{AccountID: "acc_1234", TransactionID: "tx_1234"})

        Expect(err).To(HaveOccurred())
        Expect(err.Error()).To(Equal("Error refreshing Monzo token: Non-200 Status code. Status code: (500), body: {\"error\": true}"))
      })
    })

    Context("with a non-json response whilst refreshing token", func() {
      var _ = BeforeEach(func() {
        httpmock.RegisterResponder("POST", "https://api.monzo.com/oauth2/token",
          httpmock.NewStringResponder(500, `{"error": true}`))
      })

      It("responds with the expected error", func() {
        // expect(handler.main(handler.HandlerInput{ User: dynamodb.GetUserInput{ AccountID: "acc_12345" } }))
        _, err := monzo_roundup.Handler(monzo_roundup.Request{AccountID: "acc_1234", TransactionID: "tx_1234"})

        Expect(err).To(HaveOccurred())
        Expect(err.Error()).To(Equal("Error refreshing Monzo token: Non-200 Status code. Status code: (500), body: {\"error\": true}"))
      })
    })
  })

  Context("with an unknown user", func() {
    var _ = BeforeEach(func() {
      httpmock.RegisterResponder("POST", "https://dynamodb.eu-west-1.amazonaws.com/",
        httpmock.NewStringResponder(200, `{"ConsumedCapacity": {"CapacityUnits": 25,"Table": {"CapacityUnits": 25},"TableName": "monzo-roundup"}}`))
    })

    It("responds with the expected error", func() {
      // expect(handler.main(handler.HandlerInput{ User: dynamodb.GetUserInput{ AccountID: "acc_12345" } }))
      _, err := monzo_roundup.Handler(monzo_roundup.Request{AccountID: "acc_1234", TransactionID: "tx_1234"})

      Expect(err).To(HaveOccurred())
      Expect(err.Error()).To(Equal("Unable to find a DynamoDB item for (id: acc_1234)"))
    })
  })

  Context("with an error retrieving from DynamoDB", func() {
    var _ = BeforeEach(func() {
      httpmock.RegisterResponder("POST", "https://dynamodb.eu-west-1.amazonaws.com/",
        httpmock.NewStringResponder(401, `{"error": true}`))
    })

    It("responds with the expected error", func() {
      // expect(handler.main(handler.HandlerInput{ User: dynamodb.GetUserInput{ AccountID: "acc_12345" } }))
      _, err := monzo_roundup.Handler(monzo_roundup.Request{AccountID: "acc_1234", TransactionID: "tx_1234"})

      Expect(err).To(HaveOccurred())
      Expect(err.Error()).To(Equal("Error getting item from DyanamoDB (id: acc_1234): : \n\tstatus code: 401, request id: "))
    })
  })

  Context("with a non-json response from DynamoDB", func() {
    var _ = BeforeEach(func() {
      httpmock.RegisterResponder("POST", "https://dynamodb.eu-west-1.amazonaws.com/",
        httpmock.NewStringResponder(200, `Random non-json response`))
    })

    It("responds with the expected error", func() {
      // expect(handler.main(handler.HandlerInput{ User: dynamodb.GetUserInput{ AccountID: "acc_12345" } }))
      _, err := monzo_roundup.Handler(monzo_roundup.Request{AccountID: "acc_1234", TransactionID: "tx_1234"})

      Expect(err).To(HaveOccurred())
      Expect(err.Error()).To(Equal("Error getting item from DyanamoDB (id: acc_1234): SerializationError: failed decoding JSON RPC response\ncaused by: invalid character 'R' looking for beginning of value"))
    })
  })
})
