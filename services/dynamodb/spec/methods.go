package spec

import (
  . "github.com/onsi/ginkgo"
  . "github.com/onsi/gomega"
  "gopkg.in/jarcoal/httpmock.v1"
  "github.com/mattrayner/monzo-roundup/types/services/dynamodb"
  uut "github.com/mattrayner/monzo-roundup/services/dynamodb"
)

var _ = Describe("Service", func() {
  var (
    service uut.Service
  )

  BeforeEach(func() {
    service = uut.NewService()
    service.Start()
  })

  Describe("GetUser", func() {
    Context("with a successful response from DynamoDB", func() {
      Context("with no matching user", func() {
        BeforeEach(func() {
          httpmock.RegisterResponder("POST", "https://dynamodb.eu-west-1.amazonaws.com/",
            httpmock.NewStringResponder(200, `{"ConsumedCapacity": {"CapacityUnits": 25,"Table": {"CapacityUnits": 25},"TableName": "monzo-roundup"}}`))
        })

        It("returns an error", func() {
          _, err := service.GetUser(&dynamodb.GetUserInput{ AccountID: "acc_1234" })

          Expect(err).To(HaveOccurred())
          Expect(err.Error()).To(Equal("Unable to find a DynamoDB item for (id: acc_1234)"))
        })
      })

      Context("with a matching user", func() {
        Context("with an error un-marshalling", func() {
          BeforeEach(func() {
            httpmock.RegisterResponder("POST", "https://dynamodb.eu-west-1.amazonaws.com/",
              httpmock.NewStringResponder(200, `Random non-json response`))
          })

          It("returns an error", func() {
            _, err := service.GetUser(&dynamodb.GetUserInput{ AccountID: "acc_1234" })

            Expect(err).To(HaveOccurred())
            Expect(err.Error()).To(Equal("Error getting item from DyanamoDB (id: acc_1234): SerializationError: failed decoding JSON RPC response\ncaused by: invalid character 'R' looking for beginning of value"))
          })
        })

        Context("with a successful un-marshalling", func() {
          BeforeEach(func() {
            httpmock.RegisterResponder("POST", "https://dynamodb.eu-west-1.amazonaws.com/",
              httpmock.NewStringResponder(200, `{"ConsumedCapacity": {"CapacityUnits": 25,"Table": {"CapacityUnits": 25},"TableName": "monzo-roundup"},"Item": {"id" : {"S": "acc_1234"},"authToken" : {"S": "auth_1234"},"refreshToken" : {"S": "refresh_1234"}}}`))
          })

          It("returns an error", func() {
            output, err := service.GetUser(&dynamodb.GetUserInput{ AccountID: "acc_1234" })

            Expect(err).NotTo(HaveOccurred())

            Expect(output.User.AccountID).To(Equal("acc_1234"))
            Expect(output.User.AuthToken).To(Equal("auth_1234"))
            Expect(output.User.RefreshToken).To(Equal("refresh_1234"))
          })
        })
      })
    })

    Context("with an unsuccessful response from DynamoDB", func() {
      BeforeEach(func() {
        httpmock.RegisterResponder("POST", "https://dynamodb.eu-west-1.amazonaws.com/",
          httpmock.NewStringResponder(404, `Not Found`))
      })

      It("returns an error", func() {
        _, err := service.GetUser(&dynamodb.GetUserInput{ AccountID: "acc_1234" })

        Expect(err).To(HaveOccurred())
        Expect(err.Error()).To(Equal("Error getting item from DyanamoDB (id: acc_1234): SerializationError: failed decoding JSON RPC error response\ncaused by: invalid character 'N' looking for beginning of value"))
      })
    })
  })

  Describe("UpdateUser", func() {
    Context("with a successful response from DynamoDB", func() {
      BeforeEach(func() {
        httpmock.RegisterResponder("POST", "https://dynamodb.eu-west-1.amazonaws.com/",
          httpmock.NewStringResponder(200, `{"Attributes": { "refreshToken": {"S": "refresh_1234"}, "authToken": {"S": "auth_1234"} }, "ConsumedCapacity": {"CapacityUnits": 25,"Table": {"CapacityUnits": 25},"TableName": "monzo-roundup"}}`))
      })

      It("returns an error", func() {
        output, err := service.UpdateUser(&dynamodb.UpdateUserInput{ User: &dynamodb.User{ AccountID: "acc_1234", AuthToken: "auth_1234", RefreshToken: "refresh_1234" } })

        Expect(output.User.AccountID).To(Equal(""))
        Expect(output.User.AuthToken).To(Equal("auth_1234"))
        Expect(output.User.RefreshToken).To(Equal("refresh_1234"))
        Expect(err).NotTo(HaveOccurred())
      })
    })

    Context("with an unsuccessful response from DynamoDB", func() {
      BeforeEach(func() {
        httpmock.RegisterResponder("POST", "https://dynamodb.eu-west-1.amazonaws.com/",
          httpmock.NewStringResponder(404, `Not Found`))
      })

      It("returns an error", func() {
        _, err := service.UpdateUser(&dynamodb.UpdateUserInput{ User: &dynamodb.User{ AccountID: "acc_1234", AuthToken: "auth_1234", RefreshToken: "refresh_1234" } })

        Expect(err).To(HaveOccurred())
        Expect(err.Error()).To(Equal("Error updating item in DynamoDB: SerializationError: failed decoding JSON RPC error response\ncaused by: invalid character 'N' looking for beginning of value"))
      })
    })
  })
})
